package structs

import (
	"net/http"
)

// type RequestMaker interface  {
// 	MakeHubRequest(url, token string) (*http.Response, error)
// 	MakeLabRequest(url string) (*http.Response, error)
// }

type MakeRequest struct{}

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

func (t *MakeRequest) ToGitlab(url string) (*http.Response, error) {
	return http.Get(url)
}
