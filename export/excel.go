package export

import (
	"bytes"
	"html"

	"github.com/shixinshuiyou/framework/conv"
	"github.com/xuri/excelize/v2"
)

func getAxis(rowId, colId int) string {
	colName, _ := excelize.ColumnNumberToName(colId)
	axis, _ := excelize.JoinCellName(colName, rowId)
	return axis
}

// 由于excel03 最大列个数为256，限制比较多，因此建议使用 excel导出
// data 全表数据，rowHightlight 标记某行高亮，若无需高亮，传nil即可
func Excel(data [][]interface{}, rowHighlight []bool) (*bytes.Buffer, error) {
	f := excelize.NewFile()
	//	f.SetActiveSheet(1)
	sheetName := "Sheet1"

	// 设置数据
	for rowId, row := range data {
		for colId, v := range row {
			f.SetCellValue(sheetName, getAxis(rowId+1, colId+1), v)
		}
	}

	// 设置样式
	if len(rowHighlight) != 0 {
		style, err := f.NewStyle(&excelize.Style{
			Fill:      excelize.Fill{Type: "pattern", Color: []string{"#BFBFBF"}, Pattern: 1},
			Font:      &excelize.Font{Bold: true},
			Alignment: &excelize.Alignment{Horizontal: "center"},
		})
		if err != nil {
			return nil, err
		}

		for rowId, needHighlight := range rowHighlight {
			if needHighlight {
				f.SetCellStyle(sheetName, getAxis(rowId+1, 1), getAxis(rowId+1, len(data[rowId])), style)
			}
		}
	}
	return f.WriteToBuffer()
}

// 指定某行高亮
func Excel03(data [][]interface{}, rowHighlight []bool) *bytes.Buffer {
	buffer := bytes.NewBufferString("")
	buffer.WriteString(`<html xmlns:x="urn:schemas-microsoft-com:office:excel">
	<head><meta http-equiv="Content-type" content="text/html;charset=UTF-8" /></head><style><!--
	body,table{font-size:14px;}
	table{border-collapse:collapse;} 
	th{padding:2px;border:1px solid #000;white-space:nowrap;background-color:#F2FAFF;}
    td{padding:2px;border:1px solid #000;white-space:nowrap;}--></style>
	<body><table border="1" cellspacing="0" cellpadding="0"><tbody>`)

	for i, row := range data {
		buffer.WriteString(`<tr>`)
		if len(rowHighlight) > i && rowHighlight[i] {
			for _, cell := range row {
				buffer.WriteString(`<th style="background-color:#F2FAFF;white-space:nowrap;">`)
				buffer.WriteString(html.EscapeString(conv.String(cell)))
				buffer.WriteString(`</th>`)
			}
		} else {
			for _, cell := range row {
				buffer.WriteString(`<td>`)
				buffer.WriteString(html.EscapeString(conv.String(cell)))
				buffer.WriteString(`</td>`)
			}
		}
		buffer.WriteString(`</tr>`)
	}
	buffer.WriteString(`</tbody></table></body></html>`)

	return buffer
}

func Excel03WriteFile(filename string, data [][]interface{}, rowHighlight []bool) error {
	return WriteFile(filename, Excel03(data, rowHighlight))
}

// filename 带后缀 xlsx
// data 全表数据
// rowHightlight 标记某行高亮，若无需高亮，传nil即可
func ExcelWriteFile(filename string, data [][]interface{}, rowHighlight []bool) error {
	buf, err := Excel(data, rowHighlight)
	if err != nil {
		return err
	}
	return WriteFile(filename, buf)
}
