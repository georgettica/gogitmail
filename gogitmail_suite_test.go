package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGogitmail(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gogitmail Suite")
}
