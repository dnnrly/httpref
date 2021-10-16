package common

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

//nolint: unused
// TestContext will hold all of the common data for each test
type TestContext struct {
	// BinaryPath should point to the tested binary location
	BinaryPath string
	err        error
	cmdInput   struct {
		parameters string
	}
	CmdResult struct {
		Output string
		Err    error
	}
}

// Errorf is used by the called assertion to report an error and is required to
// make testify assertions work
func (c *TestContext) Errorf(format string, args ...interface{}) {
	c.err = fmt.Errorf(format, args...)
}

func (c *TestContext) theAppRunsWithParameters(args string) error {
	c.cmdInput.parameters = args
	cmdArgs := strings.Split(args, " ")
	cmd := exec.Command(c.BinaryPath, cmdArgs...)
	output, err := cmd.CombinedOutput()
	c.CmdResult.Output = string(output)
	c.CmdResult.Err = err

	return nil
}

func (c *TestContext) theAppExitsWithoutError() error {
	assert.NoError(c, c.CmdResult.Err)
	return c.err
}

func (c *TestContext) theAppExitsWithAnError() error {
	assert.Error(c, c.CmdResult.Err)
	return c.err
}

func (c *TestContext) theAppExitCodeIs(exitCode int) error {
	var e *exec.ExitError
	if errors.As(c.CmdResult.Err, &e) {
		assert.Equal(c, exitCode, e.ExitCode())
	}
	return nil
}

func (c *TestContext) theAppExitCodeIsNot(exitCode int) error {
	var e *exec.ExitError
	if errors.As(c.CmdResult.Err, &e) {
		assert.NotEqual(c, exitCode, e.ExitCode())
	}
	return nil
}

func (c *TestContext) theAppOutputContains(expected string) error {
	expected = strings.ReplaceAll(expected, "\\\"", "\"")
	assert.Contains(c, c.CmdResult.Output, expected)
	return c.err
}

func (c *TestContext) theAppOutputDoesNotContain(unexpected string) error {
	unexpected = strings.ReplaceAll(unexpected, "\\\"", "\"")
	assert.NotContains(c, c.CmdResult.Output, unexpected)
	return c.err
}

//nolint: unused
func InitializeScenario(ctx *godog.ScenarioContext, tc TestContext) {
	ctx.BeforeScenario(func(*godog.Scenario) {})
	ctx.AfterScenario(func(s *godog.Scenario, err error) {
		if err != nil {
			fmt.Printf(
				"Command line output for \"%s\"\nUsing parameters: %s\n%s",
				s.Name,
				tc.cmdInput.parameters,
				tc.CmdResult.Output,
			)
		}
	})
	ctx.Step(`^the app runs with parameters "(.*)"$`, tc.theAppRunsWithParameters)
	ctx.Step(`^the app exits without error$`, tc.theAppExitsWithoutError)
	ctx.Step(`^the app exits with an error$`, tc.theAppExitsWithAnError)
	ctx.Step(`^the app exit code is (\d+)$`, tc.theAppExitCodeIs)
	ctx.Step(`^the app exit code is not (\d+)$`, tc.theAppExitCodeIsNot)
	ctx.Step(`^the app output contains "(.*)"$`, tc.theAppOutputContains)
	ctx.Step(`^the app output does not contain "(.*)"$`, tc.theAppOutputDoesNotContain)
}
