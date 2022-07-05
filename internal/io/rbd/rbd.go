package main

import (
	"fmt"
	"time"

	"github.com/ceph/go-ceph/rados"
	"github.com/ceph/go-ceph/rbd"
)

func main() {
	conn, err := rados.NewConn()
	if err != nil {
		fmt.Print("Error when create connection.")
		return
	}
	defer conn.Shutdown()

	conn.ReadDefaultConfigFile()
	conn.Connect()

	ioctx, err := conn.OpenIOContext("rbd")
	if err != nil {
		fmt.Print("Error when open io context.")
		return
	}
	defer ioctx.Destroy()

	image := rbd.GetImage(ioctx, "restic")
	image.Open()

	// write_buffer := make([]byte, 5*1024*1024)
	// rand.Read(write_buffer)
	// image.Write(write_buffer)
	// image.Flush()

	go func() {
		for {
			read_buffer := make([]byte, 1024)
			len, err := image.Read(read_buffer)
			if err != nil {
				fmt.Print("Error when read image.")
				return
			}
			fmt.Printf("Read len %d.\n", len)
		}
	}()

	time.Sleep(time.Minute * 10)

}
