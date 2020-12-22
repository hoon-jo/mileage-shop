package user

type buyer struct {
	point  int
	basket map[string]int
}

func createBuyer() *buyer {
	user := buyer{}
	user.point = 1000000
	user.basket = map[string]int{}
	return &user
}
