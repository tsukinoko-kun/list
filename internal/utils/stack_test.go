package utils_test

import (
	"testing"

	"github.com/Frank-Mayer/list/internal/utils"
)

func TestStack(t *testing.T) {
	t.Parallel()
	s := utils.NewStack[int]()

	if s.Len() != 0 {
		t.Errorf("Stack should be empty")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Len() != 3 {
		t.Errorf("Stack should have 3 elements")
	}

	if *s.Pop() != 3 {
		t.Errorf("Stack should return 3")
	}

	if s.Len() != 2 {
		t.Errorf("Stack should have 2 elements")
	}

	if *s.Pop() != 2 {
		t.Errorf("Stack should return 2")
	}

	if s.Len() != 1 {
		t.Errorf("Stack should have 1 element")
	}

	if *s.Pop() != 1 {
		t.Errorf("Stack should return 1")
	}

	if s.Len() != 0 {
		t.Errorf("Stack should be empty")
	}

	if s.Pop() != nil {
		t.Errorf("Stack should return nil")
	}

	if s.Len() != 0 {
		t.Errorf("Stack should be empty")
	}
}
