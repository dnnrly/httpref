package test

import (
	"regexp"

	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/spf13/pflag"

	"github.com/dnnrly/httpref/test/common"
)

var opts = godog.Options{Output: colors.Colored(os.Stdout)}

const (
	matchUrlRegex = `/docs/Web/HTTP/Headers/Accept`
	httpRegex     = `Hypertext Transfer Protocol`
)

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

func theArgsAre(available string) error {
	var err error
	if available == "" {
		Args = []string{}
	} else {
		Args = strings.Fields(available)
	}

	cmd := exec.Command("httpref", Args...)
	out, err := cmd.Output()
	if err != nil {
		var e *exec.ExitError
		if errors.As(err, &e) {
			ExitCode = e.ExitCode()
			return nil
		}
		return fmt.Errorf("unrecognized error, value is %w", err)
	}
	Output = string(out)

	if cmd.ProcessState != nil {
		ExitCode = cmd.ProcessState.ExitCode()
		return nil
	}
	return nil

}

func returnStatusCodeIsNot(num int) error {
	if ExitCode == num {
		return fmt.Errorf("Expected the exit code not to be %d, but was %d", num, ExitCode)
	}

	return nil
}

func returnStatusCodeIs(num int) error {
	if ExitCode != num {
		return fmt.Errorf("Expected the exit code to be %d, but was %d", num, ExitCode)
	}

	return nil
}

func theOutputWillNotContainTheAcceptHeaders() error {
	r := regexp.MustCompile(matchUrlRegex)
	matched := r.MatchString(Output)
	if matched {
		return fmt.Errorf("the output was expected to match %s, but didn't. raw output is %s", matchUrlRegex, Output)
	}

	return nil
}

func theOutputWillContainTheAcceptHeaders() error {
	r := regexp.MustCompile(matchUrlRegex)
	matched := r.MatchString(Output)
	if !matched {
		return fmt.Errorf("the output was not supposed to match %s, but did. raw output is %s", matchUrlRegex, Output)
	}

	return nil
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

func theOutputWillContainTheHttpString() error {
	r := regexp.MustCompile(httpRegex)
	matched := r.MatchString(Output)
	if !matched {
		return fmt.Errorf("the output was expected to match %s, but didn't. raw output is %s", httpRegex, Output)
	}

	return nil
}

func theOutputWillNotContainTheHttpString() error {
	r := regexp.MustCompile(httpRegex)
	matched := r.MatchString(Output)
	if matched {
		return fmt.Errorf("the output was expected to match %s, but didn't. raw output is %s", httpRegex, Output)

	}

	return nil
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		Args = []string{} // clean the state before every scenario
		ExitCode = -1
		Output = ""
	})
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	tc := localTestContext{TestContext: common.TestContext{BinaryPath: "httpref"}}
	common.InitializeScenario(ctx, tc.TestContext)
	//ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	//	Args = []string{} // clean the state before every scenario
	//	ExitCode = -1
	//	Output = ""
	//	return nil, nil
	//})

	//ctx.Step(`^the args are "([^"]*)"$`, theArgsAre)
	//ctx.Step(`^return status code is not (\d+)$`, returnStatusCodeIsNot)
	//ctx.Step(`^return status code is (\d+)$`, returnStatusCodeIs)
	//ctx.Step(`^the output will not contain the Accept headers$`, theOutputWillNotContainTheAcceptHeaders)
	//ctx.Step(`^the output will contain the Accept headers$`, theOutputWillContainTheAcceptHeaders)
	ctx.Step(`^each line in output is shorter than (\d+) characters$`, tc.eachLineInOutputIsShorterThanCharacters)
	//ctx.Step(`^the output will contain the http string$`, theOutputWillContainTheHttpString)
	//ctx.Step(`^the output will not contain the http string$`, theOutputWillNotContainTheHttpString)
}
