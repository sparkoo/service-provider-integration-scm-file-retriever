package gitfile

import (
	"fmt"
)

// ScmProvider defines the interface that in order to determine if URL belongs to SCM provider
type ScmProvider interface {
	// Detect will check whether the provided repository URL matches a known SCM pattern,
	// and transform input params into valid file download URL.
	// Params are repository, path to the file inside the repository, Git reference (branch/tag/commitId) and
	// set of optional parameters ScmProvider may need, such as Http headers, additional resource identifiers etc
	Detect(repoUrl, filepath, ref string, v ...interface{}) (bool, string, error)
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

func Detect(repoUrl, filepath, ref string, v ...interface{}) (string, error) {
	for _, d := range ScmProviders {
		ok, resultUrl, err := d.Detect(repoUrl, filepath, ref, v...)
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
