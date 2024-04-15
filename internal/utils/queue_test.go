package utils_test

import (
	"testing"

	"github.com/tsukinoko-kun/list/internal/utils"
)

func TestQueue(t *testing.T) {
	t.Parallel()
	q := utils.NewQueue[int]()

	if q.Len() != 0 {
		t.Errorf("Queue should be empty")
	}

	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.Len() != 3 {
		t.Errorf("Queue should have 3 elements")
	}

	if *q.Pop() != 1 {
		t.Errorf("Queue should return 1")
	}

	if q.Len() != 2 {
		t.Errorf("Queue should have 2 elements")
	}

	if *q.Pop() != 2 {
		t.Errorf("Queue should return 2")
	}

	if q.Len() != 1 {
		t.Errorf("Queue should have 1 element")
	}

	if *q.Pop() != 3 {
		t.Errorf("Queue should return 3")
	}

	if q.Len() != 0 {
		t.Errorf("Queue should be empty")
	}

	if q.Pop() != nil {
		t.Errorf("Queue should return nil")
	}

	if q.Len() != 0 {
		t.Errorf("Queue should be empty")
	}
}
