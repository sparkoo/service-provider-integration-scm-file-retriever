package gitfile

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFileHead(t *testing.T) {
	r1, err := Detect("https://github.com/redhat-appstudio/service-provider-integration-operator", "Makefile", "HEAD")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	assert.Equal(t, r1, "https://raw.githubusercontent.com/redhat-appstudio/service-provider-integration-operator/HEAD/Makefile")
}

func TestGetFileHead2(t *testing.T) {
	r1, err := Detect("https://github.com/redhat-appstudio/service-provider-integration-operator.git", "Makefile", "HEAD")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	assert.Equal(t, r1, "https://raw.githubusercontent.com/redhat-appstudio/service-provider-integration-operator/HEAD/Makefile")
}

func TestGetFileOnBranch(t *testing.T) {
	r1, err := Detect("https://github.com/redhat-appstudio/service-provider-integration-operator.git", "Makefile", "v0.1.0")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	assert.Equal(t, r1, "https://raw.githubusercontent.com/redhat-appstudio/service-provider-integration-operator/v0.1.0/Makefile")
}

func TestGetFileOnCommitId(t *testing.T) {
	r1, err := Detect("https://github.com/redhat-appstudio/service-provider-integration-operator.git", "Makefile", "efaf08a367921ae130c524db4a531b7696b7d967")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	assert.Equal(t, r1, "https://raw.githubusercontent.com/redhat-appstudio/service-provider-integration-operator/efaf08a367921ae130c524db4a531b7696b7d967/Makefile")
}

func TestGetUnexistingFile(t *testing.T) {
	_, err := Detect("https://github.com/redhat-appstudio/service-provider-integration-operator.git", "Makefile-Non-Exist", "")

	if err == nil {
		t.Error("error expected")
	}
	assert.Equal(t, "unexpected status code from GitHub API: 404. Response: {\"message\":\"Not Found\",\"documentation_url\":\"https://docs.github.com/rest/reference/repos#get-repository-content\"}", fmt.Sprint(err))
}
