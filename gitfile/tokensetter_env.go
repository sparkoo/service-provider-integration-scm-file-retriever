package gitfile

import "os"

type EnvVarTokenSetter struct{}

func (s *EnvVarTokenSetter) BuildHeader(repoUrl string) HeaderStruct {
	envToken := os.Getenv("TOKEN")
	if len(envToken) == 0 {
		return HeaderStruct{}
	}
	return HeaderStruct{
		"Bearer " + envToken,
	}
}
