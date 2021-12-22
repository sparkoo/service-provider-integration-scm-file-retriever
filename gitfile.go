package main

import (
	"bytes"
	"context"
	"github.com/imroc/req"
	"io"
)

func getFileContents(ctx context.Context, repoUrl, filepath, ref string, callback func(url string)) (io.ReadCloser, error) {

	fileUrl, err := Detect(repoUrl, filepath, ref)
	if err != nil {
		return nil, err
	}
	r, _ := req.Get(fileUrl, ctx)
	return io.NopCloser(bytes.NewBuffer(r.Bytes())), nil
}
