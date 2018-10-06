package ipfsutil

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	file, err := ioutil.TempFile("/tmp", "")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(file.Name())
	if _, err := file.Write([]byte("hello world")); err != nil {
		t.Error(err)
	}

	hash, err := Add(file)
	if hash == "" {
		t.Error("hash is empty")
	}
	if err != nil {
		t.Error(err)
	}

	t.Log(hash)
}

func TestAddBytes(t *testing.T) {
	hash, err := AddBytes([]byte("hello world"))
	if hash == "" {
		t.Error("hash is empty")
	}
	if err != nil {
		t.Error(err)
	}

	t.Log(hash)
}

func TestAddReader(t *testing.T) {
	reader := bytes.NewReader([]byte("hello world"))
	hash, err := AddReader(reader)
	if hash == "" {
		t.Error("hash is empty")
	}
	if err != nil {
		t.Error(err)
	}

	t.Log(hash)
}

func TestAddFile(t *testing.T) {
	file, err := ioutil.TempFile("/tmp", "")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(file.Name())
	if _, err := file.Write([]byte("hello world")); err != nil {
		t.Error(err)
	}

	hash, err := AddFile(file)
	if hash == "" {
		t.Error("hash is empty")
	}
	if err != nil {
		t.Error(err)
	}

	t.Log(hash)
}

func TestPublish(t *testing.T) {
	hash, err := AddBytes([]byte("hello world"))
	if err != nil {
		t.Error(err)
	}
	ipns, err := Publish(hash)
	if ipns == "" {
		t.Error("ipns hash is empty")
	}
	if err != nil {
		t.Error(err)
	}

	t.Log(ipns)
}

func TestResolve(t *testing.T) {
	hash, err := AddBytes([]byte("hello world"))
	if err != nil {
		t.Error(err)
	}
	ipns, err := Publish(hash)
	if err != nil {
		t.Error(err)
	}
	resolvedHash, err := Resolve(ipns)
	if resolvedHash != hash {
		t.Error("hash doesn't match")
	}
	if err != nil {
		t.Error(err)
	}

	t.Log(resolvedHash)
}

func TestGenKey(t *testing.T) {
	keyname := "test123"
	hash, err := GenerateKey(keyname)
	if err != nil {
		t.Error(err)
	}

	t.Log(hash)

	removedName, err := RemoveKey(keyname)
	if err != nil {
		t.Error(err)
	}

	if removedName != keyname {
		t.Error("key does not match")
	}

	t.Log(removedName)
}
