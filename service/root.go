package service

import "net/http"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// TODO
}
