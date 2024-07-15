package repo_manager_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/the-gigi/multi-git/pkg/helpers"
	"os"
	"path"
	"strings"
	"testing"
)


const baseDir = "tmp/test-multi-git"
var repoList = []string{}

var _ = Describe("Repo manager tests", func(){
	var err error

	removeAll := func(){
		err = os.RemoveAll(baseDir)
		Expect(err).To(BeNil())
	}
})