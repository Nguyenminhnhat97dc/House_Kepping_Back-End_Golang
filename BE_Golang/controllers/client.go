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
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"result": provider})
		return
	}
}

// select * from Services
func FindServices(c *gin.Context) {
	var services []models.Services
	if err := database.DBConn().Find(&services).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"result": services})
		return
	}
}

// SELECT * FROM users LIMIT 4;
func LimitServices(c *gin.Context) {
	var services []models.Services
	count := c.Param("count")
	number, _ := strconv.Atoi(count)
	if err := database.DBConn().Limit(number).Find(&services).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"result": services})
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
			return
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
					return
				} else {
					c.JSON(http.StatusOK, gin.H{"result": "Insert yêu cầu khách hàng thành công"})
					return
				}
			} else {
				if err := database.DBConn().Model(&NewRequirement).Where("customer_id = ? AND day_start = ? AND time_start = ? ", NewRequirement.CustomerID, NewRequirement.DayStart, NewRequirement.TimeStart).Update("name_services", NewRequirement.NameServices).Error; err != nil {
					c.JSON(http.StatusOK, gin.H{"result": "Không Update được"})
					return
				} else {
					c.JSON(http.StatusOK, gin.H{"result": "Update thành công"})
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
		if err := database.DBConn().Where("customer_id = ? AND day_start = ? AND time_start = ? ", NewRequirement.CustomerID, NewRequirement.DayStart, NewRequirement.TimeStart).First(&Requirement).Error; err != nil {
			if err := database.DBConn().Create(&NewRequirement).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"result": "không Insert yêu cầu khách hàng insert được"})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{"result": "Insert yêu cầu khách hàng thành công"})
				return
			}
		} else {
			if err := database.DBConn().Model(&NewRequirement).Where("customer_id = ? AND day_start = ? AND time_start = ? ", NewRequirement.CustomerID, NewRequirement.DayStart, NewRequirement.TimeStart).Update("name_services", NewRequirement.NameServices).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"result": "Không Update được"})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{"result": "Update thành công"})
				return
			}
		}
	}
}

func ServiceProvider(c *gin.Context) {
	type CheckProvider struct {
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
		c.JSON(http.StatusOK, gin.H{"result": "False"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "true"})
		return
	}
}

func RequirementsCustomer(c *gin.Context) {
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
	type Pagination struct {
		PaginationStart uint
		PaginationEnd   uint
	}
	var pagination Pagination
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &pagination)
	fmt.Println(">>>>>>>Pagination", pagination)

	if err := database.DBConn().Raw(
		"SELECT requirements_customers.id,requirements_customers.name_services,requirements_customers.day_start,requirements_customers.time_start,customers.name_customer,customers.address_customer,customers.phone_customer"+
			" FROM requirements_customers,customers WHERE requirements_customers.customer_id = customers.id and requirements_customers.status = ? LIMIT ?,?", 0, pagination.PaginationStart, pagination.PaginationEnd).Scan(&informationRequirementsCustomer).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"result": informationRequirementsCustomer})
		return
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
		DayEnd          string
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
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"result": informationLoggin})
		return
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
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"result": informationProvider})
		return
	}
}

func FindPriceOfServices(c *gin.Context) {
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
	fmt.Println("Check>>>>>>>>>>>>>>>>>", checkID.Id)
	if err := database.DBConn().Raw(
		"SELECT services.name_services, services_of_providers.price, providers.name from"+
			" services_of_providers,services,providers WHERE services_of_providers.services_id = services.id and"+
			" services_of_providers.provider_id = providers.id and providers.id = ?", checkID.Id).Scan(&price).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"result": price})
		return
	}
}

func AddPrice(c *gin.Context) {
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
	if err := database.DBConn().First(&services, "name_services=?", &checkInformation.NameServices).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
		return
	} else {
		fmt.Println(">>>>>>", services.ID, checkInformation.Price)
		if err := database.DBConn().Where("services_id=? and provider_id=?", services.ID, checkInformation.Id).First(&servicesOfProvider).Error; err != nil {
			NewServicesOfProvider := models.ServicesOfProvider{
				ServicesId: services.ID,
				ProviderID: checkInformation.Id,
				Price:      checkInformation.Price,
			}
			if err := database.DBConn().Create(&NewServicesOfProvider).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"result": "False"})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{"result": "True"})
				return
			}
		} else {
			if err := database.DBConn().Model(&servicesOfProvider).Where("services_id=?", services.ID).Update("price", checkInformation.Price).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"result": "Update thất bại"})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{"result": "Update thành công"})
				return
			}
		}
	}
}

func AddTodoList(c *gin.Context) {
	var requirementcustomer models.RequirementsCustomer
	var addTodoList models.ToDoList
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &addTodoList)
	fmt.Println("Check>>>>>>>>>>>>>>>>>", addTodoList.RequirementsCustomerID)
	if err := database.DBConn().First(&requirementcustomer, "id=?", addTodoList.RequirementsCustomerID).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "Không thấy"})
	} else {
		if requirementcustomer.Status == 0 {
			if err := database.DBConn().Create(&addTodoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"result": "False"})
			} else {
				database.DBConn().Model(&requirementcustomer).Where("id=?", addTodoList.RequirementsCustomerID).Update("status", 1)
				c.JSON(http.StatusOK, gin.H{"result": "True"})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "Công việc đã được người khác nhận"})
		}
	}
}

func CountPagination(c *gin.Context) {
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
	fmt.Println(">>>>>>>>>>Check", checkStatus)
	if err := database.DBConn().Raw("SELECT COUNT(requirements_customers.id) AS "+"Count"+" FROM `requirements_customers` WHERE requirements_customers.status = ?", checkStatus.Status).Scan(&count).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": count})
	}
}

func UpdateTodoList(c *gin.Context) {
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
	fmt.Println(getInformation)
	if err := database.DBConn().Model(&updateTodoList).Where("requirements_customer_id=? and provider_id=?", getInformation.RequirementCustomerId, getInformation.ProviderId).Update("status", 1).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "True"})
	}
}

func DeleteServicesProvider(c *gin.Context) {
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
	fmt.Println(information)
	if err := database.DBConn().Raw("DELETE FROM `services_of_providers` WHERE provider_id = ? and services_id = ?", information.ProviderId, information.ServicesId).Scan(&deleteservices).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "True"})
	}
}
