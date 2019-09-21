package urpc

import (
	"errors"
	"os"
	"testing"

	"github.com/multiverse-os/ruby/ipc/unet"
)

type test struct {
}

type testArg struct {
	StringArg string
	IntArg    int
	FilePayload
}

type testResult struct {
	StringResult string
	IntResult    int
	FilePayload
}

func (t test) Func(a *testArg, r *testResult) error {
	r.StringResult = a.StringArg
	r.IntResult = a.IntArg
	return nil
}

func (t test) Err(a *testArg, r *testResult) error {
	return errors.New("test error")
}

func (t test) FailNoFile(a *testArg, r *testResult) error {
	if a.Files == nil {
		return errors.New("no file found")
	}

	return nil
}

func (t test) SendFile(a *testArg, r *testResult) error {
	r.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
	return nil
}

func (t test) TooManyFiles(a *testArg, r *testResult) error {
	for i := 0; i <= maxFiles; i++ {
		r.Files = append(r.Files, os.Stdin)
	}
	return nil
}

func startServer(socket *unet.Socket) {
	s := NewServer()
	s.Register(test{})
	s.StartHandling(socket)
}

func testClient() (*Client, error) {
	serverSock, clientSock, err := unet.SocketPair(false)
	if err != nil {
		return nil, err
	}
	startServer(serverSock)
	return NewClient(clientSock), nil
}

func TestCall(t *testing.T) {
	c, err := testClient()
	if err != nil {
		t.Fatalf("error creating test client: %v", err)
	}
	defer c.Close()

	var r testResult
	if err := c.Call("test.Func", &testArg{}, &r); err != nil {
		t.Errorf("[error] basic call failed: %v", err)
	} else if r.StringResult != "" || r.IntResult != 0 {
		t.Errorf("[error] unexpected result, got %v expected zero value", r)
	}
	if err := c.Call("test.Func", &testArg{StringArg: "hello"}, &r); err != nil {
		t.Errorf("[error] basic call failed: %v", err)
	} else if r.StringResult != "hello" {
		t.Errorf("[error] unexpected result, got %v expected hello", r.StringResult)
	}
	if err := c.Call("test.Func", &testArg{IntArg: 1}, &r); err != nil {
		t.Errorf("basic call failed: %v", err)
	} else if r.IntResult != 1 {
		t.Errorf("[error] unexpected result, got %v expected 1", r.IntResult)
	}
}

func TestUnknownMethod(t *testing.T) {
	c, err := testClient()
	if err != nil {
		t.Fatalf("error creating test client: %v", err)
	}
	defer c.Close()

	var r testResult
	if err := c.Call("test.Unknown", &testArg{}, &r); err == nil {
		t.Errorf("[error] expected non-nil err, got nil")
	} else if err.Error() != ErrUnknownMethod.Error() {
		t.Errorf("[error] expected test error, got %v", err)
	}
}

func TestErr(t *testing.T) {
	c, err := testClient()
	if err != nil {
		t.Fatalf("[fatal] error creating test client: %v", err)
	}
	defer c.Close()

	var r testResult
	if err := c.Call("test.Err", &testArg{}, &r); err == nil {
		t.Errorf("[error] expected non-nil err, got nil")
	} else if err.Error() != "test error" {
		t.Errorf("[error] expected test error, got %v", err)
	}
}

func TestSendFile(t *testing.T) {
	c, err := testClient()
	if err != nil {
		t.Fatalf("[error] error creating test client: %v", err)
	}
	defer c.Close()

	var r testResult
	if err := c.Call("test.FailNoFile", &testArg{}, &r); err == nil {
		t.Errorf("[error] expected non-nil err, got nil")
	}
	if err := c.Call("test.FailNoFile", &testArg{FilePayload: FilePayload{Files: []*os.File{os.Stdin, os.Stdout, os.Stdin}}}, &r); err != nil {
		t.Errorf("[error] expected nil err, got %v", err)
	}
}

func TestRecvFile(t *testing.T) {
	c, err := testClient()
	if err != nil {
		t.Fatalf("error creating test client: %v", err)
	}
	defer c.Close()

	var r testResult
	if err := c.Call("test.SendFile", &testArg{}, &r); err != nil {
		t.Errorf("[error] expected nil err, got %v", err)
	}
	if r.Files == nil {
		t.Errorf("[error] expected file, got nil")
	}
}

func TestShutdown(t *testing.T) {
	serverSock, clientSock, err := unet.SocketPair(false)
	if err != nil {
		t.Fatalf("[error] error creating test client: %v", err)
	}
	clientSock.Close()

	s := NewServer()
	if err := s.Handle(serverSock); err == nil {
		t.Errorf("expected non-nil err, got nil")
	}
}

func TestTooManyFiles(t *testing.T) {
	c, err := testClient()
	if err != nil {
		t.Fatalf("[error] error creating test client: %v", err)
	}
	defer c.Close()

	var r testResult
	var a testArg
	for i := 0; i <= maxFiles; i++ {
		a.Files = append(a.Files, os.Stdin)
	}

	// Client-side error.
	if err := c.Call("test.Func", &a, &r); err != ErrTooManyFiles {
		t.Errorf("[error] expected ErrTooManyFiles, got %v", err)
	}

	// Server-side error.
	if err := c.Call("test.TooManyFiles", &testArg{}, &r); err == nil {
		t.Errorf("[error] expected non-nil err, got nil")
	} else if err.Error() != "too many files" {
		t.Errorf("[error] expected too many files, got %v", err.Error())
	}
}
