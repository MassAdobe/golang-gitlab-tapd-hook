package createXLS

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"gitlab_tapd/common"
	"strconv"
	"time"
)

func CreateExcel(excelName string) {
	f := excelize.NewFile()
	// 概览
	index := f.NewSheet("KPI-VIEW")
	f.SetCellValue("KPI-VIEW", "A1", "ID")
	f.SetCellValue("KPI-VIEW", "B1", "姓名")
	f.SetCellValue("KPI-VIEW", "C1", "用户名")
	f.SetCellValue("KPI-VIEW", "D1", "账户状态")
	f.SetCellValue("KPI-VIEW", "E1", "项目名称")
	f.SetCellValue("KPI-VIEW", "F1", "提交行数")
	f.SetCellValue("KPI-VIEW", "G1", "提交次数")
	x, y := 1, 2
	for k, v := range common.FUser {
		f.SetCellValue("KPI-VIEW", generateColumnIndex(x)+strconv.Itoa(y), k)
		x++
		f.SetCellValue("KPI-VIEW", generateColumnIndex(x)+strconv.Itoa(y), v.Name)
		x++
		f.SetCellValue("KPI-VIEW", generateColumnIndex(x)+strconv.Itoa(y), v.Username)
		x++
		f.SetCellValue("KPI-VIEW", generateColumnIndex(x)+strconv.Itoa(y), v.State)
		x++
		if len(v.ProjectInfos) == 0 {
			y++
		} else if len(v.ProjectInfos) == 1 {
			for _, v1 := range v.ProjectInfos {
				f.SetCellValue("KPI-VIEW", generateColumnIndex(x)+strconv.Itoa(y), v1.ProjectName)
				x++
				f.SetCellValue("KPI-VIEW", generateColumnIndex(x)+strconv.Itoa(y), v1.CommitColums)
				x++
				f.SetCellValue("KPI-VIEW", generateColumnIndex(x)+strconv.Itoa(y), v1.CommitTms)
			}
			y++
		} else {
			for _, v1 := range v.ProjectInfos {
				f.SetCellValue("KPI-VIEW", generateColumnIndex(x)+strconv.Itoa(y), v1.ProjectName)
				x++
				f.SetCellValue("KPI-VIEW", generateColumnIndex(x)+strconv.Itoa(y), v1.CommitColums)
				x++
				f.SetCellValue("KPI-VIEW", generateColumnIndex(x)+strconv.Itoa(y), v1.CommitTms)
				x -= 2
				y++
			}
		}
		x = 1
	}
	f.SetActiveSheet(index)
	f.Save()
	// 铺平排序
	secondIndex := f.NewSheet("KPI-SORT")
	f.SetCellValue("KPI-SORT", "A1", "ID")
	f.SetCellValue("KPI-SORT", "B1", "姓名")
	f.SetCellValue("KPI-SORT", "C1", "用户名")
	f.SetCellValue("KPI-SORT", "D1", "账户状态")
	f.SetCellValue("KPI-SORT", "E1", "项目名称")
	f.SetCellValue("KPI-SORT", "F1", "提交行数")
	f.SetCellValue("KPI-SORT", "G1", "提交次数")
	xx, yy := 1, 2
	for k, v := range common.FUser {
		if len(v.ProjectInfos) == 0 {
			f.SetCellValue("KPI-SORT", generateColumnIndex(xx)+strconv.Itoa(yy), k)
			xx++
			f.SetCellValue("KPI-SORT", generateColumnIndex(xx)+strconv.Itoa(yy), v.Name)
			xx++
			f.SetCellValue("KPI-SORT", generateColumnIndex(xx)+strconv.Itoa(yy), v.Username)
			xx++
			f.SetCellValue("KPI-SORT", generateColumnIndex(xx)+strconv.Itoa(yy), v.State)
			xx = 1
			yy++
			continue
		}
		for _, v1 := range v.ProjectInfos {
			f.SetCellValue("KPI-SORT", generateColumnIndex(xx)+strconv.Itoa(yy), k)
			xx++
			f.SetCellValue("KPI-SORT", generateColumnIndex(xx)+strconv.Itoa(yy), v.Name)
			xx++
			f.SetCellValue("KPI-SORT", generateColumnIndex(xx)+strconv.Itoa(yy), v.Username)
			xx++
			f.SetCellValue("KPI-SORT", generateColumnIndex(xx)+strconv.Itoa(yy), v.State)
			xx++
			f.SetCellValue("KPI-SORT", generateColumnIndex(xx)+strconv.Itoa(yy), v1.ProjectName)
			xx++
			f.SetCellValue("KPI-SORT", generateColumnIndex(xx)+strconv.Itoa(yy), v1.CommitColums)
			xx++
			f.SetCellValue("KPI-SORT", generateColumnIndex(xx)+strconv.Itoa(yy), v1.CommitTms)
			xx = 1
			yy++
		}
	}
	f.SetActiveSheet(secondIndex)
	f.Save()
	// 生成
	nameExtend := ""
	if len(excelName) == 0 {
		nameExtend = time.Now().Format(common.FORMAT_TIME)
	} else {
		nameExtend = excelName
	}
	if err := f.SaveAs(fmt.Sprintf(common.SAVE_PATH_FILE, nameExtend)); err != nil {
		panic(err.Error())
	}
}

// 生成横坐标
func generateColumnIndex(n int) string {
	res := ""
	for n > 0 {
		m := n % 26
		if m == 0 {
			m = 26
		}
		res = string(m+64) + res
		n = (n - m) / 26
	}
	return res
}
