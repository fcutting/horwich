package testReader

import (
	"fmt"
	"path/filepath"
	"testing"
)

// getTestdataPath return an absolute path that points to the
// internal/testdata folder of this project
func getTestdataPath() (testdataPath string, err error) {
	var currentPath string
	currentPath, err = filepath.Abs("")

	if err != nil {
		err = fmt.Errorf("failed to get current path: %w", err)
		return
	}

	return filepath.Join(currentPath, "..", "..", "internal", "testdata"), nil
}

func Test_getTestdataPath(t *testing.T) {
	testdataPath, err := getTestdataPath()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// This is going to be different between environments. For this reason I'm not
	// checking the output of thefunction, so to check the resutl of this test
	// you can run the following in your terminal:
	// go test -v -run ^Test_getTestdataPath$ ./...
	t.Log("testdataPath:", testdataPath)
}
