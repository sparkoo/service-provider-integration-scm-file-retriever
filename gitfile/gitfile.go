package gitfile

import (
	"bytes"
	"context"
	"github.com/imroc/req"
	"io"
)

func GetFileContents(ctx context.Context, repoUrl, filepath, ref string, callback func(url string)) (io.ReadCloser, error) {

	header := BuildAuthHeader(repoUrl)
	authHeader := req.HeaderFromStruct(header)
	fileUrl, err := Detect(repoUrl, filepath, ref, authHeader)
	if err != nil {
		return nil, err
	}

	r, _ := req.Get(fileUrl, ctx, authHeader)
	return io.NopCloser(bytes.NewBuffer(r.Bytes())), nil
}
