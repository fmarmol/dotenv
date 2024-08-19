package dotenv

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestLoadFileFromEnv(t *testing.T) {
	home, ok := os.LookupEnv("HOME")
	if !ok {
		home = "local"
	}
	os.Clearenv()
	os.Setenv("HOME", home)
	file := `
		TEST_DOTENV=$HOME/test
		TEST_DOTENV2=$TEST_DOTENV
	`
	buf := bytes.NewBufferString(file)

	err := LoadFromReader(buf)
	if err != nil {
		t.Fatalf("ERROR: %v\n", err)
	}
	os.Unsetenv("HOME")
	envs := os.Environ()
	if len(envs) != 2 {
		t.Fatalf("ERROR: env should have only 2 entries\n")
	}
	if envs[0] != fmt.Sprintf("TEST_DOTENV=%v/test", home) {
		t.Fatalf("ERROR: first entry %q is not correct", envs[0])
	}
	if envs[1] != fmt.Sprintf("TEST_DOTENV2=%v/test", home) {
		t.Fatalf("ERROR: second entry %q is not correct", envs[1])
	}
}
