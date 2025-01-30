package prompt

import (
	"testing"

	"github.com/manifoldco/promptui"

	"github.com/stretchr/testify/assert"
)

func runWithPromptMock(input string, fn func()) {
	orig := PromptRunFunc
	PromptRunFunc = func(_ promptui.Prompt) (string, error) {
		return input, nil
	}
	defer func() { PromptRunFunc = orig }()
	fn()
}

func TestGetPassword_PromptsUser(t *testing.T) {
	runWithPromptMock("testpassword", func() {
		password := GetPassword("Enter password:", false, "")
		assert.Equal(t, "testpassword", password, "GetPassword should return the mocked user input")
	})
}
