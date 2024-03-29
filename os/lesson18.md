# 장치 컨트롤러와 장치드라이버

## 장치 컨트롤러

입출력 장치의 종류가 너무 다양하고, CPU와의 데이터 전송률이 크게 차이가 나기 때문에 반드시 필요하다

입출력장치의 상태를 보여주는 상태 레지스터
입출력장치와 CPU가 주고받는 데이터가 임시로 저장되는 데이터 레지스터
각종 명령어를 저장해두는 제어 레지스터로 구성된다.

- 통신 중개
- 오류 검출
- 버퍼링: CPU와의 입출력 속도를 맞추기 위해 데이터를 모아두었다가 한번에 주거나, 모아두고 조금씩 내보내는 기법이다.

## 장치 드라이버

장치 컨트롤러라는 하드웨어를 컴퓨터 내부, 즉 CPU와 통신할 수 있도록 도와주는 소프트웨어다.
이것이 없으면 입출력이 불가능

장치 드라이버는 보통 운영체제가 인식한다. 운영체제가 기본으로 제공하기도 하지만 프로그램 제작자가 공급할 수도 있음

# 입출력 방법

3가지 방법이 있다

1. 프로그램 입출력
   프로그램 속 명령어로 입출력을 제어한다.

장치 컨트롤러에 있는 3개의 레지스터를 사용함.
쓰기: 제어 레지스터에 쓰기 명령 -> 상태 레지스터 확인 후 가용 상태인 경우 데이터를 제어 레지스터에 등록.

각 입출력장치의 레지스터에 주는 명령어 표현 방식:

- 메모리 맵 입출력: CPU의 메모리에 메모리 주소 접근 공간 + 입출력 접근을 위한 주소 공간을 하나로 간주한다.
  즉 여러개의 주소공간을 나눠서 쓴다.
  각 메모리 주소마다 장치 컨트롤러를 가리키는 정보를 담고 있으므로 해당 메모리 주소로 장치 컨트롤러와 상호작용이 가능해진다.
  따라서 같은 종류의 명령어로 메모리와 입출력장치를 조작할 수 있다.

- 고립형 입출력: 메모리 주소 접근 공간과 입출력 접근을 위한 주소공간을 분리하여 메모리 공간이 줄어들지 않는다는 장점이 있다.
  이는 제어 버스에 메모리 읽기 쓰기 선에 추가로 입출력장치 읽기 쓰기 선이 있을 때 가능하다.
  즉 메모리 읽기 쓰기 명령어를 실행할때는 메모리에 접근하고, 입출력장치 읽기 쓰기 명령어 실행시에는 입출력장치에 접근할 수 있다.
  이때 각각에 대해서 사용하는 명령어가 다르다. 즉 입출력장치에 대해서는 전용 명령어를 쓴다.

2. 인터럽트 기반 입출력
   CPU가 항상 입출력장치의 일을 기다리고 있을 수는 없기 때문에 나온 방식이다.
   장치 컨트롤러가 인터럽트를 발생시키는 주체이며, 입출력장치가 맡은 일을 완료했을 때 장치 컨트롤러는 CPU에 인터럽트 신호를 보내서
   CPU가 해당 작업을 처리할 수 있도록 한다.

   - 폴링 방식은 CPU가 주기적으로 입출력장치로부터 오는 정보를 확인하는 방식이라 위 방법보다 리소스 소모가 큰 편이다.
   - 여러개의 인터럽트가 발생하는 경우 우선순위에 따라 처리된다.
   - NMI, (non maskable interrupt) 는 모든 인터럽트 중 가장 우선한다. 그 외에는 PIC를 통해서 우선순위가 결정된다.
     PIC는 입출력장치에서 발생한 여러 입터럽트들의 우선순위를 계산하여 CPU에게 우선 처리해야 할 인터럽트가 무엇인지를 알려주는 부품이며
     이 부품들을 계층적으로 조직화해서 좀 더 많고 복잡한 인터럽트들을 처리할 수도 있다.
     PIC는 데이터 버스를 통해서 CPU에 인터럽트 벡터를 보낸다. 이로서 CPU는 인터럽트를 보낸 주체가 누구인지를 알게 된다.

3. DMA(direct memory access, CPU 없이 바로 메모리 접근) 입출력
   입출력장치와 메모리 간의 데이터 읽기 쓰기는 CPU를 거치도록 설계되어 있었는데,
   이렇게 하면 매번 CPU를 거쳐야 하기 때문에 그 부담이 커진다.
   이를 피하기 위해 시스템 버스에 DMA 컨트롤러를 연결하게 된다.
   DMA 컨트롤러는 CPU 대신 입출력장치와 상호작용을 하며 이때 시스템 버스를 사용한다.

(CPU가 DMA 컨트롤러에 명령을 내리면 바로 메모리에 접근 그 후 장치 컨트롤러에 내보낸다. 작업 완료 후 DMA 컨트롤러가 CPU에게 인터럽트를 보내준다.)

이때 문제가 되는데,
시스템 버스는 한번에 한 주체만 사용을 할 수 있다는 것이다.
따라서 DMA가 시스템 버스를 쓰려면 CPU에게 양해를 구해야 한다는 문제가 있다.
이를 해결하기 위해, 입출력 버스라는 것을 두고, 시스템 버스 대신 사용하도록 한다. 현재 대부분의 컴퓨터에 존재한다.
예) PCI 버스

요즘은.. 입출력장치 전용 CPU인 입출력 프로세서(입출력 채널)이 만들어져서 CPU의 부담을 완전 덜어내기도 한다.
예) 듀얼코어 CPU를 탑재한 레이저프린터
이때 CPU는 입출력 채널에 명령만 내리고 그 이후에는 입출력채널이 알아서 작업 후 결과를 인터럽트로 CPU에게 준다.
