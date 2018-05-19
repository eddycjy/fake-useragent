package browser

import (
	"github.com/EDDYCJY/fake-useragent/setting"
	"github.com/EDDYCJY/fake-useragent/useragent"
)

type Browser struct {}

var DefaultBrowser Browser

func Random() string {
	return DefaultBrowser.Random()
}

func Chrome() string {
	return DefaultBrowser.Chrome()
}

func InternetExplorer() string {
	return DefaultBrowser.InternetExplorer()
}

func Opera() string {
	return DefaultBrowser.Opera()
}

func Firefox() string {
	return DefaultBrowser.Firefox()
}

func UC() string {
	return DefaultBrowser.UC()
}

func Safari() string {
	return DefaultBrowser.Safari()
}

func Android() string {
	return DefaultBrowser.Android()
}

func MacOSX() string {
	return DefaultBrowser.MacOSX()
}

func IOS() string {
	return DefaultBrowser.IOS()
}

func Linux() string {
	return DefaultBrowser.Linux()
}

func IPhone() string {
	return DefaultBrowser.IPhone()
}

func IPad() string {
	return DefaultBrowser.IPad()
}

func (b *Browser) Random() string {
	return useragent.UA.GetAllRandom()
}

func (b *Browser) Chrome() string {
	return useragent.UA.GetRandom(setting.CHROME)
}

func (b *Browser) InternetExplorer() string {
	return useragent.UA.GetRandom(setting.INTERNET_EXPLORER)
}

func (b *Browser) Opera() string {
	return useragent.UA.GetRandom(setting.OPERA)
}

func (b *Browser) Firefox() string {
	return useragent.UA.GetRandom(setting.FIREFOX)
}

func (b *Browser) UC() string {
	return useragent.UA.GetRandom(setting.UC)
}

func (b *Browser) Safari() string {
	return useragent.UA.GetRandom(setting.SAFARI)
}

func (b *Browser) Android() string {
	return useragent.UA.GetRandom(setting.ANDROID)
}

func (b *Browser) MacOSX() string {
	return useragent.UA.GetRandom(setting.MAC_OS_X)
}

func (b *Browser) IOS() string {
	return useragent.UA.GetRandom(setting.IOS)
}

func (b *Browser) Linux() string {
	return useragent.UA.GetRandom(setting.LINUX)
}

func (b *Browser) IPhone() string {
	return useragent.UA.GetRandom(setting.IPHONE)
}

func (b *Browser) IPad() string {
	return useragent.UA.GetRandom(setting.IPAD)
}