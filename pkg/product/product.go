package product

import (
	"fmt"
	"reflect"
)

type Item struct {
	name   string
	price  int
	amount int
}

func GetField(v *Item, field string) int {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return int(f.Int())
}
func CreateProducts() []Item {
	items := make([]Item, 5) // 물품 목록

	items[0] = Item{"텀블러", 10000, 30}
	items[1] = Item{"롱패딩", 500000, 20}
	items[2] = Item{"투미 백팩", 400000, 20}
	items[3] = Item{"나이키 운동화", 150000, 50}
	items[4] = Item{"빼빼로", 1200, 500}
	return items
}

func GetProducts(products []Item) {
	for _, product := range products {
		fmt.Printf("%s 남은수량 : %d\n", product.name, product.amount)
	}
}
func GetProductFields(product Item) (string, int, int) {

	return product.name, product.price, product.amount
}

func CalculProductAmount(products []Item, index int, bought int) int {
	products[index].amount = products[index].amount - bought

	return products[index].price
}
