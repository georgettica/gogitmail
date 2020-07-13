package remotetype

import (
	"strings"

	"github.com/go-git/go-git"
)

type RemoteType int

const (
	GitLabRemote RemoteType = iota
	GitHubRemote
	NoRemote
)

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
