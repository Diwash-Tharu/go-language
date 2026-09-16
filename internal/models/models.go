// package models
// import (

// 	"fmt"
// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
	
// )

// type DBModel struct {
// 	Order OrderModel
// 	User  UserModel
// 	DB    *gorm.DB
// }


// //  connection to the database
// func InitDB(dataSourceName string) (*DBModel, error) {
// 	db, err := gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{})

// 	if err != nil {
// 		return nil, fmt.Errorf("failed to connect database: %v", err)
// 	}
// 	//  thes is migrate 
// 	err = db.AutoMigrate(&Order{}, &OrderItem{}, &User{})
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to migrate database: %v", err)
// 	}

// 	dbModel :=&DBModel{
// 		DB: db,
// 		OrderModel: OrderModel{
// 			BD: db,
// 		},
// 		User: UserModel{
// 			BD: db,
// 		},
		
// 	}


// 	return &dbModel, nil
// }




package models

import (
	"fmt"

	// "gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type DBModel struct {
	Order OrderModel
	User  UserModel
	DB    *gorm.DB
}

func InitDB(dataSourceName string) (*DBModel, error) {
	db, err := gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	err = db.AutoMigrate(&Order{}, &OrderItem{}, &User{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database %v", err)
	}

	dbModel := &DBModel{
		DB:    db,
		Order: OrderModel{DB: db},
		User:  UserModel{DB: db},
	}

	return dbModel, nil
}