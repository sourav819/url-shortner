package utils

import (
	"url-shortner/pkg/logger"

	"github.com/jaevor/go-nanoid"
)

const DefaultAlphabet = "0123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

func CreateNanoID(length int) (string, error) {
	nanoID, err := nanoid.CustomASCII(DefaultAlphabet, length)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	return nanoID(), nil
}
