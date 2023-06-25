package core

import "sync"

type Cache struct {
	Visited map[string]bool
	Lock    sync.Mutex
}

func (c *Cache) IsVisited(url string) bool {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	_, ok := c.Visited[url]
	return ok
}

func (c *Cache) AddVisited(url string) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.Visited[url] = true
}

func (c *Cache) Flush() {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.Visited = make(map[string]bool)
}
