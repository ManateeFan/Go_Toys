package queue

type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() (ret int) {
	ret = (*q)[0]
	*q = (*q)[1:]
	return
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
