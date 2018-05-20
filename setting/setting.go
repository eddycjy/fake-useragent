package setting

import "time"

const (
	VERSION = "0.1.0"

	BROWSER_URL            = "https://developers.whatismybrowser.com/useragents/explore/%s/%s/%d"
	BROWSER_MAX_PAGE       = 2
	BROWSER_ALLOW_MAX_PAGE = 5

	HTTP_TIMEOUT         = 15 * time.Second
	HTTP_DELAY           = 100 * time.Millisecond
	HTTP_ALLOW_MIN_DELAY = 100 * time.Millisecond

	TEMP_FILE_NAME      = "fake_useragent_%s.json"
	TEMP_FILE_TEST_NAME = "fake_useragent_test_%s.json"
)
