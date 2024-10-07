package bridge

import "fmt"

type computer interface {
	Print()
	SetPrinter(printer)
}

type printer interface {
	PrintFile()
}

type Mac struct {
	printer printer
}

func (m *Mac) Print() {
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p printer) {
	m.printer = p
}

type Windows struct {
	printer printer
}

func (w *Windows) Print() {
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p printer) {
	w.printer = p
}

type LaserPrinter struct{}

func (lp *LaserPrinter) PrintFile() {
	fmt.Println("laserPrinter")
}

type InkInjectPrinter struct{}

func (ip *InkInjectPrinter) PrintFile() {
	fmt.Println("inkInjectPrinter")
}
