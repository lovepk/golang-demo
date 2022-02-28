package main

import (
	"familyCard/customerManage/service"
	"fmt"
)

func main() {
	customer := customerView{
		key:             "",
		lool:            false,
		customerService: service.NewCustomerService(),
	}
	customer.mainMenu()
}

type customerView struct {
	//	定义必要字段
	key             string // 接受用户输入
	lool            bool   // 是否循环展示菜单
	customerService *service.CustomerService
}

// 显示所有用户信息
func (cv *customerView) list() {
	customers := cv.customerService.List()
	fmt.Println("---------客户列表开始-------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t手机号\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("---------客户列表完成-------------")

}

func (cv *customerView) add() {
	cv.customerService.Add()
}

func (cv *customerView) delete() {
	fmt.Println("---------删除客户-------------")
	fmt.Println("请选择待删除客户编号（-1退出）")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}

	if cv.customerService.DeleteById(id) {
		fmt.Println("删除成功！")
	} else {
		fmt.Println("删除失败！客户编号不存在")
	}
}

func (cv *customerView) mainMenu() {
	for {
		fmt.Println("---------客户信息管理软件-------------")
		fmt.Println("----------1.添加客户-------------")
		fmt.Println("----------2.修改客户-------------")
		fmt.Println("----------3.删除客户-------------")
		fmt.Println("----------4.客户列表-------------")
		fmt.Println("----------5.退   出-------------")
		fmt.Print("请选择（1-5）：")

		fmt.Scanln(&cv.key)

		switch cv.key {
		case "1":
			cv.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			cv.delete()
		case "4":
			fmt.Println("客户列表")
			cv.list()
		case "5":
			cv.lool = true
		}
		if cv.lool {
			break
		}
	}
	fmt.Println("已退出客户关系管理系统")
}
