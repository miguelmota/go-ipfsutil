package ipfsutil

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"os/exec"
	"strings"
)

// Add ...
func Add(data interface{}) (string, error) {
	switch v := data.(type) {
	case []byte:
		return AddBytes(v)
	case *os.File:
		return AddFile(v)
	case multipart.File:
		return AddReader(v)
	case io.Reader:
		return AddReader(v)
	default:
		return "", errors.New("not supported")
	}
}

// AddBytes ...
func AddBytes(data []byte) (string, error) {
	file, err := ioutil.TempFile("/tmp", "")
	if err != nil {
		return "", err
	}
	defer os.Remove(file.Name())
	if _, err := file.Write(data); err != nil {
		return "", err
	}
	out, _ := exec.Command("ipfs", "add", "-q", file.Name()).Output()
	hash := strings.TrimSpace(string(out))
	return hash, nil
}

// AddReader ...
func AddReader(r io.Reader) (string, error) {
	file, err := ioutil.TempFile("/tmp", "")
	if err != nil {
		return "", err
	}
	defer os.Remove(file.Name())
	if _, err := io.Copy(file, r); err != nil {
		return "", err
	}

	out, _ := exec.Command("ipfs", "add", "-q", file.Name()).Output()
	hash := strings.TrimSpace(string(out))
	return hash, nil
}

// AddFile ...
func AddFile(file *os.File) (string, error) {
	out, _ := exec.Command("ipfs", "add", "-q", file.Name()).Output()
	hash := strings.TrimSpace(string(out))
	return hash, nil
}

// Publish returns IPNS name
func Publish(hash string) (string, error) {
	out, _ := exec.Command("bash", "-c", fmt.Sprintf("ipfs name publish --key=default /ipfs/%s | awk '{print $3}' | sed 's/://'", hash)).Output()
	ipns := strings.TrimSpace(string(out))
	return ipns, nil
}

// Resolve resolves IPNS name
func Resolve(ipns string) (string, error) {
	out, _ := exec.Command("bash", "-c", fmt.Sprintf(`ipfs name resolve %s | sed 's/\/ipfs\///'`, ipns)).Output()
	hash := strings.TrimSpace(string(out))
	return hash, nil
}

// GenerateKey generate a new key
func GenerateKey(name string) (string, error) {
	out, _ := exec.Command("ipfs", "key", "gen", name, "--type=rsa", "--size=2048").Output()
	hash := strings.TrimSpace(string(out))
	return hash, nil
}

// RemoveKey deletes a key
func RemoveKey(name string) (string, error) {
	out, _ := exec.Command("ipfs", "key", "rm", name).Output()
	removedName := strings.TrimSpace(string(out))
	return removedName, nil
}
