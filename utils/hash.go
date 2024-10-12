package utils

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func Hash(str string) ([]byte, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("error while hashing a password")
	}

	return pass, nil
}

func VerifyHashed(hashedStr []byte, candidateStr string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedStr), []byte(candidateStr))

	return err == nil
}
