package prompt

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/niespodd/nanana/internal/crypt"
)

var PromptRunFunc = func(p promptui.Prompt) (string, error) {
	return p.Run()
}

func GetPassword(label string, decryptMode bool, encryptedData string) string {
	storedPasswords := ListStoredPasswords()

	if decryptMode && len(storedPasswords) > 0 {
		fmt.Println("🔍 Trying stored passwords...")
		for _, pass := range storedPasswords {
			_, err := crypt.Decrypt(encryptedData, pass)
			if err == nil {
				fmt.Println("✅ Successfully used stored password!")
				return pass
			}
		}
		fmt.Println("❌ No stored passwords worked.")
	}

	prompt := promptui.Prompt{
		Label: label,
		Mask:  '•',
	}

	// Use function variable instead of direct `Run()`
	password, err := PromptRunFunc(prompt)
	if err != nil {
		log.Fatalf("❌ Password input failed: %v", err)
	}

	if !decryptMode {
		StorePassword(password)
	}

	return password
}
