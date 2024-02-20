package main

import "fmt"

/*
맥과 윈도우 유형의 두 컴퓨터 유형이 있다고 가정해 봅시다. 또한 엡손과 HP 유형의 두 프린터 유형이 있다고 가정해 봅시다. 컴퓨터와 프린터 유형들은 모든 조합으로 서로 작동해야 합니다. 클라이언트는 프린터와 컴퓨터를 연결하는 세부 사항에 대하여 걱정하고 싶어 하지 않습니다.

또 우리는 새 프린터를 도입할 때 코드가 기하급수적으로 증가하는 것을 원하지 않습니다. 따라서 2*2 조합에 대해 4개의 구조체를 만드는 대신 우리는 두 개의 계층구조를 만듭니다.

추상화 계층구조: 이것은 우리의 컴퓨터가 될 것입니다.
구현 계층구조: 이것은 우리의 프린터가 될 것입니다.
이 두 계층구조는 브리지를 통해 서로 통신하며, 여기서 추상화​(컴퓨터)​에는 구현​(프린터)​에 대한 참조가 포함됩니다. 추상화와 구현은 서로 영향을 미치지 않고 독립적으로 개발될 수 있습니다.

*/

func main() {

	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}

type Computer interface {
	Print()
	SetPrinter(Printer)
}

type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

//브리지
func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

//브리지
func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

type Printer interface {
	PrintFile()
}

type Epson struct {
}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

type Hp struct {
}

func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}
