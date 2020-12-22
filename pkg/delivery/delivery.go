package delivery

import (
	"fmt"
	"time"
)

type Delivery struct {
	status   string
	stations map[string]int
}

func NewDelivery(deliverystart chan bool, deliveryCount int) []Delivery {
	dele := Delivery{}
	dele.stations = map[string]int{}
	var deliveryList = make([]Delivery, 5)
	for i := 0; i < 5; i++ {
		deliveryList[i] = dele
	}
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond) //고루틴 순서대로 실행되도록 약간 딜레이
		go DeliveryStatus(deliverystart, i, deliveryList, &deliveryCount)
	}
	return deliveryList
}

func DeliveryStatus(deliveryState chan bool, i int, deliveryList []Delivery, numbuy *int) {
	for {
		if <-deliveryState {
			deliveryList[i].status = "주문접수"
			time.Sleep(time.Second * 10)

			deliveryList[i].status = "배송중"
			time.Sleep(time.Second * 30)

			deliveryList[i].status = "배송완료"
			time.Sleep(time.Second * 10)

			deliveryList[i].status = ""
			*numbuy--
		}
	}
}
func GetDeliveryStatus(deliveryList []Delivery) {
	for i := 0; i < len(deliveryList); i++ {
		fmt.Printf("배송상황: %s\n", deliveryList[i].status)
	}

	fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
	fmt.Scanln()
}
