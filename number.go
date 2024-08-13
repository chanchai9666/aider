package aider

import (
	"fmt"
	"regexp"
	"strconv"
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

// เติม 0 ด้านหน้าตัวเลข
func PadZeros(width int, number int) string {
	numberStr := fmt.Sprintf("%d", number)
	padding := width - len(numberStr)
	if padding <= 0 {
		return numberStr
	}
	return strings.Repeat("0", padding) + numberStr
}

// แปลงค่าเป็น string
func ToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%v", v)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// แปลงค่าเป็น int
func ToInt(value interface{}) int {
	switch v := value.(type) {
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case string:
		result, err := strconv.Atoi(v)
		if err != nil {
			// fmt.Println("cannot convert %q to int: %v", v, err)
			return 0
		}
		return result
	default:
		// fmt.Println("cannot convert %v to int", value)
		return 0
	}
}

// แปลงค่าเป็น float64
func ToFloat64(value interface{}) float64 {
	switch v := value.(type) {
	case float32:
		return float64(v)
	case float64:
		return v
	case int:
		return float64(v)
	case int8:
		return float64(v)
	case int16:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case string:
		result, err := strconv.ParseFloat(v, 64)
		if err != nil {
			// fmt.Println("cannot convert %q to float64: %v", v, err)
			return 0
		}
		return result
	default:
		// fmt.Println("cannot convert %v of type %s to float64", value, reflect.TypeOf(value))
		return 0
	}
}
