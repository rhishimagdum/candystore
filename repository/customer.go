package repository

import "github.com/rhishimagdum/candystore/entity"

type Customer interface {
	GetAll() []entity.CustomerRecord
}
