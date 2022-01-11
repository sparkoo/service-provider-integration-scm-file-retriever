package gitfile

// HeaderStruct is the simple struct to carry authentication string from different suppliers
type HeaderStruct struct {
	Authorization string `json:"Authorization"`
}

// TokenFetcher is the interface for the authentication token suppliers which are provides tokens as a HeaderStruct
// instances
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
