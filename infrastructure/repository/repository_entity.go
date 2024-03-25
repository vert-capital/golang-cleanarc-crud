package repository

import (
	"app/entity"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repository{{entityUpCase}} struct {
	DB *gorm.DB
}

func New{{entityUpCase}}Postgres(DB *gorm.DB) *Repository{{entityUpCase}} {
	return &Repository{{entityUpCase}}{DB: DB}
}

func (r *Repository{{entityUpCase}}) GetFromID(id int) (ityUpCase %>, err error) {
	r.DB.First(&

	return
}

func (r *Repository{{entityUpCase}}) GetAll(searchParams entity.SearchEntity{{entityUpCase}}Params) (response []entity.Entity{{entityUpCase}}, totalRegisters int64, err error) {
	offset := (searchParams.Page) * searchParams.PageSize

	qry := r.DB.Model(entity.Entity{{entityUpCase}}{})

	if gin.IsDebugging() {
		qry = qry.Debug()
	}

	if searchParams.Q != "" {
		qry = qry.Where("name LIKE ?", "%"+searchParams.Q+"%")
	}

	if searchParams.CreatedAt != "" {
		dates := strings.Split(searchParams.CreatedAt, ",")
		if len(dates) == 2 {
			_, err1 := time.Parse("2006-01-02", dates[0])
			_, err2 := time.Parse("2006-01-02", dates[1])
			if err1 == nil && err2 == nil {
				qry = qry.Where("created_at BETWEEN ? AND ?", dates[0], dates[1])
			}
		}
	}

	err = qry.Count(&totalRegisters).Error

	if err != nil {
		return nil, 0, err
	}

	if searchParams.OrderBy == "" {
		searchParams.OrderBy = "id"
	}

	if searchParams.SortOrder == "" {
		searchParams.SortOrder = "asc"
	}

	qry = qry.Order(searchParams.OrderBy + " " + searchParams.SortOrder).
		Offset(offset).
		Limit(searchParams.PageSize)

	err = qry.Find(&response).Error
	if err != nil {
		return nil, 0, err
	}

	return response, totalRegisters, nil
}

func (r *Repository{{entityUpCase}}) Create(ityUpCase %>) (err error) {
	err = r.DB.Create(&

	return err
}

func (r *Repository{{entityUpCase}}) Update(ityUpCase %>) (err error) {
	_, err = r.GetFromID(int(

	if err != nil {
		return err
	}

	err = r.DB.Save(&

	return err
}

func (r *Repository{{entityUpCase}}) Delete(id int) (err error) {
	err = r.DB.Delete(&entity.Entity{{entityUpCase}}{}, id).Error

	return err
}
