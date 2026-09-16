package models


import (
	
	
		"time"
		"gorm.io/gorm"
		"github.com/teris-io/shortid"


		
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


	PizzaSize =[] string{
		"Small",
		"Medium",
		"Large",
		"Extra Large",}
	)


	//  creating the struct and Struct are the same as the object in the java

	type OrderModel struct {
		DB *gorm.DB

	}

	type Order struct {
		ID  string `gorm:"primaryKey;size:14" json:"id"`
		Status string `gorm:"not null" json:"status"`
		CustomerName string `gorm:"not null" json:"customerName"`
		Phone string `gorm:"not null" json:"phone"`
		Address string `gorm:"not null" json:"address"`
		Items []OrderItem `gorm:"foreignKey:OrderID" json:"pizzas"`
		CreatedAt time.Time `json:"createdAt"`

	
	}

	type OrderItem struct {
		ID string `gorm:"primaryKey;size:14" json:"id"`
		OrderID string `gorm:"index;size:14;not nill" json:"orderId"`
		Size string `gorm:"not null" json:"size"`
		Pizza string `gorm:"not null" json:"pizza"`
		Instruction string  `json:"instruction"`


	}

	//  checking the order status and if the order status is not in the order status slice then it will return an error
func (o *Order) BeforeCreated(tx *gorm.DB) error{
	if o.ID == "" {
		o.ID = shortid.MustGenerate()
	}
	return nil
}

func (oi *OrderItem) BeforeCreated(tx *gorm.DB) error{
	if oi.ID == "" {
		oi.ID = shortid.MustGenerate()
	}
	return nil
}
func (o *OrderModel) CreateOrder(order *Order) error {
	//  checking the order status and if the order status is not in the order status slice then it will return an error
	
	return o.DB.Create(order).Error
}

func (o *OrderModel) GetOrder (id string)(*Order, error) {
	//  checking the order status and if the order status is not in the order status slice then it will return an error
	
	var order Order
	err :=o.DB.Preload("Items").First(&order, "id = ?",id).Error
	return &order, err

	
}



