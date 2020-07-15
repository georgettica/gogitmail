package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/adammck/venv"
	"github.com/georgettica/local_git_email/pkg/gomarshal/interfaces"
	"github.com/georgettica/local_git_email/pkg/gomarshal/structs"
	//"github.com/go-git/go-git/config"
)

var e venv.Env

func main() {
	e = venv.OS()

	gitlabEmail := LabEmail()
	githubEmail := HubEmail()

	fmt.Printf("lab ::: %v\nhub ::: %v\n", gitlabEmail, githubEmail)

	//conf, err := config.LoadConfig(config.SystemScope)
	//if err != nil {
	//	panic(err.Error())
	//}
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
