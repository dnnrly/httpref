package test

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

//nolint: unused
type testContext struct {
	err      error
	cmdInput struct {
		parameters string
	}
	cmdResult struct {
		Output string
		Err    error
		Code   int
	}
}

// Errorf is used by the called assertion to report an error and is required to
// make testify assertions work
func (c *testContext) Errorf(format string, args ...interface{}) {
	c.err = fmt.Errorf(format, args...)
}

func (c *testContext) theAppRunsWithoutParameters() error {
	cmd := exec.Command("httpref")
	output, err := cmd.CombinedOutput()
	c.cmdResult.Output = string(output)
	c.cmdResult.Err = err
	c.cmdResult.Code = cmd.ProcessState.ExitCode()

	return nil
}

func (c *testContext) theAppRunsWithParameters(args string) error {
	c.cmdInput.parameters = args
	cmdArgs := strings.Split(args, " ")
	cmd := exec.Command("httpref", cmdArgs...)
	output, err := cmd.CombinedOutput()
	c.cmdResult.Output = string(output)
	c.cmdResult.Err = err
	c.cmdResult.Code = cmd.ProcessState.ExitCode()

	return nil
}

func (c *testContext) theAppExitsWithoutError() error {
	assert.NoError(c, c.cmdResult.Err)
	if c.err != nil {
		return c.err
	}

	assert.Equal(c, 0, c.cmdResult.Code)
	if c.err != nil {
		return c.err
	}

	return nil
}

func (c *testContext) theAppExitsWithAnError() error {
	assert.Error(c, c.cmdResult.Err)
	if c.err != nil {
		return c.err
	}

	assert.NotEqual(c, 0, c.cmdResult.Code)
	if c.err != nil {
		return c.err
	}

	return nil
}

func (c *testContext) theAppOutputContains(expected string) error {
	expected = strings.ReplaceAll(expected, "\\\"", "\"")
	assert.Contains(c, c.cmdResult.Output, expected)
	return c.err
}

func (c *testContext) theAppOutputDoesNotContain(unexpected string) error {
	unexpected = strings.ReplaceAll(unexpected, "\\\"", "\"")
	assert.NotContains(c, c.cmdResult.Output, unexpected)
	return c.err
}

func (c *testContext) eachOutputLineShorterThan(expectedLineLength int) error {
	lines := strings.Split(strings.Replace(c.cmdResult.Output, "\r\n", "\n", -1), "\n")
	for i := range lines {
		element := lines[i]
		lineLength := len(element)
		if lineLength > expectedLineLength {
			return fmt.Errorf("line number %d was expected to be shorter than %d characters, but was %d", i, expectedLineLength, lineLength)
		}
	}
	return nil
}

//nolint: unused
func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {})
}


func InitializeScenario(ctx *godog.ScenarioContext) {
	tc := testContext{}
	ctx.Before(func(c context.Context, s *godog.Scenario) (context.Context, error) { return c, nil })
	ctx.After(func(ctx context.Context, s *godog.Scenario, err error) (context.Context, error) {
		if err != nil {
			fmt.Printf(
				"Command line output for \"%s\"\nUsing parameters: %s\n%s",
				s.Name,
				tc.cmdInput.parameters,
				tc.cmdResult.Output,
			)
			return ctx, err
		}

		return ctx, nil
	})

	ctx.Step(`^the app runs without parameters$`, tc.theAppRunsWithoutParameters)
	ctx.Step(`^the app runs with parameters "(.*)"$`, tc.theAppRunsWithParameters)
	ctx.Step(`^the app exits without error$`, tc.theAppExitsWithoutError)
	ctx.Step(`^the app exits with an error$`, tc.theAppExitsWithAnError)
	ctx.Step(`^the app output contains "(.*)"$`, tc.theAppOutputContains)
	ctx.Step(`^the app output does not contain "(.*)"$`, tc.theAppOutputDoesNotContain)
	ctx.Step(`^each line in output is shorter than (\d+) characters$`, tc.eachOutputLineShorterThan)
}
