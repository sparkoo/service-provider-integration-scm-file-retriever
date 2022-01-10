package gitfile

import "os"

type EnvVarTokenFetcher struct{}

func (s *EnvVarTokenFetcher) BuildHeader(repoUrl string) HeaderStruct {
	envToken := os.Getenv("TOKEN")
	if len(envToken) == 0 {
		return HeaderStruct{}
	}
	return HeaderStruct{
		"Bearer " + envToken,
	}
}
