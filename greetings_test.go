package greetings

import (
	"regexp"
	"testing"
)

func TextHelloName(t *testing.T) {
	name := "Miguel"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello("Miguel")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`hello("Miguel") = %q, /%v, quiere coincidencia para %#q, nil`, msg, err, want)
	}

}

func TextHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`hello("Miguel") = %q, /%v, quiere "", error`, msg, err)
	}
}
