# 트랜잭션의 잘못된 사용으로 인한 잠금 문제.

현재 사용중인 트랜잭션의 목록 확인하고, process id를 찾아서
kill process id 까지는 했지만 왜 잠금이 자주 생겼는지를 확인해야 한다.

## 어떻게 해결했냐면 ..

일단 급했기 때문에, POST API에 걸려있던 모든 트랜잭션을 삭제했다.

해당 로직은 다음과 같았다:

1.트랜잭션 시작
2.defer 무조건 트랜잭션 반납 또는 롤백 3.트랜잭션을 이용해 데이터 생성한다, RPC 통신도 한다.

```
	newConsultationList := []rpcPayload.Consultation{}
	tx, err := dbm.WithTx(edbm.WriteReadDB, context.Background())
	var errList []error
	if err != nil {
		errList = append(errList, err)
		return err
	}

	defer func() {
		fmt.Println("에러가 발생했다면 무조건 여기서 실행이 되어ㅇ햐안다.")
		if tx != nil {
			fmt.Println("에러가 발생했다면 무조건 롤백이 되어야한다")
			dbm.CommitOrRollBack(tx, errList)
		}
	}()

	for _, e := range consultationList {
		id, err := createConsultation(tx, orgId, edbm.Consultation(e))
		if err != nil {
			fmt.Println("에러가 발생했다")
			errList = append(errList, err)
			return err
		}
```

실제 동작 확인을 위해 위 로직에 대해 테스트코드를 실행해봤다.

```
func TestPostConsultation(t *testing.T) {
	var result string
	cs := &ConsultRpc{}

	cs.PostConsultation(rpcPayload.Envelop{
		From:   rpcPayload.Ehr,
		To:     rpcPayload.Registeration,
		OrgId:  "s00001",
		Method: rpcPayload.Post,
		Card:   []byte("{\"scheduleId\":5,\"medicalType\":9,\"doctorId\":1,\"consultTime\":\"20231211134513\",\"subjective\":\"아파보인다.\",\"objective\":\"OOOO\",\"assessment\":\"침술\",\"plan\":\"침술장기간으로\",\"stage\":1,\"result\":\"1\",\"createdAt\":\"\",\"discarded\":0,\"patientId\":22,\"instype\":10,\"addition\":\"0|0\",\"bookmark\":1}"),
	}, &result)

	t.Log(result, "결과")
}
```

그리고 출력된 로그는 다음과 같았다.

```
=== RUN   TestPostConsultation
working.....
에러가 발생했다
에러가 발생했다면 무조건 여기서 실행이 되어ㅇ햐안다.
에러가 발생했다면 무조건 롤백이 되어야한다
[Error 1146 (42S02): Table 's00001_ehr.consultation111' doesn't exist]
roll back
```

일단 이 부분에서는 이상한게 없는것 같다.
rpc 요청을 의심해봐야했다.
일부러 에러를 내본 다음 테스트코드를 실행해봤다.

```
func (mc *MockRpcClient) PostBoardConsultatonListViaRPC(orgId, schid string, consultationList []rpcPayload.Consultation) (string, error) {
	return "", errors.New("ㅡㄷ개ㅜㅎ")
}
```

```
=== RUN   TestPostConsultation
working.....
에러가 발생했다면 무조건 여기서 실행이 되어ㅇ햐안다.
에러가 발생했다면 무조건 롤백이 되어야한다
[sql: no rows in result set]
roll back
```

롤백되는건 여전했다.
그러면 왜 이 요청 이후에 바로 PATCH 요청을 보낸게 lock이 걸려버린건지..

트랜잭션 목록 확인했을 때는 이름 없는 trx와, 이름 있는 trx 2개가 생성되어 있어서 항상 2개를 kill 해야 했다.
그리고 이름없는 trx가 항상 먼저 생겨서 이것 때문에 뒤의 trx가 한참이나 기다리다가 실패하는 패턴이었다.

잠금에 대한 이해가 좀더 필요할거같다.

일단 문제가 발생했던 상황은
POST 요청에서도 trx 사용
그 이후 바로 PATCH 요청을 보내는데서 trx 사용
=> 개별적으로는 문제가 전혀 없었는데 연달아 보내면서 잠금과 관한 로직이 꼬인것 같다.

# 잠금?

8.0버전으로 넘어오게 되면서 data_locks, data_lock_waits 테이블로 대체되고 있음.

SELECT \* FROM perfomance_schema.data_locks 확인하여 가지고 있는 잠금 확인가능하다.
