package main

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slices"
	"github.com/gin-gonic/gin/binding"
	"pizza-tracker-go/internal/models"
	
)

//  creatinf the custom validator

func RegisterCustomValidators(){
	//  checking the order status and if the order status is not in the order status slice then it will return an error
	//  here we are creating if condion in the cursuary 
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("valid_pizza_type", createSliceValidator(models.PizzaType))
		v.RegisterValidation("valid_pizza_size", createSliceValidator(models.PizzaSize))
		
	}
}

func createSliceValidator(allowedValues [] string) validator.Func {
	return func(fl validator.FieldLevel) bool {
		return slices.Contains(allowedValues,fl.Field().String())
	}
}

