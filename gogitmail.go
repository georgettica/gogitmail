package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/url"

	"os"
	"path"

	"github.com/adammck/venv"
	"github.com/georgettica/gogitmail/interfaces"
	"github.com/georgettica/gogitmail/structs"
	"github.com/georgettica/gogitmail/structs/remotetype"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

var e venv.Env

type Config struct {
	env venv.Env
}

func main() {
	e = venv.OS()
	repository, err := GetRepository()
	if err != nil {
		panic(err.Error())
	}

	localEmail := PrintGitEmail(repository)

	if localEmail != "" {
		panic("email already set, no action to do")
	}

	gitlabEmail := LabEmail()
	githubEmail := HubEmail()
	var email string

	currentRepoRemoteType := remotetype.GetRepoType(repository)
	switch currentRepoRemoteType {
	case remotetype.GitLabRemote:
		email = gitlabEmail
	case remotetype.GitHubRemote:
		email = githubEmail
	case remotetype.NoRemote:
		panic("No Remote found, exiting")
	}

	cfg, err := repository.Config()
	if err != nil {
		panic(err.Error())
	}
	cfg.User.Email = email
	err = repository.SetConfig(cfg)
	if err != nil {
		panic(err.Error())
	}
}

var (
	LocalRequestMaker interfaces.RequestMaker
)

func init() {
	LocalRequestMaker = &structs.MakeRequest{}

}

func SetEnv(inputEnv venv.Env) {
	e = inputEnv
}

// HubEmail gets the users email from github
func HubEmail() string {
	const githubURL string = "github.com"
	token := e.Getenv("GITHUB_TOKEN")
	bearer := fmt.Sprintf("token %v", token)

	resp, err := LocalRequestMaker.ToGithub("https://api."+githubURL+"/user", bearer)
	if err != nil {
		panic(err.Error())
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	github := structs.HubUser{}

	if err := json.Unmarshal(data, &github); err != nil {
		panic(err.Error())
	}
	i := interfaces.GitRemoteUser(github)

	return fmt.Sprintf("%v@users.noreply.github.com", i.GetID())
}

// LabEmail gets the users email from gitlab
func LabEmail() string {
	token := e.Getenv("GITLAB_TOKEN")
	gitlabPrivateURL := e.Getenv("GITLAB_HOSTNAME")

	resp, err := LocalRequestMaker.ToGitlab(fmt.Sprintf("https://%v/api/v4/user?access_token=%v", gitlabPrivateURL, token))
	if err != nil {
		newErr := errors.Unwrap(err)
		fmt.Println(newErr)
		newErr1 := errors.Unwrap(newErr)
		fmt.Println(newErr1)
		//var dnsT net.DNSError
		if _, ok := newErr1.(*net.DNSError); ok {
			fmt.Println("hi")
		}
		if _, ok := err.(*url.Error); ok {
			fmt.Printf("Cannot access %v, maybe it's behind a VPN\n", gitlabPrivateURL)
			panic(0)
		}
		panic(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	gitlab := structs.LabUser{}

	if err := json.Unmarshal(body, &gitlab); err != nil {
		panic(err.Error())
	}
	i := interfaces.GitRemoteUser(gitlab)

	return fmt.Sprintf("%v@users.noreply.%v", i.GetID(), gitlabPrivateURL)
}

func GetRepository() (*git.Repository, error) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	return git.PlainOpen(path.Join(wd, ".git"))
}
func PrintGitEmail(repository *git.Repository) string {

	cfg, err := repository.ConfigScoped(config.SystemScope)
	if err != nil {
		panic(err.Error())
	}

	return cfg.User.Email
}
