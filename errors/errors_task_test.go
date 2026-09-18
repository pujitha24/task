package errors_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-task/task/v3/errors"
)

func TestTaskRunErrorTaskExitCode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		err      error
		expected int
	}{
		{
			name:     "cancelled by user",
			err:      &errors.TaskCancelledByUserError{TaskName: "foo"},
			expected: errors.CodeTaskCancelled,
		},
		{
			name:     "cancelled, no terminal",
			err:      &errors.TaskCancelledNoTerminalError{TaskName: "foo"},
			expected: errors.CodeTaskCancelled,
		},
		{
			name:     "missing required vars",
			err:      &errors.TaskMissingRequiredVarsError{TaskName: "foo"},
			expected: errors.CodeTaskMissingRequiredVars,
		},
		{
			name:     "timeout still reports the timeout(1) convention code",
			err:      &errors.TaskTimeoutError{TaskName: "foo"},
			expected: errors.TimeoutExitCode,
		},
		{
			name:     "unrecognised error falls back to the generic task run code",
			err:      errors.New("some command failed"),
			expected: errors.CodeTaskRunError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			runErr := &errors.TaskRunError{TaskName: "foo", Err: test.err}
			assert.Equal(t, test.expected, runErr.TaskExitCode())
		})
	}
}
