package util

import (
	"testing"
)

func TestArgon2idHashGenerateHashAndCompare(t *testing.T) {
	time := uint32(1)
	saltLen := uint32(16)
	memory := uint32(64 * 1024) // 64MB
	threads := uint8(4)
	keyLen := uint32(32)

	argon2idHash := NewArgon2idHash(time, saltLen, memory, threads, keyLen)

	// Test case 1: Generate hash and compare
	password := []byte("password123")
	salt := []byte("randomsalt")
	hashSalt, err := argon2idHash.GenerateHash(password, salt)
	if err != nil {
		t.Errorf("Error generating hash: %v", err)
	}

	err = argon2idHash.Compare(hashSalt.Hash, hashSalt.Salt, password)
	if err != nil {
		t.Errorf("Error comparing hash: %v", err)
	}

	// Test case 2: Compare with incorrect password
	incorrectPassword := []byte("incorrect123")
	err = argon2idHash.Compare(hashSalt.Hash, hashSalt.Salt, incorrectPassword)
	if err == nil {
		t.Error("Expected error for mismatched password, but got nil")
	} else if err.Error() != "hash doesn't match" {
		t.Errorf("Expected 'hash doesn't match' error, but got: %v", err)
	}
}
