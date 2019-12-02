package useragent

import (
	"math/rand"
	"sync"
	"time"
)

type useragent struct {
	data map[string][]string
	lock sync.RWMutex
}

var (
	UA = useragent{data: make(map[string][]string)}
	r  = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func (u *useragent) Get(key string) []string {
	u.lock.RLock()
	defer u.lock.RUnlock()
	return u.data[key]
}

func (u *useragent) GetAll() map[string][]string {
	u.lock.RLock()
	defer u.lock.RUnlock()
	return u.data
}

func (u *useragent) GetRandom(key string) string {
	u.lock.RLock()
	defer u.lock.RUnlock()
	browser := u.Get(key)
	len := len(browser)
	if len < 1 {
		return ""
	}

	n := r.Intn(len)
	return browser[n]
}

func (u *useragent) GetAllRandom() string {
	u.lock.RLock()
	defer u.lock.RUnlock()
	browsers := u.data
	totalLen := 0
	for _, uas := range browsers {
		totalLen += len(uas)
	}

	if totalLen < 1 {
		return ""
	}

	index := r.Intn(totalLen)
	var ua string
	for _, uas := range browsers {
		if index+1 <= len(uas) {
			ua = uas[index]
			break
		}
		index = index - len(uas)
	}

	return ua
}

func (u *useragent) Set(key, value string) {
	u.lock.Lock()
	u.data[key] = append(u.data[key], value)
	u.lock.Unlock()
}

func (u *useragent) SetData(data map[string][]string) {
	u.lock.Lock()
	u.data = data
	u.lock.Unlock()
}
