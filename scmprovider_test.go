package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFile(t *testing.T) {
	r1, err := Detect("https://github.com/redhat-appstudio/service-provider-integration-operator", "Makefile", "HEAD")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	assert.Equal(t, r1, "https://raw.githubusercontent.com/redhat-appstudio/service-provider-integration-operator/HEAD/Makefile")
}

func TestGetFile2(t *testing.T) {
	r1, err := Detect("https://github.com/redhat-appstudio/service-provider-integration-operator.git", "Makefile", "HEAD")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	assert.Equal(t, r1, "https://raw.githubusercontent.com/redhat-appstudio/service-provider-integration-operator/HEAD/Makefile")
}
