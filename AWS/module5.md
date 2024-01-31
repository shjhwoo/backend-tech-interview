# EC2 크기 조정

확장성과 탄력성

수요에 맞게끔 프로비저닝을 가능하게 한다.

확장성

확장성을 위해서는 필요한 리소스만으로 시작하고 확장 및 축소를 통해 수요 변화에 자동으로 대응하도록 아키텍처를 설계해야 합니다. 그 결과, 사용한 리소스에 대해서만 비용을 지불합니다. 컴퓨팅 용량 부족 때문에 고객의 요구 사항을 충족할 수 없을지 걱정할 필요가 없습니다.

이 조정 프로세스가 자동으로 수행되도록 하려면 어떤 AWS 서비스를 사용해야 할까요? Amazon EC2 인스턴스에 이 기능을 제공하는 AWS 서비스가 Amazon EC2 Auto Scaling입니다.

Amazon EC2 Auto Scaling

잘 로드되지 않고 빈번히 시간 초과되는 웹 사이트에 액세스하려고 한 적이 있다면 이 웹 사이트가 처리할 수 있는 것보다 많은 요청을 수신한 것일 수 있습니다. 이는 커피숍에 고객의 주문을 처리할 바리스타가 한 명밖에 없을 때 길게 줄을 서서 기다리는 상황과 비슷합니다.

Amazon EC2 Auto Scaling을 사용하면 변화하는 애플리케이션 수요에 따라 Amazon EC2 인스턴스를 자동으로 추가하거나 제거할 수 있습니다. 필요에 따라 인스턴스를 자동으로 조정하여 애플리케이션 가용성을 효과적으로 유지할 수 있습니다.

Amazon EC2 Auto Scaling에서는 동적 조정과 예측 조정이라는 2가지 접근 방식을 사용할 수 있습니다.

동적 조정은 수요 변화에 대응합니다.
예측 조정은 예측된 수요에 따라 적절한 수의 Amazon EC2 인스턴스를 자동으로 예약합니다.

예: Amazon EC2 Auto Scaling

클라우드에서는 컴퓨팅 파워가 프로그래밍 방식의 리소스이므로 더 유연한 크기 조정 방식을 사용할 수 있습니다. Amazon EC2 Auto Scaling을 애플리케이션에 추가하면 필요할 때 새 인스턴스를 애플리케이션에 추가했다가 더 이상 필요하지 않으면 종료할 수 있습니다.

Amazon EC2 인스턴스에서 애플리케이션을 시작할 준비를 하고 있다고 가정해 보겠습니다. Auto Scaling 그룹의 크기를 구성할 때 최소 Amazon EC2 인스턴스 수를 1로 설정할 수 있습니다. 즉, 하나 이상의 Amazon EC2 인스턴스가 항상 실행 중이어야 합니다.

Auto Scaling 그룹의 일부로 확장 및 축소되는 Amazon EC2 인스턴스.
Auto Scaling 그룹을 생성할 때 최소 Amazon EC2 인스턴스 수를 설정할 수 있습니다. 최소 용량은 Auto Scaling 그룹을 생성한 직후 시작되는 Amazon EC2 인스턴스의 수입니다. 이 예에서 Auto Scaling 그룹의 최소 용량은 Amazon EC2 인스턴스 1개입니다.

그런 다음 애플리케이션을 실행하려면 최소 하나의 Amazon EC2 인스턴스가 필요하더라도 희망 용량을 Amazon EC2 인스턴스 2개로 설정할 수 있습니다.

Auto Scaling 그룹에서 희망 Amazon EC2 인스턴스 수를 지정하지 않으면 희망 용량은 기본적으로 최소 용량으로 설정됩니다.

Auto Scaling 그룹에서 설정할 수 있는 세 번째 구성은 최대 용량입니다. 예를 들어 수요 증가에 대응하여 확장하도록 Auto Scaling 그룹을 구성하되 Amazon EC2 인스턴스 수를 최대 4개로 제한할 수 있습니다.

Amazon EC2 Auto Scaling은 Amazon EC2 인스턴스를 사용하므로 사용하는 인스턴스에 대해서만 비용을 지불하면 됩니다. 이제 여러분은 비용을 줄이면서도 최상의 고객 경험을 제공하는 비용 효율적인 아키텍처를 갖게 되었습니다.