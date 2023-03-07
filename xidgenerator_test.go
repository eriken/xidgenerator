package xidgenerator

import (
	"testing"
)

func TestDefaults(t *testing.T) {
	xid := Generate()

	actualPrefix := xid[:len(defaultPrefix)]
	if actualPrefix != defaultPrefix {
		t.Errorf("default prefix should be %v, but is: %v", defaultPrefix, actualPrefix)
	}

	actualLength := len(xid)
	expectedLength := len(defaultPrefix) + defaultLength
	if actualLength != expectedLength {
		t.Errorf("default length should be %v, but is: %v", expectedLength, actualLength)
	}
}

func TestCustom(t *testing.T) {
	customPrefix := "custom-"
	customLength := 16
	Init(customLength, customPrefix)

	xid := Generate()

	actualPrefix := xid[:len(customPrefix)]
	if actualPrefix != customPrefix {
		t.Errorf("custom prefix should be %v, but is: %v", customPrefix, actualPrefix)
	}

	actualLength := len(xid)
	expectedLength := len(customPrefix) + customLength
	if actualLength != expectedLength {
		t.Errorf("custom length should be %v, but is: %v", expectedLength, actualLength)
	}
}
