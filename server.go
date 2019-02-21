package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func (c *Cache) BuildRouter() *fasthttprouter.Router {
	var router = fasthttprouter.New()

	router.GET("/cache/:key", c.getFromCache)
	router.PUT("/cache/:key", c.setFromCache)
	router.DELETE("/cache/:key", c.deleteFromCache)
	return router
}

func (c *Cache) getFromCache(ctx *fasthttp.RequestCtx) {
	if val, ok := c.Get(ctx.UserValue("key")); ok {
		ctx.SetContentTypeBytes(val.Content)
		fmt.Fprint(ctx, val.Data)
		return
	}
	ctx.Error("Key cache value not found", 404)
}

func (c *Cache) setFromCache(ctx *fasthttp.RequestCtx) {
	var ttl int64
	if ttlVal, err := strconv.Atoi(string(ctx.FormValue("ttl"))); err == nil {
		ttl = int64(ttlVal)
	} else {
		ttl = DefaultTTL
	}

	c.Set(ctx.UserValue("key"), &CacheElement{
		TTL:     time.Now().Add(time.Duration(ttl) * time.Second),
		Data:    ctx.Request.Body(),
		Content: ctx.Request.Header.ContentType(),
	})
	return
}

func (c *Cache) deleteFromCache(ctx *fasthttp.RequestCtx) {
	c.Del(ctx.UserValue("key"))
}

func (c *Cache) StartCleanUpWorker() {
	go func() {
		for {
			c.DelExpired()
			time.Sleep(time.Second)
		}
	}()
}
