package setting

const (
	CHROME            = "chrome"
	INTERNET_EXPLORER = "internet-explorer"
	OPERA             = "opera"
	FIREFOX           = "firefox"
	UC                = "uc-browser"
	SAFARI            = "safari"

	ANDROID  = "android"
	MAC_OS_X = "mac-os-x"
	IOS      = "ios"
	LINUX    = "linux"

	IPHONE = "iphone"
	IPAD   = "ipad"
)

var (
	BrowserUserAgentMaps = map[string][]string{
		"software_name": {
			CHROME,
			INTERNET_EXPLORER,
			FIREFOX,
			UC,
			OPERA,
			SAFARI,
		},
		"operating_system_name": {
			ANDROID,
			MAC_OS_X,
			IOS,
			LINUX,
		},
		"operating_platform": {
			IPHONE,
			IPAD,
		},
	}
)
