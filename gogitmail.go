package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"os"
	"path"
	"strings"

	"github.com/adammck/venv"
	"github.com/georgettica/gogitmail/interfaces"
	"github.com/georgettica/gogitmail/structs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

var e venv.Env

func main() {
	e = venv.OS()
	PrintGitRemotes()

	gitlabEmail := LabEmail()
	githubEmail := HubEmail()

	fmt.Printf("lab ::: %v\nhub ::: %v\n", gitlabEmail, githubEmail)

	//fmt.Printf("email ::: %v\n", conf.User.Email)

}

var (
	// LocalRequestMaker jj
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
	token, exists := e.LookupEnv("GITHUB_TOKEN")
	if !exists {
		panic(errors.New("GITHUB_TOKEN env doesnt exist"))
	}
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
	token, exists := e.LookupEnv("GITLAB_TOKEN")
	if !exists {
		panic(errors.New("GITLAB_TOKEN env doesnt exist"))
	}
	gitlabPrivateURL := e.Getenv("GITLAB_HOSTNAME")

	resp, err := LocalRequestMaker.ToGitlab("https://" + gitlabPrivateURL + "/api/v4/user?access_token=" + token)
	if err != nil {
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

func GetRepoType(repository *git.Repository) string {
	remotes, _ := repository.Remotes()
	for _, remote := range remotes {
		for _, remoteURL := range remote.Config().URLs {
			if strings.Contains(remoteURL, "gitlab") {
				return "gitlab"
			} else if strings.Contains(remoteURL, "github") {
				return "github"
			}
		}
	}
	return ""
}

func PrintGitRemotes() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	repository, err := git.PlainOpen(path.Join(wd, ".git"))
	if err != nil {
		panic(err.Error())
	}

	cfg, err := repository.ConfigScoped(config.SystemScope)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("email merged ::: %v\n", cfg.User.Email)
	return cfg.User.Email
}
