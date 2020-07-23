package helper

import (
    "fmt"
)

type Printer struct {
    rows [][]string
}

func NewPrinter(numRows int) *Printer {
    var printer Printer
    printer.rows = make([]string, 0, numRows)

    return &printer
}

func (printer *Printer) AddRow(row ...string) {
    fmt.Printf("%v\n", row)
    printer.rows = append(printer.rows, row...)
    fmt.Printf("%v\n", printer.rows[0])
}

func (printer *Printer) Print() {
    for _, row := range printer.rows {
        for _, cell := range row {
            fmt.Printf("%s\n", cell)
        }
    }
}
