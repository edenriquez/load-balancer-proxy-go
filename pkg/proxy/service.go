package proxy

import (
	"github.com/edenriquez/load-balancer-proxy-go/pkg/entity"
)

//Service service interface
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//Find layer for find action
func (s *Service) Find(service string) (*entity.Proxy, error) {
	return s.repo.Find(service)
}

//Migrate layer for find action
func (s *Service) Migrate() {
	s.repo.Migrate()
}
