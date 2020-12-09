# lru_cache
lru_cache Concurrency security

##install
```
go get github.com/duckbb/lru_cache
```

##quick start
```api
import(
    lru "github.com/duckbb/lru_cache"
)
cache := lru.New(100)
cache.Add("key1","value1")
```
###option init
```api
l := list.New() 
cache := lru.New(WithMaxCapactity(100),WithCacheList(l))
```
###operate
```api
//1.add one element
cache.Add("key","value")
//2.remove one element
cache.Remove("key")
//3.get element
value,ok := cache.Get("key")
//4.get cache length
length := cache.Len()
```
