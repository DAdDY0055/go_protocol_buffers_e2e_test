package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	pb "github.com/DAdDY0055/go_protocol_buffers_e2e_test/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestHelloHandler(t *testing.T) {
	// 想定する値を定義
	want := &pb.Hello{
		Hello: "Hello, World!",
	}

	router := NewRouter()

	req := httptest.NewRequest("GET", "/hello", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	res := &pb.Hello{}
	proto.Unmarshal(rec.Body.Bytes(), res)
	// assert.Equal(t, want, res) そのまま比較すると一致しない
	assert.Equal(t, want.Hello, res.Hello)

	fmt.Println("res", res) // 味気ないので出力
}
