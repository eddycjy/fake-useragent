package downloader

import (
	"testing"
	"github.com/EDDYCJY/fake-useragent/setting"
	"time"
)

func TestDownload_Get(t *testing.T) {
	downloader := Download{
		Delay: setting.HTTP_DELAY * time.Millisecond,
		Timeout: setting.HTTP_TIMEOUT * time.Second,
	}

	_, err := downloader.Get("https://developers.whatismybrowser.com")
	if err != nil {
		t.Errorf("downloader.Get err: %v", err)
	}
}