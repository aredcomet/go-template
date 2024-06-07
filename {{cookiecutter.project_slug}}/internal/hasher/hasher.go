package hasher

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

var (
	ErrIncorrectPassword = errors.New("passwords do not match")
	ErrInvalidHash       = errors.New("hashed password is incorrectly formatted")
)

type PasswordHasher struct {
	Time       uint32
	Memory     uint32
	Threads    uint8
	KeyLength  uint32
	SaltLength uint32
}

func NewHasher() *PasswordHasher {
	return &PasswordHasher{
		Time:       2,
		Memory:     102400,
		Threads:    8,
		KeyLength:  32,
		SaltLength: 16,
	}
}

func (ph *PasswordHasher) GenerateHash(password string) (string, error) {
	salt := make([]byte, ph.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, ph.Time, ph.Memory, ph.Threads, ph.KeyLength)
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)

	return fmt.Sprintf("argon2$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		ph.Memory, ph.Time, ph.Threads, encodedSalt, encodedHash), nil
}

func (ph *PasswordHasher) Compare(password, hash string) error {
	parts := strings.Split(hash, "$")
	if len(parts) != 6 {
		return ErrInvalidHash
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return err
	}

	comparisonHash := argon2.IDKey([]byte(password), salt, ph.Time, ph.Memory, ph.Threads, ph.KeyLength)
	if subtle.ConstantTimeCompare(decodedHash, comparisonHash) != 1 {
		return ErrIncorrectPassword
	}

	return nil
}
