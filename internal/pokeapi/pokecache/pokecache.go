package pokecache

import (
  "fmt"
  "sync"
  "time"
)

type Cache struct {
  data map[string]cacheEntry
  cacheLock *sync.Mutex
  interval time.Duration
}

type cacheEntry struct {
  createdAt time.Time
  val []byte
}

func NewCache(interval time.Duration) Cache {
  cache := Cache {
    data: make(map[string]cacheEntry),
    cacheLock: &sync.Mutex{},
    interval: interval,
  }

  cache.reapLoop()
  return cache
}

func (cache *Cache) Add(key string, newVal []byte) {
  cache.cacheLock.Lock()

  cache.data[key] = cacheEntry{
    createdAt: time.Now(),
    val: newVal,
  }

  cache.cacheLock.Unlock()
}

func (cache *Cache) Get(key string) ([]byte, bool) {
  entry, ok := cache.data[key]
  return entry.val, ok
}

func (cache *Cache) reapLoop() {
  ticker := time.NewTicker(cache.interval)

  go func() {
    for {
      <-ticker.C

      cache.cacheLock.Lock()

      for key, entry := range cache.data {
        age := time.Now().Sub(entry.createdAt)

        if age > cache.interval {
          fmt.Printf("deleting %s\n", key)
          fmt.Print("Pokedex > ")
          delete(cache.data, key)
        }
      }

      cache.cacheLock.Unlock()
    }
  }()
}
