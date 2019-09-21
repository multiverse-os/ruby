// Package server provides a basic control server interface.
// Note that no objects are registered by default. Users must provide their own
// implementations of the control interface.
package server

import (
	"os"
	"sync"

	"github.com/multiverse-os/ruby/ipc/unet"
	"github.com/multiverse-os/ruby/ipc/urpc"
)

var curUID = os.Getuid()

type Server struct {
	socket *unet.ServerSocket
	server *urpc.Server
	wg     sync.WaitGroup
}

func New(socket *unet.ServerSocket) *Server {
	return &Server{
		socket: socket,
		server: urpc.NewServer(),
	}
}

func (s *Server) FD() int { return s.socket.FD() }
func (s *Server) Wait()   { s.wg.Wait() }

func (s *Server) Stop() {
	s.socket.Close()
	s.wg.Wait()
	s.server.Stop()
}

func (s *Server) StartServing() error {
	if err := s.socket.Listen(); err != nil {
		return err
	}
	s.wg.Add(1)
	go func() { // S/R-SAFE: does not impact state directly.
		s.serve()
		s.wg.Done()
	}()

	return nil
}

func (s *Server) serve() {
	for {
		conn, err := s.socket.Accept()
		if err != nil {
			return
		}

		ucred, err := conn.GetPeerCred()
		if err != nil {
			fmt.Printf("[warn] control couldn't get credentials: %s", err.Error())
			conn.Close()
			continue
		}

		if int(ucred.Uid) != curUID && ucred.Uid != 0 {
			fmt.Printf("[warn] ontrol auth failure: other UID = %d, current UID = %d", ucred.Uid, curUID)
			conn.Close()
			continue
		}
		s.server.StartHandling(conn)
	}
}

// Register registers a specific control interface with the server.
func (s *Server) Register(obj interface{}) {
	s.server.Register(obj)
}

// CreateFromFD creates a new control bound to the given 'fd'. It has no
// registered interfaces and will not start serving until StartServing is
// called.
func CreateFromFD(fd int) (*Server, error) {
	socket, err := unet.NewServerSocket(fd)
	if err != nil {
		return nil, err
	}
	return New(socket), nil
}

// Create creates a new control server with an abstract unix socket
// with the given address, which must must be unique and a valid
// abstract socket name.
func Create(addr string) (*Server, error) {
	socket, err := unet.Bind(addr, false)
	if err != nil {
		return nil, err
	}
	return New(socket), nil
}

// CreateSocket creates a socket that can be used with control server,
// but doesn't start control server.  'addr' must be a valid and unique
// abstract socket name.  Returns socket's FD, -1 in case of error.
func CreateSocket(addr string) (int, error) {
	socket, err := unet.Bind(addr, false)
	if err != nil {
		return -1, err
	}
	return socket.Release()
}
