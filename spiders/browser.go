package spiders

import (
	"time"
	"log"
	"sync"
	"fmt"

	"github.com/PuerkitoBio/goquery"

	"github.com/EDDYCJY/fake-useragent/setting"
	"github.com/EDDYCJY/fake-useragent/scheduler"
	"github.com/EDDYCJY/fake-useragent/useragent"
	"github.com/EDDYCJY/fake-useragent/downloader"
)

type Spider struct {
	Attribute
	FullUrl  string
}

type Attribute struct {
	Tag      string
	Category string
	Page     int
}

var urlAttributeResults = make(map[string]Attribute)

func NewBrowserSpider() *Spider {
	return &Spider{}
}

func (a *Attribute) GetSpider() *Spider {
	return &Spider{
		Attribute: Attribute{
			Tag: a.Tag,
			Category: a.Category,
			Page: a.Page,
		},
		FullUrl: fmt.Sprintf(setting.BROWSER_URL, a.Tag, a.Category, a.Page),
	}
}

func (s *Spider) AppendBrowser(maxPage int) {
	for tag, categorys := range setting.BrowserUserAgentMaps {
		for _, category := range categorys {
			for page := 1; page <= maxPage; page++ {
				attribute := Attribute{Tag: tag, Category: category, Page: page}
				urlAttributeResults[attribute.GetSpider().FullUrl] = attribute
				scheduler.AppendUrl(attribute.GetSpider().FullUrl)
			}
		}
	}
}

func (s *Spider) StartBrowser(delay time.Duration, timeout time.Duration) {
	var wg sync.WaitGroup
	count := scheduler.CountUrl()
	for i := 0; i <= count; i ++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if url := scheduler.PopUrl(); url != "" {
				downloader := downloader.Download{Delay: delay, Timeout: timeout}
				body, err := downloader.Get(url)
				if err != nil {
					log.Fatalf("downloader.Get err: %v", err)
				}

				doc, err := goquery.NewDocumentFromReader(body)
				if err != nil {
					log.Fatalf("goquery.NewDocumentFromReader err: %v", err)
				}

				doc.Find("td.useragent a").Each(func(i int, selection *goquery.Selection) {
					if value := selection.Text(); value != "" {
						useragent.UA.Set(urlAttributeResults[url].Category, value)
					}
				})
			}
		}()
	}

	wg.Wait()
}