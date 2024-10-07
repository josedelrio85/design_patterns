package chainofresponsability

type chain struct {
	Links []Link
}

func NewChain(links []Link) *chain {
	lastElement := len(links) - 1
	for i, h := range links {
		if i < lastElement {
			nextHandler := links[i+1]
			h.SetNext(nextHandler)
		}
	}
	if links[0] != nil {
		links[0].Handle()
	}
	return &chain{
		Links: links,
	}
}
