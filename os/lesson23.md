# 프로세스 상태와 계층 구조

프로세스의 상태는 다음 단계로 나눌 수 있음

## 생성 상태

프로세스를 생성 중인 상태

## 준비 상태

CPU 자원을 할당받기 위해 기다리고 있는 상태
준비상태를 실행상태로 넘기는 것을 디스패치 한다고 한다.

## 실행 상태

CPU 자원을 할당받아 주어진 프로그램을 실행중인 상태

## 대기상태

(대부분의 경우) 입출력장치 접근으로, 입출력장치 작업이 완료될 때까지 기다리는 상태.
대기상태가 끝나면 다시 준비상태로 돌아간다

# 종료 상태

사용중이던 프로세스를 메모리 영역에서 제거하고 CPU 자원을 수거한 상태.

=> 운영체제는 프로세스들을 PCB를 통해서 인식하고 관리한다. 즉 PCB는 옷가게의 옷들에 붙어있는 태그와도 같은 것임)

## 프로세스 계층 구조:

트리 구조를 이룬다.
부모 프로세스가 실행되고, 부모 프로세스가 자신의 정보를 복사해서 자식 프로세스를 만들고,
이러한 과정이 재귀적으로 일어나면서 트리 구조를 만들어나간다.

프로세스마다 고유의 아이디를 가지는데 이를 PID 라고 한다.
자식 프로세스는 자신의 부모 프로세스가 누군지 알수있도록 PPID 즉 부모의 PID 정보도 같이 가지고 있다.

운영체제의 최초 프로세스 ID는 항상 1번으로 시작한다.
리눅스, macOS에서는 프로세스 계층 구조를 pstree 명령어로 확인할 수 있다

## 프로세스 생성기법

fork() 명령어를 통해서 복사본을 자식 프로세스로 생성하고,(이때 부모 프로세스 자원들이 상속된다고 한다 => 그러나 저장된 위치는 다르다)
exec() 명령어를 통해서 생성한 자식 프로세스의 스택, 코드, 데이터, 메모리 영역에 새로운 프로그램을 덮어씌운다

예) bash 쉘에서 ls 라는 명령어 실행 시 :
쉘 프로세스는 fork 명령어로 새로운 자식 프로세스 생성하고,
자식 프로세스는 exec 명령어로 ls 명령어를 실햏하는 프로그램으로 덮어씌워진다

만일, 부모 자식 프로세스가 생성만되고 exec를 안 한다면 하나의 프로그램을 부모와 자식이 병행하여 실행하게 된다.

### GO 언어로 된 예제

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

    fmt.Println("Command finished successfully.")
}
```
