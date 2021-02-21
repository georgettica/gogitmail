package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/adammck/venv"
	"github.com/georgettica/gogitmail/pkg/interfaces"
	"github.com/georgettica/gogitmail/pkg/structs"
	"github.com/georgettica/gogitmail/pkg/structs/remotetype"
	"github.com/georgettica/gogitmail/pkg/utils"

	"github.com/urfave/cli/v2"
)

type GogitmailConfig struct {
	env          venv.Env
	requestMaker interfaces.RequestMaker
}

func NewGogitmailConfig(e venv.Env, rm interfaces.RequestMaker) *GogitmailConfig {
	return &GogitmailConfig{
		env:          e,
		requestMaker: rm,
	}
}

func main() {
	app := &cli.App{
		Name:  "gogitmail",
		Usage: "run inside a git repository (top level) and the private git email will be added to your local repository",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "force",
				Aliases: []string{"f"},
				Value:   false,
				Usage:   "rewrite git config even if it exists",
			},
		},
		Action: func(c *cli.Context) error {
			runGogitmail(c)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func runGogitmail(c *cli.Context) {
	conf := GogitmailConfig{
		env:          venv.OS(),
		requestMaker: &structs.MakeRequest{},
	}

	repository, err := utils.GetRepository()
	if err != nil {
		panic(err.Error())
	}

	localEmail := utils.GetGitEmail(repository)

	if localEmail != "" && !c.Bool("force") {
		panic("email already set, no action to do")
	} else if c.Bool("force") {
		fmt.Println("'force' flag triggered, overwriting local config")
	}

	var email string

	currentRepoRemoteType := remotetype.GetRepoType(repository)
	switch currentRepoRemoteType {
	case remotetype.GitLabRemote:
		email = conf.LabEmail()
	case remotetype.GitHubRemote:
		email = conf.HubEmail()
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

func (conf GogitmailConfig) HubEmail() string {
	const githubURL string = "github.com"
	token := conf.env.Getenv("GOGITMAIL_GITHUB_TOKEN")
	bearer := fmt.Sprintf("token %v", token)

	resp, err := conf.requestMaker.ToGithub(fmt.Sprintf("https://api.%v/user", githubURL), bearer)
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
	id := i.GetID()
	if id == "0+" {
		panic("ID is 0, probbably because of revoked token")
	}

	return fmt.Sprintf("%v@users.noreply.github.com", id)
}

func (conf GogitmailConfig) LabEmail() string {
	token := conf.env.Getenv("GOGITMAIL_GITLAB_TOKEN")
	gitlabPrivateURL := conf.env.Getenv("GOGITMAIL_GITLAB_HOSTNAME")

	resp, err := conf.requestMaker.ToGitlab(fmt.Sprintf("https://%v/api/v4/user?access_token=%v", gitlabPrivateURL, token))
	if err != nil {
		var dnsErr *net.DNSError
		if errors.As(err, &dnsErr) {
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

	id := interfaces.GitRemoteUser(gitlab).GetID()
	if id == "0+" {
		panic("ID is 0, probbably because of revoked token")
	}

	return fmt.Sprintf("%v@users.noreply.%v", id, gitlabPrivateURL)
}
