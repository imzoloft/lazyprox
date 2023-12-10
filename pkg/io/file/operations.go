package file

import (
	"errors"
	"log"
	"os"
	"strings"
	"sync"

	inout "github.com/imzoloft/lazyprox/internal/io"
)

var fileMutex sync.Mutex

func ReadFile(fileName string) ([]string, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, errors.New("unable to read file")
	}

	lines := strings.Split(string(file), "\n")

	return lines, nil
}

func WriteToFile(fileName string, message string) {
	fileMutex.Lock()
	defer fileMutex.Unlock()

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.WriteString(message + "\n"); err != nil {
		inout.FatalError(err.Error())
	}
}
