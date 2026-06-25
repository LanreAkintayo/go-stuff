/*
Challenge 3: The Safe Cache Memory Leak (Hard)
Scenario: You are building an in-memory cache system for your distributed backend service to store temporary session tokens.

Your Task: Wrap Go's built-in map[string]string inside a SecureCache struct protected by a sync.RWMutex. Implement Set(key, value string) and Get(key string) (string, bool).

The Catch: Implement a third method: DeleteAfter(key string, delay time.Duration). This method must trigger an asynchronous worker (a background goroutine) that sleeps for the duration and then cleanly removes that specific key from the map without causing a deadlock or data race with active readers or writers.

Output;
[Cache Set] Key 'token_123' stored successfully.
[Cache Get] Key 'token_123' found. Value: active_user
[Cache Expiry initiated] Key 'token_123' scheduled for deletion in 50ms.
[Cache Get] Key 'token_123' found. Value: active_user
... (50ms passes) ...
[Background Eviction] Key 'token_123' has expired and was removed from memory.
[Cache Get] Key 'token_123' -> Cache Miss (Not Found).
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

type SecureCache struct {
	cache map[string]string
	lock  sync.RWMutex
}

func (sc *SecureCache) Set(key, value string) {
	sc.lock.Lock()
	defer sc.lock.Unlock()

	sc.cache[key] = value
	fmt.Printf("[Cache Set] Key '%s' stored successfully\n", key)
}

func (sc *SecureCache) Get(key string) (string, bool) {
	sc.lock.RLock()
	value, exist := sc.cache[key]
	sc.lock.RUnlock()

	if !exist {
		fmt.Printf("[Cache Get] Key '%s' -> Cache Miss (Not Found).\n", key)
		return "", false
	}

	fmt.Printf("[Cache Get] Key '%s' found. Value: %s\n", key, value)
	return value, true
}

func (sc *SecureCache) DeleteAfter(key string, sleepTime time.Duration) {
	fmt.Printf("[Cache Expiry initiated] Key '%s' scheduled for deletion in %v.\n", key, sleepTime)
	go func() {
		time.Sleep(sleepTime)
		sc.lock.Lock()
		defer sc.lock.Unlock()

		delete(sc.cache, key)
		fmt.Printf("[Background Eviction] Key '%s' has expired and was removed from memory.\n", key)
	}()
}

func Challenge3() {
	secureCache := SecureCache{cache: make(map[string]string)}

	secureCache.Set("token_123", "active_user")
	secureCache.Get("token_123")
	secureCache.DeleteAfter("token_123", 50*time.Millisecond)
	secureCache.Get("token_123")

	time.Sleep(10 * time.Millisecond)
	fmt.Println("... (50ms passes) ...")
	time.Sleep(50 * time.Millisecond)

	secureCache.Get("token_123")
}
