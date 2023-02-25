package models

type Product struct {
	Id			int64		`gorm:"primaryKey" json:"id"`
	Name		string		`gorm:"type:varchar(75)" json:"name" binding:"required,min=3,max=75"`
	Merchant	string		`gorm:"type:varchar(75)" json:"merchant" binding:"required,min=3,max=75"`
	Desc		string		`gorm:"type:varchar(100)" json:"desc" binding:"required,min=3,max=100"`
	Stock		int			`gorm:"type:integer" json:"stock" binding:"required,number,gte=0"`
	Price		int			`gorm:"type:integer" json:"price" binding:"required,number,gte=1000"`
}