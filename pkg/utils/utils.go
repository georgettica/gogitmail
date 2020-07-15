package utils

import (
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

func GetRepository() (*git.Repository, error) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	return git.PlainOpen(path.Join(wd, ".git"))
}

func GetGitEmail(repository *git.Repository) string {

	cfg, err := repository.ConfigScoped(config.SystemScope)
	if err != nil {
		panic(err.Error())
	}

	return cfg.User.Email
}
