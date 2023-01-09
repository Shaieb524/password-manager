package appcache

import (
	"errors"
	"sync"
	"time"

	"password-manager/src/models/database/accountPassword"
)

type LocalCache struct {
	stop chan struct{}

	wg           sync.WaitGroup
	mu           sync.RWMutex
	AccPasswords map[string]accountPassword.CachedAccountPassword
}

func NewLocalCache(cleanupInterval time.Duration) *LocalCache {
	lc := &LocalCache{
		AccPasswords: make(map[string]accountPassword.CachedAccountPassword),
		stop:         make(chan struct{}),
	}

	lc.wg.Add(1)
	go func(cleanupInterval time.Duration) {
		defer lc.wg.Done()
		lc.cleanupLoop(cleanupInterval)
	}(cleanupInterval)

	return lc
}

func (lc *LocalCache) cleanupLoop(interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-lc.stop:
			return
		case <-t.C:
			lc.mu.Lock()
			for sName, accPass := range lc.AccPasswords {
				if accPass.ExpireAtTimestamp <= time.Now().Unix() {
					delete(lc.AccPasswords, sName)
				}
			}
			lc.mu.Unlock()
		}
	}
}

var (
	errUserNotInCache = errors.New("the user isn't in cache")
)

func (lc *LocalCache) stopCleanup() {
	close(lc.stop)
	lc.wg.Wait()
}

func (lc *LocalCache) Update(accPass accountPassword.AccountPasswordInputDto, expireAtTimestamp int64) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.AccPasswords[accPass.Service] = accountPassword.CachedAccountPassword{
		AccountPasswordInputDto: accPass,
		ExpireAtTimestamp:       expireAtTimestamp,
	}
}

func (lc *LocalCache) Read(serviceName string) (accountPassword.CachedAccountPassword, error) {
	lc.mu.RLock()
	defer lc.mu.RUnlock()

	cu, ok := lc.AccPasswords[serviceName]
	if !ok {
		return accountPassword.CachedAccountPassword{}, errUserNotInCache
	}

	return cu, nil
}

func (lc *LocalCache) delete(serviceName string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	delete(lc.AccPasswords, serviceName)
}
