package handlers

import "testing"

type FakeUserHandlerStruct struct {
	functionCalls int
}

func (user *FakeUserHandlerStruct) Register() {
	user.functionCalls++
}

func TestRegister(t *testing.T) {
	user := &FakeUserHandlerStruct{}
	RegisterHandler()
}
