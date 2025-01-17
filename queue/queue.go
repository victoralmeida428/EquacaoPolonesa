package queue

type IQueue interface {
	Add(interface{})
	Remove()
	Front() interface{}
	GetData() []interface{}
}

type Queue struct {
	data []interface{}
}

func (q *Queue) Add(data interface{}) {
	q.data = append(q.data, data)

}

func (q *Queue) Remove() {
	if len(q.data) > 0 {
		q.data = q.data[1:]
	}
}

func (q *Queue) Front() interface{} {
	if len(q.data)>0 {
		return q.data[0]
	}
	return nil
}

func (q *Queue) GetData() []interface{} {
	return q.data
}

func New() IQueue {
	return &Queue{}
}
