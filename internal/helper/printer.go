package helper

import (
    "fmt"
)

type Printer struct {
    rows [][]string
    columnLengths []int
}

func NewPrinter(numRows int) *Printer {
    return &Printer{}
}

func (printer *Printer) AddRow(row ...string) {
    printer.rows = append(printer.rows, row)

    if printer.columnLengths == nil {
        printer.columnLengths = make([]int, len(row))
    }

    var columnLength int
    for i, cell := range row {
        columnLength = len(cell) + 2
        if columnLength > printer.columnLengths[i] {
            printer.columnLengths[i] = columnLength 
        }
    }
}

func (printer *Printer) Print() {
    for _, row := range printer.rows {
        fmt.Println()
        for i, cell := range row {
            fmt.Printf("%-*s ", printer.columnLengths[i], cell)
        }
    }
    fmt.Println()
}
