package controllers

import (
	"API_House_Kepping/BE_Golang/BE_Golang/create_database/models"
	connectdatabase "API_House_Kepping/BE_Golang/Be_Golang/connectDatabase"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/* var host = "ec2-34-230-153-41.compute-1.amazonaws.com"
var user = "gqkzhktjpbsnix"
var password = "3cc9ee2fd230e1696ee764c83ef829474e27577be64388c849031eb618a637ab"
var dbname = "d2u77vk80vvs75"
var dsn = "host=" + host + "user=" + user + "password=" + password + "dbname=" + dbname + "port=5432 sslmode=disable TimeZone=Asia/Shanghai"
*/
//var dsn = "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

var dsn = "sql6496052:JVUfiJ9mBJ@tcp(sql6.freemysqlhosting.net:3306)/sql6496052?charset=utf8mb4&parseTime=True&loc=Local"
var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// select * from Provider
/* func FindProvider(c *gin.Context) {
	var provider []models.Provider
	if err := database.DBConn().Find(&provider).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"result": provider})
		return
	}
} */

// select * from Services
func FindServices(c *gin.Context) {
	dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	}
	var services []models.Services
	if err := dbConnect.Find(&services).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
		sqlDB, err := dbConnect.DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"result": services})
		sqlDB, err := dbConnect.DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()
		return
	}
}

// SELECT * FROM users LIMIT 4;
func LimitServices(c *gin.Context) {
	dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	}
	var services []models.Services
	count := c.Param("count")
	number, _ := strconv.Atoi(count)
	if err := dbConnect.Limit(number).Find(&services).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
		sqlDB, err := dbConnect.DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"result": services})
		sqlDB, err := dbConnect.DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()
		return
	}

}

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

//
func AddRequirementCustomer(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	type CheckCustomer struct {
		Name         string
		Address      string
		Phone        string
		NameServices string
		DayStart     string
		TimeStart    string
	}
	var checkCustomer CheckCustomer
	var Customer models.Customer
	var Requirement models.RequirementsCustomer

	// convert string
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newStr := buf.String()
	// convert Json
	res, err := PrettyString(newStr)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(res), &checkCustomer)
	abc := checkCustomer.Name
	if err := connectdatabase.DBConn().Where("name_customer = ? AND  address_customer = ?", abc, checkCustomer.Address).First(&Customer).Error; err != nil {
		NewCustomer := models.Customer{
			NameCustomer:    checkCustomer.Name,
			AddressCustomer: checkCustomer.Address,
			PhoneCustomer:   checkCustomer.Phone,
		}
		fmt.Println("THEM KH", NewCustomer)
		if err := connectdatabase.DBConn().Create(&NewCustomer).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"result": "Không  insert Khách Hàng được"})
			sqlDB, err := connectdatabase.DBConn().DB()
			if err != nil {
				log.Fatalln(err)
			}
			defer sqlDB.Close()
			return
		} else {
			connectdatabase.DBConn().Where("name_customer = ? AND  address_customer = ?", abc, checkCustomer.Address).First(&Customer)
			NewRequirement := models.RequirementsCustomer{
				CustomerID:   Customer.ID,
				NameServices: checkCustomer.NameServices,
				DayStart:     checkCustomer.DayStart,
				TimeStart:    checkCustomer.TimeStart,
			}
			fmt.Println(">>requirement", NewRequirement)
			if err := connectdatabase.DBConn().Where("customer_id = ? AND day_start = ? AND time_start = ? ", NewRequirement.CustomerID, NewRequirement.DayStart, NewRequirement.TimeStart).First(&Requirement).Error; err != nil {
				if err := connectdatabase.DBConn().Create(&NewRequirement).Error; err != nil {
					c.JSON(http.StatusOK, gin.H{"result": "create - không Insert yêu cầu khách hàng insert được"})
					sqlDB, err := connectdatabase.DBConn().DB()
					if err != nil {
						log.Fatalln(err)
					}
					defer sqlDB.Close()
					return
				} else {
					c.JSON(http.StatusOK, gin.H{"result": "Insert yêu cầu khách hàng thành công"})
					sqlDB, err := connectdatabase.DBConn().DB()
					if err != nil {
						log.Fatalln(err)
					}
					defer sqlDB.Close()
					return
				}
			} else {
				if err := connectdatabase.DBConn().Model(&NewRequirement).Where("customer_id = ? AND day_start = ? AND time_start = ? ", NewRequirement.CustomerID, NewRequirement.DayStart, NewRequirement.TimeStart).Update("name_services", NewRequirement.NameServices).Error; err != nil {
					c.JSON(http.StatusOK, gin.H{"result": "Không Update được"})
					sqlDB, err := connectdatabase.DBConn().DB()
					if err != nil {
						log.Fatalln(err)
					}
					defer sqlDB.Close()
					return
				} else {
					c.JSON(http.StatusOK, gin.H{"result": "Update thành công"})
					sqlDB, err := connectdatabase.DBConn().DB()
					if err != nil {
						log.Fatalln(err)
					}
					defer sqlDB.Close()
					return
				}
			}
		}
	} else {
		NewRequirement := models.RequirementsCustomer{
			CustomerID:   Customer.ID,
			NameServices: checkCustomer.NameServices,
			DayStart:     checkCustomer.DayStart,
			TimeStart:    checkCustomer.TimeStart,
		}
		if err := connectdatabase.DBConn().Where("customer_id = ? AND day_start = ? AND time_start = ? ", NewRequirement.CustomerID, NewRequirement.DayStart, NewRequirement.TimeStart).First(&Requirement).Error; err != nil {
			if err := connectdatabase.DBConn().Create(&NewRequirement).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"result": "không Insert yêu cầu khách hàng insert được"})
				sqlDB, err := connectdatabase.DBConn().DB()
				if err != nil {
					log.Fatalln(err)
				}
				defer sqlDB.Close()
				return
			} else {
				c.JSON(http.StatusOK, gin.H{"result": "Insert yêu cầu khách hàng thành công"})
				sqlDB, err := connectdatabase.DBConn().DB()
				if err != nil {
					log.Fatalln(err)
				}
				defer sqlDB.Close()
				return
			}
		} else {
			if err := connectdatabase.DBConn().Model(&NewRequirement).Where("customer_id = ? AND day_start = ? AND time_start = ? ", NewRequirement.CustomerID, NewRequirement.DayStart, NewRequirement.TimeStart).Update("name_services", NewRequirement.NameServices).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"result": "Không Update được"})
				sqlDB, err := connectdatabase.DBConn().DB()
				if err != nil {
					log.Fatalln(err)
				}
				defer sqlDB.Close()
				return
			} else {
				c.JSON(http.StatusOK, gin.H{"result": "Update thành công"})
				sqlDB, err := connectdatabase.DBConn().DB()
				if err != nil {
					log.Fatalln(err)
				}
				defer sqlDB.Close()
				return
			}
		}
	}
}

