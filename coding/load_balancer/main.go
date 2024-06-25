package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

/*
load balancer recevies requests from clients
and forwards/proxies the request/response to/from api sever
*/
func main() {
	InitLoadBalancer()
	go lb.Run()

	process_cli()

}

const CMD_Exit = "exit"

func process_cli() {
	for {
		var command string
		fmt.Print(">>> ")
		fmt.Scanf("%s", &command)
		switch command {
		case CMD_Exit:
			lb.events <- Event{EventName: CMD_Exit}
		default:
			fmt.Printf("List of command available:%s \n", CMD_Exit)
		}
	}
}

type APIServer struct {
	Host        string
	Port        int
	isHealthy   bool
	NumRequests int
}

func (apiServer *APIServer) toString() string {
	return fmt.Sprintf("%s:%d", apiServer.Host, apiServer.Port)
}

type Event struct {
	EventName string
	Data      interface{}
}

type IncomingRequest struct {
	srcConn net.Conn
	reqId   string
}

type LoadBalancer struct {
	apiServers []*APIServer
	events     chan Event
	strategy   BalancingStrategy
}

var lb *LoadBalancer

func NewRRBalancingStrategy(apiServers []*APIServer) *RRBalancingStrategy {
	strategy := new(RRBalancingStrategy)
	strategy.Init(apiServers)
	return strategy
}

func InitLoadBalancer() {
	api_servers := []*APIServer{
		&APIServer{Host: "localhost", Port: 8081, isHealthy: true},
		&APIServer{Host: "localhost", Port: 8082, isHealthy: true},
		&APIServer{Host: "localhost", Port: 8083, isHealthy: true},
	}

	lb = &LoadBalancer{
		apiServers: api_servers,
		events:     make(chan Event),
		strategy:   NewRRBalancingStrategy(api_servers),
	}
}

func (lb *LoadBalancer) proxy(request IncomingRequest) {
	apiServer := lb.strategy.GetNextAPIServer(request)

	log.Printf("in-request: %s, out-request: $s", request.reqId, apiServer.toString())
	// set up connection to api server
	apiServerConnection, err := net.Dial("tcp", fmt.Sprintf("%s:%d", apiServer.Host, apiServer.Port))
	if err != nil {
		log.Printf("Error connecting to the backend: %s", err.Error())

		//relay error message back to client
		request.srcConn.Write([]byte("Api server not available."))
		//close connection
		request.srcConn.Close()
		panic(err)
	}

	apiServer.NumRequests++

	//set up layer 4 connection (network layer)
	// COPY bytes from client to api server
	go io.Copy(apiServerConnection, request.srcConn)

	// COPY bytes from api server client
	go io.Copy(request.srcConn, apiServerConnection)

}

func (lb *LoadBalancer) Run() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("Load balancer listening on port 9090...")

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Printf("Unable to accept connection: $s", err.Error)
			panic(err)
		}
		fmt.Println("New connection request received...")
		go lb.proxy(IncomingRequest{
			srcConn: connection,
			reqId:   uuid.NewString(),
		})
	}
}

func startApiServer1() {
	router := gin.Default()
	router.Run("localhost:8081")
}

func startApiServer2() {
	router := gin.Default()
	router.Run("localhost:8082")
}

func startApiServer3() {
	router := gin.Default()
	router.Run("localhost:8083")
}
