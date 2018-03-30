package main

import (
	"encoding/xml"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pallat/go-wsdl"
)

type envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    body     `xml:"Body"`
}

type body struct {
	ProtoType         *Request  `xml:"ProtoType,omitempty"`
	ProtoTypeResponse *Response `xml:"ProtoTypeResponse,omitempty"`
}

type Request struct {
	ID       string `xml:"ID"`
	RowID    string `xml:"RowID"`
	CustNo   string `xml:"CustNo"`
	SubrNo   string `xml:"SubrNo"`
	ListName string `xml:"ListName"`
}

type Response struct {
	Index   string `xml:"Index"`
	Type    string `xml:"Type"`
	ID      string `xml:"ID"`
	Version string `xml:"Version"`
	Created string `xml:"Created"`
}

type Fault struct {
	En   string
	TH   string
	Code string
}

func Service(c echo.Context) error {
	var e envelope
	err := c.Bind(&e)
	if err != nil {
		return c.XMLBlob(500, []byte(`<Error>`+err.Error()+"</Error>"))
	}

	res := envelope{
		Body: body{
			ProtoTypeResponse: &Response{
				Index:   e.Body.ProtoType.ID,
				Type:    "test",
				ID:      e.Body.ProtoType.RowID,
				Version: "1",
				Created: "now",
			},
		},
	}

	return c.XML(http.StatusOK, res)
}

func ServiceWSDL(c echo.Context) error {
	oper := wsdl.NewOperation(ProtoType{})
	wsdlString, err := wsdl.WSDL(oper)
	if err != nil {
		return c.XMLBlob(500, []byte(`<Error>`+err.Error()+"</Error>"))
	}
	return c.XMLBlob(200, []byte(wsdlString))
}

type ProtoType struct{}

func (ProtoType) Location() string {
	return "http://localhost:1323/service"
}
func (ProtoType) OperationName() string {
	return "ProtoType"
}

func (ProtoType) InputType() wsdl.Type {
	return ProtoTypeInput{}
}
func (ProtoType) OutputType() wsdl.Type {
	return ProtoTypeOutput{}
}
func (ProtoType) ErrorType() wsdl.Type {
	return ProtoTypeError{}
}

type ProtoTypeInput struct{}

func (ProtoTypeInput) MessageName() string {
	return "ProtoTypeInput"
}
func (ProtoTypeInput) TypeName() string {
	return "ProtoType"
}
func (ProtoTypeInput) SingleFields() []string {
	return []string{"ID", "RowID", "CustNo", "SubrNo", "ListName"}
}

type ProtoTypeOutput struct{}

func (ProtoTypeOutput) MessageName() string {
	return "ProtoTypeOutput"
}
func (ProtoTypeOutput) TypeName() string {
	return "ProtoTypeResponse"
}
func (ProtoTypeOutput) SingleFields() []string {
	return []string{"Index", "Type", "ID", "Version", "Created"}
}

type ProtoTypeError struct{}

func (ProtoTypeError) MessageName() string {
	return "ProtoTypeError"
}
func (ProtoTypeError) TypeName() string {
	return "ProtoTypeFault"
}
func (ProtoTypeError) SingleFields() []string {
	return []string{"En", "Th", "Code"}
}
