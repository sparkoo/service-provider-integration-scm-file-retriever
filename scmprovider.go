package main

import "fmt"

// ScmProvider defines the interface that in order to determine if URL belongs to SCM provider
type ScmProvider interface {
	// Detect will detect whether the string matches a known SCM file pattern
	// and transform it to valid https url.
	Detect(repoUrl, filepath, ref string) (bool, string, error)
}

// ScmProviders is the list of detectors that are tried on an SCM URL.
// This is also the order they're tried (index 0 is first).
var ScmProviders []ScmProvider

func init() {
	ScmProviders = []ScmProvider{
		new(GitLabScmProvider),
		new(GitHubScmProvider),
	}
}

func Detect(repoUrl, filepath, ref string) (string, error) {
	for _, d := range ScmProviders {
		ok, resultUrl, err := d.Detect(repoUrl, filepath, ref)
		if err != nil {
			return "", err
		}
		if !ok {
			continue
		}
		return resultUrl, nil
	}
	return "", fmt.Errorf("invalid source string: %s for %s", repoUrl, filepath)
}
