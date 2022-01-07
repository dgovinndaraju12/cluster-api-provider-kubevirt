package e2e_tests

import (
	"flag"
)

var KubectlPath string
var ClusterctlPath string

func init() {
	flag.StringVar(&KubectlPath, "kubectl-path", "", "Path to the kubectl binary")
	flag.StringVar(&ClusterctlPath, "clusterctl-path", "", "Path to the clusterctl binary")
}

// BeforeTestSuiteSetup is the function called to initialize the test environment
func BeforeTestSuiteSetup() {
	flag.Parse()
}

// AfterTestSuiteCleanup is the function called to tear down the test environment
func AfterTestSuiteCleanup() {
	flag.Parse()
}
