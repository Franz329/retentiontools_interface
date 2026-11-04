package retentiontools

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

type RetentionTools interface {
	Store(data StoreRequest) StoreResponse
	Status(request StatusRequest) StatusResponse
	Search(query SearchRequest) SearchResponse
}

type RetentionToolsRPC struct{ client *rpc.Client }

func (e *RetentionToolsRPC) Store(data StoreRequest) StoreResponse {
	var response StoreResponse

	err := e.client.Call("Plugin.Store", data, &response)
	if err != nil {
		panic(err)
	}

	return response
}

func (e *RetentionToolsRPC) Status(request StatusRequest) StatusResponse {
	var response StatusResponse

	err := e.client.Call("Plugin.Status", request, &response)
	if err != nil {
		panic(err)
	}

	return response
}

func (e *RetentionToolsRPC) Search(query SearchRequest) SearchResponse {
	var response SearchResponse

	err := e.client.Call("Plugin.Search", query, &response)
	if err != nil {
		panic(err)
	}

	return response
}

type RetentionToolsRPCServer struct {
	Impl RetentionTools
}

func (s *RetentionToolsRPCServer) Store(data StoreRequest, response *StoreResponse) error {
	*response = s.Impl.Store(data)
	return nil
}

func (s *RetentionToolsRPCServer) Status(request StatusRequest, response *StatusResponse) error {
	*response = s.Impl.Status(request)
	return nil
}

func (s *RetentionToolsRPCServer) Search(query SearchRequest, response *SearchResponse) error {
	*response = s.Impl.Search(query)
	return nil
}

type RetentionToolsPlugin struct {
	Impl RetentionTools
}

func (p *RetentionToolsPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &RetentionToolsRPCServer{Impl: p.Impl}, nil
}

func (RetentionToolsPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &RetentionToolsRPC{client: c}, nil
}
