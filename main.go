package main

import (
	"campaing/internal/domain/campaign"

	"github.com/go-playground/validator/v10"
)

func main() {
	contact := []campaign.Contact{{Email: ""}, {Email: ""}}
	campaign := campaign.Campaign{Contact: contact}
	validate := validator.New()
	err := validate.Struct(campaign)

	if err == nil {
		println("Nenhum erro")
	} else {
		validationErrors := err.(validator.ValidationErrors)
		for _, v := range validationErrors {
			switch v.Tag() {
			case "required":
				println(v.StructField() + " is required with min " + v.Param())
			case "min":
				println(v.StructField() + " is required with min")
			case "max":
				println(v.StructField() + " is required with max")
			case "email":
				println(v.StructField() + " is invalid!")
			}
		}
	}
}
