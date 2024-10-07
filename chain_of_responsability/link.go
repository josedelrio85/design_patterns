package chainofresponsability

import (
	"fmt"
)

//go:generate mockgen -source=$GOFILE -package $GOPACKAGE -destination link_mock.go
type Link interface {
	Handle()
	SetNext(next Link) Link
}

type ReceiverLink struct {
	next Link
}

func NewReceiverLink() *ReceiverLink {
	return &ReceiverLink{}
}

func (r *ReceiverLink) Handle() {
	fmt.Println("ReceiverLink")
	if r.next != nil {
		r.next.Handle()
	}
}
func (r *ReceiverLink) SetNext(next Link) Link {
	r.next = next
	return r
}

type ConversorLink struct {
	next Link
}

func NewConversorLink() *ConversorLink {
	return &ConversorLink{}
}

func (c *ConversorLink) Handle() {
	fmt.Println("ConversorLink")
	if c.next != nil {
		c.next.Handle()
	}
}
func (c *ConversorLink) SetNext(next Link) Link {
	c.next = next
	return c
}

type SenderLink struct {
	next Link
}

func NewSenderLink() *SenderLink {
	return &SenderLink{}
}

func (s *SenderLink) Handle() {
	fmt.Println("SenderLink")
	if s.next != nil {
		s.next.Handle()
	}
}
func (s *SenderLink) SetNext(next Link) Link {
	s.next = next
	return next
}
