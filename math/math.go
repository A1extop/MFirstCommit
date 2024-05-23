package main

import "fmt"

type Printer interface {
	Print(string)
}

type Scanner interface {
	Scan() string
}

type MultiFunctionDevice interface {
	Printer
	Scanner
}

type SimplePrinter struct{}

func (p *SimplePrinter) Print(doc string) {
	fmt.Println("Printing:", doc)
}

type SimpleScanner struct{}

func (s *SimpleScanner) Scan() string {
	return "Scanning document"
}

type MultiFunctionDeviceImpl struct {
	printer Printer
	scanner Scanner
}

func (m *MultiFunctionDeviceImpl) Print(doc string) {
	m.printer.Print(doc)
}

func (m *MultiFunctionDeviceImpl) Scan() string {
	return m.scanner.Scan()
}

func main() {
	simplePrinter := &SimplePrinter{}
	simpleScanner := &SimpleScanner{}

	multiFunctionDevice := &MultiFunctionDeviceImpl{printer: simplePrinter, scanner: simpleScanner}
	multiFunctionDevice.Print("Hello, world!")
	fmt.Println(multiFunctionDevice.Scan())
}
