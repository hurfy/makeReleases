package main

import "testing"

var config = Config{
	inputPath:  "D:\\test_releases",
	outputPath: "D:\\test_releases\\output",
}

func TestOk(t *testing.T) {
	if err := processFiles(&config); err != nil {
		t.Errorf("Test failed")
	}
}
