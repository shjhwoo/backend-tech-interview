# 채널 다루기

<잘 안되었던 부분1>
아래 두개의 고루틴으로 처리한 비동기 작업의 결과물은 어느거나 먼저 받아와줘도 상관없다고 생각해서 채널에서 값을 빼오는 과정에
select, case 문을 사용했는데, 첫번째 DiagnosisList만 받아와졌고
두번째는 계속 안 받아와졌다.
그래서 어쩔수없이 아래 코드로 수정해서 했더니 두개 다 받아와졌다.

내께 안되었던 이유는? for {} 로 안 감싸줘서임.
select... case는 반복되지 않고 단 한번만 실행되니깐.

```

	var diagChan = make(chan []*rpcPayload.DocDiagnosis)
	var payChan = make(chan []*rpcPayload.PaymentModel)

	go func() {
		err = rpcClient.IToEhr.GetAllDiagnosesOfConsultations(&rpcPayload.DiagnosesFilter{
			OrgId:              mcd.CalcReq.OrgId,
			ConsultationIdList: mcd.CalcResult.ConsultationIdList,
			DiagnosisTypes:     []int{1},
			SelectColumns:      []string{"consultationId", "disease", "kcd"},
		}, diagChan)
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	go func() {
		err = rpcClient.IToPay.GetPaidAmountList(mcd.CalcReq.OrgId, &rpcPayload.PaymentFilter{
			PatientId:          mcd.CalcReq.PatientId,
			ConsultationIdList: mcd.CalcResult.ConsultationIdList,
		}, payChan)
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	mcd.DiagnosisList = <-diagChan
	mcd.PaidAmtList = <-payChan

    --- 고친다면:

for {
    select {
    case mcd.DiagnosisList = <-diagChan:
    case mcd.PaidAmtList = <-payChan:
    }
}

근데 고쳤더니 실패했다; (waitGroup넣었는데도!)
그 이유는.


```

<잘 안되었던 부분2>
range 를 통해, 채널에서 값을 추출하는 경우에는
채널에 값을 전달해준 다음에 꼭 그 채널을 close 해 줘야 한다.
그렇지 않으면 데드락이 발생한다

```
틀린 예시
ch := make(chan int, 1)
ch <- 101
for value := range ch {
	fmt.Println(value)
}

```

<잘 안되었던 부분3>
