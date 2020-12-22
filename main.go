package main

import (
	"fmt"

	deli "github.com/hoon-jo/mileage-shop/pkg/delivery"
	item "github.com/hoon-jo/mileage-shop/pkg/product"
	user "github.com/hoon-jo/mileage-shop/pkg/user"
)

func main() {
	deliveryCount := 0
	buyer := user.CreateBuyer()
	products := item.CreateProducts()
	deliverystart := make(chan bool)
	deliveryList := deli.NewDelivery(deliverystart, deliveryCount)

	for {
		menu := 0 // 첫 메뉴

		fmt.Println("1. 구매")
		fmt.Println("2. 잔여 수량 확인")
		fmt.Println("3. 잔여 마일리지 확인")
		fmt.Println("4. 배송 상태 확인")
		fmt.Println("5. 장바구니 확인")
		fmt.Println("6. 프로그램 종료")
		fmt.Print("실행할 기능을 입력하시오 :")

		fmt.Scanln(&menu)
		fmt.Println()

		switch menu {

		case 1:
			user.BuyProduct(products, buyer, &deliveryCount, deliverystart)
		case 2:
			item.GetProducts(products)
		case 3:
			user.GetMaileage(buyer)
		case 4:
			deli.GetDeliveryStatus(deliveryList)
		case 5:
			user.ViewMyCart(products, buyer)
		case 6:
			fmt.Println("프로그램을 종료합니다.")
			return
		default:
			fmt.Println("잘못된 입력입니다. 다시 입력해주세요.")
		}
		if menu > 0 && menu < 7 {
			fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
		}

	}

}
