package repository

import (
	models "go-postgres/model"
)

type UserRepo interface {
	Select()
	Create(*[]models.User) 
	Delete()
	Update()
}
