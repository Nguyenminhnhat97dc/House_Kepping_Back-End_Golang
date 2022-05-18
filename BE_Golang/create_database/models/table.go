package models

import "gorm.io/gorm"

type Provider struct {
	gorm.Model
	Name      string `gorm:"type: nvarchar(50); not null"`
	Address   string `gorm:"type: nvarchar(100); not null"`
	Phone     string `gorm:"type: nvarchar(11); not null"`
	Introduce string `gorm:"type: varchar(10); not null"`
}

type User struct {
	gorm.Model
	UserName   string   `gorm:"type: varchar(20); not null"`
	Password   string   `gorm:"type: varchar(20); not null"`
	Provider   Provider `gorm:"references:id"`
	ProviderID uint
}

type Services struct {
	gorm.Model
	NameServices string `gorm:"type: nvarchar(20); not null"`
	Image        string
	Introduce    string `gorm:"type: varchar(10); not null"`
}

type ServicesOfProvider struct {
	gorm.Model
	Services   Services `gorm:"references:id"`
	ServicesId uint
	Provider   Provider `gorm:"references:id"`
	ProviderID uint
	Price      int64 `gorm:"type: int; not null"`
}

type Customer struct {
	gorm.Model
	NameCustomer    string `gorm:"type:nvarchar(50);COLLATE utf8_unicode_ci"`
	AddressCustomer string `gorm:"type:nvarchar(100);COLLATE utf8_unicode_ci"`
	PhoneCustomer   string `gorm:"type: varchar(11); not null"`
}

type RequirementsCustomer struct {
	gorm.Model
	Customer     Customer `gorm:"foreignKey:CustomerID"`
	CustomerID   uint
	NameServices string `gorm:"type: nvarchar(200); not null"`
	DayStart     string `gorm:"type:varchar(50)"`
	TimeStart    string `gorm:"type:varchar(10)"`
	Status       int    `gorm:"type: boolean; default 0"`
}

type ToDoList struct {
	gorm.Model
	RequirementsCustomer   RequirementsCustomer `gorm:"foreignKey:RequirementsCustomerID"`
	RequirementsCustomerID uint
	Provider               Provider `gorm:"foreignKey:ProviderID"`
	ProviderID             uint
	DayEnd                 string `gorm:"type: varchar(10); default null"`
	Status                 int    `gorm:"type: boolean; default 0"`
}
