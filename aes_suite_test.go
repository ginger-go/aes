package aes_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Aes Suite")
}
