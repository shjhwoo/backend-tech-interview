# defer, channel, deadlock

```
package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

var medicalItems = []string{"이체자락관법-3850", "진찰료-9810", "이상한옹더-%#^$", "경혈침술-3860", "경피경근온열요법-2000", "추나요법-25000"}

func main() {
	//각각의 오더에 대해서 금액계산 요청을 고루틴으로 동시에 보내버리자

	calcResChan := make(chan *CalcResult, len(medicalItems))

	var wg sync.WaitGroup
	for _, item := range medicalItems {
		wg.Add(1)
		go func(item string) {
			defer wg.Done()
			RPCcalcSingleMedicalItem(item, calcResChan)
		}(item)
	}

	/*
		By using a separate goroutine for close(calcResChan),
		you ensure that the main goroutine doesn't close the channel
		until all worker goroutines have completed their tasks.
		This helps avoid any race conditions
		and ensures that the channel is closed only after all values have been sent.
	*/
	wg.Wait()
	close(calcResChan)

	for calcR := range calcResChan {
		fmt.Println(calcR)
	}
}

type Resp struct {
	CalcResult *CalcResult
	ErrContent error
}

type CalcResult struct {
	Name      string
	UnitPrice string
	Copay     string
	Claim     string
}

func RPCcalcSingleMedicalItem(item string, calcResChan chan *CalcResult) {
	name := strings.Split(item, "-")[0]

	fmt.Printf("%s 의 계산을 시작합니다", name)

	unitPrice := strings.Split(item, "-")[1]
	unitPriceInt, err := strconv.ParseInt(unitPrice, 10, 64)
	if err != nil {
		calcResChan <- nil
		return
	}

	copay := float64(unitPriceInt) * 0.3
	copayStr := fmt.Sprintf("%f", copay)

	claim := float64(unitPriceInt) - copay
	claimStr := fmt.Sprintf("%f", claim)

	calcResChan <- &CalcResult{
		Name:      name,
		UnitPrice: unitPrice,
		Copay:     copayStr,
		Claim:     claimStr,
	}
}

```
