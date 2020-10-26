package helpers

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// ReadFile takes a Document to ([]string, error)
func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// FileToString reads a File to a string without \n
func FileToString(path string) (string, error) {
	lines, err := ReadFile(path)
	var answer string
	if err != nil {
		return "", err
	}
	for _, line := range lines {
		answer += line
	}
	return answer, nil
}

// AppendToFile appends text to a File
// Creates File if it doesnt exists
// returns written Bytes and Error
func AppendToFile(path, text string) (int, error) {
	file, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	writtenBytes, err := file.WriteString(text)
	if err != nil {
		return writtenBytes, err
	}
	return writtenBytes, nil
}

// WriteToFile creates a textfile
// returns written Bytes and Error
func WriteToFile(path, text string) (int, error) {
	file, err := os.Create(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	writtenBytes, err := file.WriteString(text)
	if err != nil {
		return writtenBytes, err
	}
	return writtenBytes, nil
}

// MakeExecutable chages the file Permissions to 0755
func MakeExecutable(path string) error {
	return os.Chmod(path, 0755)
}

func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func haveReadPermissions(path string) bool {
	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		if os.IsPermission(err) {
			return false
		}
	}
	return true
}

func haveWritePermissions(path string) bool {
	file, err := os.OpenFile("test.txt", os.O_RDONLY, 0666)
	defer file.Close()
	if err != nil {
		if os.IsPermission(err) {
			return false
		}
	}
	return true
}

func CheckFile(path string) bool {
	return haveReadPermissions(path) && haveWritePermissions(path) && FileExists(path)
}
