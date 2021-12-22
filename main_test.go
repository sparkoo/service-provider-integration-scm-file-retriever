package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

//func TestGetFile(t *testing.T) {
//	dst := tempTestFile(t)
//	defer os.RemoveAll(filepath.Dir(dst))
//	if err := getter.GetFile(dst, "github.com/redhat-appstudio/service-provider-integration-operator?filename=Makefile"); err != nil {
//		t.Fatalf("err: %s", err)
//	}
//
//	// Verify the main file exists
//	assertContents(t, dst, "Hello\n")
//}

func tempTestFile(t *testing.T) string {
	dir := tempDir(t)
	return filepath.Join(dir, "foo")
}

func tempDir(t *testing.T) string {
	dir, err := ioutil.TempDir("", "tf")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if err := os.RemoveAll(dir); err != nil {
		t.Fatalf("err: %s", err)
	}

	return dir
}

func assertContents(t *testing.T, path string, contents string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if !reflect.DeepEqual(data, []byte(contents)) {
		t.Fatalf("bad. expected:\n\n%s\n\nGot:\n\n%s", contents, string(data))
	}
}
