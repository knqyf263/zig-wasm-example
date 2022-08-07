package main

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test_main ensures the following will work:
//
//	go run greet.go wazero
func Test_main(t *testing.T) {
	stdout, _ := testMain(t, main, "greet", "wazero")
	require.Equal(t, `wasm >> Hello, wazero!
go >> Hello, wazero!
`, stdout)
}

// ported from https://github.com/tetratelabs/wazero/blob/ef190285134cc15e5c69f7183206f5bc7e434ee4/internal/testing/maintester/maintester.go#L12-L56
func testMain(t *testing.T, main func(), args ...string) (stdout, stderr string) {
	// Setup files to capture stdout and stderr
	tmp := t.TempDir()

	stdoutPath := path.Join(tmp, "stdout.txt")
	stdoutF, err := os.Create(stdoutPath)
	require.NoError(t, err)

	stderrPath := path.Join(tmp, "stderr.txt")
	stderrF, err := os.Create(stderrPath)
	require.NoError(t, err)

	// Save the old os.XXX and revert regardless of the outcome.
	oldArgs := os.Args
	os.Args = args
	oldStdout := os.Stdout
	os.Stdout = stdoutF
	oldStderr := os.Stderr
	os.Stderr = stderrF
	revertOS := func() {
		os.Args = oldArgs
		_ = stdoutF.Close()
		os.Stdout = oldStdout
		_ = stderrF.Close()
		os.Stderr = oldStderr
	}
	defer revertOS()

	// Run the main command.
	main()

	// Revert os.XXX so that test output is visible on failure.
	revertOS()

	// Capture any output and return it in a portable way (ex without windows newlines)
	stdoutB, err := os.ReadFile(stdoutPath)
	require.NoError(t, err)
	stdout = strings.ReplaceAll(string(stdoutB), "\r\n", "\n")

	stderrB, err := os.ReadFile(stderrPath)
	require.NoError(t, err)
	stderr = strings.ReplaceAll(string(stderrB), "\r\n", "\n")

	return
}
