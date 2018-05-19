package browser

import "testing"

func TestRandom(t *testing.T) {
	if Random() == "" {
		t.Error("browser.Random is empty")
	}
}

func TestChrome(t *testing.T) {
	if Chrome() == "" {
		t.Error("browser.Chrome is empty")
	}
}

func TestInternetExplorer(t *testing.T) {
	if InternetExplorer() == "" {
		t.Error("browser.InternetExplorer is empty")
	}
}

func TestOpera(t *testing.T) {
	if Opera() == "" {
		t.Error("browser.Opera is empty")
	}
}

func TestFirefox(t *testing.T) {
	if Firefox() == "" {
		t.Error("browser.Firefox is empty")
	}
}

func TestUC(t *testing.T) {
	if UC() == "" {
		t.Error("browser.UC is empty")
	}
}

func TestSafari(t *testing.T) {
	if Safari() == "" {
		t.Error("browser.Safari is empty")
	}
}

func TestAndroid(t *testing.T) {
	if Android() == "" {
		t.Error("browser.Android is empty")
	}
}

func TestMacOSX(t *testing.T) {
	if MacOSX() == "" {
		t.Error("browser.MacOSX is empty")
	}
}

func TestIOS(t *testing.T) {
	if IOS() == "" {
		t.Error("browser.IOS is empty")
	}
}

func TestLinux(t *testing.T) {
	if Linux() == "" {
		t.Error("browser.IOS is empty")
	}
}

func TestIPhone(t *testing.T) {
	if IPhone() == "" {
		t.Error("browser.IPhone is empty")
	}
}

func TestIPad(t *testing.T) {
	if IPad() == "" {
		t.Error("browser.IPad is empty")
	}
}