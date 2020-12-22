package product

type item struct {
	name   string
	price  int
	amount int
}

func CreateProducts() []item {
	items := make([]item, 5) // 물품 목록

	items[0] = item{"텀블러", 10000, 30}
	items[1] = item{"롱패딩", 500000, 20}
	items[2] = item{"투미 백팩", 400000, 20}
	items[3] = item{"나이키 운동화", 150000, 50}
	items[4] = item{"빼빼로", 1200, 500}
	return items
}
