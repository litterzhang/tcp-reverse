# tcp-reverse

Steps:
1. Start a http server listenes on port 3001 by: `go run http-server.go`

2. Start relay server by `go run relay-server.go`

relay
```
2024/03/21 15:07:53 Start listening on port 8000 for realy client connection
```

3. Start relay client by `go run relay-client.go`

client
```
2024/03/21 15:09:13 Connected to port 8000
```

relay
```
2024/03/21 15:07:53 Start listening on port 8000 for realy client connection
2024/03/21 15:09:13 Start listening on port 8001 for client connections
```

4. curl the relay server by `curl -sv http://localhost:8001`


```
* Rebuilt URL to: http://localhost:8001/
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8001 (#0)
> GET / HTTP/1.1
> Host: localhost:8001
> User-Agent: curl/7.61.1
> Accept: */*
> 
< HTTP/1.1 200 OK
< Date: Thu, 21 Mar 2024 07:10:32 GMT
< Content-Length: 13
< Content-Type: text/plain; charset=utf-8
< 
* Connection #0 to host localhost left intact
Hello, World!% 
```


