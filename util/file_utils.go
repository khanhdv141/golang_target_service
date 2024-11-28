package util

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(source string, dest string) error {
	sourceFileStat, err := os.Stat(source)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", source)
	}

	sourceData, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceData.Close()

	destinationFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceData)
	return err
}
