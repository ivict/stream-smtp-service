package stream

import (
	"strings"
)

type CircularBuffer struct {
	buffer       []string
	size         int
	readPointer  int
	writePointer int
	count        int
}

func NewCircularBuffer(size int) *CircularBuffer {
	return &CircularBuffer{
		buffer:       make([]string, size),
		size:         size,
		readPointer:  0,
		writePointer: 0,
		count:        0,
	}
}

func (c *CircularBuffer) String() string {
	items := make([]string, 0, 10)
	items = append(items, c.buffer[c.readPointer:]...)
	items = append(items, c.buffer[:c.readPointer]...)
	content := strings.Join(items, ",")
	if len(content) > 10 {
		return content[:10] + "..."
	}
	return content
}

func (c *CircularBuffer) Put(data ...string) string {
	for _, r := range data {
		c.push(r)
	}
	return data[0]
}

func (c *CircularBuffer) push(data string) {
	if c.count == c.size {
		c.readPointer = (c.readPointer + 1) % c.size
	} else if c.count < c.size {
		c.count++
	}
	c.buffer[c.writePointer] = data
	c.writePointer = (c.writePointer + 1) % c.size
}

func (c *CircularBuffer) Has(element string) bool {
	for _, item := range c.buffer[c.readPointer:] {
		if item == element {
			return true
		}
	}
	for _, item := range c.buffer[:c.readPointer] {
		if item == element {
			return true
		}
	}
	return false
}
