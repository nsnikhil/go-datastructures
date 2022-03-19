package queue

import (
	"errors"
	"github.com/nsnikhil/go-datastructures/internal"
	"github.com/nsnikhil/go-datastructures/list"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateNewBlockingQueue(t *testing.T) {
	actualResult := NewBlockingQueue[int]()

	lq := NewLinkedQueue[int]()

	expectedResult := &BlockingQueue[int]{lq}

	assert.Equal(t, expectedResult, actualResult)
}

func TestBlockingQueueRemove(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() ([]int, error, Queue[int])
		expectedResult  func() Queue[int]
		expectedElement []int
		expectedError   error
	}{
		{
			name: "test block remove until an element is added",
			actualResult: func() ([]int, error, Queue[int]) {
				var err error
				var q Queue[int]
				var e int

				q = NewBlockingQueue[int]()

				c := make(chan int)

				go func(q Queue[int], e int, err error) {
					e, err = q.Remove()
					c <- e
				}(q, e, err)

				go func(q Queue[int]) {
					time.AfterFunc(time.Millisecond*5, func() {
						q.Add(1)
					})
				}(q)

				return []int{<-c}, err, q

			},
			expectedElement: []int{1},
			expectedResult: func() Queue[int] {
				ll := list.NewLinkedList(1)
				ll.Clear()

				return &BlockingQueue[int]{&LinkedQueue[int]{ll: ll}}
			},
		},
		{
			name: "test do not block when an element are already present added",
			actualResult: func() ([]int, error, Queue[int]) {
				var err error
				var q Queue[int]
				var e int

				q = NewBlockingQueue[int]()

				q.Add(1)
				q.Add(2)

				c := make(chan int)

				go func(q Queue[int], e int, err error) {
					e, err = q.Remove()
					c <- e
				}(q, e, err)

				return []int{<-c}, err, q
			},
			expectedElement: []int{1},
			expectedResult: func() Queue[int] {
				ll := list.NewLinkedList(2)

				return &BlockingQueue[int]{&LinkedQueue[int]{ll: ll}}
			},
		},
		{
			name: "test remove multiple elements",
			actualResult: func() ([]int, error, Queue[int]) {
				var err error
				var q Queue[int]
				var e int
				var res []int

				c := make(chan int)

				remove := func(q Queue[int], e int, err error) {
					e, err = q.Remove()
					c <- e
				}

				add := func(q Queue[int], e int) {
					time.AfterFunc(time.Millisecond*5, func() {
						q.Add(e)
					})
				}

				q = NewBlockingQueue[int]()

				q.Add(1)
				q.Add(2)

				go remove(q, e, err)
				res = append(res, <-c)

				go remove(q, e, err)
				res = append(res, <-c)

				go add(q, 4)
				go add(q, 3)

				go remove(q, e, err)
				res = append(res, <-c)

				go remove(q, e, err)
				res = append(res, <-c)

				return res, err, q
			},
			expectedElement: []int{1, 2, 3, 4},
			expectedResult: func() Queue[int] {
				ll := list.NewLinkedList(1)
				ll.Clear()

				return &BlockingQueue[int]{&LinkedQueue[int]{ll: ll}}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			e, err, res := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)

			if testCase.expectedError == nil {
				assert.Equal(t, len(testCase.expectedElement), len(e))
				assert.Equal(t, testCase.expectedResult(), res)
			}
		})
	}
}

