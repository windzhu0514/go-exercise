package server

import request "go-exercise/test/common"

type server struct {
}

func NewServer() *server {
	return &server{}
}

func (s *server) ServeRequest(r *request.Request) {

}
