package service

import (
	"familyCard/customerManage/model"
	"fmt"
)

// CustomerService 完成对Customer的增删改查
type CustomerService struct {
	customers []model.Customer
	//	声明当前切片有多少个客户
	customerNumber int
}

//编写方法返回CustomerService实例
func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNumber = 1
	customer := model.NewCustomerById(1, "张三", "男", 20, "18021192217", "2275701932@qq.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}

func (cs *CustomerService) DeleteById(id int) bool {
	index := cs.FindById(id)
	if index == -1 {
		return false
	}
	cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
	return true
}

func (cs *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(cs.customers); i++ {
		if cs.customers[i].Id == id {
			index = i
			break
		}
	}
	return index
}

// 添加客户
func (c *CustomerService) Add() {
	c.customerNumber++
	fmt.Println("---------添加客户-------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("手机：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)
	//	构建一个实例
	customer := model.NewCustomer(name, gender, age, phone, email)
	customer.Id = c.customerNumber
	c.customers = append(c.customers, customer)
}
