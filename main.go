package main

import (
	chainofresponsability "github.com/josedelrio85/design_patterns/chain_of_responsability"
)

func main() {
	chainOfResponsability()
}

func chainOfResponsability() {
	receiver := chainofresponsability.NewReceiverLink()
	conversor := chainofresponsability.NewConversorLink()
	sender := chainofresponsability.NewSenderLink()

	links := []chainofresponsability.Link{receiver, conversor, sender}
	chainofresponsability.NewChain(links)
}