func ServiceProvider(c *gin.Context) {
	dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	}
	//Upgrade get request to webSocket protocol
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
		log.Fatal(err)
	}
	defer ws.Close()
	var data struct {
		Id string `json:"Id"`
	}

	err = ws.ReadJSON(&data)
	if err != nil {
		log.Println("error read json")
		log.Fatal(err)
	}
	type GetServices struct {
		ServicesId   uint
		NameServices string
		Price        string
		ProviderId   uint
	}
	var getServices []GetServices
	for {

		dbConnect.Raw("SELECT services_of_providers.services_id,services.name_services, services_of_providers.price, services_of_providers.provider_id FROM"+
			" `services_of_providers` LEFT JOIN services on services_of_providers.services_id = services.id"+
			" WHERE services_of_providers.provider_id = ?", data.Id).Scan(&getServices)

		err = ws.WriteJSON(getServices)
		if err != nil {
			log.Println("Lỗi ở đây nè error write json: " + err.Error())
		}

		time.Sleep(1 * time.Second)
	}

	/* type CheckProvider struct {
		Id string
	}
	type GetServices struct {
		ServicesId   uint
		NameServices string
		Price        string
		ProviderId   uint
	}
	var checkProvider CheckProvider
	var getServices []GetServices
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &checkProvider)
	if err := database.DBConn().Raw("SELECT services_of_providers.services_id,services.name_services, services_of_providers.price, services_of_providers.provider_id FROM"+
		" `services_of_providers` LEFT JOIN services on services_of_providers.services_id = services.id"+
		" WHERE services_of_providers.provider_id = ?", checkProvider.Id).Scan(&getServices).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "Không tìm thấy"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"result": getServices})
		return
	} */
}

