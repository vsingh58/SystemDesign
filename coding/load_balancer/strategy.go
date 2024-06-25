package main

import "fmt"

type BalancingStrategy interface {
	Init([]*APIServer)
	GetNextAPIServer(IncomingRequest) *APIServer
	RegisterAPIServer(*APIServer)
	PrintTopology()
}

type RRBalancingStrategy struct {
	Index      int
	ApiServers []*APIServer
}

/*type RandomBalancingStrategy struct {
	ApiServers []*APIServer
}

type HashBalancingStrategy struct {
	Index      int
	ApiServers []*APIServer
}*/

func (strategy *RRBalancingStrategy) Init(apiServers []*APIServer) {
	strategy.ApiServers = apiServers
	strategy.Index = 0
}

func (strategy *RRBalancingStrategy) GetNextAPIServer(req IncomingRequest) *APIServer {
	strategy.Index = (strategy.Index + 1) % len(strategy.ApiServers)
	return strategy.ApiServers[strategy.Index]
}

func (strategy *RRBalancingStrategy) RegisterAPIServer(apiServer *APIServer) {
	strategy.ApiServers = append(strategy.ApiServers, apiServer)
}

func (strategy *RRBalancingStrategy) PrintTopology() {
	for index, apiServer := range strategy.ApiServers {
		fmt.Printf("     [%d] %s", index, apiServer)
	}
}
