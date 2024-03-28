# 스레드

스레드는 실행 흐름의 단위이다.
프로세스를 구성한다.

단 , 리눅스 운영체제에서는 스레드와 프로세스를 명확히 구분짓지 않고, 태스크라는 개념으로 취급한다.

멀티스레드: 하나의 프로세스를 여러개의 스레드로 실행하는 방법이다. 프로세스가 가진 코드 데이터 힙 파일 영역은 공유하지만
각 스레드마다는 서로 다른 프로그램 카운터, 스택, 레지스터를 가지고 있다.
메모리는 아낄 수 있지만 공유 영역에 문제가 생기면 모든 스레드에 영향이 간다.

멀티프로세스: 여러개의 프로세스를 실행하는 개념이다.
프로세스 간에는 기본적으로 메모리 공유가 불가능하다는 점에서 스레드와는 다르다.
따라서 일반적으로는 멀티스레드보다 메모리 차지를 많이하게 된다.

하지만, 공유 메모리라는 것을 둘 수도 있고, 소켓이나 파이프를 통해서 프로세스 간 통신을 구현할 수도 있다.

## 멀티프로세스 예시:

```
package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    // Fork a new process
    cmd := exec.Command("/bin/ls", "-l") // Command to list files in the current directory

    // Set up pipes for reading stdout and stderr
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    // Start the command
    err := cmd.Start()
    if err != nil {
        fmt.Printf("Error starting command: %s\n", err)
        return
    }

    // Wait for the command to finish
    err = cmd.Wait()
    if err != nil {
        fmt.Printf("Error waiting for command to finish: %s\n", err)
        return
    }

    fmt.Println("Parent process finished.")
}

```

메인 프로세스에서 ls 명령어를 실행하는 새로운 프로세스를 포크한다. => 멀티프로세스의 예시

## 멀티스레드 예시:

```
package main

import (
    "fmt"
    "sync"
)

func printNumbers(wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}

func main() {
    var wg sync.WaitGroup
    wg.Add(2) // Number of goroutines to wait for

    // Launch two goroutines
    go printNumbers(&wg)
    go printNumbers(&wg)

    // Wait for both goroutines to finish
    wg.Wait()

    fmt.Println("All goroutines finished.")
}

```

하나의(또는 여러개의)스레드를 여러 고루틴이 왔다갔다 공유하고 있음.
프로세스는 하나지만 프로세스 실행 시 스레드에 접근하는 고루틴을 통해 프로그램 실행이 이루어진다.
