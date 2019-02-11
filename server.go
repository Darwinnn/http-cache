package main

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (c *Cache) BuildRouter() *httprouter.Router {
	var router = httprouter.New()

	router.GET("/cache/:key", c.getFromCache)
	router.PUT("/cache/:key", c.setFromCache)
	router.DELETE("/cache/:key", c.deleteFromCache)

	return router
}

func (c *Cache) getFromCache(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if val, ok := c.Get(ps.ByName("key")); ok {
		w.Header().Set("Content-Type", val.Content)
		w.Write(val.Data)
		return
	}
	http.Error(w, "Key cache value not found", http.StatusNotFound)
}

func (c *Cache) setFromCache(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var ttl int64
	if ttlVal, err := strconv.Atoi(r.FormValue("ttl")); err == nil {
		ttl = int64(ttlVal)
	} else {
		ttl = DefaultTTL
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	elem := CacheElement{
		TTL:     time.Now().Add(time.Duration(ttl) * time.Second),
		Data:    data,
		Content: r.Header.Get("Content-Type"),
	}

	c.Set(ps.ByName("key"), elem)
	w.WriteHeader(http.StatusOK)
	return
}

func (c *Cache) deleteFromCache(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c.Del(ps.ByName("key"))
}
