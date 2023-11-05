package main

import (
	"context"
	"log"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/lithammer/shortuuid/v4"
	"github.com/ruinshe/kratos-examples/path-variable-missing/apis"
)

func main() {
	conn, err := http.NewClient(
		context.Background(),
		http.WithEndpoint("127.0.0.1:8080"),
	)
	if err != nil {
		panic(err)
	}
	client := apis.NewUserServiceHTTPClient(conn)
	ctx := context.Background()

	id := shortuuid.New()
	_, err = client.UpdateUser(ctx, &apis.UpdateUserRequest{
		UserId: id,
		User:   &apis.User{Id: id, Name: "Jack"},
	})
	if err != nil {
		log.Fatal(err)
	}
}
