package repository

import entity "github.com/edenriquez/load-balancer-proxy-go/pkg/entity"

//Repository repository interface
type Repository interface {
	Find(id string) (*entity.Proxy, error)
}
