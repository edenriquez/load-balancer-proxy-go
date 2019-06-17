package proxy

import "github.com/edenriquez/load-balancer-proxy-go/pkg/entity"

//Reader interface
type Reader interface {
	Find(service string) (*entity.Proxy, error)
	Migrate()
}

//Repository repository interface
type Repository interface {
	Reader
}

//UseCase use case interface
type UseCase interface {
	Reader
}
