package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test3WriteAppendTextFile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "3WriteAppendTextFile Suite")
}
