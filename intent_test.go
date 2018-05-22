package fsm_test

import (
	"strconv"
	"testing"

	"github.com/fsm/fsm"
)

var (
	TypeGender = &fsm.Type{
		Slug: "gender",
		Options: []string{
			"male",
			"female",
			"other",
		},
	}

	TypeInteger = &fsm.Type{
		Slug: "integer",
		IsValid: func(input string) bool {
			_, err := strconv.ParseInt(input, 10, 64)
			return err == nil
		},
	}

	SampleIntent = &fsm.Intent{
		Slug: "sample-intent",
		Slots: map[string]*fsm.Type{
			"gender": TypeGender,
			"age":    TypeInteger,
		},
		Utterances: []string{
			"I am a {age} year old {gender}.",
			"I am a {gender} and am {age}.",
			"The {gender} is {age}.",
		},
	}
)

func TestStandardParse(t *testing.T) {
	testParse(t, SampleIntent, "I am a 29 year old male.",
		true,
		map[string]string{
			"age":    "29",
			"gender": "male",
		},
	)
}

func TestCapsParse(t *testing.T) {
	testParse(t, SampleIntent, "I AM A 29 YEAR OLD MALE!",
		true,
		map[string]string{
			"age":    "29",
			"gender": "male",
		},
	)
}

func TestDoubleSpacingParse(t *testing.T) {
	testParse(t, SampleIntent, "I am a 29 year  old male.",
		true,
		map[string]string{
			"age":    "29",
			"gender": "male",
		},
	)
}

func TestUtteranceParse(t *testing.T) {
	testParse(t, SampleIntent, "I am a 29 year old male.",
		true,
		map[string]string{
			"age":    "29",
			"gender": "male",
		},
	)

	testParse(t, SampleIntent, "I am a female and am 30",
		true,
		map[string]string{
			"age":    "30",
			"gender": "female",
		},
	)

	testParse(t, SampleIntent, "The male is 30.",
		true,
		map[string]string{
			"age":    "30",
			"gender": "male",
		},
	)
}

func TestUtteranceNoMatch(t *testing.T) {
	testParse(t, SampleIntent, "I am 29.",
		false,
		nil,
	)

	testParse(t, SampleIntent, "",
		false,
		nil,
	)
}

func TestSlotNoMatch(t *testing.T) {
	testParse(t, SampleIntent, "I am a 29a year old male.",
		false,
		nil,
	)

	testParse(t, SampleIntent, "I am a 29 year old dog.",
		false,
		nil,
	)
}

func testParse(t *testing.T, intent *fsm.Intent, input string, expectMatch bool, expectParams map[string]string) {
	// Parse the input
	matches, params := intent.Parse(fsm.CleanInput(input))

	// Validate Match
	if matches != expectMatch {
		expectedMatchStr := "not match"
		if expectMatch {
			expectedMatchStr = "match"
		}
		t.Errorf("Expected `%v` to %v the intent `%v`.", input, expectedMatchStr, intent.Slug)
	}

	// Validate Parameters
	for k, value := range params {
		if expectedValue, ok := expectParams[k]; ok {
			if value != expectedValue {
				t.Errorf("Expected `%v` parameter `%v` to be `%v`", input, k, expectedValue)
			}
		} else {
			t.Errorf("Expected `%v` to contain parameter `%v`", input, k)
		}
	}
}
