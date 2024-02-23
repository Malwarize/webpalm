package shared_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Malwarize/webpalm/v2/shared"
)

func TestIsValidDomain(t *testing.T) {
	exportFile := "8.8.8.8"
	isValidDomain := shared.IsValidDomain(exportFile)
	assert.True(t, isValidDomain, "Valid Domain!")
}

func TestInvalidParseRegexes(t *testing.T) {
	data := "\bpassword\b.{0,10}"
	_, err := shared.ParseRegexes(data)
	assert.NotEmpty(t, err, "Error exists!")
}
