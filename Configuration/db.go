package configuration

import (
	"log"
	productsentity "visiku-restapi/Entity/products"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConn() {

	// user: golang, password: golang
	user := "golang:golang@/visiku-db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(user), &gorm.Config{})

	if err != nil {
		log.Println("error db")
	}

	// Unit Testing
	var count, counts int64
	db.Model(&productsentity.Product_Categories{}).Count(&count)
	db.Model(&productsentity.Product_Categories{}).Count(&counts)

	if count == 0 && counts == 0 {
		ctg := []productsentity.Product_Categories{
			{
				Name: "Electronic",
			},

			{
				Name: "Knowledge",
			},
		}

		prd := []productsentity.Products{
			{
				Name:        "Komputer",
				Description: "Ini komputer limited",
				Category_ID: 1,
			},

			{
				Name:        "Majalah Bobo",
				Description: "Majalah tahun 1950-an",
				Category_ID: 2,
			},

			{
				Name:        "Laptop",
				Description: "Ini laptop limited",
				Category_ID: 1,
			},

			{
				Name:        "Atlas",
				Description: "Gambar peta bumi secara keseluruhan",
				Category_ID: 2,
			},

			{
				Name:        "RAM-8GB",
				Description: "Tempat penyimpan sederhana",
				Category_ID: 1,
			},

			{
				Name:        "Printer",
				Description: "untuk membantu hal percetakan",
				Category_ID: 1,
			},

			{
				Name:        "Spidol",
				Description: "Menulis di papan tulis dan memiliki tinta yang tebal",
				Category_ID: 2,
			},

			{
				Name:        "Scanner",
				Description: "Melakukan scan terhadap dokumen tertentu",
				Category_ID: 1,
			},
		}

		db.Create(ctg)
		db.Create(prd)
	}

	db.AutoMigrate(&productsentity.Product_Categories{},
		&productsentity.Products{})

	DB = db
	log.Println("DB Operating")
}
