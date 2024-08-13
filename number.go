package aider

import (
	"regexp"
	"strings"

	"github.com/shopspring/decimal"
)

// ลบตัวอัพษรออก มีได้เฉพาะตัวเลข และ จุดทศนิยม 1 จุด
func removeNonNumeric(s string) string {
	// ลบอักษรทั้งหมด ยกเว้นตัวเลขและจุด
	re := regexp.MustCompile(`[^\d.]`)
	processedString := re.ReplaceAllString(s, "")

	// หากมีมากกว่าหนึ่งจุด ให้เก็บเฉพาะจุดแรกและลบส่วนที่เหลือออก
	if strings.Count(processedString, ".") > 1 {
		// ค้นหาตำแหน่งของจุดทศนิยมตัวแรก
		dotIndex := strings.Index(processedString, ".")
		// เก็บเฉพาะส่วนที่อยู่ก่อนและหลังจุดทศนิยมตัวแรก
		processedString = processedString[:strings.Index(processedString[dotIndex+1:], ".")+dotIndex+1]
	}

	return processedString
}

// value=ค่าที่ต้องการ d=จำนวนทศนิยม
func Round2d(value float64, d int32) float64 {
	v := decimal.NewFromFloat(value)
	v2 := v.RoundFloor(d) //ใช้อันนี้ปัดเศษ
	v3, _ := v2.Float64() //แปลงเป็น float64

	return v3
}

// แปลง string เป็น Float64
func StringToFloat64(value string) float64 {
	data := removeNonNumeric(value)
	if data == "" {
		data = "0"
	}
	num, _ := decimal.NewFromString(data)
	conv, _ := num.Float64()
	return conv
}
