package browser

import (
	"github.com/EDDYCJY/fake-useragent/setting"
	"github.com/EDDYCJY/fake-useragent/useragent"
)

func Random() string {
	return defaultBrowser.Random()
}

func Chrome() string {
	return defaultBrowser.Chrome()
}

func InternetExplorer() string {
	return defaultBrowser.InternetExplorer()
}

func Firefox() string {
	return defaultBrowser.Firefox()
}

func Safari() string {
	return defaultBrowser.Safari()
}

func Android() string {
	return defaultBrowser.Android()
}

func MacOSX() string {
	return defaultBrowser.MacOSX()
}

func IOS() string {
	return defaultBrowser.IOS()
}

func Linux() string {
	return defaultBrowser.Linux()
}

func IPhone() string {
	return defaultBrowser.IPhone()
}

func IPad() string {
	return defaultBrowser.IPad()
}

func Computer() string {
	return defaultBrowser.Computer()
}

func Mobile() string {
	return defaultBrowser.Mobile()
}

func (b *browser) Random() string {
	return useragent.UA.GetAllRandom()
}

func (b *browser) Chrome() string {
	return useragent.UA.GetRandom(setting.CHROME)
}

func (b *browser) InternetExplorer() string {
	return useragent.UA.GetRandom(setting.INTERNET_EXPLORER)
}

func (b *browser) Firefox() string {
	return useragent.UA.GetRandom(setting.FIREFOX)
}

func (b *browser) Safari() string {
	return useragent.UA.GetRandom(setting.SAFARI)
}

func (b *browser) Android() string {
	return useragent.UA.GetRandom(setting.ANDROID)
}

func (b *browser) MacOSX() string {
	return useragent.UA.GetRandom(setting.MAC_OS_X)
}

func (b *browser) IOS() string {
	return useragent.UA.GetRandom(setting.IOS)
}

func (b *browser) Linux() string {
	return useragent.UA.GetRandom(setting.LINUX)
}

func (b *browser) IPhone() string {
	return useragent.UA.GetRandom(setting.IPHONE)
}

func (b *browser) IPad() string {
	return useragent.UA.GetRandom(setting.IPAD)
}

func (b *browser) Computer() string {
	return useragent.UA.GetRandom(setting.COMPUTER)
}

func (b *browser) Mobile() string {
	return useragent.UA.GetRandom(setting.MOBILE)
}