package test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/spf13/pflag"

	"github.com/dnnrly/httpref/test/common"
)

var opts = godog.Options{Output: colors.Colored(os.Stdout)}

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
}

func TestMain(m *testing.M) {
	pflag.Parse()
	opts.Paths = pflag.Args()

	status := godog.TestSuite{
		Name:                 "godogs",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()

	os.Exit(status)
}

type localTestContext struct {
	common.TestContext
}

func (c *localTestContext) eachLineInOutputIsShorterThanCharacters(expectedLineLength int) error {
	lines := strings.Split(strings.Replace(c.CmdResult.Output, "\r\n", "\n", -1), "\n")
	for i := range lines {
		element := lines[i]
		assert.LessOrEqual(c, expectedLineLength, len(element))
	}
	return nil
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {})
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	tc := localTestContext{TestContext: common.TestContext{BinaryPath: "httpref"}}
	common.InitializeScenario(ctx, tc.TestContext)
	ctx.Step(`^each line in output is shorter than (\d+) characters$`, tc.eachLineInOutputIsShorterThanCharacters)
}
