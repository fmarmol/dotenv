package dotenv

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func LoadFromReader(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		values := strings.Split(line, "=")
		if len(values) != 2 {
			return fmt.Errorf("Invalide line %d: %q", lineNumber, line)
		}
		key, value := values[0], os.ExpandEnv(values[1])
		os.Setenv(key, value)
		lineNumber++
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func LoadFromFile(path string) error {
	fd, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fd.Close()
	return LoadFromReader(fd)
}
