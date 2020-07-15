package main_test

import (
	"encoding/json"

	. "github.com/georgettica/gogitmail"
	"github.com/georgettica/gogitmail/pkg/structs"
	. "github.com/georgettica/gogitmail/pkg/structs"

	"github.com/adammck/venv"
	"github.com/golang/mock/gomock"

	"net/http"

	mockhttp "github.com/karupanerura/go-mock-http-response"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func mockResponse(statusCode int, headers map[string]string, body []byte) {
	http.DefaultClient = mockhttp.NewResponseMock(statusCode, headers, body).MakeClient()
}

func init() {
	e := venv.Mock()
	e.Setenv("GITHUB_TOKEN", "aaaa")
	e.Setenv("GITLAB_TOKEN", "bbbb")
	e.Setenv("GITLAB_HOSTNAME", "test.example.com")
	conf = NewGogitmailConfig(e, &structs.MakeRequest{})
}

var conf *GogitmailConfig

var _ = Describe("GoGitMail", func() {

	Describe("LabEmail", func() {
		Describe("with external http mocking package", func() {
			BeforeEach(func() {
				labUser := LabUser{
					ID:       1234,
					Username: "username",
				}
				bytes, _ := json.Marshal(labUser)
				mockResponse(http.StatusOK, map[string]string{"Content-Type": "text/plain"}, bytes)
			})
			It("should return user id for email", func() {
				Expect(conf.LabEmail()).To(Equal("1234+username@users.noreply.test.example.com"))
			})
		})
		Context("With valid response", func() {
			BeforeEach(func() {
				labUser := LabUser{
					ID:       1234,
					Username: "username",
				}
				bytes, _ := json.Marshal(labUser)
				mockResponse(http.StatusOK, map[string]string{"Content-Type": "text/plain"}, bytes)
			})
			It("should return user id for email", func() {
				Expect(conf.LabEmail()).To(Equal("1234+username@users.noreply.test.example.com"))
			})
		})
	})
	Describe("HubEmail", func() {
		Context("With valid response", func() {
			BeforeEach(func() {
				hubUser := HubUser{
					ID:    1234,
					Login: "username",
				}

				bytes, _ := json.Marshal(hubUser)
				mockResponse(http.StatusOK, map[string]string{"Content-Type": "text/plain"}, bytes)
			})
			It("should return user id for email", func() {
				Expect(conf.HubEmail()).To(Equal("1234+username@users.noreply.github.com"))
			})
		})
		Context("With gomock", func() {
			BeforeEach(func() {
				mockCtrl := gomock.NewController(GinkgoT())
				defer mockCtrl.Finish()

				hubUser := HubUser{
					ID:    1234,
					Login: "username",
				}

				bytes, _ := json.Marshal(hubUser)
				mockResponse(http.StatusOK, map[string]string{"Content-Type": "text/plain"}, bytes)
			})
			It("should return user id for email", func() {
				Expect(conf.HubEmail()).To(Equal("1234+username@users.noreply.github.com"))
			})
		})
	})
})
