package browser

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/EDDYCJY/fake-useragent/setting"
	"github.com/EDDYCJY/fake-useragent/spiders"
	"github.com/EDDYCJY/fake-useragent/useragent"
	"github.com/EDDYCJY/fake-useragent/useragent/cache"
)

type browser struct {
	Client
	Cache
}

type Client struct {
	MaxPage int
	Delay   time.Duration
	Timeout time.Duration
}

type Cache struct {
	CloseFile  bool
	UpdateFile bool
}

var defaultBrowser = NewBrowser(Client{
	MaxPage: setting.BROWSER_MAX_PAGE,
	Delay:   setting.HTTP_DELAY,
	Timeout: setting.HTTP_TIMEOUT,
}, Cache{
	CloseFile: false,
})

func NewBrowser(client Client, cache Cache) *browser {
	maxPage := client.MaxPage
	if maxPage > setting.BROWSER_ALLOW_MAX_PAGE || maxPage == 0 {
		maxPage = setting.BROWSER_MAX_PAGE
	}
	delay := client.Delay
	if delay < setting.HTTP_ALLOW_MIN_DELAY {
		delay = setting.HTTP_ALLOW_MIN_DELAY
	}
	timeout := client.Timeout
	if timeout == 0 {
		timeout = setting.HTTP_TIMEOUT
	}

	b := browser{
		Client: Client{
			MaxPage: maxPage,
			Delay:   delay,
			Timeout: timeout,
		},
		Cache: Cache{
			CloseFile:  cache.CloseFile,
			UpdateFile: cache.UpdateFile,
		},
	}
	return b.load()
}

func (b *browser) load() *browser {
	fileCache := file.NewFileCache(file.GetTempDir(), fmt.Sprintf(setting.TEMP_FILE_NAME, setting.VERSION))
	fileCacheExist, err := fileCache.IsExist()
	if err != nil {
		log.Fatalf("fileCache.IsExist err: %v", err)
	}

	if b.CloseFile == false && b.UpdateFile == false && fileCacheExist == true {
		fileRead, err := fileCache.Read()
		if err != nil {
			log.Fatalf("fileCache.Read err: %v", err)
		}

		m := make(map[string][]string)
		json.Unmarshal(fileRead, &m)
		useragent.UA.SetData(m)

		return b
	}

	s := spiders.NewBrowserSpider()
	s.AppendBrowser(b.MaxPage)
	s.StartBrowser(b.Delay, b.Timeout)

	if b.CloseFile == true {
		return b
	}

	if fileCacheExist == true && b.UpdateFile == true {
		err := fileCache.Remove()
		if err != nil {
			log.Fatalf("fileCache.Remove err: %v", err)
		}
	}

	uasJson, err := json.Marshal(useragent.UA.GetAll())
	if err != nil {
		log.Fatalf("ua.json.Marshal err: %v", err)
	}

	fileCache.Write(uasJson)
	return b
}
