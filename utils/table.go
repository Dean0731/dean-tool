package utils

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// PrintSliceAsTable 打印结构体数组为表格形式
func PrintSliceAsTable[T any](title string, data []T) {
	table := tablewriter.NewTable(os.Stdout)
	table.Bulk(data)
	table.Render()
	os.Stdout.Write([]byte(title))
}
