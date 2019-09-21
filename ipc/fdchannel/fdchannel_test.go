package fdchannel

import (
	"io/ioutil"
	"os"
	"sync"
	"syscall"
	"testing"
	"time"
)

func TestSendRecvFD(t *testing.T) {
	sendFile, err := ioutil.TempFile("", "go_ruby_")
	if err != nil {
		t.Fatalf("[fatal] failed to create temporary file: %v", err)
	}
	defer sendFile.Close()

	chanFDs, err := NewConnectedSockets()
	if err != nil {
		t.Fatalf("[fatal] failed to create fdchannel sockets: %v", err)
	}
	sendEP := NewEndpoint(chanFDs[0])
	defer sendEP.Destroy()
	recvEP := NewEndpoint(chanFDs[1])
	defer recvEP.Destroy()

	recvFD, err := recvEP.RecvFDNonblock()
	if err != syscall.EAGAIN && err != syscall.EWOULDBLOCK {
		// TODO: This is too long
		t.Errorf("[fatal] RecvFDNonblock before SendFD: got (%d, %v), wanted (<unspecified>, EAGAIN or EWOULDBLOCK", recvFD, err)
	}

	if err := sendEP.SendFD(int(sendFile.Fd())); err != nil {
		t.Fatalf("[fatal] SendFD failed: %v", err)
	}
	recvFD, err = recvEP.RecvFD()
	if err != nil {
		t.Fatalf("[fatal] RecvFD failed: %v", err)
	}
	recvFile := os.NewFile(uintptr(recvFD), "received file")
	defer recvFile.Close()

	sendInfo, err := sendFile.Stat()
	if err != nil {
		t.Fatalf("[fatal] failed to stat sent file: %v", err)
	}
	sendInfoSys := sendInfo.Sys()
	sendStat, ok := sendInfoSys.(*syscall.Stat_t)
	if !ok {
		t.Fatalf("[fatal] sent file's FileInfo is backed by unknown type %T", sendInfoSys)
	}

	recvInfo, err := recvFile.Stat()
	if err != nil {
		t.Fatalf("[fatal] failed to stat received file: %v", err)
	}
	recvInfoSys := recvInfo.Sys()
	recvStat, ok := recvInfoSys.(*syscall.Stat_t)
	if !ok {
		t.Fatalf("[fatal] received file's FileInfo is backed by unknown type %T", recvInfoSys)
	}

	if sendStat.Dev != recvStat.Dev || sendStat.Ino != recvStat.Ino {
		t.Errorf("[fatal] sent file (dev=%d, ino=%d) does not match received file (dev=%d, ino=%d)", sendStat.Dev, sendStat.Ino, recvStat.Dev, recvStat.Ino)
	}
}

func TestShutdownThenRecvFD(t *testing.T) {
	sendFile, err := ioutil.TempFile("", "go_ruby_")
	if err != nil {
		t.Fatalf("[fatal] failed to create temporary file: %v", err)
	}
	defer sendFile.Close()

	chanFDs, err := NewConnectedSockets()
	if err != nil {
		t.Fatalf("[fatal] failed to create fdchannel sockets: %v", err)
	}
	sendEP := NewEndpoint(chanFDs[0])
	defer sendEP.Destroy()
	recvEP := NewEndpoint(chanFDs[1])
	defer recvEP.Destroy()

	recvEP.Shutdown()
	if _, err := recvEP.RecvFD(); err == nil {
		t.Error("[fatal] RecvFD succeeded unexpectedly")
	}
}

func TestRecvFDThenShutdown(t *testing.T) {
	sendFile, err := ioutil.TempFile("", "go_ruby_")
	if err != nil {
		t.Fatalf("[fatal] failed to create temporary file: %v", err)
	}
	defer sendFile.Close()

	chanFDs, err := NewConnectedSockets()
	if err != nil {
		t.Fatalf("[fatal] failed to create fdchannel sockets: %v", err)
	}
	sendEP := NewEndpoint(chanFDs[0])
	defer sendEP.Destroy()
	recvEP := NewEndpoint(chanFDs[1])
	defer recvEP.Destroy()

	var receiverWG sync.WaitGroup
	receiverWG.Add(1)
	go func() {
		defer receiverWG.Done()
		if _, err := recvEP.RecvFD(); err == nil {
			t.Error("[fatal] RecvFD succeeded unexpectedly")
		}
	}()
	defer receiverWG.Wait()
	time.Sleep(time.Second) // to ensure recvEP.RecvFD() has blocked
	recvEP.Shutdown()
}
