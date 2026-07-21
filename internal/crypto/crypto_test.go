package crypto_test

import (
	"testing"

	"github.com/Fracizz/sshctl/internal/crypto"
)

func TestEncryptDecrypt(t *testing.T) {
	enc, err := crypto.Encrypt("hello")
	if err != nil {
		t.Fatal(err)
	}
	if !crypto.IsEncrypted(enc) {
		t.Fatalf("prefix missing: %s", enc)
	}
	plain, err := crypto.Decrypt(enc)
	if err != nil || plain != "hello" {
		t.Fatalf("got %q err=%v", plain, err)
	}
	plain, err = crypto.Decrypt("not-encrypted")
	if err != nil || plain != "not-encrypted" {
		t.Fatalf("passthrough: %q %v", plain, err)
	}
}
