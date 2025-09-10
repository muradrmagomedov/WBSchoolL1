package main

type Customer interface {
	ContactCustomer()
}

type CRMClient interface {
	Call()
	SendSMS()
}

func NewCRMClient() CRMClient {
	return CRMClient{}
}

type Adapter struct {
	CRMClient CRMClient
}

func (a *Adapter) ContactCustomer() {
	a.CRMClient.SendSMS()
}

func main() {
	var customer Customer
	var adapter Adapter
	adapter.CRMClient := NewCRMClient()
}
