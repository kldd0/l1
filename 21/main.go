package main

import "fmt"

type Service interface {
	ProcessJSON()
}

type Client struct{}

func (c *Client) SendJSONDataInMainService(srv Service) {
	fmt.Println("client sends JSON data in main service")
	srv.ProcessJSON()
}

type ServiceWithMainBusinessLogic struct{}

func (s *ServiceWithMainBusinessLogic) ProcessJSON() {
	fmt.Println("JSON Data is processed by MainService")
}

// external API that return data in XML format
type ServiceThatReturnDataInXML struct{}

func (s ServiceThatReturnDataInXML) GetDataInXML() {
	fmt.Println("ExternalService returns XML data")
}

// Adapter that converts XML into JSON
type XMLServiceAdapter struct {
	extSrv *ServiceThatReturnDataInXML
}

func (xs *XMLServiceAdapter) ProcessJSON() {
    // converting XML data...
	fmt.Println("Adapter converts XML to JSON")
    xs.extSrv.GetDataInXML()
	fmt.Println("JSON data is processed by main service")
}

func main() {
	client := &Client{}
	mainSrv := &ServiceWithMainBusinessLogic{}

	client.SendJSONDataInMainService(mainSrv)

	xmlSrv := &ServiceThatReturnDataInXML{}
	xmlSrvAdapter := &XMLServiceAdapter{
		extSrv: xmlSrv,
	}

	client.SendJSONDataInMainService(xmlSrvAdapter)
}
