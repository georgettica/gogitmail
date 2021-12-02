// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/georgettica/gogitmail/pkg/interfaces (interfaces: RequestMaker)

// Package mock is a generated GoMock package.
package mock

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRequestMaker is a mock of RequestMaker interface.
type MockRequestMaker struct {
	ctrl     *gomock.Controller
	recorder *MockRequestMakerMockRecorder
}

// MockRequestMakerMockRecorder is the mock recorder for MockRequestMaker.
type MockRequestMakerMockRecorder struct {
	mock *MockRequestMaker
}

// NewMockRequestMaker creates a new mock instance.
func NewMockRequestMaker(ctrl *gomock.Controller) *MockRequestMaker {
	mock := &MockRequestMaker{ctrl: ctrl}
	mock.recorder = &MockRequestMakerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestMaker) EXPECT() *MockRequestMakerMockRecorder {
	return m.recorder
}

// ToGithub mocks base method.
func (m *MockRequestMaker) ToGithub(arg0, arg1 string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToGithub", arg0, arg1)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToGithub indicates an expected call of ToGithub.
func (mr *MockRequestMakerMockRecorder) ToGithub(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToGithub", reflect.TypeOf((*MockRequestMaker)(nil).ToGithub), arg0, arg1)
}

// ToGitlab mocks base method.
func (m *MockRequestMaker) ToGitlab(arg0 string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToGitlab", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToGitlab indicates an expected call of ToGitlab.
func (mr *MockRequestMakerMockRecorder) ToGitlab(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToGitlab", reflect.TypeOf((*MockRequestMaker)(nil).ToGitlab), arg0)
}
