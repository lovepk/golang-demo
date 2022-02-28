package model

import "fmt"

// 声明一个customer 一个结构体
type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

//使用工厂模式，返回一个Customer的实例
func NewCustomerById(id int, name string, gender string, age int,
	phone string, email string) Customer{
	return Customer{
		Id: id,
		Name: name,
		Gender: gender,
		Age: age,
		Phone: phone,
		Email: email,
	}
}

func NewCustomer(name string, gender string, age int,
	phone string, email string) Customer{
	return Customer {
		Name: name,
		Gender: gender,
		Age: age,
		Phone: phone,
		Email: email,
	}
}

//返回当前用户信息
func (c *Customer)GetInfo() string {
	str := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", c.Id, c.Name, c.Gender, c.Age, c.Phone, c.Email)
	return str
}
