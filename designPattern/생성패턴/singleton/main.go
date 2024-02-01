package main

import (
	"fmt"
	"sync"
)

/*
싱글턴 구조체​(struct)​는 여러 고루틴들이 해당 인스턴스에 접근하려고 할 때마다
 반드시 같은 인스턴스를 반환해야 합니다.
이 때문에 싱글턴 디자인 패턴을 잘못 구현하기 매우 쉽습니다.

처음에 single­Instance가 비어 있는지 확인하기 위한 nil 검사가 있습니다.
이 검사는 getinstance 메서드가 호출될 때마다 시간을 소모하는 잠금​(lock) 작업을 방지하기 위한 것이며
이 검사가 실패하면 single­Instance 필드가 이미 채워져 있음을 의미합니다.

single­Instance 구조체​(struct)​는 잠금​(lock) 내에서 생성됩니다.

잠금​(lock)​을 획득한 후 또 다른 nil 확인이 있습니다.
이는 둘 이상의 고루틴이 첫 번째 검사를 우회하는 경우
하나의 고루틴만 싱글턴 인스턴스를 생성할 수 있도록 하기 위함입니다.
그렇게 하지 않으면 모든 고루틴들은 싱글턴 구조체​(struct)​의 자체 인스턴스들을 생성할 것입니다.
*/

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("creating single instance now")
			singleInstance = &single{}
		} else {
			fmt.Println("single instance alreay exists")
		}
	}

	return singleInstance
}

func main() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}

/*싱글턴 인스턴스를 생성하는 다른 방법*/
/*
init 함수
init 함수 내에서 단일 인스턴스를 생성할 수 있습니다. 이는 인스턴스의 초기 초기화가 괜찮은 경우에만 적용됩니다.
init 함수는 패키지의 파일당 한 번만 호출되므로 단일 인스턴스만 생성될 수 있습니다.


sync.Once
sync.Once는 작업을 한 번만 수행합니다.

var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance() *single {
    if singleInstance == nil {
        once.Do(
            func() {
                fmt.Println("Creating single instance now.")
                singleInstance = &single{}
            })
    } else {
        fmt.Println("Single instance already created.")
    }

    return singleInstance
}
*/
