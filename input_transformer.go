package fsm

import (
	"regexp"
	"strings"
)

// InputTransformer converts the input of a platform to an *Intent.
type InputTransformer func(input interface{}, validIntents []*Intent) (*Intent, map[string]string)

// TextInputTransformer is an implementation of an InputTransformer
// which handles text as the input type.
func TextInputTransformer(input interface{}, validIntents []*Intent) (*Intent, map[string]string) {
	inputString := CleanInput(input.(string))
	for _, intent := range validIntents {
		matches, params := intent.Parse(inputString)
		if matches {
			return intent, params
		}
	}
	return nil, nil
}

var (
	cleanInputRegex  = regexp.MustCompile("[^a-z0-9 ]+")
	doubleSpaceRegex = regexp.MustCompile(" +")
)

// CleanInput converts the input string to only the following:
// - Lowercase Letters (a-z)
// - Numbers (0-9)
// - Spaces ( )
//
// Uppercase letters are converted to lowercase letters, but any character outside
// of what is noted above is stripped from the string. Double (or more) spaces are
// converted into a single space.
func CleanInput(input string) string {
	return doubleSpaceRegex.ReplaceAllString(cleanInputRegex.ReplaceAllString(strings.ToLower(input), ""), " ")
}
