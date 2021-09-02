package remotetype

import (
	"strings"

	"github.com/go-git/go-git/v5"
)

// RemoteType allows for a general enum with all values provided
type RemoteType int

const (
	// GitLabRemote is the value for the gitlab remote type
	GitLabRemote RemoteType = iota
	// GitHubRemote is the value for the github remote type
	GitHubRemote
	// NoRemote holds the default value when a remote cannot be parsed
	NoRemote
)

// GetRepoType takes a repository and searches through all of the remotes to find
// if it connects to github or gitlab (usually doesn't do both so if it finds one it just returns it)
func GetRepoType(repository *git.Repository) RemoteType {
	remotes, _ := repository.Remotes()
	for _, remote := range remotes {
		for _, remoteURL := range remote.Config().URLs {
			if strings.Contains(remoteURL, "gitlab") {
				return GitLabRemote
			} else if strings.Contains(remoteURL, "github") {
				return GitHubRemote
			}
		}
	}
	return NoRemote
}
