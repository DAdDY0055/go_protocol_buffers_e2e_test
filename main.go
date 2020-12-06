package main

import (
	"net/http"

	pb "github.com/DAdDY0055/go_protocol_buffers_e2e_test/proto"
	"github.com/labstack/echo/v4"

	"google.golang.org/protobuf/proto"
)

func main() {
	e := echo.New()
	e.GET("/hello", func(c echo.Context) error {
		hello := &pb.Hello{Hello: "Hello, World!"}
		data, _ := proto.Marshal(hello)
		return c.Blob(http.StatusOK, "application/protobuf", data)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
