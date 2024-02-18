package shared_tests

import (
	"testing"

	"github.com/Malwarize/webpalm/v2/shared"
	"github.com/stretchr/testify/assert"
)

func TestIsValidDomain(t *testing.T) {
	var exportFile = "8.8.8.8"
	isValidDomain := shared.IsValidDomain(exportFile)
	assert.True(t, isValidDomain, "Valid Domain!")
}

func TestInvalidParseRegexes(t *testing.T) {
	var data = "\bpassword\b.{0,10}"
	_, err := shared.ParseRegexes(data)
	assert.NotEmpty(t, err, "Error exists!")
}