func AddServiceProvider(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, err := dbConnect.DB()
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	type GetServicesOfProvider struct {
		ServicesId uint
		ProviderId uint
		Price      int64
	}
	var getServicesOfProvider GetServicesOfProvider
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &getServicesOfProvider)
	AddNewServiceProvider := models.ServicesOfProvider{
		ServicesId: getServicesOfProvider.ServicesId,
		ProviderID: getServicesOfProvider.ProviderId,
		Price:      getServicesOfProvider.Price,
	}
	if err := connectdatabase.DBConn().Create(&AddNewServiceProvider).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
		sqlDB, err := connectdatabase.DBConn().DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "true"})
		sqlDB, err := connectdatabase.DBConn().DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()
		return
	}
}

func RequirementsCustomer(c *gin.Context) {
	dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	}
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
		log.Fatal(err)
	}
	defer ws.Close()
	type InformationRequirementsCustomer struct {
		Id              uint
		NameServices    string
		DayStart        string
		TimeStart       string
		NameCustomer    string
		AddressCustomer string
		PhoneCustomer   string
	}
	var informationRequirementsCustomer []InformationRequirementsCustomer

	for {
		if err := dbConnect.Raw(
			"SELECT requirements_customers.id,requirements_customers.name_services,requirements_customers.day_start,requirements_customers.time_start,customers.name_customer,customers.address_customer,customers.phone_customer"+
				" FROM requirements_customers,customers WHERE requirements_customers.customer_id = customers.id and requirements_customers.status = ?", 0).Scan(&informationRequirementsCustomer).Error; err != nil {
			err = ws.WriteJSON("False")
			if err != nil {
				log.Println("error write json: " + err.Error())
			}
		} else {
			err = ws.WriteJSON(informationRequirementsCustomer)
			if err != nil {
				log.Println("error write json: " + err.Error())
			}
		}
		time.Sleep(500 * time.Millisecond)
	}

	/* type InformationRequirementsCustomer struct {
		Id              uint
		NameServices    string
		DayStart        string
		TimeStart       string
		NameCustomer    string
		AddressCustomer string
		PhoneCustomer   string
	}
	var informationRequirementsCustomer []InformationRequirementsCustomer
	type Pagination struct {
		PaginationStart uint
		PaginationEnd   uint
	}
	var pagination Pagination
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &pagination)
	if err := dbConnect.Raw(
		"SELECT requirements_customers.id,requirements_customers.name_services,requirements_customers.day_start,requirements_customers.time_start,customers.name_customer,customers.address_customer,customers.phone_customer"+
			" FROM requirements_customers,customers WHERE requirements_customers.customer_id = customers.id and requirements_customers.status = ? LIMIT ?,?", 0, pagination.PaginationStart, pagination.PaginationEnd).Scan(&informationRequirementsCustomer).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"result": informationRequirementsCustomer})
		return
	} */
}

