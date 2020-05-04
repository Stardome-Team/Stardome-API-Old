package utils

import (
	"github.com/alexedwards/argon2id"
)

// GenerateHashFromPassword :
func GenerateHashFromPassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)

	if err != nil {
		// TODO: Log Error
		return "", err
	}

	return hash, nil

}

// CompareHashWithPassword :
func CompareHashWithPassword(password string, hash string) (match bool, err error) {

	match, err = argon2id.ComparePasswordAndHash(password, hash)

	if err != nil {
		// TODO: Log Error
		return match, err
	}

	return match, nil
}
