package utils_test

import (
	"testing"

	"github.com/Frank-Mayer/list/internal/utils"
)

func TestTable(t *testing.T) {
	table := utils.NewTable(3)
	var err error

	e1 := table.NewEntry()
	err = e1.AddCell("1")
	if err != nil {
		t.Errorf("Error adding cell: %s", err)
	}
	err = e1.AddCell("2")
	if err != nil {
		t.Errorf("Error adding cell: %s", err)
	}
	err = e1.AddCell("3")
	if err != nil {
		t.Errorf("Error adding cell: %s", err)
	}

	e2 := table.NewEntry()
	err = e2.AddCell("4")
	if err != nil {
		t.Errorf("Error adding cell: %s", err)
	}
	err = e2.AddCell("5000")
	if err != nil {
		t.Errorf("Error adding cell: %s", err)
	}
	err = e2.AddCell("6000000")
	if err != nil {
		t.Errorf("Error adding cell: %s", err)
	}

	e3 := table.NewEntry()
	err = e3.AddCell("foo")
	if err != nil {
		t.Errorf("Error adding cell: %s", err)
	}
	err = e3.AddCell("bar")
	if err != nil {
		t.Errorf("Error adding cell: %s", err)
	}
	err = e3.AddCell("baz")
	if err != nil {
		t.Errorf("Error adding cell: %s", err)
	}
}
