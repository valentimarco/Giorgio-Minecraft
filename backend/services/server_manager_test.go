package services

import (
	"testing"
)

func TestJavaVersionFromManifest(t *testing.T) {

	result, err := javaVersionFromManifest("1.20.1")
	expected := "17"
	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
