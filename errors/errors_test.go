package errors

import (
	stderrors "errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnsureError(t *testing.T) {
	assert, require := assert.New(t), require.New(t)
	rawErr := stderrors.New("HELLO FUCK ERROR")
	err := rawErr

	if Ensure(&err) {
		require.Error(err)
		assert.NotEqual(rawErr, err)
	}
}
