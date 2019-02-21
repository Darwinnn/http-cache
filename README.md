# http-cache

A fast in-memory REST HTTP key/value database/cache.
Written with fasthttp and fasthttp-router.

Based on the benchmarks I've run this is capable of delivering around *100000* requests per second, which is probably more than you'd ever need for an app like this.

I build it because I couldn't find an open-source solution for key/value database with REST API.

## The http-cache could be used for

- Caching gigabytes of data in RAM and handing them off by native HTTP with low CPU footprint (no RAM-DISKS, no Redis/memcache wrappers are needed)
- As a key/value database with native HTTP interface (no external libs needed)
- As a fast file-server for temporary files

Be aware that http-cache does not persist data on disk, so if it's restarted the data is lost.

## Install

```bash
git clone https://github.com/Darwinnn/http-cache
cd http-cache
make
```

## Run

```bash
Usage of ./app:
  -addr string
        address to listen on (default ":8080")
  -maxsize int
        maximum upload size (in bytes) (default 4294967295)
  -memopt
        Memory optimization (might increase CPU usage)
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

All test are run with 32 byte data, performed on a 4 cores 2.3Ghz i5

PUT value in cache

```bash
baton -u http://localhost:8080/cache/32b -m PUT -f 32b -c 100 -t 60
Configuring to send PUT requests to: http://localhost:8080/cache/32b
Generating the requests...
Finished generating the requests
Sending the requests to the server...
Finished sending the requests
Processing the results...

=========================== Results ========================================

Total requests:                               5899890
Time taken to complete requests:       1m0.002815914s
Requests per second:                            98327

========= Percentage of responses by status code ==========================

Number of connection errors:                        0
Number of 1xx responses:                            0
Number of 2xx responses:                      5899890
Number of 3xx responses:                            0
Number of 4xx responses:                            0
Number of 5xx responses:                            0

===========================================================================
```

GET value from cache: 

```bash
baton -u http://localhost:8080/cache/32b -c 100 -t 60
Configuring to send GET requests to: http://localhost:8080/cache/32b
Generating the requests...
Finished generating the requests
Sending the requests to the server...
Finished sending the requests
Processing the results...

=========================== Results ========================================

Total requests:                               6107886
Time taken to complete requests:       1m0.001823743s
Requests per second:                           101795

========= Percentage of responses by status code ==========================

Number of connection errors:                        0
Number of 1xx responses:                            0
Number of 2xx responses:                      6107886
Number of 3xx responses:                            0
Number of 4xx responses:                            0
Number of 5xx responses:                            0

===========================================================================
```

GET not found key:

```bash
baton -u http://localhost:8080/cache/32b123 -c 100 -t 60
Configuring to send GET requests to: http://localhost:8080/cache/32b123
Generating the requests...
Finished generating the requests
Sending the requests to the server...
Finished sending the requests
Processing the results...

=========================== Results ========================================

Total requests:                               6259580
Time taken to complete requests:       1m0.001837238s
Requests per second:                           104323

========= Percentage of responses by status code ==========================

Number of connection errors:                        0
Number of 1xx responses:                            0
Number of 2xx responses:                            0
Number of 3xx responses:                            0
Number of 4xx responses:                      6259580
Number of 5xx responses:                            0

===========================================================================
```
