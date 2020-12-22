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

	name, _, _ := item.GetProductFields(product)
	user.basket[name] = amount
}

func BuyProduct(products []item.Item, user *buyer, BuyProduct int) {
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
				if BuyProduct > 5 {
					fmt.Println("배송 한도를 초과했습니다. 배송이 완료되면 주문하세요.")
					break
				}
				price := item.CalculProductAmount(products, index-1, inputAmount)
				CalculUserPoint(user, price*inputAmount)
				BuyProduct++
				break
			} else if buyOrAddToCart == 2 {
				AddToCart(user, products[index-1], inputAmount)
				fmt.Println("성공적으로 장바구니에 담겼습니다.")
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
func ViewMyCart(products []item.Item, user *buyer) {
	if len(user.basket) == 0 {
		fmt.Println("장바구니가 비었습니다.")
	} else {
		for index, val := range user.basket {
			fmt.Printf("%s, 수량: %d\n", index, val)
		}
		var orderOrNot int
		for {
			fmt.Println("1. 장바구니 상품 주문")
			fmt.Println("2. 메뉴로 돌아가기")
			fmt.Print("실행할 기능을 입력하시오 :")
			fmt.Scanln(&orderOrNot)
			fmt.Println()
			if orderOrNot == 1 {
				BuyAllItemsInCart(products, user)
				break
			} else if orderOrNot == 2 {
				break
			} else {
				fmt.Println("잘못된 입력입니다 다시 입력하세요")
			}
		}
	}
}

func BuyAllItemsInCart(products []item.Item, user *buyer) bool {
	totalPrice := 0
	ableToOrder := true

	for _, product := range products {
		var itemName, itemPrice, itemAmount = item.GetProductFields(product)
		for cartItemName, cartItemAmount := range user.basket {

			if cartItemName == itemName {
				if itemAmount < cartItemAmount {
					fmt.Println("주문할 수량이 재고보다 많습니다.")
					ableToOrder = false
					break
				}
				// cart[cartItemName]
				totalPrice += itemPrice * cartItemAmount
			}
		}
		if ableToOrder == false {
			break
		}
	}
	fmt.Printf("필요 마일리지 : %d\n", totalPrice)
	fmt.Printf("보유 마일리지 : %d\n", user.point)

	fmt.Println()
	if user.point < totalPrice {
		fmt.Println("마일리지가 %d점 부족합니다.", totalPrice-user.point)
		ableToOrder = false
	}

	if ableToOrder {
		CalculUserPoint(user, totalPrice)
		for i := 0; i < len(products); i++ {
			var itemName, _, _ = item.GetProductFields(products[i])
			for cartItemName, cartItemAmount := range user.basket {
				if cartItemName == itemName {
					price := item.CalculProductAmount(products, i, cartItemAmount)
					totalPrice = price
				}
			}

		}
		user.basket = map[string]int{}

	}
	return ableToOrder
}
