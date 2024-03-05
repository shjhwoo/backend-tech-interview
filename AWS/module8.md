# 추가 컴퓨팅 서비스

유연
안정
확정

- 서버리스 컴퓨팅 옵션을 제공.
  알아서 관리해주기 때문에, 기본 인프라를 보거나 액세스가 불가하다.
  사용자는 앱만 신경쓰면 된다.

예: aws lambda
트리거를 구성한다 => 트리거 감지하면 자동으로 코드를 실행한다.

람다 실행시간은 15분 미만으로, 즉 딥러닝에는 적합하지 않다. 이미지 리사이징에 적합.

예: ECS, EKS
컨테이너 오케스트레이션 도구이다.
도커 컨테이너(코드 패키지)를 의미한다.

즉 클러스터라는 인스턴스 모음에 대해 시작, 중지, 재시작, 모니터링을 도와줌.

ECS는 컨테이너화 된 앱을 대규모 실행

EKS는 위와 비슷한 기능을 하지만 ECS와는 다른 도구를 쓴다

Fargate도 서버리스 컴퓨팅 플랫폼이다.

> > > > > > 즉 정리하면

EC2는 운영체제에 대한 완전 관리 직접할 때
EC2 인스턴스는 최소한의 손실로 프로비저닝해서 AWS에서 가동하고, 실행할 수 있는 가상 머신입니다. EC2는 기본적인 웹 서버 실행에서부터 고성능 컴퓨팅 클러스터까지 다양한 사용 사례에 적합하게 사용하실 수 있습니다. EC2는 놀라울 정도로 유연하고 안정적이며 확장 가능하지만 사용 사례에 따라 대안을 찾아야 할 수도 있습니다.

EC2를 사용하면 사용자가 시간에 따라 인스턴스 플릿을 직접 설정하고 관리해야 합니다. 또한 EC2를 사용할 때는 사용자가 새 소프트웨어 패키지가 출시되면 인스턴스 패치를 내가 직접 책임지고 인스턴스의 규모 조정을 설정하고, 솔루션이 가용성이 높은 방식으로 호스팅되도록 아키텍처를 설계했는지 확인해야 합니다. 물론 온프레미스에서 호스팅할 때만큼 많은 관리가 필요하지는 않습니다. 그렇지만 여전히 이러한 관리 프로세스가 필요합니다.

AWS는 다양한 서버리스 컴퓨팅 옵션을 제공합니다. 서버리스는 애플리케이션을 호스팅하는 기본 인프라나 인스턴스를 마치 서버가 없는 것처럼 관리할 필요가 없다는 뜻입니다. 필요한 프로비저닝, 규모 조정, 고가용성 및 유지 관리와 관련한 모든 기본적인 환경 관리를 AWS가 대신 처리해줍니다. 그렇기 때문에 사용자는 애플리케이션만 신경쓰시면 되겠죠. 나머지는 AWS에서 자동으로 처리해주니까요.

\*\*\*서버리스 컴퓨팅의 또 다른 이점은 서버리스 애플리케이션을 자동으로 확장할 수 있는 유연성입니다. 서버리스 컴퓨팅은 처리량 및 메모리와 같은 소비 단위를 수정하여 애플리케이션의 용량을 조정할 수 있습니다.
AWS Lambda를 사용하는 경우 사용한 컴퓨팅 시간에 대해서만 비용을 지불합니다. 코드를 실행하는 동안에만 요금이 부과됩니다. 사실상 모든 유형의 애플리케이션 또는 백엔드 서비스 코드를 실행할 수 있으며 이를 관리할 필요는 전혀 없습니다.

람다는 단기 실행 함수 호스팅, 서비스 중심으로 관리할 때.
AWS Lambda는 이러한 서버리스 컴퓨팅 옵션 중에 대표적인 서비스라고 보실 수 있습니다. Lambda는 사용자가 코드를 Lambda 함수라는 곳에 업로드할 수 있게 도와주는 서비스입니다. 먼저 트리거를 구성하면 Lambda 함수에서 서비스가 트리거를 기다립니다. 그리고 트리거가 감지되면 코드가 관리형 환경에서 자동으로 실행됩니다. 자동으로 규모가 조절되고 가용성이 높으며 환경 내 모든 유지관리를 AWS가 수행하기 때문에 사용자가 걱정할 거리가 거의 없는 환경이죠. 트리거가 하나만 있든 천 개가 있든 Lambda는 수요에 맞게 함수의 규모를 조정합니다. 또한 Lambda는 코드를 15분 미만으로 실행하도록 설계되었습니다. 따라서 딥 러닝 같은 장기 실행 프로세스에는 적합하지 않죠.
웹 서비스의 백엔드나 요청 처리, 백엔드 비용 보고 처리 서비스처럼 각 처리를 완료하는 데 15분이 걸리지 않는 빠른 처리에 더 적합

