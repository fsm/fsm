package fsm_test

import (
	"testing"

	"github.com/fsm/fsm"
)

var (
	validIntents = []*fsm.Intent{
		SampleIntent,
	}
)

func TestCleanInput(t *testing.T) {
	if fsm.CleanInput("Hello!") != "hello" {
		t.Fail()
	}
	if fsm.CleanInput("Hello  World") != "hello world" {
		t.Fail()
	}
	if fsm.CleanInput("Hello, World") != "hello world" {
		t.Fail()
	}
}

func TestValidTextInputTransformer(t *testing.T) {
	intent, params := fsm.TextInputTransformer("I am a 29 year old male.", validIntents)
	if intent != SampleIntent {
		t.Fail()
	}
	if params["gender"] != "male" {
		t.Fail()
	}
	if params["age"] != "29" {
		t.Fail()
	}
}

func TestInvalidTextInputTransformer(t *testing.T) {
	intent, params := fsm.TextInputTransformer("hello world", validIntents)
	if intent != nil {
		t.Fail()
	}
	if len(params) != 0 {
		t.Fail()
	}
}
