package main

import (
	"bytes"
	"fmt"
)

type Pipeline interface {
	do()
}

type Pipeline1 struct{}
type Pipeline2 struct{}

func (pipeline Pipeline1) do(in chan bytes.Buffer, out chan bytes.Buffer) {
	var buffio bytes.Buffer
	buffio = <-in
	buffio.WriteString("STAGE1 : ")
	close(in)
	out <- buffio
}
func (pipeline Pipeline2) do(in chan bytes.Buffer, out chan bytes.Buffer) {
	var buffio bytes.Buffer
	buffio = <-in
	buffio.WriteString("STAGE2 : ")
	close(in)
	out <- buffio
}

func main() {
	var buffer bytes.Buffer
	buffer.WriteString("START : ")

	var chans [3]chan bytes.Buffer
	for i := range chans {
		chans[i] = make(chan bytes.Buffer)
	}

	p1 := Pipeline1{}
	p2 := Pipeline2{}

	go p1.do(chans[0], chans[1])
	go p2.do(chans[1], chans[2])

	chans[0] <- buffer

	var buffio bytes.Buffer
	buffio = <-chans[2]
	buffio.WriteString("END")

	fmt.Println(buffio.String())
}
