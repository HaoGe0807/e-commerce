package utils

import (
	"bytes"
	"e-commerce/service/infra/log"
	"github.com/xuri/excelize/v2"
	"sync"
)

type ExcelWriter struct {
	f        *excelize.File
	r        *excelize.Rows
	lock     sync.RWMutex
	sheetMap map[string]int
}

type ExcelWriterOption func(e *ExcelWriter)

func NewExcelWriter(opts ...ExcelWriterOption) *ExcelWriter {
	e := new(ExcelWriter)
	e.f = excelize.NewFile()
	e.lock = sync.RWMutex{}
	e.sheetMap = make(map[string]int)
	return e
}

func (e *ExcelWriter) WriteRows(sheetName string, rows [][]string) *ExcelWriter {
	e.lock.Lock()
	defer e.lock.Unlock()

	if _, ok := e.sheetMap[sheetName]; !ok {
		e.sheetMap[sheetName] = e.f.NewSheet(sheetName)
	}

	for i, row := range rows {
		axis, err := excelize.CoordinatesToCellName(1, i+1)
		if err != nil {
			log.Errorf("excel CoordinatesToCellName err: %v, col: %d, row: %d", err, 1, i+1)
		}
		err = e.f.SetSheetRow(sheetName, axis, &row)
		if err != nil {
			log.Errorf("excel SetSheetRow err: %v, row: %v", err, row)
		}
	}
	return e
}

func (e *ExcelWriter) WriteRowsWithStyleForFailList(sheetName string, rows [][]string) *ExcelWriter {
	e.lock.Lock()
	defer e.lock.Unlock()

	if _, ok := e.sheetMap[sheetName]; !ok {
		e.sheetMap[sheetName] = e.f.NewSheet(sheetName)
	}
	// 设置标题样式
	styleId, _ := e.f.NewStyle(`{"fill":{"type":"pattern","color":["#C2CCD0"],"pattern":1},"font":{"bold":true},"alignment":{"wrap_text":true,"vertical":"top"}}`)
	_ = e.f.SetCellStyle(sheetName, "A1", "G1", styleId) // 设置标题样式
	//设置格式
	_ = e.f.SetRowHeight(sheetName, 1, 40)
	_ = e.f.SetColWidth(sheetName, "A", "G", 50)

	for i, row := range rows {
		axis, err := excelize.CoordinatesToCellName(1, i+1)
		if err != nil {
			log.Errorf("excel CoordinatesToCellName err: %v, col: %d, row: %d", err, 1, i+1)
		}
		err = e.f.SetSheetRow(sheetName, axis, &row)
		if err != nil {
			log.Errorf("excel SetSheetRow err: %v, row: %v", err, row)
		}
	}
	e.f.SetSheetName("Sheet1", "失败设备列表")
	return e
}

func (e *ExcelWriter) WriteRowsWithStyle(sheetName string, rows [][]string) *ExcelWriter {
	e.lock.Lock()
	defer e.lock.Unlock()

	if _, ok := e.sheetMap[sheetName]; !ok {
		e.sheetMap[sheetName] = e.f.NewSheet(sheetName)
	}
	// 设置标题样式
	styleId, _ := e.f.NewStyle(`{"fill":{"type":"pattern","color":["#C2CCD0"],"pattern":1},"font":{"bold":true},"alignment":{"wrap_text":true,"vertical":"top"}}`)
	_ = e.f.SetCellStyle(sheetName, "A1", "F1", styleId) // 设置标题样式
	//设置格式
	_ = e.f.SetRowHeight(sheetName, 1, 40)
	_ = e.f.SetColWidth(sheetName, "A", "F", 50)

	for i, row := range rows {
		axis, err := excelize.CoordinatesToCellName(1, i+1)
		if err != nil {
			log.Errorf("excel CoordinatesToCellName err: %v, col: %d, row: %d", err, 1, i+1)
		}
		err = e.f.SetSheetRow(sheetName, axis, &row)
		if err != nil {
			log.Errorf("excel SetSheetRow err: %v, row: %v", err, row)
		}
	}
	e.f.SetSheetName("Sheet1", "成功设备列表")
	return e
}

