package files

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

func CopyFile(filePath, outputPath string) error {
	// open the file
	sourceFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// create a new file
	newFile, err := os.Create(filepath.Join(outputPath, strings.ReplaceAll(filePath, "/", "+")))
	if err != nil {
		return err
	}
	defer newFile.Close()

	// copying data from the original file to a new file
	_, err = io.Copy(newFile, sourceFile)
	if err != nil {
		return err
	}

	// set access parameters
	sourceInfo, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	err = os.Chmod(outputPath, sourceInfo.Mode())
	if err != nil {
		return err
	}

	return nil
}
