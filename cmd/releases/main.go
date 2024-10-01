package main

import (
	"flag"
	"os"
	"releases/internal/app/files"
)

type Config struct {
	inputPath  string
	outputPath string
}

func configureFlags() *Config {
	var config = new(Config)

	flag.StringVar(&config.inputPath, "i", "./", "Root directory")
	flag.StringVar(&config.outputPath, "o", "./output", "Output directory")

	return config
}

func processFiles(config *Config) error {
	// create output folder
	if err := os.MkdirAll(config.outputPath, os.ModePerm); err != nil {
		return err
	}

	// recursive process files
	if err := files.ProcessFiles(config.inputPath, config.outputPath); err != nil {
		return err
	}

	return nil
}

func main() {
	var config = configureFlags()

	if err := processFiles(config); err != nil {
		panic(err.Error())
	}
}
