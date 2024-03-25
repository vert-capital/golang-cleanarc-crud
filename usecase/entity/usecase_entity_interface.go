package usecase_

import "app/entity"

//go:generate mockgen -destination=../../mocks/mock_usecase_repository_/usecase/{pCase %>
type IRepository{{entityUpCase}} interface {
	GetFromID(id int) (*entity.Entity{{entityUpCase}}, error)
	GetAll(searchParams entity.SearchEntity{{entityUpCase}}Params) (response []entity.Entity{{entityUpCase}}, totalRegisters int64, err error)
	Create(*entity.Entity{{entityUpCase}}) error
	Update(*entity.Entity{{entityUpCase}}) error
	Delete(id int) error
}

//go:generate mockgen -destination=../../mocks/mock_usecase_/usecase/{se %>
type IUsecase{{entityUpCase}} interface {
	Get(id int) (*entity.Entity{{entityUpCase}}, error)
	GetAll(searchParams entity.SearchEntity{{entityUpCase}}Params) (response []entity.Entity{{entityUpCase}}, totalRegisters int64, err error)
	Create(*entity.Entity{{entityUpCase}}) error
	Update(*entity.Entity{{entityUpCase}}) error
	Delete(id int) error
}