//TODO: FIX TEST
func NotTestBlockingQueueRemoveWithTimeout(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() ([]int, error, Queue[int])
		expectedResult  func() Queue[int]
		expectedElement []int
		expectedError   error
	}{
		{
			name: "test remove element with max wait of 10 milli sec",
			actualResult: func() ([]int, error, Queue[int]) {
				var err error
				var q *BlockingQueue[int]
				var e int

				q = NewBlockingQueue[int]()

				c := make(chan int)

				go func(q *BlockingQueue[int], e int, err error) {
					e, err = q.RemoveWithTimeout(time.Millisecond * 10)
					c <- e
				}(q, e, err)

				go func(q Queue[int]) {
					time.AfterFunc(time.Second, func() {
						q.Add(1)
					})
				}(q)

				return []int{<-c}, err, q

			},
			expectedElement: []int{1},
			expectedResult: func() Queue[int] {
				ll := list.NewLinkedList[int]()

				return &BlockingQueue[int]{&LinkedQueue[int]{ll: ll}}
			},
		},
		{
			name: "test remove element fails when max wait is 1 milli sec",
			actualResult: func() ([]int, error, Queue[int]) {
				var err error
				var q *BlockingQueue[int]
				var e int

				q = NewBlockingQueue[int]()

				c := make(chan int)
				done := make(chan bool)

				go func(q *BlockingQueue[int], e int, err *error) {
					e, *err = q.RemoveWithTimeout(time.Millisecond)
					c <- e
				}(q, e, &err)

				go func(q Queue[int]) {
					time.AfterFunc(time.Second, func() {
						q.Add(1)
						done <- true
					})
				}(q)

				<-done
				return []int{<-c}, err, q

			},
			expectedElement: []int{1},
			expectedResult: func() Queue[int] {
				ll := list.NewLinkedList(1)

				return &BlockingQueue[int]{&LinkedQueue[int]{ll: ll}}
			},
			expectedError: errors.New("timed out"),
		},
		{
			name: "test do not block when an element are already present added",
			actualResult: func() ([]int, error, Queue[int]) {
				var err error
				var q *BlockingQueue[int]
				var e int

				q = NewBlockingQueue[int]()

				q.Add(1)
				q.Add(2)

				c := make(chan int)

				go func(q *BlockingQueue[int], e int, err error) {
					e, err = q.RemoveWithTimeout(time.Millisecond)
					c <- e
				}(q, e, err)

				return []int{<-c}, err, q
			},
			expectedElement: []int{1},
			expectedResult: func() Queue[int] {
				ll := list.NewLinkedList(2)

				return &BlockingQueue[int]{&LinkedQueue[int]{ll: ll}}
			},
		},
		{
			name: "test remove multiple elements with max wait of 10 milli sec",
			actualResult: func() ([]int, error, Queue[int]) {
				var err error
				var q *BlockingQueue[int]
				var e int
				var res []int

				c := make(chan int)

				remove := func(q *BlockingQueue[int], e int, err *error) {
					e, *err = q.RemoveWithTimeout(time.Millisecond * 10)
					c <- e
				}

				add := func(q *BlockingQueue[int], e int) {
					time.AfterFunc(time.Millisecond*5, func() {
						q.Add(e)
					})
				}

				q = NewBlockingQueue[int]()

				q.Add(1)
				q.Add(2)

				go remove(q, e, &err)
				res = append(res, <-c)

				go remove(q, e, &err)
				res = append(res, <-c)

				go add(q, 4)
				go add(q, 3)

				go remove(q, e, &err)
				res = append(res, <-c)

				go remove(q, e, &err)
				res = append(res, <-c)

				return res, err, q
			},
			expectedElement: []int{1, 2, 3, 4},
			expectedResult: func() Queue[int] {
				ll := list.NewLinkedList(1)
				ll.Clear()

				return &BlockingQueue[int]{&LinkedQueue[int]{ll: ll}}
			},
		},
		{
			name: "test remove multiple elements fails when max wait is 1 milli sec",
			actualResult: func() ([]int, error, Queue[int]) {
				var err error
				var q *BlockingQueue[int]
				var e int
				var res []int

				c := make(chan int)

				remove := func(q *BlockingQueue[int], e int, err *error) {
					e, *err = q.RemoveWithTimeout(time.Millisecond)
					c <- e
				}

				add := func(q *BlockingQueue[int], e int) {
					time.AfterFunc(time.Second, func() {
						q.Add(e)
					})
				}

				q = NewBlockingQueue[int]()

				q.Add(1)
				q.Add(2)

				go remove(q, e, &err)
				res = append(res, <-c)

				go remove(q, e, &err)
				res = append(res, <-c)

				go add(q, 4)
				go add(q, 3)

				go remove(q, e, &err)
				res = append(res, <-c)

				go remove(q, e, &err)
				res = append(res, <-c)

				return res, err, q
			},
			expectedElement: []int{},
			expectedResult: func() Queue[int] {
				ll := list.NewLinkedList(1)
				ll.Clear()

				return &BlockingQueue[int]{&LinkedQueue[int]{ll: ll}}
			},
			expectedError: errors.New("timed out"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			e, err, res := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)

			assert.Equal(t, len(testCase.expectedElement), len(e))
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}
