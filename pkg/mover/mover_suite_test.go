package mover_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMoverPackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mover Package Suite")
}
