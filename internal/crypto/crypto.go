package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"os/user"
	"runtime"
	"strings"
)

const prefix = "enc:v1:"

// IsEncrypted reports whether value uses the sshctl ciphertext format.
func IsEncrypted(value string) bool {
	return strings.HasPrefix(value, prefix)
}

// Encrypt encrypts plaintext with a machine-derived AES-256-GCM key.
func Encrypt(plaintext string) (string, error) {
	key, err := machineKey()
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	out := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return prefix + base64.StdEncoding.EncodeToString(out), nil
}

// Decrypt decrypts an enc:v1: value. Plaintext is returned unchanged.
func Decrypt(value string) (string, error) {
	if !IsEncrypted(value) {
		return value, nil
	}
	raw, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(value, prefix))
	if err != nil {
		return "", fmt.Errorf("decode ciphertext: %w", err)
	}
	key, err := machineKey()
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	if len(raw) < gcm.NonceSize() {
		return "", errors.New("ciphertext too short")
	}
	nonce, ciphertext := raw[:gcm.NonceSize()], raw[gcm.NonceSize():]
	plain, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("decrypt failed (wrong machine or corrupted data): %w", err)
	}
	return string(plain), nil
}

func machineKey() ([]byte, error) {
	host, _ := os.Hostname()
	u, err := user.Current()
	username := "unknown"
	if err == nil {
		username = u.Username
	}
	material := strings.Join([]string{
		"sshctl-v1",
		runtime.GOOS,
		runtime.GOARCH,
		host,
		username,
		machineID(),
	}, "|")
	sum := sha256.Sum256([]byte(material))
	return sum[:], nil
}
