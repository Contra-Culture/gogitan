package gogitan_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGogitan(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gogitan Suite")
}
