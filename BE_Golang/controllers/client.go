package controllers

import (
	database "BE_Golang/BE_Golang/connect_database"
	"BE_Golang/BE_Golang/create_database/models"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// select * from Provider
func FindProvider(c *gin.Context) {
	var provider []models.Provider
	if err := database.DBConn().Find(&provider).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": provider})
	}
}

// select * from Services
func FindServices(c *gin.Context) {
	var services []models.Services
	if err := database.DBConn().Find(&services).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": services})
	}
}

// SELECT * FROM users LIMIT 4;
func LimitServices(c *gin.Context) {
	var services []models.Services
	count := c.Param("count")
	number, _ := strconv.Atoi(count)
	if err := database.DBConn().Limit(number).Find(&services).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": services})
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
func UpdateRequirementCustomer(c *gin.Context) {
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
	fmt.Println(checkCustomer)
	abc := checkCustomer.Name
	if err := database.DBConn().Where("name = ? AND  address_customer = ?", abc, checkCustomer.Address).First(&Customer).Error; err != nil {
		NewCustomer := models.Customer{
			NameCustomer:    checkCustomer.Name,
			AddressCustomer: checkCustomer.Address,
			PhoneCustomer:   checkCustomer.Phone,
		}
		fmt.Println("THEM KH", NewCustomer)
		if err := database.DBConn().Create(&NewCustomer).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"result": "Không  insert Khách Hàng được"})
		} else {
			database.DBConn().Where("name_customer = ? AND  address_customer = ?", abc, checkCustomer.Address).First(&Customer)
			NewRequirement := models.RequirementsCustomer{
				CustomerID:   Customer.ID,
				NameServices: checkCustomer.NameServices,
				DayStart:     checkCustomer.DayStart,
				TimeStart:    checkCustomer.TimeStart,
			}
			fmt.Println(">>requirement", NewRequirement)
			if err := database.DBConn().Where("customer_id = ? AND day_start = ? AND time_start = ? ", NewRequirement.CustomerID, NewRequirement.DayStart, NewRequirement.TimeStart).First(&Requirement).Error; err != nil {
				if err := database.DBConn().Create(&NewRequirement).Error; err != nil {
					c.JSON(http.StatusOK, gin.H{"result": "create - không Insert yêu cầu khách hàng insert được"})
				} else {
					c.JSON(http.StatusOK, gin.H{"result": "Insert yêu cầu khách hàng thành công"})
				}
			} else {
				if err := database.DBConn().Model(&NewRequirement).Where("customer_id = ? AND day_start = ? AND time_start = ? ", NewRequirement.CustomerID, NewRequirement.DayStart, NewRequirement.TimeStart).Update("name_services", NewRequirement.NameServices).Error; err != nil {
					c.JSON(http.StatusOK, gin.H{"result": "Không Update được"})
				} else {
					c.JSON(http.StatusOK, gin.H{"result": "Update thành công"})
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
		if err := database.DBConn().Where("customer_id = ? AND day_start = ? AND time_start = ? ", NewRequirement.CustomerID, NewRequirement.DayStart, NewRequirement.TimeStart).First(&Requirement).Error; err != nil {
			if err := database.DBConn().Create(&NewRequirement).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"result": "không Insert yêu cầu khách hàng insert được"})
			} else {
				c.JSON(http.StatusOK, gin.H{"result": "Insert yêu cầu khách hàng thành công"})
			}
		} else {
			if err := database.DBConn().Model(&NewRequirement).Where("customer_id = ? AND day_start = ? AND time_start = ? ", NewRequirement.CustomerID, NewRequirement.DayStart, NewRequirement.TimeStart).Update("name_services", NewRequirement.NameServices).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"result": "Không Update được"})
			} else {
				c.JSON(http.StatusOK, gin.H{"result": "Update thành công"})
			}
		}
	}
}

func ServiceProvider(c *gin.Context) {
	type CheckProvider struct {
		Id string
	}
	type GetServices struct {
		NameServices string
		Price        string
	}
	var checkProvider CheckProvider
	var getServices []GetServices
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &checkProvider)
	if err := database.DBConn().Table("services_of_providers").Select(
		"services.name_services,price").Joins(
		"LEFT JOIN services on services_of_providers.services_id = services.id ").Where("services_id = ?", checkProvider.Id).Scan(&getServices).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "Không tìm thấy"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": getServices})
	}
}

func AddServiceProvider(c *gin.Context) {
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
	if err := database.DBConn().Create(&AddNewServiceProvider).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "false"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "true"})
	}
}

func RequirementsCustomer(c *gin.Context) {
	type InformationRequirementsCustomer struct {
		NameServices    string
		DayStart        string
		TimeStart       string
		NameCustomer    string
		AddressCustomer string
		PhoneCustomer   string
	}
	var informationRequirementsCustomer []InformationRequirementsCustomer
	if err := database.DBConn().Raw(
		"SELECT requirements_customers.name_services,requirements_customers.day_start,requirements_customers.time_start,customers.name_customer,customers.address_customer,customers.phone_customer"+
			" FROM requirements_customers,customers WHERE requirements_customers.customer_id = customers.id and requirements_customers.status = ?", 0).Scan(&informationRequirementsCustomer).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "false"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": informationRequirementsCustomer})
	}
}

func TodoList(c *gin.Context) {
	type CheckProvider struct {
		Id     string
		Status int
	}
	var checkProvider CheckProvider
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
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &checkProvider)
	fmt.Println(">>>>>>>ToDo", checkProvider)
	if err := database.DBConn().Raw(
		"SELECT to_do_lists.id,requirements_customers.name_services,to_do_lists.status,requirements_customers.day_start,requirements_customers.time_start,"+
			" customers.name_customer,customers.address_customer,customers.phone_customer"+
			" FROM `to_do_lists`,requirements_customers,customers,providers WHERE to_do_lists.requirements_customer_id = requirements_customers.id and"+
			" requirements_customers.customer_id = customers.id and to_do_lists.provider_id = providers.id and to_do_lists.status = ? and providers.id = ?", checkProvider.Status, checkProvider.Id).Scan(&todoList).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "false"})
	} else {
		if len(todoList) > 0 {
			c.JSON(http.StatusOK, gin.H{"result": todoList})
		} else {
			if checkProvider.Status == 1 {
				c.JSON(http.StatusOK, gin.H{"result": []string{"Bạn không có lịch sử công việc"}})
			} else {
				c.JSON(http.StatusOK, gin.H{"result": []string{"Bạn không có việc cần làm"}})
			}
		}
	}
}

func Loggin(c *gin.Context) {
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
	if err := database.DBConn().Where("user_name = ? and password = ?", checkLoggin.User, checkLoggin.Password).First(&informationLoggin).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": informationLoggin})
	}
}

func FindProviderID(c *gin.Context) {
	type CheckID struct {
		Id uint
	}
	var checkID CheckID
	var informationProvider models.Provider
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &checkID)
	if err := database.DBConn().First(&informationProvider, "id=?", checkID.Id).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": informationProvider})
	}
}
