package gitfile

import (
	"context"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestGetFilePrivateRepo(t *testing.T) {
	if os.Getenv("TOKEN") == "" {
		t.Skip("test skipped as no token found")
	}
	r1, err := GetFileContents(context.Background(), "https://github.com/mshaposhnik/test1", ".devfile.yaml", "", func(url string) {
	})

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	b, err := ioutil.ReadAll(r1)
	assert.Equal(t, "apiVersion: 1.0.0\nmetadata:\n  name: che-github-demo\n\n", string(b))
}
