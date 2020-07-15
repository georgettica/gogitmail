package interfaces

import "net/http"

//go:generate mockgen -package mock -destination mock/requestmaker.go  github.com/georgettica/local_git_email/pkg/gomarshal/interfaces RequestMaker

// RequestMaker holds the gitlab and the github requests
type RequestMaker interface {
	ToGithub(url, token string) (*http.Response, error)
	ToGitlab(url string) (*http.Response, error)
}
