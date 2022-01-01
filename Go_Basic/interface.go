package main

import (
	"bytes"
	"fmt"
)

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct {
}

func (cw *ConsoleWriter) Write(data []byte) (int, error) {
	res, err := fmt.Println(string(data))
	return res, err
}

// interface without struct

type Incrementer interface {
	Increment() int
}

type Intcounter int

// because no struct -> diff memory -> must use pointers
func (ic *Intcounter) Increment() int {
	*ic++
	return int(*ic)
}

// Interface nested
type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err = bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferWriterCloser() *BufferWriterCloser {
	return &BufferWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

func main() {
	var w Writer = &ConsoleWriter{}
	w.Write([]byte("Hello Long"))

	myInt := Intcounter(0)
	var inc Incrementer = &myInt
	for i := 1; i <= 5; i++ {
		fmt.Println(inc.Increment())
	}

	var wc WriterCloser = NewBufferWriterCloser()
	wc.Write([]byte("Hello every one, welcome to GO"))
	wc.Close()

	// empty interface
	var obj interface{} = NewBufferWriterCloser()
	object, err := obj.(WriterCloser)
	if !err {
		fmt.Println("Error!!!")
	}
	object.Write([]byte("Hello every one, welcome to GO"))
	object.Close()

	// interface is parent data type
	var objParent interface{} = 1
	switch objParent.(type) {
	case int:
		fmt.Println("Int")
	case float64:
		fmt.Println("Float64")
	default:
		fmt.Println("Not exist")
	}
}
