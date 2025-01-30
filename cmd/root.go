package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/niespodd/nanana/internal/crypt"
	"github.com/niespodd/nanana/internal/prompt"
	"github.com/spf13/cobra"
)

var decryptFlag bool
var fileFlag string

// RootCmd is the main CLI command
var RootCmd = &cobra.Command{
	Use:   "nanana [text] or -f [file]",
	Short: "🔥 encrypt & decrypt secrets",
	Long: `Nananasecret CLI encrypts or decrypts sensitive text using AES-GCM.

Usage:
  nanana "my secret text"        -> Encrypts text
  nanana -d "my encrypted text"  -> Decrypts text
  nanana -f file.txt             -> Encrypts file contents
  nanana -d -f file.txt.enc      -> Decrypts file contents`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if fileFlag != "" {
			handleFileEncryptionOrDecryption()
			return
		}

		if len(args) == 0 {
			fmt.Println("❌ Error: No input provided. Use `nanana -f file.txt` or `nanana \"text\"`.")
			os.Exit(1)
		}

		input := args[0]

		if decryptFlag {
			password := prompt.GetPassword("🔑 Enter password to decrypt", true, input)
			decrypted, err := crypt.Decrypt(input, password)
			if err != nil {
				log.Fatalf("❌ Decryption failed: %v", err)
			}

			fmt.Printf("\n🔓 Decrypted: %s\n", decrypted)
		} else {
			password := prompt.GetPassword("🔑 Enter password to encrypt", false, "")
			encrypted, err := crypt.Encrypt(input, password)
			if err != nil {
				log.Fatalf("❌ Encryption failed: %v", err)
			}

			fmt.Printf("\n🎉 Encrypted: %s\n", encrypted)
		}
	},
}

// Handle file encryption or decryption
func handleFileEncryptionOrDecryption() {
	data, err := os.ReadFile(fileFlag)
	if err != nil {
		log.Fatalf("❌ Failed to read file %s: %v", fileFlag, err)
	}

	if decryptFlag {
		password := prompt.GetPassword(fmt.Sprintf("🔑 Enter password to decrypt %s", fileFlag), true, string(data))
		decrypted, err := crypt.Decrypt(string(data), password)
		if err != nil {
			log.Fatalf("❌ Decryption failed: %v", err)
		}

		fmt.Printf("\n🔓 Decrypted: %s\n", decrypted)
	} else {
		password := prompt.GetPassword(fmt.Sprintf("🔑 Enter password to encrypt %s", fileFlag), false, "")
		encrypted, err := crypt.Encrypt(string(data), password)
		if err != nil {
			log.Fatalf("❌ Encryption failed: %v", err)
		}

		encFile := fileFlag + ".enc"
		err = os.WriteFile(encFile, []byte(encrypted), 0644)
		if err != nil {
			log.Fatalf("❌ Failed to write encrypted file %s: %v", encFile, err)
		}

		fmt.Printf("\n🎉 Encrypted file saved as: %s\n", encFile)
	}
}

func init() {
	RootCmd.Flags().BoolVarP(&decryptFlag, "decrypt", "d", false, "Decrypt input")
	RootCmd.Flags().StringVarP(&fileFlag, "file", "f", "", "Encrypt or decrypt file contents")
}

// Execute runs the CLI
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println("❌ Error:", err)
		os.Exit(1)
	}
}
