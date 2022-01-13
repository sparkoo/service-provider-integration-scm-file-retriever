package gitfile

import (
	"bytes"
	"context"
	"github.com/imroc/req"
	"io"
)

type GitFile struct {
	fetcher TokenFetcher
}

var gitFile = This()

// GetFileContents is a main entry function allowing to retrieve file content from the SCM provider.
// It expects three file location parameters, from which the repository URL and path to the file are mandatory,
// and optional Git reference for the branch/tags/commidIds.
// Function type parameter is a callback used when user authentication is needed in order to retrieve the file,
// that function will be called with the URL to OAuth service, where user need to be redirected.
func GetFileContents(ctx context.Context, repoUrl, filepath, ref string, callback func(url string)) (io.ReadCloser, error) {
	return gitFile.GetFileContents(ctx, repoUrl, filepath, ref, callback)
}

func (g *GitFile) GetFileContents(ctx context.Context, repoUrl, filepath, ref string, callback func(url string)) (io.ReadCloser, error) {
	headerStruct := BuildAuthHeader(repoUrl, g.fetcher)
	authHeader := req.HeaderFromStruct(headerStruct)
	fileUrl, err := detect(repoUrl, filepath, ref, authHeader)
	if err != nil {
		return nil, err
	}

	response, _ := req.Get(fileUrl, ctx, authHeader)
	return io.NopCloser(bytes.NewBuffer(response.Bytes())), nil
}

func (g *GitFile) SetTokenFetcher(fetcher TokenFetcher) {
	g.fetcher = fetcher
}

// This creates a new *GitFile instance
func This() *GitFile {
	return &GitFile{fetcher: &EnvVarTokenFetcher{}}
}
