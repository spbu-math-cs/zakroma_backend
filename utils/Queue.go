package utils

type Queue struct {
	Data    []interface{}
	Head    int
	Tail    int
	Size    int
	MaxSize int
}

func CreateQueue(maxSize int) *Queue {
	queue := Queue{Data: make([]interface{}, maxSize), Head: 0, Tail: 0, MaxSize: maxSize, Size: 0}
	return &queue
}

func (queue *Queue) Push(value interface{}) (interface{}, bool) {
	var removed interface{}
	flag := false
	if queue.Size == queue.MaxSize {
		removed = queue.Pop()
		flag = true
	}
	queue.Data[queue.Tail] = value
	queue.Tail += 1
	queue.Size += 1
	return removed, flag
}

func (queue *Queue) Pop() interface{} {
	value := queue.Data[queue.Head]
	queue.Head += 1
	queue.Size -= 1
	return value
}
