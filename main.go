package main

import (
	"fmt"

	bridge "github.com/josedelrio85/design_patterns/bridge"
	chainofresponsability "github.com/josedelrio85/design_patterns/chain_of_responsability"
)

func main() {
	call_chainOfResponsability()
	call_bridge()
}

func call_chainOfResponsability() {
	receiver := chainofresponsability.NewReceiverLink()
	conversor := chainofresponsability.NewConversorLink()
	sender := chainofresponsability.NewSenderLink()

	links := []chainofresponsability.Link{receiver, conversor, sender}
	chainofresponsability.NewChain(links)
}

func call_bridge() {
	printer1 := bridge.LaserPrinter{}
	printer2 := bridge.InkInjectPrinter{}

	mac := bridge.Mac{}
	mac.SetPrinter(&printer1)
	mac.Print()
	fmt.Println()

	windows := bridge.Windows{}
	windows.SetPrinter(&printer2)
	windows.Print()
	fmt.Println()
}
