package conservservice

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

//

type DataToStore struct {
	Id      string
	Content string
}

//interface

type ConservService interface {
	Store(data DataToStore) string
	Status() string
	Search() string
}

type ConservServiceRPC struct{ client *rpc.Client }

func (e *ConservServiceRPC) Store(data DataToStore) string {
	var response string

	err := e.client.Call("Plugin.Store", data, &response)
	if err != nil {
		panic(err)
	}

	return response
}

func (e *ConservServiceRPC) Status() string {
	var response string

	err := e.client.Call("Plugin.Status", new(interface{}), &response)
	if err != nil {
		panic(err)
	}

	return response
}

func (e *ConservServiceRPC) Search() string {
	var response string

	err := e.client.Call("Plugin.Search", new(interface{}), &response)
	if err != nil {
		panic(err)
	}

	return response
}

type ConservServiceRPCServer struct {
	Impl ConservService
}

func (s *ConservServiceRPCServer) Store(data DataToStore, response *string) error {
	*response = s.Impl.Store(data)
	return nil
}

func (s *ConservServiceRPCServer) Status(args interface{}, response *string) error {
	*response = s.Impl.Status()
	return nil
}

func (s *ConservServiceRPCServer) Search(args interface{}, response *string) error {
	*response = s.Impl.Search()
	return nil
}

type ConservServicePlugin struct {
	Impl ConservService
}

func (p *ConservServicePlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &ConservServiceRPCServer{Impl: p.Impl}, nil
}

func (ConservServicePlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ConservServiceRPC{client: c}, nil
}
