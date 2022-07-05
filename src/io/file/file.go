package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filepath := "/root/fake.dat"
	fi, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)

	chunks := make([]byte, 0)
	buf := make([]byte, 1024) //一次读取多少个字节
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		fmt.Println(string(buf[:n]))
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	fmt.Println(string(chunks))
}
