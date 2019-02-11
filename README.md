# http-cache

A pretty fast (~35k rps on my old macbook) KEY/VALUE cache with REST-like API

## Install

```bash
git clone https://github.com/Darwinnn/http-cache
cd http-cache
make
```

## Run

```bash
./app --help
Usage of ./app:
  -addr string
    	address to listen on (default ":8080")
  -ttl int
    	default time-to-live of cache objects (default 4294967295)
```

## Run with docker

### Build a container

```bash
make docker
```

### Start a container

```bash
docker-compose up -d
```

## Usage

- *GET* `/cache/`*key* - get value from cache stored by *key*
- *PUT* `/cache/`*key*`?ttl=seconds` - put value in cache (ttl argument is optional, of ommited the default value is used)
- *DELETE* `/cache/`*key* - delete value stored by key


## Examples

- Put a value:

```bash
curl -X PUT -H "Content-Type: Content-type is also cached" http://localhost:8080/cache/Hello -d "World"
```

- Get a value:

```bash
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
< Date: Sun, 10 Feb 2019 20:09:33 GMT
< Content-Length: 5
< Content-Type: Content-type is also cached
<
* Connection #0 to host localhost left intact
World
```

Note that the Content-Type header is also cached, so wichever contet-type is set when PUTting an object, the same will be set when GETting it.

## Benchmarks

PUT value in cache (this is the slowest method, since it's done in serial, not parallel)

```bash
go/bin/baton -u http://localhost:8080/cache/test -m PUT -b "HELLO" -t 60 -c 100
Configuring to send PUT requests to: http://localhost:8080/cache/test
Generating the requests...
Finished generating the requests
Sending the requests to the server...
Finished sending the requests
Processing the results...

=========================== Results ========================================

Total requests:                               1625914
Time taken to complete requests:       1m0.008456565s
Requests per second:                            27095

========= Percentage of responses by status code ==========================

Number of connection errors:                        0
Number of 1xx responses:                            0
Number of 2xx responses:                      1625914
Number of 3xx responses:                            0
Number of 4xx responses:                            0
Number of 5xx responses:                            0

===========================================================================
```

GET value from cache: 

```bash
go/bin/baton -u http://localhost:8080/cache/test -t 60 -c 100
Configuring to send GET requests to: http://localhost:8080/cache/test
Generating the requests...
Finished generating the requests
Sending the requests to the server...
Finished sending the requests
Processing the results...

=========================== Results ========================================

Total requests:                               2092374
Time taken to complete requests:       1m0.004429216s
Requests per second:                            34870

========= Percentage of responses by status code ==========================

Number of connection errors:                        0
Number of 1xx responses:                            0
Number of 2xx responses:                      2092374
Number of 3xx responses:                            0
Number of 4xx responses:                            0
Number of 5xx responses:                            0

===========================================================================
```

GET not found key:

```bash
go/bin/baton -u http://localhost:8080/cache/NOT_FOUND -t 60 -c 100
Configuring to send GET requests to: http://localhost:8080/cache/NOT_FOUND
Generating the requests...
Finished generating the requests
Sending the requests to the server...
Finished sending the requests
Processing the results...

=========================== Results ========================================

Total requests:                               1855883
Time taken to complete requests:         1m0.0065943s
Requests per second:                            30928

========= Percentage of responses by status code ==========================

Number of connection errors:                        0
Number of 1xx responses:                            0
Number of 2xx responses:                            0
Number of 3xx responses:                            0
Number of 4xx responses:                      1855883
Number of 5xx responses:                            0

===========================================================================
```
