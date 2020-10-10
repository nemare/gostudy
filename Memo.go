package main

//import (
//	"io/ioutil"
//	"net/http"
//	"sync"
//)
//
//type Memo struct {
//	f     Func
//	cache map[string]result
//	mu sync.Mutex
//}
//
//type Func func(key string) (interface{}, error)
//
//type result struct {
//	value interface{}
//	err   error
//}
//
//func New(f Func) *Memo {
//	return &Memo{f: f, cache: make(map[string]result)}
//}
//
//func (memo *Memo) Get(key string) (interface{}, error) {
//	res, ok := memo.cache[key]
//	if !ok {
//		res.value, res.err = memo.f(key)
//		memo.cache[key] = res
//	}
//	return res.value, res.err
//}
//
//func httpGetBody(url string) (interface{}, error) {
//	resp, err := http.Get(url)
//	if err != nil {
//		return nil, err
//	}
//
//	defer resp.Body.Close()
//	return ioutil.ReadAll(resp.Body)
//}
//
