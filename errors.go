package argparse

type subCommandError struct {
	error
	cmd *command
}

func (e subCommandError) Error() string {
	return "[sub]command required"
}

func newSubCommandError(cmd *command) error {
	return subCommandError{cmd: cmd}
}
