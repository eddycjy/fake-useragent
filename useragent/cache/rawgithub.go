package cache

import (
	"io/ioutil"
	"github.com/EDDYCJY/fake-useragent/downloader"
	"github.com/EDDYCJY/fake-useragent/setting"
	"time"
	"io"
	"net/http"
)

type raw struct {
	Dir          string
	Name         string
	CompletePath string
}

func NewRawCache(dir string, name string) *raw {
	return &raw{
		Dir:          dir,
		Name:         name,
		CompletePath: dir + name,
	}
}

func (f *raw) Get() (*http.Response, bool, error) {
	downloader := downloader.Download{
		Delay: setting.GetDelay(time.Duration(0)),
		Timeout: setting.GetTimeout(time.Duration(0)),
	}

	resp, err := downloader.Get(f.CompletePath)
	if err != nil {
		return nil, false, err
	}

	return resp, f.IsExist(resp), nil
}

func (f *raw) Read(body io.ReadCloser) ([]byte, error) {
	return ioutil.ReadAll(body)
}

func (f *raw) IsExist(resp *http.Response) bool {
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNotModified {
		return true
	}

	return false
}