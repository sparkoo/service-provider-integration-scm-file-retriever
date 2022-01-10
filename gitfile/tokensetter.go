package gitfile

type HeaderStruct struct {
	Authorization string `json:"Authorization"`
}

type TokenSetter interface {
	BuildHeader(repoUrl string) HeaderStruct
}

var TokenSetters []TokenSetter

func init() {
	TokenSetters = []TokenSetter{
		new(EnvVarTokenSetter),
		//new(SecretTokenSetter),
	}
}

func BuildAuthHeader(repoUrl string) HeaderStruct {
	for _, s := range TokenSetters {
		headerStruct := s.BuildHeader(repoUrl)
		if len(headerStruct.Authorization) > 0 {
			return headerStruct
		}
	}
	return HeaderStruct{}
}
