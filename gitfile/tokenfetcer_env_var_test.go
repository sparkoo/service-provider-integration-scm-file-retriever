package gitfile

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateHeaderStructFromEnv(t *testing.T) {
	t.Setenv("TOKEN", "abcd_foo")
	headerStruct := new(EnvVarTokenFetcher).BuildHeader("https://github.com/any/test.git")
	assert.Equal(t, "Bearer abcd_foo", headerStruct.Authorization, "Authorization header value mismatch")
}
