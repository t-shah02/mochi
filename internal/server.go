package internal

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/t-shah02/mochi/internal/persistence"
)

type MochiServer struct {
	port          uint16
	memoryManager *persistence.MemoryManager
	diskManager   *persistence.DiskManager
}

func NewMochiServer() *MochiServer {
	rawPort, found := os.LookupEnv(PORT_ENV_NAME)
	if !found {
		log.Fatalf("required env: %s was not provided", PORT_ENV_NAME)
	}

	port, err := strconv.Atoi(rawPort)
	if err != nil {
		log.Fatalf("%s was either unspecified or not in a valid format", PORT_ENV_NAME)
	}

	dataVolumePath, found := os.LookupEnv(DATA_VOLUME_PATH_ENV_NAME)
	if !found {
		log.Fatalf("required env: %s was not provided", DATA_VOLUME_PATH_ENV_NAME)
	}

	memoryManager := persistence.NewMemoryManager()
	diskManager := persistence.NewDiskManager(dataVolumePath)

	return &MochiServer{
		port:          uint16(port),
		memoryManager: memoryManager,
		diskManager:   diskManager,
	}
}

func (server *MochiServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	connectionReadDeadline := time.Now().Add(time.Duration(BUFFER_READ_TIME_LIMIT) * time.Second)
	conn.SetReadDeadline(connectionReadDeadline)

	connBuffer := make([]byte, MAXIMUM_BUFFER_READ_SIZE)
	n, err := conn.Read(connBuffer)
	if err != nil {
		log.Println(err)
		return
	}

	if n > MAXIMUM_BUFFER_READ_SIZE {
		connectionWriteDeadline := time.Now().Add(time.Duration(BUFFER_WRITE_TIME_LIMIT) * time.Second)
		conn.SetWriteDeadline(connectionWriteDeadline)

		conn.Write([]byte("the provided request buffer exceeded the maximum size"))
		return
	}

	data := string(connBuffer)
	log.Println(data)

	connectionWriteDeadline := time.Now().Add(time.Duration(BUFFER_WRITE_TIME_LIMIT) * time.Second)
	conn.SetWriteDeadline(connectionWriteDeadline)

	conn.Write([]byte("got it"))
}

func (server *MochiServer) Init() {
	server.diskManager.Init()

}

func (server *MochiServer) Serve() {
	sAddr := fmt.Sprintf(":%d", server.port)
	ln, err := net.Listen(SERVER_PROTOCOL, sAddr)
	if err != nil {
		log.Fatal("server was unable to bind to the specified port")
	}

	log.Printf("server is now listening to connections at %s", ln.Addr())

	for {
		incomingConn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}

		log.Printf("received an incoming connection from: %s", incomingConn.RemoteAddr())
		go server.handleConnection(incomingConn)
	}
}
