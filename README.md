# http-cache
Pretty darn fast (~30k rps) in-memory cache with REST interface


# Install
```
git clone https://github.com/Darwinnn/http-cache
cd http-cache
go build
```
<br>

# Run

```
./http-cache --help
Usage of ./http-cache:
  -addr string
    	address to listen on (default ":8080")
  -ttl int
    	default time-to-live of cache objects (default 4294967295)
```
<br>

# Usage

GET /cache/*key* - get value from cache stored by *key*<br>
PUT /cache/*key*?ttl=seconds - put value in cache (ttl argument is optional, of ommited the default value is used)<br>
DELETE /cache/*key* - delete value stored by key<br>
<br>
# Examples
To put a value: <br>
```curl -X PUT -H "Content-Type: Content-type is also cached" http://localhost:8080/cache/Hello -d "World"```

<br>To get a value:<br>

```
curl -v http://localhost:8080/cache/Hello
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> GET /cache/Hello HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content: Content-type is also cached
< Date: Sun, 10 Feb 2019 20:09:33 GMT
< Content-Length: 5
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host localhost left intact
World
```
Note that the Content-Type header is also cached, so wichever contet-type is set when PUTting an object, the same will be set when GETting it.
