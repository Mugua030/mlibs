package mlibs

import (
	"mlibs/transport/grpc"
	"mlibs/transport/http"
	"testing"
)

func TestApp(t *testing.T) {
	httpSrv := http.NewServer()
	grpcSrv := grpc.NewServer()

	app := NewApp(
		Name("mlibs"),
		Version("v1.0.0"),
		Server(httpSrv, grpcSrv),
	)

	if err := app.Run(); err != nil {
		t.Fatal(err)
	}
}
