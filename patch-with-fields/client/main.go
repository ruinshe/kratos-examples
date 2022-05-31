package main

import (
	"context"
	"log"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/lithammer/shortuuid/v4"
	"github.com/ruinshe/kratos-examples/patch-with-fields/apis"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
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

	fields, _ := fieldmaskpb.New(&apis.User{}, "name", "age")

	_, err = client.UpdateUser(ctx, &apis.UpdateUserRequest{
		User:   &apis.User{Id: shortuuid.New(), Name: "Jack"},
		Fields: fields,
	})
	if err != nil {
		log.Fatal(err)
	}
}
