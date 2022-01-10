package gitfile

type HeaderStruct struct {
	Authorization string `json:"Authorization"`
}

type TokenFetcher interface {
	BuildHeader(repoUrl string) HeaderStruct
}

var TokenFetchers []TokenFetcher

func init() {
	TokenFetchers = []TokenFetcher{
		new(EnvVarTokenFetcher),
		//new(SecretTokenFetcher),
	}
}

func BuildAuthHeader(repoUrl string) HeaderStruct {
	for _, s := range TokenFetchers {
		headerStruct := s.BuildHeader(repoUrl)
		if len(headerStruct.Authorization) > 0 {
			return headerStruct
		}
	}
	return HeaderStruct{}
}
