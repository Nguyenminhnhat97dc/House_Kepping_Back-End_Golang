package main

import (
	"BE_Golang/BE_Golang/create_database/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	}
	// Create Table
	db.AutoMigrate(models.Provider{})
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Services{})
	db.AutoMigrate(models.ServicesOfProvider{})
	db.AutoMigrate(models.Customer{})
	db.AutoMigrate(models.RequirementsCustomer{})
	db.AutoMigrate(models.ToDoList{})
	// insert Provider
	provider := []models.Provider{
		{Name: "Nguyễn Minh Nhật", Address: "Cam Ranh - Khánh Hòa", CCCD: "037153000257", Phone: "0975661107", Introduce: "abc"},
		{Name: "Hà Anh Tuấn", Address: "Đức Trọng - Lâm Đồng", CCCD: "065842695782", Phone: "0867676745", Introduce: "abc"},
		{Name: "Trần Minh Tuấn", Address: "Đức Trọng - Lâm Đồng", CCCD: "052478632591", Phone: "0961055050", Introduce: "abc"},
		{Name: "Võ Tiến Hải", Address: "Cam Ranh - Khánh Hòa", CCCD: "021589635874", Phone: "0971150202", Introduce: "abc"},
		{Name: "Đào Quốc Sang", Address: "Đức Trọng - Lâm Đồng", CCCD: "096358614752", Phone: "0961165050", Introduce: "abc"},
		{Name: "Trần Minh Tiến", Address: "Cam Ranh - Khánh Hòa", CCCD: "052651472369", Phone: "0961980303", Introduce: "abc"},
		{Name: "Nguyễn Hữu Lộc", Address: "Đức Trọng - Lâm Đồng", CCCD: "036541258520", Phone: "0961853030", Introduce: "abc"},
		{Name: "Võ Duy Tôn", Address: "Cam Ranh - Khánh Hòa", CCCD: "036589741252", Phone: "0961232200", Introduce: "abc"},
		{Name: "Lê Trần Khánh Vy", Address: "Đức Trọng - Lâm Đồng", CCCD: "041259875602", Phone: "0971170011", Introduce: "abc"},
		{Name: "Hoàng Minh Vương", Address: "Cam Ranh - Khánh Hòa", CCCD: "065842675105", Phone: "0971230077", Introduce: "abc"},
		{Name: "Ưng Vi Vương", Address: "Đức Trọng - Lâm Đồng", CCCD: "0751423654892", Phone: "0961854040", Introduce: "abc"},
	}
	db.Create(&provider)
	//insert User
	user := []models.User{
		{UserName: "nguyenminhnhat", Password: "123456", ProviderID: 1},
		{UserName: "haanhtuan", Password: "123456", ProviderID: 2},
		{UserName: "tranminhtuan", Password: "123456", ProviderID: 3},
		{UserName: "votienhai", Password: "123456", ProviderID: 4},
		{UserName: "daoquocsang", Password: "123456", ProviderID: 5},
		{UserName: "tranminhtien", Password: "123456", ProviderID: 6},
		{UserName: "nguyenhuuloc", Password: "123456", ProviderID: 7},
		{UserName: "voduyton", Password: "123456", ProviderID: 8},
		{UserName: "letrankhanhvy", Password: "123456", ProviderID: 9},
		{UserName: "hoangminhvuong", Password: "123456", ProviderID: 10},
		{UserName: "ungvivuong", Password: "123456", ProviderID: 11},
	}
	db.Create(&user)
	// inser Services
	services := []models.Services{
		{NameServices: "Dọn nhà", Image: "don-nha-theo-gio", Introduce: "link_gioithieu"},
		{NameServices: "Quét nhà", Image: "abc.jpg", Introduce: "link_gioithieu"},
		{NameServices: "Lau nhà", Image: "lau_nha.jpg", Introduce: "link_gioithieu"},
		{NameServices: "Vệ sinh nhà", Image: "ve_sinh_nha", Introduce: "link_gioithieu"},
		{NameServices: "Giặt quần áo", Image: "link_image", Introduce: "link_gioithieu"},
		{NameServices: "Chăm sóc cây cảnh", Image: "link_image", Introduce: "link_gioithieu"},
	}
	db.Create(&services)
	//insert ServicesOfProvider
	servicesOfProvider := []models.ServicesOfProvider{
		{ServicesId: 1, ProviderID: 1, Price: 500.000},
		{ServicesId: 2, ProviderID: 1, Price: 500.000},
		{ServicesId: 3, ProviderID: 1, Price: 500.000},
		{ServicesId: 4, ProviderID: 1, Price: 500.000},
		{ServicesId: 5, ProviderID: 1, Price: 500.000},
		{ServicesId: 6, ProviderID: 1, Price: 500.000},

		{ServicesId: 1, ProviderID: 2, Price: 400.000},
		{ServicesId: 2, ProviderID: 2, Price: 400.000},
		{ServicesId: 3, ProviderID: 2, Price: 500.000},
		{ServicesId: 4, ProviderID: 2, Price: 500.000},
		{ServicesId: 5, ProviderID: 2, Price: 400.000},
		{ServicesId: 6, ProviderID: 2, Price: 500.000},

		{ServicesId: 1, ProviderID: 3, Price: 500.000},
		{ServicesId: 2, ProviderID: 3, Price: 500.000},
		{ServicesId: 3, ProviderID: 3, Price: 500.000},
		{ServicesId: 4, ProviderID: 3, Price: 500.000},
		{ServicesId: 5, ProviderID: 3, Price: 500.000},
		{ServicesId: 6, ProviderID: 3, Price: 500.000},

		{ServicesId: 1, ProviderID: 4, Price: 500.000},
		{ServicesId: 2, ProviderID: 4, Price: 500.000},
		{ServicesId: 3, ProviderID: 4, Price: 500.000},
		{ServicesId: 4, ProviderID: 4, Price: 500.000},
		{ServicesId: 5, ProviderID: 4, Price: 500.000},
		{ServicesId: 6, ProviderID: 4, Price: 500.000},

		{ServicesId: 1, ProviderID: 5, Price: 500.000},
		{ServicesId: 2, ProviderID: 5, Price: 500.000},
		{ServicesId: 3, ProviderID: 5, Price: 500.000},
		{ServicesId: 4, ProviderID: 5, Price: 500.000},
		{ServicesId: 5, ProviderID: 5, Price: 500.000},
		{ServicesId: 6, ProviderID: 5, Price: 500.000},

		{ServicesId: 1, ProviderID: 6, Price: 500.000},
		{ServicesId: 2, ProviderID: 6, Price: 500.000},
		{ServicesId: 3, ProviderID: 6, Price: 500.000},
		{ServicesId: 4, ProviderID: 6, Price: 500.000},
		{ServicesId: 5, ProviderID: 6, Price: 500.000},
		{ServicesId: 6, ProviderID: 6, Price: 500.000},

		{ServicesId: 1, ProviderID: 7, Price: 500.000},
		{ServicesId: 2, ProviderID: 7, Price: 500.000},
		{ServicesId: 3, ProviderID: 7, Price: 500.000},
		{ServicesId: 4, ProviderID: 7, Price: 500.000},
		{ServicesId: 5, ProviderID: 7, Price: 500.000},
		{ServicesId: 6, ProviderID: 7, Price: 500.000},
	}
	db.Create((&servicesOfProvider))

	//insert customer
	customer := []models.Customer{
		{NameCustomer: "Mai Hoàng Hương", AddressCustomer: "Cam Ranh - Khánh Hòa", PhoneCustomer: "0867676745"},
		{NameCustomer: "Nguyễn Hưu Nam", AddressCustomer: "Cam Ranh - Khánh Hòa", PhoneCustomer: "0961980303"},
		{NameCustomer: "Trần Hoành Phong", AddressCustomer: "Cam Ranh - Khánh Hòa", PhoneCustomer: "0961055050"},
		{NameCustomer: "Mai Nhật Tuấn", AddressCustomer: "Cam Ranh - Khánh Hòa", PhoneCustomer: "0971230077"},
		{NameCustomer: "Phong Bạch Vũ", AddressCustomer: "Cam Ranh - Khánh Hòa", PhoneCustomer: "0976359836"},
	}
	db.Create(&customer)

	//inser requirement
	requirement := []models.RequirementsCustomer{
		{
			CustomerID:   1,
			NameServices: "Quét nhà, Lau nhà, Dọn nhà, Vệ sinh nhà, Giặt quần áo, Chăm sóc cây cảnh",
			DayStart:     "11/5/2022, 12/5/2022",
			TimeStart:    "8:00",
		},
		{
			CustomerID:   2,
			NameServices: "Quét nhà, Lau nhà, Dọn nhà, Vệ sinh nhà, Giặt quần áo, Chăm sóc cây cảnh",
			DayStart:     "12/5/2022, 13/5/2022",
			TimeStart:    "8:00",
		},
		{
			CustomerID:   3,
			NameServices: "Quét nhà, Lau nhà, Dọn nhà, Vệ sinh nhà, Giặt quần áo, Chăm sóc cây cảnh",
			DayStart:     "15/5/2022, 15/5/2022",
			TimeStart:    "8:00",
		},
		{
			CustomerID:   5,
			NameServices: "Quét nhà, Lau nhà, Dọn nhà, Vệ sinh nhà, Giặt quần áo, Chăm sóc cây cảnh",
			DayStart:     "16/5/2022, 17/5/2022",
			TimeStart:    "8:00",
		},
		{
			CustomerID:   5,
			NameServices: "Quét nhà, Lau nhà, Dọn nhà, Vệ sinh nhà, Giặt quần áo, Chăm sóc cây cảnh",
			DayStart:     "18/5/2022, 19/5/2022",
			TimeStart:    "8:00",
		},
	}
	db.Create(&requirement)

	//insert TodoList
	todoList := []models.ToDoList{
		{RequirementsCustomerID: 1, ProviderID: 1},
		{RequirementsCustomerID: 2, ProviderID: 1},
		{RequirementsCustomerID: 3, ProviderID: 1},
		{RequirementsCustomerID: 4, ProviderID: 1},
		{RequirementsCustomerID: 5, ProviderID: 1},
	}
	db.Create((&todoList))
}
