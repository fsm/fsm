package fsm

import (
	"regexp"
	"strings"
)

var (
	strippedSlotMatchRegex = regexp.MustCompile(`\\\{[^ ]*\\\}`)
	cleanUtteranceRegex    = regexp.MustCompile("[^a-z0-9 {}]+")
	slotsRegex             = regexp.MustCompile(`\{([^ ]*)\}`)
)

// InputToIntentTransformer converts the input of a platform to an *Intent.
type InputToIntentTransformer func(input interface{}, validIntents []*Intent) (*Intent, map[string]string)

// Intent is an event that occurs that can trigger a transition
type Intent struct {
	Slug          string
	PlatformSlugs map[string]string
	Slots         map[string]*Type
	Utterances    []string
}

// Type is a definition of what a Intent slot value can be
type Type struct {
	Slug          string
	PlatformSlugs map[string]string
	Options       []string
	IsValid       func(string) bool
}

// Parse checks if an input string matches this intent. If the input string
// matches this intent, any parameters are also returned.
func (intent *Intent) Parse(input string) (bool, map[string]string) {
Utterance:
	for _, utterance := range intent.Utterances {
		// Generate regex for the utterance
		cleanedUtterance := cleanUtteranceRegex.ReplaceAllString(strings.ToLower(utterance), "")
		regexStr := "^" + strippedSlotMatchRegex.ReplaceAllString(regexp.QuoteMeta(cleanedUtterance), "(.*)") + "$"
		utteranceRegex := regexp.MustCompile(regexStr)

		// Check if it matches
		matches := utteranceRegex.FindStringSubmatch(input)
		if len(matches) < 1 {
			continue
		}

		// Prepare slotValues
		slotValues := map[string]string{}

		// Get utterance slots
		utteranceMatches := slotsRegex.FindAllStringSubmatch(utterance, -1)
		slots := []string{}
		for _, match := range utteranceMatches {
			slots = append(slots, match[1])
		}

		// Validate all Parameters
		for k, slot := range slots {
			inputValue := matches[k+1]
			slotType := intent.Slots[slot]

			// Make sure it appears in the Options if we have a whitelist
			if len(slotType.Options) > 0 {
				valid := false
				for _, option := range slotType.Options {
					cleanedOption := CleanInput(option)
					if cleanedOption == inputValue {
						valid = true
						break
					}
				}
				if !valid {
					continue Utterance
				}
			} else {
				// If we don't have a whitelist, validate via the IsValid function
				if !slotType.IsValid(inputValue) {
					continue Utterance
				}
			}

			// Valid if we reach this point, add to our slotValues
			slotValues[slot] = inputValue
		}

		// OK
		return true, slotValues
	}

	// We didn't match any of the utterances
	return false, nil
}
