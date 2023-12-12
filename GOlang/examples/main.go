package main

import (
	"fmt"
	"sync"
	"time"
)

//컨베이어 벨트 시스템을 만들자.

//금액 계산 요청 빌드 단계
//금액 계산 rpc 요청 보내서 결과 받아오는단계
//계산 결과를 편집하여 빌드하는 단계

// 그리고 나중에는 이 부분을 금액 관련 서식에서 공통으로 쓸 수 있어야 한다
type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func main() {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Printf("Start Factory\n")

	wg.Add(3)
	go makeBody(tireCh)
	go installTire(tireCh, paintCh)
	go paintCar(paintCh)

	wg.Wait()
	fmt.Println("Close the factory")
}

func makeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			car := &Car{
				Body: "sports car",
			}

			tireCh <- car
		case <-after:
			close(tireCh)
			wg.Done()
			return
		}
	}
}

func installTire(tireCh, paintCh chan *Car) {
	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "winter tire"
		paintCh <- car
	}
	wg.Done()
	close(paintCh)
}

func paintCar(paintCh chan *Car) {
	for car := range paintCh {
		time.Sleep(time.Second)
		car.Color = "blue"
		duration := time.Since(startTime)

		fmt.Printf("%.2f Complete car: %s %s %s \n", duration.Seconds(), car.Body, car.Color, car.Tire)
	}
	wg.Done()
}
