package productsentity

type Product_Categories struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(50)" json:"name"`
}

type Products struct {
	ID          uint               `gorm:"primaryKey" json:"id" `
	Name        string             `gorm:"type:varchar(100)" json:"name"`
	Description string             `json:"description"`
	Category_ID uint               `json:"category_id"`                                   //foreign key column
	Category    Product_Categories `gorm:"foreignKey:Category_ID;references:id" json:"-"` //foreign key value
}
