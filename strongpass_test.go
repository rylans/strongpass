package strongpass

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZeroValidationRules(t *testing.T) {
	validator := NewValidator()

	result := validator.Validate("abc123")
	assert.Equal(t, 0, result.strength)
}

func TestCommonPasswordsRule(t *testing.T) {
	validator := NewValidator()
	validator.NoCommonPasswords()

	assert.Equal(t, "Password contains string 'letmein'", validator.Validate("letmein").errors[0])
	assert.Equal(t, "Password contains string 'password'", validator.Validate("password").errors[0])
	assert.Equal(t, "Password contains string 'hello'", validator.Validate("hello").errors[0])
}