도커 기반:
컨테이너는 애플리케이션의 코드와 종속성을 하나의 객체로 패키징하는 표준 방식을 제공합니다. 보안성, 신뢰성, 확장성 요구 사항이 매우 중요한 프로세스 및 워크플로에도 컨테이너를 사용합니다.

ECS, EKS를 선택 => 플랫폼은 EC2가 할지, Fargate가 할지 지정해준다.
두 서비스 모두 컨테이너 오케스트레이션 도구입니다. 그렇지만 자세히 짚기 전에 말씀을 드리면 여기서 이야기하는 컨테이너는 Docker 컨테이너를 의미합니다. Docker는 많은 곳에서 사용하는 플랫폼으로 운영 체제 수준에서의 가상화를 사용하여 컨테이너에 소프트웨어를 제공합니다. 여기서 말하는 컨테이너는 애플리케이션과 애플리케이션에서 실행해야 하는 모든 구성을 모아 놓은 코드 패키지가 다 컨테이너라고 할 수 있습니다.

ECS는 자체 컨테이너 오케스트레이션 소프트웨어를 관리하는 번거로움 없이도 컨테이너화된 애플리케이션을 대규모로 실행하는 데 도움이 되도록 설계되었습니다. 그리고 EKS는 비슷한 작업을 수행하지만 다른 도구와 다른 기능을 사용합니다. Amazon ECS와 Amazon EKS는 모두 EC2에서 실행할 수 있습니다. 하지만 기본 OS에 액세스할 필요가 없거나 EC2 인스턴스를 직접 컨트롤하면서 컨테이너를 호스팅하지 않아도 되는 경우에는 AWS Fargate라는 컴퓨팅 플랫폼을 사용하는 게 더 좋습니다.
Fargate는 ECS 또는 EKS용 서버리스 컴퓨팅 플랫폼입니다.
AWS에서 Docker 컨테이너 기반 워크로드를 실행하고 싶다면 먼저 오케스트레이션 도구를 선택해야 합니다. Amazon ECS나 Amazon EKS 중 무엇을 사용하고 싶은지를 먼저 정해야겠죠? 적절한 도구를 선택한 후에는 플랫폼을 선택해야 합니다. 컨테이너를 내가 관리하는 EC2 인스턴스에서 실행하고 싶은지 아니면 나 대신 AWS가 모두 관리해주는 AWS Fargate 같은 서버리스 환경에서 실행하고 싶은지를요.

- Amazon Elastic Container Service(Amazon ECS)
  Amazon Elastic Container Service(ECS)는 AWS에서 컨테이너식 애플리케이션을 실행하고 확장할 수 있는 확장성이 뛰어난 고성능 컨테이너 관리 시스템입니다.

Amazon ECS는 Docker 컨테이너를 지원합니다. Docker는 애플리케이션을 신속하게 구축, 테스트, 배포할 수 있는 소프트웨어 플랫폼입니다. AWS는 오픈 소스 Docker Community Edition 및 구독 기반 Docker Enterprise Edition의 사용을 지원합니다. Amazon ECS에서는 API 호출을 사용하여 Docker 지원 애플리케이션을 시작 및 중지할 수 있습니다.

- Amazon Elastic Kubernetes Service(Amazon EKS)
  Amazon Elastic Kubernetes Service(Amazon EKS)는 AWS에서 Kubernetes를 실행하는 데 사용할 수 있는 완전관리형 서비스입니다.

Kubernetes는 컨테이너식 애플리케이션을 대규모로 배포하고 관리하는 데 사용할 수 있는 오픈 소스 소프트웨어입니다. 자원자로 구성된 대규모 커뮤니티에서 Kubernetes를 유지 관리하며, AWS는 Kubernetes 커뮤니티와 적극적으로 협력합니다. Kubernetes 애플리케이션의 새로운 기능이 릴리스되면 Amazon EKS로 관리되는 애플리케이션에 이러한 업데이트를 손쉽게 적용할 수 있습니다.

- AWS Fargate
  AWS Fargate는 컨테이너용 서버리스 컴퓨팅 엔진입니다. Amazon ECS와 Amazon EKS에서 작동합니다.

AWS Fargate를 사용하는 경우 서버를 프로비저닝하거나 관리할 필요가 없습니다. AWS Fargate는 자동으로 서버 인프라를 관리합니다. 애플리케이션 혁신과 개발에 더 집중할 수 있으며, 컨테이너를 실행하는 데 필요한 리소스에 대해서만 비용을 지불하면 됩니다.