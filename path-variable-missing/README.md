This demo reproduce [Kratos Issue 2761](https://github.com/go-kratos/kratos/issues/2761)

## Background

```log
go: go version go1.23.3 linux/amd64
kratos: v2.8.4
```

First, run the server using the command below:

```sh
go run server/main.go
```

Start TCP dump using command:

```sh
sudo tcpdump -i lo dst 127.0.0.1 and port 8080 -A
```

Then, run the client using the command below:

```sh
go run client/main.go
```

The dumping HTTP request:

```http
GET /api/v1/users/actions:updateUser HTTP/1.1
Host: 127.0.0.1:8080
User-Agent: Go-http-client/1.1
Content-Type: application/json
Referer: http://127.0.0.1:8080/api/v1/users//actions:updateUser
Accept-Encoding: gzip
```

In this case, the `user_id` URL parameter is not encoded.
