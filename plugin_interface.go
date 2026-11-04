package retentionservice

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

//inputs and outputs of the functions

type StoreRequest struct {
	Id       string
	FileName string
	MetaData string
	Payload  []byte
	Hash     string
}

type StoreResponse struct {
	Id    string
	Error int
}

type StatusRequest struct {
	Id string
}

type StatusResponse struct {
	Status string
	Error  int
}

type SearchRequest struct {
	Id       string
	FileName string
	MetaData string
}

type SearchResponse struct {
	Payload []byte
	Hash    string
	Error   int
}

//interface

type RetentionService interface {
	Store(data StoreRequest) StoreResponse
	Status(request StatusRequest) StatusResponse
	Search(query SearchRequest) SearchResponse
}

type RetentionServiceRPC struct{ client *rpc.Client }

func (e *RetentionServiceRPC) Store(data StoreRequest) StoreResponse {
	var response StoreResponse

	err := e.client.Call("Plugin.Store", data, &response)
	if err != nil {
		panic(err)
	}

	return response
}

func (e *RetentionServiceRPC) Status(request StatusRequest) StatusResponse {
	var response StatusResponse

	err := e.client.Call("Plugin.Status", request, &response)
	if err != nil {
		panic(err)
	}

	return response
}

func (e *RetentionServiceRPC) Search(query SearchRequest) SearchResponse {
	var response SearchResponse

	err := e.client.Call("Plugin.Search", query, &response)
	if err != nil {
		panic(err)
	}

	return response
}

type RetentionServiceRPCServer struct {
	Impl RetentionService
}

func (s *RetentionServiceRPCServer) Store(data StoreRequest, response *StoreResponse) error {
	*response = s.Impl.Store(data)
	return nil
}

func (s *RetentionServiceRPCServer) Status(request StatusRequest, response *StatusResponse) error {
	*response = s.Impl.Status(request)
	return nil
}

func (s *RetentionServiceRPCServer) Search(query SearchRequest, response *SearchResponse) error {
	*response = s.Impl.Search(query)
	return nil
}

type RetentionServicePlugin struct {
	Impl RetentionService
}

func (p *RetentionServicePlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &RetentionServiceRPCServer{Impl: p.Impl}, nil
}

func (RetentionServicePlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &RetentionServiceRPC{client: c}, nil
}