func TodoList(c *gin.Context) {
	dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	}
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
		log.Fatal(err)
	}
	defer ws.Close()
	type CheckProvider struct {
		Id string
	}
	var checkProvider CheckProvider

	err = ws.ReadJSON(&checkProvider)

	if err != nil {
		log.Println("error read json")
		log.Fatal(err)
	}
	type TodoList struct {
		Id              uint
		NameServices    string
		Status          string
		DayStart        string
		TimeStart       string
		NameCustomer    string
		AddressCustomer string
		PhoneCustomer   string
	}
	var todoList []TodoList

	for {

		if err := dbConnect.Raw(
			"SELECT requirements_customers.id,requirements_customers.name_services,to_do_lists.status,requirements_customers.day_start,requirements_customers.time_start,"+
				" customers.name_customer,customers.address_customer,customers.phone_customer"+
				" FROM `to_do_lists`,requirements_customers,customers,providers WHERE to_do_lists.requirements_customer_id = requirements_customers.id and"+
				" requirements_customers.customer_id = customers.id and to_do_lists.provider_id = providers.id and to_do_lists.status = 0 and providers.id = ?", checkProvider.Id).Scan(&todoList).Error; err != nil {

			err = ws.WriteJSON("False")
			if err != nil {
				log.Println("error write json: " + err.Error())
			}
		} else {
			if len(todoList) > 0 {

				err = ws.WriteJSON(todoList)
				if err != nil {
					log.Println("error write json: " + err.Error())
				}
			} else {
				err = ws.WriteJSON("Bạn không có việc cần làm")
				if err != nil {
					log.Println("error write json: " + err.Error())
				}
			}
		}
		time.Sleep(500 * time.Millisecond)

	}

	/* type CheckProvider struct {
		Id     string
		Status int
		PaginationStart uint
		PaginationEnd   uint
	}
	var checkProvider CheckProvider
	type TodoList struct {
		Id              uint
		NameServices    string
		Status          string
		DayStart        string
		TimeStart       string
		DayEnd          string
		NameCustomer    string
		AddressCustomer string
		PhoneCustomer   string
	}

	var pagination Pagination
	var todoList []TodoList
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &checkProvider)
	if err := dbConnect.Raw(
		"SELECT to_do_lists.id,requirements_customers.name_services,to_do_lists.status,requirements_customers.day_start,requirements_customers.time_start,"+
			" customers.name_customer,customers.address_customer,customers.phone_customer"+
			" FROM `to_do_lists`,requirements_customers,customers,providers WHERE to_do_lists.requirements_customer_id = requirements_customers.id and"+
			" requirements_customers.customer_id = customers.id and to_do_lists.provider_id = providers.id and to_do_lists.status = ? and providers.id = ? LIMIT ?,?", checkProvider.Status, checkProvider.Id, checkProvider.PaginationStart, checkProvider.PaginationEnd).Scan(&todoList).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
		return
	} else {
		if len(todoList) > 0 {
			c.JSON(http.StatusOK, gin.H{"result": todoList})
		} else {
			if checkProvider.Status == 1 {
				c.JSON(http.StatusOK, gin.H{"result": []string{"Bạn không có lịch sử công việc"}})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{"result": []string{"Bạn không có việc cần làm"}})
				return
			}
		}
	} */
}

func Loggin(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, err := dbConnect.DB()
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	type CheckLoggin struct {
		User     string
		Password string
	}
	var checkLoggin CheckLoggin
	var informationLoggin models.User
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &checkLoggin)
	if err := connectdatabase.DBConn().Where("user_name = ? and password = ?", checkLoggin.User, checkLoggin.Password).First(&informationLoggin).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
		sqlDB, err := connectdatabase.DBConn().DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"result": informationLoggin})
		sqlDB, err := connectdatabase.DBConn().DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()
		return
	}
}

func FindProviderID(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, err := dbConnect.DB()
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	type CheckID struct {
		Id uint
	}
	var checkID CheckID
	var informationProvider models.Provider
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &checkID)
	if err := connectdatabase.DBConn().First(&informationProvider, "id=?", checkID.Id).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
		sqlDB, err := connectdatabase.DBConn().DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"result": informationProvider})
		sqlDB, err := connectdatabase.DBConn().DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()
		return
	}
}

func FindPriceOfServices(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, err := dbConnect.DB()
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	type CheckID struct {
		Id string
	}
	type Price struct {
		NameServices string
		Price        string
		Name         string
	}
	var checkID CheckID
	var price []Price
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &checkID)
	if err := connectdatabase.DBConn().Raw(
		"SELECT services.name_services, services_of_providers.price, providers.name from"+
			" services_of_providers,services,providers WHERE services_of_providers.services_id = services.id and"+
			" services_of_providers.provider_id = providers.id and providers.id = ?", checkID.Id).Scan(&price).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
		sqlDB, err := connectdatabase.DBConn().DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"result": price})
		sqlDB, err := connectdatabase.DBConn().DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()
		return
	}
}

