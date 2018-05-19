package browser

import (
	"fmt"
	"log"
	"encoding/json"
	"time"

	"github.com/EDDYCJY/fake-useragent/setting"
	"github.com/EDDYCJY/fake-useragent/useragent/cache"
	"github.com/EDDYCJY/fake-useragent/spiders"
	"github.com/EDDYCJY/fake-useragent/useragent"
)

func init() {
 	fileCache := file.NewFileCache(file.GetTempDir(), fmt.Sprintf(setting.TEMP_FILE_NAME, setting.VERSION))
	fileCacheExist, err := fileCache.IsExist()
	if err != nil {
		log.Fatalf("fileCache.IsExist err: %v", err)
	}

	if fileCacheExist == false {
		s := spiders.NewBrowserSpider()
		s.AppendBrowser(setting.BROWSER_MAX_PAGE)
		s.StartBrowser(setting.HTTP_DELAY * time.Millisecond, setting.HTTP_TIMEOUT * time.Second)

		uasJson, err := json.Marshal(useragent.UA.GetAll())
		if err != nil {
			log.Fatalf("ua.json.Marshal err: %v", err)
		}

		fileCache.Write(uasJson)
		return
	}

	fileRead, err := fileCache.Read()
	if err != nil {
		log.Fatalf("fileCache.Read err: %v", err)
	}

	m := make(map[string][]string)
	json.Unmarshal(fileRead, &m)
	useragent.UA.SetData(m)
}
