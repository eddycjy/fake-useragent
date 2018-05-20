package file

import (
	"io/ioutil"
	"os"
	"strings"
)

type file struct {
	Dir          string
	Name         string
	CompletePath string
}

func NewFileCache(dir string, name string) *file {
	return &file{
		Dir:          dir,
		Name:         name,
		CompletePath: dir + name,
	}
}

func (f *file) Read() ([]byte, error) {
	return ioutil.ReadFile(f.CompletePath)
}

func (f *file) Write(data []byte) error {
	return ioutil.WriteFile(f.CompletePath, data, 0644)
}

func (f *file) Remove() error {
	err := os.Remove(f.CompletePath)
	if err != nil {
		return err
	}

	return nil
}

func (f *file) IsExist() (bool, error) {
	_, err := os.Stat(f.CompletePath)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func GetTempDir() string {
	tempDir := os.TempDir()
	if exist := strings.HasSuffix(tempDir, "/"); exist == false {
		tempDir = tempDir + "/"
	}

	return tempDir
}
