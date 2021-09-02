package structs

import (
	"net/http"
)

// MakeRequest is a type with no struct to encapsulate the request commands
// in a unified place
type MakeRequest struct{}

// ToGithub runs a request to github
func (t *MakeRequest) ToGithub(url, token string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}
	req.Header.Add("Authorization", token)
	req.Header.Add("Accept", "application/json")

	client := http.DefaultClient
	return client.Do(req)
}

// ToGitlab runa a request to gitlab
func (t *MakeRequest) ToGitlab(url string) (*http.Response, error) {
	return http.Get(url)
}
