package gqlstruct_test

import (
	"github.com/jamillosantos/macchiato"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"log"
	"testing"
)

func TestGqlStruct(t *testing.T) {
	log.SetOutput(ginkgo.GinkgoWriter)
	gomega.RegisterFailHandler(ginkgo.Fail)
	macchiato.RunSpecs(t, "gqlstruct: GraphQL Struct Test Suite")
}
