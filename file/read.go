package file

import (
	"bufio"
	"os"
	"strings"
)

func GetFileContents(path string) string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	out := []string{}
	for s.Scan() {
		out = append(out, s.Text())
	}

	return strings.Join(out, "\n")
}
