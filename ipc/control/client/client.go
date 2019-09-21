// Package client provides a basic control client interface.
package client

import (
	unet "github.com/multiverse-os/ruby/ipc/unet"
	urpc "github.com/multiverse-os/ruby/ipc/urpc"
)

func ConnectTo(addr string) (*urpc.Client, error) {
	conn, err := unet.Connect(addr, false)
	if err != nil {
		return nil, err
	}
	return urpc.NewClient(conn), nil
}
