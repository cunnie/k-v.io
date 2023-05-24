package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestKVIO(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "k-v.io Suite")
}