func AddPrice(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, err := dbConnect.DB()
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	type CheckInformation struct {
		Id           uint
		NameServices string
		Price        int64
	}
	var services models.Services
	var checkInformation CheckInformation
	var servicesOfProvider models.ServicesOfProvider
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &checkInformation)
	if err := connectdatabase.DBConn().First(&services, "name_services=?", &checkInformation.NameServices).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
		return
	} else {
		fmt.Println(">>>>>>", services.ID, checkInformation.Price)
		if err := connectdatabase.DBConn().Where("services_id=? and provider_id=?", services.ID, checkInformation.Id).First(&servicesOfProvider).Error; err != nil {
			NewServicesOfProvider := models.ServicesOfProvider{
				ServicesId: services.ID,
				ProviderID: checkInformation.Id,
				Price:      checkInformation.Price,
			}
			if err := connectdatabase.DBConn().Create(&NewServicesOfProvider).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"result": "False"})
				sqlDB, err := connectdatabase.DBConn().DB()
				if err != nil {
					log.Fatalln(err)
				}
				defer sqlDB.Close()
				return
			} else {
				c.JSON(http.StatusOK, gin.H{"result": "True"})
				sqlDB, err := connectdatabase.DBConn().DB()
				if err != nil {
					log.Fatalln(err)
				}
				defer sqlDB.Close()
				return
			}
		} else {
			if err := connectdatabase.DBConn().Model(&servicesOfProvider).Where("services_id=?", services.ID).Update("price", checkInformation.Price).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"result": "Update thất bại"})
				sqlDB, err := connectdatabase.DBConn().DB()
				if err != nil {
					log.Fatalln(err)
				}
				defer sqlDB.Close()
				return
			} else {
				c.JSON(http.StatusOK, gin.H{"result": "Update thành công"})
				sqlDB, err := connectdatabase.DBConn().DB()
				if err != nil {
					log.Fatalln(err)
				}
				defer sqlDB.Close()
				return
			}
		}
	}
}

func AddTodoList(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, err := dbConnect.DB()
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	var requirementcustomer models.RequirementsCustomer
	var addTodoList models.ToDoList
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &addTodoList)
	if err := connectdatabase.DBConn().First(&requirementcustomer, "id=?", addTodoList.RequirementsCustomerID).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "Không thấy"})
		sqlDB, err := connectdatabase.DBConn().DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()
	} else {
		if requirementcustomer.Status == 0 {
			if err := connectdatabase.DBConn().Create(&addTodoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"result": "False"})
				sqlDB, err := connectdatabase.DBConn().DB()
				if err != nil {
					log.Fatalln(err)
				}
				defer sqlDB.Close()
			} else {
				connectdatabase.DBConn().Model(&requirementcustomer).Where("id=?", addTodoList.RequirementsCustomerID).Update("status", 1)
				c.JSON(http.StatusOK, gin.H{"result": "True"})
				sqlDB, err := connectdatabase.DBConn().DB()
				if err != nil {
					log.Fatalln(err)
				}
				defer sqlDB.Close()
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "Công việc đã được người khác nhận"})
			sqlDB, err := connectdatabase.DBConn().DB()
			if err != nil {
				log.Fatalln(err)
			}
			defer sqlDB.Close()
		}
	}
}

func CountPaginationRequirement(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, err := dbConnect.DB()
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	type Count struct {
		Count uint
	}
	type CheckStatus struct {
		Status uint
	}
	var count Count
	var checkStatus CheckStatus
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &checkStatus)
	if err := connectdatabase.DBConn().Raw("SELECT COUNT(requirements_customers.id) AS "+"Count"+" FROM `requirements_customers` WHERE requirements_customers.status = ?", checkStatus.Status).Scan(&count).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
		sqlDB, err := connectdatabase.DBConn().DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()
	} else {
		c.JSON(http.StatusOK, gin.H{"result": count})
		sqlDB, err := connectdatabase.DBConn().DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()
	}
}
func CountPaginationToDoList(c *gin.Context) {
	dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	}
	type Count struct {
		Count uint
	}
	type CheckStatus struct {
		Status     uint
		ProviderId uint
	}
	var count Count
	var checkStatus CheckStatus
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &checkStatus)
	fmt.Println(checkStatus)
	if err := dbConnect.Raw("SELECT COUNT(to_do_lists.id) AS "+"Count"+" FROM to_do_lists WHERE to_do_lists.status = ? AND to_do_lists.provider_id = ?", checkStatus.Status, checkStatus.ProviderId).Scan(&count).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
		sqlDB, err := dbConnect.DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()
	} else {
		c.JSON(http.StatusOK, gin.H{"result": count})
		sqlDB, err := dbConnect.DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()
	}
}

func UpdateTodoList(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, err := dbConnect.DB()
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	type GetInformation struct {
		ProviderId            uint
		RequirementCustomerId uint
	}
	var getInformation GetInformation
	var updateTodoList models.ToDoList
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &getInformation)
	fmt.Println("><<<<<<<<<<<<<<<<<<<<<", getInformation)
	if err := connectdatabase.DBConn().Raw("UPDATE to_do_lists set status = 1, day_end = CURRENT_DATE WHERE provider_id = ? and requirements_customer_id  = ?", getInformation.ProviderId, getInformation.RequirementCustomerId).Scan(&updateTodoList).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
		sqlDB, err := connectdatabase.DBConn().DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "True"})
		sqlDB, err := connectdatabase.DBConn().DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()
	}
}

