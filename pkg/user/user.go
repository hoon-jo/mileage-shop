package user

import (
	"fmt"
	"reflect"

	item "github.com/hoon-jo/mileage-shop/pkg/product"
)

type buyer struct {
	point  int
	basket map[string]int
}

func CreateBuyer() *buyer {
	user := buyer{}
	user.point = 1000000
	user.basket = map[string]int{}
	return &user
}

func GetMaileage(user *buyer) {

	fmt.Printf("현재 잔여 마일리지는 %d점입니다.\n", user.point)

}
func CalculUserPoint(user *buyer, price int) {

	user.point -= price
}
func AddToCart(user *buyer, product item.Item, amount int) {

	name, price, amount := item.GetProductFields(product)
	user.basket[name] = price
}

func BuyProduct(products []item.Item, user *buyer) {
	buy := func(index int) {
		var inputAmount int
		var buyOrAddToCart int
		fmt.Print("수량을 입력하시오 :")
		fmt.Scanln(&inputAmount)

		if inputAmount <= 0 {
			panic("올바른 수량을 입력하세요.")
		}
		fmt.Println()
		for {
			fmt.Print("1. 바로 주문\n2. 장바구니에 담기\n")
			fmt.Print("실행할 기능을 입력하시오 :")
			fmt.Scanln(&buyOrAddToCart)
			fmt.Println()
			if buyOrAddToCart == 1 {
				price := item.CalculProductAmount(products, index-1, inputAmount)
				CalculUserPoint(user, price*inputAmount)
				break
			} else if buyOrAddToCart == 2 {
				AddToCart(user, products[index-1], inputAmount)
				break
			} else {
				fmt.Println("잘못된 입력입니다. 다시 입력 해주세요.")
			}
		}

	}

	for {

		itemchoice := 0
		for i, product := range products {
			v := reflect.ValueOf(product)

			fmt.Printf("물품%d: %s,  가격: %d,  잔여 수량: %d\n", i+1, v.FieldByName("name"), v.FieldByName("price"), v.FieldByName("amount"))
		}
		fmt.Print("구매할 물품을 선택하세요 :")
		fmt.Scanln(&itemchoice)

		if itemchoice > 0 && itemchoice <= len(products) {
			buy(itemchoice)
			break
		} else {
			fmt.Println("잘못된 입력입니다. 다시 입력해주세요.\n")
		}
	}
}
func ViewMyCart(user buyer) {
	if len(user.basket) == 0 {
		fmt.Println("장바구니가 비었습니다.")
	} else {
		for index, val := range user.basket {
			fmt.Printf("%s, 수량: %d\n", index, val)
		}
	}
	fmt.Println()
	fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
	fmt.Scanln()
}
