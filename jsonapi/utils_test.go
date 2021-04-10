package jsonapi

import (
	"fmt"
	"strings"
	"testing"
)

type Command struct {
	Name      string
	PayerType string
	PayerID   string
	PayeeType string
	PayeeID   string
	OrderID   string
}

func TestUnmarshalCommand(t *testing.T) {
	data := `{"data":{"type":"items","attributes":{"tenantId":"unitTests","name":"item1","quantity":2,"ActivityTime":1492074177,"BookingTime":1492074177,"unitPrice":{"amount":1000,"currency":"USD"}},"relationships":{"payer":{"data":{"type":"operators","id":"o1"}},"payee":{"data":{"type":"distributors","id":"d1"}},"order":{"data":{"type":"orders","id":"11603"}}}}}`
	io := strings.NewReader(data)
	c := new(Command)
	e := UnmarshalCommand(io, c)
	if e != nil {
		panic(e)
	}

	if c.PayerType != "operators" {
		t.Errorf("jsonapi: expect operators got %s", c.PayerType)
	}
	if c.PayerID != "o1" {
		t.Errorf("jsonapi: expect o1 got %s", c.PayerID)
	}

	if c.PayeeType != "distributors" {
		t.Errorf("jsonapi: expect distributors got %s", c.PayeeType)
	}
	if c.PayeeID != "d1" {
		t.Errorf("jsonapi: expect d1 got %s", c.PayeeID)
	}
	fmt.Printf("%#v", c)
}
