package utils

import (
	"github.com/artdarek/go-unzip"
)

func UnZip(source, destination string) error {

	uz := unzip.New(source, destination)
	err := uz.Extract()
	if err != nil {
		return err
	}

	return nil
}
