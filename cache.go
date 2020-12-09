package lru_cache

import (
	"container/list"
	"sync"
)

type Option func(c *Cache)

func WithMaxCapactity(capcatity int) Option {
	return func(c *Cache) {
		c.maxCapacity = capcatity
	}
}

func WithCacheList(list *list.List) Option {
	return func(c *Cache) {
		c.cacheList = list
	}
}

func WithCache(cache map[interface{}]*list.Element) Option {
	return func(c *Cache) {
		c.cache = cache
	}
}

type Cache struct {
	sync.Mutex
	cacheList   *list.List
	maxCapacity int
	cache       map[interface{}]*list.Element
}

func NewCache(maxCapacity int) *Cache {
	return &Cache{
		cacheList:   list.New(),
		maxCapacity: maxCapacity,
		cache:       make(map[interface{}]*list.Element),
	}
}

func New(options ...Option) *Cache {
	lru := new(Cache)
	for _, o := range options {
		o(lru)
	}
	return lru
}

// key - value
type Entry struct {
	key   interface{}
	value interface{}
}

//get value
func (c *Cache) Get(key interface{}) (interface{}, bool) {
	c.Lock()
	defer c.Unlock()
	if elem, ok := c.cache[key]; ok {
		c.cacheList.MoveToFront(elem)
		return elem.Value.(*Entry).value, true
	}
	return nil, false
}

//add value
func (c *Cache) Add(key, value interface{}) bool {
	c.Lock()
	defer c.Unlock()
	//cache has exist
	//update and move to front
	if elem, ok := c.cache[key]; ok {
		c.cacheList.MoveToFront(elem)
		elem.Value.(*Entry).value = value
		return false
	}
	//no exist
	elem := c.cacheList.PushFront(&Entry{
		key:   key,
		value: value,
	})
	c.cache[key] = elem

	if c.cacheList.Len() > c.maxCapacity {
		e := c.cacheList.Back()
		c.cacheList.Remove(e)
		delete(c.cache, e.Value.(*Entry).key)
		return true
	}
	return false
}

//rm one
func (c *Cache) Remove(key interface{}) bool {
	c.Lock()
	defer c.Unlock()
	if elem, ok := c.cache[key]; ok {
		c.cacheList.Remove(elem)
		delete(c.cache, elem.Value.(*Entry).key)
		return true
	}
	return false
}

//get cache length
func (c *Cache) Len() int {
	c.Lock()
	defer c.Unlock()
	return c.cacheList.Len()
}
