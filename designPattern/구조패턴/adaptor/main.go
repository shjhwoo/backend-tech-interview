package main

import "fmt"

//클라이언트 코드: 프로그램의 기존 비즈니스 로직을 포함
type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}

//클라이언트 인터페이스: 다른 클래스들이 클라이언트 코드와 공동 작업할 수 있도록 따라야 하는 프로토콜
type Computer interface {
	InsertIntoLightningPort()
}

//서비스: 일반적으로 타사 또는 레거시의 유용한 클래스
type Mac struct {
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

//알려지지 않은 서비스
type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

//어댑터: 클라이언트와 서비스 양쪽에서 작동할 수 있는 클래스로, 서비스 객체를 래핑하는 동안 클라이언트 인터페이스를 구현
type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}

func main() {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
