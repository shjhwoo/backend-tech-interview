package method

import "fmt"

/*
OTP(One Time Password) 일회용 비밀번호 기능의 예시를 살펴봅시다. 사용자에 OTP를 전달하는 방법은 여러 가지가 있습니다​(SMS, 이메일 등). 그러나 방법​(SMS 또는 이메일)​과 관계없이 전체 OTP 프로세스는 같습니다.

임의의 n자리 숫자를 생성합니다.
이 숫자를 나중에 확인할 수 있도록 캐시에 저장합니다.
콘텐츠를 준비합니다.
알림을 보냅니다.
미래에 소개될 새로운 OTP 유형들은 아마도 여전히 위의 단계들을 거칠 것입니다.

우리는 이제 특정 작업의 단계들은 같지만 이러한 단계들의 구현이 다를 수 있는 상황에 부닥쳤습니다. 이것은 템플릿 메서드 패턴의 사용을 고려하기에 적절한 상황입니다.

먼저 고정된 수의 메서드들로 구성된 기초 템플릿 알고리즘을 정의합니다. 그것이 우리의 템플릿 메서드가 될 것입니다. 그런 다음 우리는 각 단계 메서드들을 구현하지만 템플릿 메서드는 변경되지 않은 상태로 둘 것입니다.
*/

//템플릿 메서드
type IOtp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
}

type Otp struct {
	iOtp IOtp
}

func (o *Otp) genAndSendOTP(otpLength int) error {
	otp := o.iOtp.genRandomOTP(otpLength)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	err := o.iOtp.sendNotification(message)
	if err != nil {
		return err
	}
	return nil
}

//구상 구현
type Sms struct {
	Otp
}

func (s *Sms) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("SMS: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *Sms) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *Sms) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *Sms) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}

//구상 구현2
type Email struct {
	Otp
}

func (s *Email) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("EMAIL: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *Email) saveOTPCache(otp string) {
	fmt.Printf("EMAIL: saving otp: %s to cache\n", otp)
}

func (s *Email) getMessage(otp string) string {
	return "EMAIL OTP for login is " + otp
}

func (s *Email) sendNotification(message string) error {
	fmt.Printf("EMAIL: sending email: %s\n", message)
	return nil
}

func main() {
	smsOTP := &Sms{}
	o := Otp{
		iOtp: smsOTP,
	}
	o.genAndSendOTP(4)

	fmt.Println("")
	emailOTP := &Email{}
	o = Otp{
		iOtp: emailOTP,
	}
	o.genAndSendOTP(4)
}
