package utils

import (
	"fmt"
	"project-e-commerces/configs"
	"project-e-commerces/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {

	connectionString :=
		fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
			config.Database.Username,
			config.Database.Password,
			config.Database.Address,
			config.Database.DB_Port,
			config.Database.Name,
		)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	InitialMigration(db)
	return db
}

func InitialMigration(db *gorm.DB) {

	db.Migrator().DropTable(&entities.Detail_transaction{})
	db.Migrator().DropTable(&entities.Transaction{})
	db.Migrator().DropTable(&entities.Detail_cart{})
	db.Migrator().DropTable(&entities.Cart{})
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().DropTable(&entities.Stock{})
	db.Migrator().DropTable(&entities.Product{})
	db.Migrator().DropTable(&entities.Category{})

	db.AutoMigrate(entities.Category{})
	db.AutoMigrate(entities.Product{})
	db.AutoMigrate(entities.Stock{})

	db.AutoMigrate(entities.User{})
	db.AutoMigrate(entities.Cart{})
	db.AutoMigrate(entities.Detail_cart{})
	db.AutoMigrate(entities.Transaction{})
	db.AutoMigrate(entities.Detail_transaction{})

}
