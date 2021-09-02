package utils

import (
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

// GetRepository extracts the repository from the current folder
func GetRepository() (*git.Repository, error) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	return git.PlainOpen(path.Join(wd, ".git"))
}

// GetGitEmail tries to get the user's config email (might be emptystring)
func GetGitEmail(repository *git.Repository) string {

	cfg, err := repository.ConfigScoped(config.SystemScope)
	if err != nil {
		panic(err.Error())
	}

	return cfg.User.Email
}