func DeleteServicesProvider(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, err := dbConnect.DB()
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	type Information struct {
		ProviderId uint
		ServicesId uint
	}
	var information Information
	var deleteservices models.ServicesOfProvider
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &information)
	if err := connectdatabase.DBConn().Raw("DELETE FROM `services_of_providers` WHERE provider_id = ? and services_id = ?", information.ProviderId, information.ServicesId).Scan(&deleteservices).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
		sqlDB, err := connectdatabase.DBConn().DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "True"})
		sqlDB, err := connectdatabase.DBConn().DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()
	}
}

//webSocket returns json format
func AABBCC(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, err := dbConnect.DB()
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	//Upgrade get request to webSocket protocol
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
		log.Fatal(err)
	}
	defer ws.Close()
	var data struct {
		A string `json:"Id"`
		B int    `json:"b"`
	}

	err = ws.ReadJSON(&data)
	if err != nil {
		log.Println("error read json")
		log.Fatal(err)
	}
	type GetServices struct {
		ServicesId   uint
		NameServices string
		Price        string
		ProviderId   uint
	}
	var getServices []GetServices
	for {

		connectdatabase.DBConn().Raw("SELECT services_of_providers.services_id,services.name_services, services_of_providers.price, services_of_providers.provider_id FROM"+
			" `services_of_providers` LEFT JOIN services on services_of_providers.services_id = services.id"+
			" WHERE services_of_providers.provider_id = ?", data.A).Scan(&getServices)

		err = ws.WriteJSON(getServices)
		if err != nil {
			log.Println("error write json: " + err.Error())
		}
		time.Sleep(1 * time.Second)
	}
}
func HistoryList(c *gin.Context) {
	dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	}
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
		log.Fatal(err)
	}
	defer ws.Close()
	type CheckProvider struct {
		Id string
	}
	var checkProvider CheckProvider

	err = ws.ReadJSON(&checkProvider)

	if err != nil {
		log.Println("error read json")
		log.Fatal(err)
	}
	type TodoList struct {
		NameServices    string
		Status          string
		DayStart        string
		TimeStart       string
		DayEnd          string
		NameCustomer    string
		AddressCustomer string
		PhoneCustomer   string
	}
	var todoList []TodoList

	for {

		if err := dbConnect.Raw(
			"SELECT requirements_customers.name_services,to_do_lists.status,requirements_customers.day_start,requirements_customers.time_start,to_do_lists.day_end,"+
				" customers.name_customer,customers.address_customer,customers.phone_customer"+
				" FROM `to_do_lists`,requirements_customers,customers,providers WHERE to_do_lists.requirements_customer_id = requirements_customers.id and"+
				" requirements_customers.customer_id = customers.id and to_do_lists.provider_id = providers.id and to_do_lists.status = 1 and providers.id = ?", checkProvider.Id).Scan(&todoList).Error; err != nil {

			err = ws.WriteJSON("False")
			if err != nil {
				log.Println("error write json: " + err.Error())
			}
		} else {
			if len(todoList) > 0 {

				err = ws.WriteJSON(todoList)
				if err != nil {
					log.Println("error write json: " + err.Error())
				}
			} else {
				err = ws.WriteJSON("Bạn không có lịch sử công việc")
				if err != nil {
					log.Println("error write json: " + err.Error())
				}
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func UpdateInformationProvider(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, err := dbConnect.DB()
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	type GetInformation struct {
		ProviderId uint
		Name       string
		Address    string
		CCCD       string
		Phone      string
		introduce  string
	}
	var getInformation GetInformation
	var updateProvider models.Provider
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &getInformation)
	if err := connectdatabase.DBConn().Raw("UPDATE providers set name= ?, address= ?, cccd= ?, phone= ?, introduce= ? WHERE id = ?", getInformation.Name, getInformation.Address, getInformation.Phone, getInformation.CCCD, getInformation.introduce, getInformation.ProviderId).Scan(&updateProvider).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
		sqlDB, err := connectdatabase.DBConn().DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()
	} else {
		c.JSON(http.StatusOK, gin.H{"result": getInformation})
		sqlDB, err := connectdatabase.DBConn().DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()
	}
}
