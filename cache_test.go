package lru_cache

import "testing"

var testKey = []string{
	"key1",
	"key2",
	"key3",
	"key4",
	"key5",
	"key6",
}

var testValue = []string{
	"v1",
	"v2",
	"v3",
	"v4",
	"v5",
	"v6",
}

//test init cache normal function
func getCache() *Cache {
	lru := NewCache(len(testKey))
	for i := 0; i < len(testKey); i++ {
		lru.Add(testKey[i], testValue[i])
	}
	return lru
}

//test get
func TestGet(t *testing.T) {
	cache := getCache()
	for i, key := range testKey {
		value, _ := cache.Get(key)
		if value != testValue[i] {
			t.Errorf("key:%v,value:%v not equal realValue:%v", key, value, testValue[i])
		}
	}
}

//test add
func TestAdd(t *testing.T) {
	cache := getCache()
	//more than capacity old will remove
	cache.Add("k8", "v8")
	_, ok := cache.Get("k8")
	if !ok {
		t.Error("add element fail")
	}
	_, ok = cache.Get("key1")
	if ok {
		t.Error("old key can not remove")
	}

	//add exist key will update value
	cache.Add("k8", "v88")
	v, ok := cache.Get("k8")
	if !ok {
		t.Errorf("not get key:k8")
	}
	if v != "v88" {
		t.Error("get k8 value not equal v88")
	}
}

//test Remove key
func TestRemove(t *testing.T) {
	cache := getCache()
	for _, key := range testKey {
		cache.Remove(key)
		if _, ok := cache.Get(key); ok {
			t.Errorf("remove key:%v fail\n", key)
		}
	}

}

func TestLen(t *testing.T) {
	cache := getCache() //capacity:6
	cache.Add("key7", "value7")
	cache.Add("key8", "value8")
	t.Log("cache length:", cache.Len())
	if cache.Len() != 6 {
		t.Error("len not equal")
	}
	cache.Remove("key7")
	cache.Remove("key8")
	if cache.Len() != 4 {
		t.Error("len not equal")
	}
}
