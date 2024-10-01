package files

import (
	"io/fs"
	"os"
)

func ProcessFiles(inputPath, outputPath string) error {
	if err := fs.WalkDir(
		os.DirFS(inputPath),
		".",
		func(path string, entry fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			// skip output folder || .exe file
			if (entry.IsDir() && entry.Name() == "output") || entry.Name() == "releases.exe" {
				return fs.SkipDir
			}

			// copy if file
			if !entry.IsDir() {
				if err := CopyFile(path, outputPath); err != nil {
					return err
				}
			}

			return nil

		},
	); err != nil {
		return err
	}

	return nil
}
