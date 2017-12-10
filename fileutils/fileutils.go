package fileutils

import (
	"bufio"
	"io/ioutil"
	"os"
)

func ByLine(filename string) (<-chan string, <-chan error) {
	lines := make(chan string)
	errc := make(chan error)

	go func() {
		defer close(lines)
		defer close(errc)

		file, err := os.Open(filename)
		if err != nil {
			errc <- err
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			errc <- err
			return
		}
	}()

	return lines, errc
}

func WholeFile(filename string) (string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
