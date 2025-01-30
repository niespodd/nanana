package crypt

import (
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	plaintext := "this is a secret"
	password := "supersecure"
	encrypted, err := Encrypt(plaintext, password)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	decrypted, err := Decrypt(encrypted, password)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	if decrypted != plaintext {
		t.Errorf("Decryption failed, expected %q but got %q", plaintext, decrypted)
	}
}

func TestDecryptWrongPassword(t *testing.T) {
	plaintext := "top secret"
	password := "correct-password"
	wrongPassword := "wrong-password"

	encrypted, err := Encrypt(plaintext, password)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	_, err = Decrypt(encrypted, wrongPassword)
	if err == nil {
		t.Fatal("Decryption should have failed with wrong password, but it didn't")
	}
}

func TestDeriveKey(t *testing.T) {
	password := "testpassword"
	salt := []byte("randomsaltvalue123")

	key1, err1 := DeriveKey(password, salt)
	key2, err2 := DeriveKey(password, salt)

	if err1 != nil || err2 != nil {
		t.Fatalf("Key derivation failed: %v, %v", err1, err2)
	}

	if string(key1) != string(key2) {
		t.Errorf("Key derivation is inconsistent, keys do not match")
	}
}
