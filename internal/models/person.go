package models

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type Person struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Patronymic  string `json:"patronymic"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

type PersonCreateInput struct {
	FirstName  string `json:"first_name" validate:"required,gte=2"`
	LastName   string `json:"last_name"  validate:"required,gte=2"`
	Patronymic string `json:"patronymic"`
}

func (s *PersonCreateInput) Validate() error {
	return validate.Struct(s)
}

type PersonUpdateInput struct {
	FirstName   string `json:"first_name" validate:"gte=2"`
	LastName    string `json:"last_name"  validate:"gte=2"`
	Patronymic  string `json:"patronymic"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"      validate:"gte=2"`
	Nationality string `json:"nationality" validate:"gte=2"`
}

func (s *PersonUpdateInput) Validate() error {
	return validate.Struct(s)
}

func init() {
	validate = validator.New()
}
