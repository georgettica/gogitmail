package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGomarshal(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gomarshal Suite")
}
