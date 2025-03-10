package main

import (
	"context"
	"log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/encoding/json"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/ruinshe/kratos-examples/path-variable-missing/apis"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	apis.UnimplementedUserServiceServer
}

func (s *server) UpdateUser(ctx context.Context, req *apis.UpdateUserRequest) (*apis.User, error) {
	log.Printf("request: %+v", req)
	return &apis.User{}, nil
}

func main() {
	json.MarshalOptions.EmitUnpopulated = false

	s := &server{}
	httpSrv := http.NewServer(
		http.Address(":8080"),
		http.Middleware(recovery.Recovery()),
	)
	// register http hanlder to http server
	apis.RegisterUserServiceHTTPServer(httpSrv, s)

	app := kratos.New(
		kratos.Name("path-with-fields"),
		kratos.Server(
			httpSrv,
		),
	)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
