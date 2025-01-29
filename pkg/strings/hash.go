package strings

import (
	"crypto/sha1"
	"fmt"
)

const (
	salt = "asoif30FJ#(F_#IJfolf)#(FPSldjfPO85fdsf" //must be secure
)

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