func (e *ExcelWriter) WriteRowsWithStyle2(sheetName string, rows [][]string) *ExcelWriter {
	e.lock.Lock()
	defer e.lock.Unlock()

	if _, ok := e.sheetMap[sheetName]; !ok {
		e.sheetMap[sheetName] = e.f.NewSheet(sheetName)
	}

	// 设置标题样式
	styleId, _ := e.f.NewStyle(`{"fill":{"type":"pattern","color":["#C2CCD0"],"pattern":1},"font":{"bold":true},"alignment":{"wrap_text":true,"vertical":"top"}}`)
	_ = e.f.SetCellStyle(sheetName, "A1", "A1", styleId) // 设置标题样式
	//设置格式
	_ = e.f.SetRowHeight(sheetName, 1, 40)
	_ = e.f.SetColWidth(sheetName, "A", "A", 60)

	for i, row := range rows {
		axis, err := excelize.CoordinatesToCellName(1, i+1)
		if err != nil {
			log.Errorf("excel CoordinatesToCellName err: %v, col: %d, row: %d", err, 1, i+1)
		}
		err = e.f.SetSheetRow(sheetName, axis, &row)
		if err != nil {
			log.Errorf("excel SetSheetRow err: %v, row: %v", err, row)
		}
	}
	e.f.SetSheetName("Sheet2", "失败原因")
	return e
}

func (e *ExcelWriter) WriteRowsWithStyle3(sheetName string, rows [][]string) *ExcelWriter {
	e.lock.Lock()
	defer e.lock.Unlock()

	if _, ok := e.sheetMap[sheetName]; !ok {
		e.sheetMap[sheetName] = e.f.NewSheet(sheetName)
	}

	// 设置标题样式
	styleId, _ := e.f.NewStyle(`{"fill":{"type":"pattern","color":["#C2CCD0"],"pattern":1},"font":{"bold":true},"alignment":{"wrap_text":true,"vertical":"top"}}`)
	_ = e.f.SetCellStyle(sheetName, "A1", "B1", styleId) // 设置标题样式
	//设置格式
	_ = e.f.SetRowHeight(sheetName, 1, 30)
	_ = e.f.SetColWidth(sheetName, "A", "B", 30)

	for i, row := range rows {
		axis, err := excelize.CoordinatesToCellName(1, i+1)
		if err != nil {
			log.Errorf("excel CoordinatesToCellName err: %v, col: %d, row: %d", err, 1, i+1)
		}
		err = e.f.SetSheetRow(sheetName, axis, &row)
		if err != nil {
			log.Errorf("excel SetSheetRow err: %v, row: %v", err, row)
		}
	}
	e.f.SetSheetName("Sheet3", "")
	return e
}

/**
  分步写入
//*/
//func (e *ExcelWriter) WriteRowsByIndex(sheetName string, rows [][]string, index int) *ExcelWriter {
//	e.lock.Lock()
//	defer e.lock.Unlock()
//
//	if _, ok := e.sheetMap[sheetName]; !ok {
//		e.sheetMap[sheetName] = e.f.NewSheet(sheetName)
//	}
//	// index+2 的意思是由于表格的表头存在，所以加2
//	for i, row := range rows {
//		axis, err := excelize.CoordinatesToCellName(1, i+index+2)
//		if err != nil {
//			slog.Errorf("excel CoordinatesToCellName err: %v, col: %d, row: %d", err, 1, i+2+index)
//		}
//		err = e.f.SetSheetRow(sheetName, axis, &row)
//		if err != nil {
//			slog.Errorf("excel SetSheetRow err: %v, row: %v", err, row)
//		}
//	}
//	return e
//}

func (e *ExcelWriter) ToBuffer() (*bytes.Buffer, error) {
	e.lock.Lock()
	defer e.lock.Unlock()

	return e.f.WriteToBuffer()
}
