package models


import (
	
		"fmt"
		"time"

		"gorm.io/gorm"
		

		
	)

var (
	//  thes is the string slice 
	
	OrderStatus =[]string {
		"Order placed",
		"Preparing",
		"Baking",
		"Quality check",
		"Out for delivery",
		"Delivered",
		"Cancelled"}

	PizzaType =[] string{
		"Margherita",
		"Pepperoni",
		"Vegetarian",
		"Meat Lovers",
		"BBQ Chicken",
		"Hawaiian",
		"Supreme",
		"Four Cheese",
		"Caprese",
		"Mexican",
		"Buffalo Chicken",
		"Spinach and Feta",
		"Seafood",}


	pizzaSize =[] string{
		"Small",
		"Medium",
		"Large",
		"Extra Large",}
	)


	//  creating the struct and Struct are the same as the object in the java

	type OrderModel struct {
		BD *gorm.DB

	}

	type Order struct {
		ID  string `gorm:"primaryKey;size:14" json:"id"`

	}

