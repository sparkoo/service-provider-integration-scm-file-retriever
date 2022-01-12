package gitfile

import (
	"strings"
)

// GitLabScmProvider implements Detector to detect Gitlab URLs.
type GitLabScmProvider struct{}

func (d *GitLabScmProvider) detect(repoUrl, filepath, ref string, v ...interface{}) (bool, string, error) {
	if len(repoUrl) == 0 {
		return false, "", nil
	}

	if strings.HasPrefix(repoUrl, "https://gitlab.com/") {
		return true, "", nil
	}

	return false, "", nil
}
