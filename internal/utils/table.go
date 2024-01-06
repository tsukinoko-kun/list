package utils

import (
	"fmt"
	"github.com/fatih/color"
)

type TableEntry struct {
	table *Table
	index int
	style *color.Color
	data  []string
}

type Table struct {
	columnCount int
	entries     []*TableEntry
}

func NewTable(columnCount int) *Table {
	t := new(Table)
	t.columnCount = columnCount
	return t
}

func (t *Table) NewEntry() *TableEntry {
	entry := new(TableEntry)
	entry.table = t
	entry.index = 0
	entry.data = make([]string, t.columnCount)
	t.entries = append(t.entries, entry)
	return entry
}

func (te *TableEntry) AddCell(cell string) error {
	if te.index > te.table.columnCount {
		return fmt.Errorf("Failed to add cell: TableEntry already has %d cells", te.index)
	}
	te.data[te.index] = cell
	te.index++
	return nil
}

func (te *TableEntry) SetStyle(style *color.Color) {
	te.style = style
}

func (te *TableEntry) Style() *color.Color {
	return te.style
}

// String creates a formatted string representation of the table.
func (t *Table) String() string {
	// find the longest string in each column
	maxLengths := make([]int, t.columnCount)
	for _, entry := range t.entries {
		for i, cell := range entry.data {
			if len(cell) > maxLengths[i] {
				maxLengths[i] = len(cell)
			}
		}
	}

	// create a format string for each column
	formats := make([]string, t.columnCount)
	for i, length := range maxLengths {
		formats[i] = fmt.Sprintf("%%-%ds", length)
	}

	// create the table
	var table string
	for _, entry := range t.entries {
		for i, cell := range entry.data {
			table += fmt.Sprintf(formats[i], cell)
		}
	}
	return table
}

func (t *Table) Print() {
	// find the longest string in each column
	maxLengths := make([]int, t.columnCount)
	for _, entry := range t.entries {
		for i, cell := range entry.data {
			if len(cell) > maxLengths[i] {
				maxLengths[i] = len(cell)
			}
		}
	}

	// create a format string for each column
	formats := make([]string, t.columnCount)
	for i, length := range maxLengths {
		formats[i] = fmt.Sprintf("%%-%ds", length)
	}

	// create the table
	for _, entry := range t.entries {
		for i, cell := range entry.data {
			entry.style.Printf(formats[i], cell)
			print(" ")
		}
		println()
	}
}
