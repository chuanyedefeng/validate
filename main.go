package main

import (
	"fmt"
	"validate/commonvalid"
	"validate/customvalid"
	"validate/customvalid/dto"
)

func main() {
	customvalid.InitValid()
	fmt.Println("自定义校验数字类型")
	//自定义校验数字类型
	order := &dto.OrderDTO{
		PayMoney: "1.2",
	}
	err := customvalid.ValidateStruct(order)
	fmt.Print(err)

	fmt.Println("校验枚举类型")
	// 自定义校验枚举类型
	car := &dto.Car{
		Color: 100,
	}
	errCar := customvalid.ValidateStruct(car)
	fmt.Println(errCar)

	commonvalid.ValidTestSimple()
}
