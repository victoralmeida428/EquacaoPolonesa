package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := New()
	queue.Add("a")
	if queue.Front() != "a"{ 
		t.Error("Error on Add")
	}
	queue.Remove()
	fmt.Println(queue.Front(), queue.GetData())
	queue.Remove()
	fmt.Println(queue.Front(), queue.GetData())
	queue.Remove()
	fmt.Println(queue.Front(), queue.GetData())
	t.Error("error")
}