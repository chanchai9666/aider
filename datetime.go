package aider

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	languageTh     string = "th"
	dateLayout     string = "2006-01-02"          //รูปแบบของวันที่ใน Go
	timeLayout     string = "15:04:05"            //รูปแบบของเวลาใน Go
	datetimeLayout string = "2006-01-02 15:04:05" //รูปแบบวันเวลาใน Go
)

var thaiMonths = []string{
	"ม.ค.", "ก.พ.", "มี.ค.", "เม.ย.", "พ.ค.", "มิ.ย.",
	"ก.ค.", "ส.ค.", "ก.ย.", "ต.ค.", "พ.ย.", "ธ.ค.",
}

var allowedLanguages = map[string]bool{
	"en": true,
	"th": true,
}

// LoadLocation returns The time zone.
func loadLocation() *time.Location {
	timeZone, _ := time.LoadLocation("Asia/Bangkok")
	return timeZone
}

// FormatDate format date
func FormatDate(dt *time.Time) (string, string) {
	dateFmt := dt.In(loadLocation()).Format("02 January 2006")
	return dateFormat(dateFmt, "en"), dateFormat(dateFmt, "th")
}

func dateFormat(dateFmt, language string) string {
	if language == languageTh {
		switch getMonth(dateFmt) {
		case "January":
			dateFmt = strings.Replace(dateFmt, "January", "มกราคม", 1)

		case "February":
			dateFmt = strings.Replace(dateFmt, "February", "กุมภาพันธ์", 1)

		case "March":
			dateFmt = strings.Replace(dateFmt, "March", "มีนาคม", 1)

		case "April":
			dateFmt = strings.Replace(dateFmt, "April", "เมษายน", 1)

		case "May":
			dateFmt = strings.Replace(dateFmt, "May", "พฤษภาคม", 1)

		case "June":
			dateFmt = strings.Replace(dateFmt, "June", "มิถุนายน", 1)

		case "July":
			dateFmt = strings.Replace(dateFmt, "July", "กรกฎาคม", 1)

		case "August":
			dateFmt = strings.Replace(dateFmt, "August", "สิงหาคม", 1)

		case "September":
			dateFmt = strings.Replace(dateFmt, "September", "กันยายน", 1)

		case "October":
			dateFmt = strings.Replace(dateFmt, "October", "ตุลาคม", 1)

		case "November":
			dateFmt = strings.Replace(dateFmt, "November", "พฤศจิกายน", 1)

		case "December":
			dateFmt = strings.Replace(dateFmt, "December", "ธันวาคม", 1)
		}
	}

	return fmt.Sprintf("%s %s %s", getDay(dateFmt), getMonth(dateFmt), getYear(dateFmt, language))
}

func getDay(dateFmt string) string {
	return strings.Split(dateFmt, " ")[0]
}

func getMonth(dateFmt string) string {
	return strings.Split(dateFmt, " ")[1]
}

func getYear(dateFmt, language string) string {
	year, _ := strconv.Atoi(strings.Split(dateFmt, " ")[2])
	if language == languageTh {
		return strconv.Itoa(year + 543)
	}
	return strconv.Itoa(year)
}

// GetToday get datetime today timezone (th)
func GetToday() time.Time {
	timeNow := TimeNowLocationTH()
	return time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, timeNow.Location())
}

// GetYesterday get datetime yesterday timezone (th)
func GetYesterday() time.Time {
	timeNow := TimeNowLocationTH()
	return time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day()-1, 0, 0, 0, 0, timeNow.Location())
}

// GetNextDateTime get datetime tomorrow timezone (th)
func GetNextDateTime(day int) time.Time {
	timeNow := TimeNowLocationTH()
	return time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day()+day, 0, 0, 0, 0, timeNow.Location())
}

// GetPreviousDateTime get datetime yesterday timezone (th)
func GetPreviousDateTime(day int) time.Time {
	timeNow := TimeNowLocationTH()
	return time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day()-day, 0, 0, 0, 0, timeNow.Location())
}

// GetDateTimeByDate get datetime by date timezone (th)
func GetDateTimeByDate(day time.Time) time.Time {
	timeNow := TimeNowLocationTH()
	return time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, timeNow.Location())
}

// TimeNowLocationTH get time location thai
func TimeNowLocationTH() time.Time {
	return time.Now().In(loadLocation())
}

// ShortDate short date time.Time To 21 Feb. 2025 OR 21 ก.พ. 2568
func ShortDate(date time.Time, language string) string {
	// ตรวจสอบว่า language ต้องเป็น "en" หรือ "th" เท่านั้น
	if !allowedLanguages[language] {
		return "invalid language"
	}

	dateFmt := date.Format("02 Jan. 2006")
	if language == "en" {
		return dateFmt
	}

	day := strings.Split(dateFmt, " ")[0]
	month := getMonthThai(date.Month())
	year := getYear(dateFmt, language)

	rr := fmt.Sprintf("%s %s %s", day, month, year)
	fmt.Println("====>", rr)
	return rr
}

// ฟังก์ชันคืนค่าชื่อเดือนเป็นภาษาไทย
func getMonthThai(month time.Month) string {
	// time.Month มีค่าเริ่มต้นจาก 1 (January = 1) ดังนั้น index ต้องลบ 1
	if month >= time.January && month <= time.December {
		return thaiMonths[month-1]
	}
	return "" // คืนค่าว่างหากค่าเดือนผิดพลาด
}

// ใช้สำหรับ บวก (เพิ่ม) หรือ ลบ (ลด) ค่าของ ปี, เดือน, วัน, ชั่วโมง, นาที และวินาที ไปยังวันที่ที่ระบุในรูปแบบของสตริง (datetime) และคืนค่าวันที่ที่ถูกปรับแล้วกลับมาในรูปแบบเดิม
func ModifyDatetime(datetime string, year, month, day, hour, min, sec int) string {
	layout := datetimeLayout
	dt, _ := time.Parse(layout, datetime)
	if dt.IsZero() {
		return ""
	}

	dt = time.Date(
		dt.Year()+year,
		dt.Month()+time.Month(month),
		dt.Day()+day,
		dt.Hour()+hour,
		dt.Minute()+min,
		dt.Second()+sec,
		0,
		loadLocation(),
	)

	return dt.Format(layout)

	/*
		EX
		originalDate := "2025-02-21 14:30:00"
		ทดสอบเพิ่ม 1 ปี, 2 เดือน, 3 วัน, 4 ชั่วโมง, 5 นาที, 6 วินาที
		newDate := ModifyDatetime(originalDate, 1, 2, 3, 4, 5, 6)

		ลดลง 1 ปี, 2 เดือน, 3 วัน, 4 ชั่วโมง, 5 นาที, 6 วินาที
		newDate := ModifyDatetime(originalDate, -1, -2, -3, -4, -5, -6)

		fmt.Println(ModifyDatetime("2025-02-21 14:30:00", 1, 2, 3, 4, 5, 6))  // +1 ปี +2 เดือน +3 วัน +4 ชั่วโมง +5 นาที +6 วินาที
		fmt.Println(ModifyDatetime("2025-02-21 14:30:00", -1, -2, -3, -4, -5, -6)) // -1 ปี -2 เดือน -3 วัน -4 ชั่วโมง -5 นาที -6 วินาที
	*/
}

// ShortYearMonth short month
func ShortYearMonth(date time.Time, language string) string {
	if !allowedLanguages[language] {
		return "invalid language"
	}
	dateFmt := date.Format("Jan. 2006")
	if language == "en" {
		return dateFmt
	}
	month := getMonthThai(date.Month())
	year := date.Year()

	return fmt.Sprintf("%s %d", month, year)
}

// ShortYearMonth short month
// time.Tim TO Feb. 2025 , กุมภาพันธ์ 2025
func ShortMonth(date time.Time, language string) string {
	if !allowedLanguages[language] {
		return "invalid language"
	}
	dateFmt := date.Format("Jan.")
	if language == "en" {
		return dateFmt
	}
	month := getMonthThai(date.Month())

	return month
}

// GetDate get date from string datetime format
// time.Time TO Feb ,กุมภาพันธ์
func GetDate(date string) string {
	if date != "" {
		s := strings.Split(date, "T")
		if len(s) > 1 {
			date = s[0]
		}
		if date == "0001-01-01" {
			date = ""
		}
	}
	return date
}

// time.Time คืนค่าเฉพาะเวลา
func DateTimeToTime(datetime string) string {
	if datetime != "" {
		s := strings.Split(datetime, "T")
		if len(s) > 1 {

			split1 := strings.Split(s[1], "Z")
			split2 := strings.Split(s[1], "z")
			split3 := strings.Split(s[1], "+07:00")

			if len(split1) > 1 {
				s[1] = split1[0]
			}
			if len(split2) > 1 {
				s[1] = split2[0]
			}
			if len(split3) > 1 {
				s[1] = split3[0]

			}

			datetime = s[1]
		}
	}

	return datetime
}

// แปลง format วันที่ 2025-02-03T10:15:30Z TO 2025-02-03 10:15:3
func FormatISOToDatetime(datetime string) string {
	if datetime != "" {
		s := strings.Split(datetime, "T")
		if len(s) > 1 {
			date := s[0]
			if date == "0001-01-01" {
				date = "0000-00-00"
			}

			split1 := strings.Split(s[1], "Z")
			split2 := strings.Split(s[1], "z")
			split3 := strings.Split(s[1], "+07:00")

			if len(split1) > 1 {
				s[1] = split1[0]
			}
			if len(split2) > 1 {
				s[1] = split2[0]
			}
			if len(split3) > 1 {
				s[1] = split3[0]
			}

			datetime = date + " " + s[1]
		}
	}

	return datetime
	/*
		Ex.
		fmt.Println(FormatISOToDatetime("2025-02-03T10:15:30Z"))       // "2025-02-03 10:15:30"
		fmt.Println(FormatISOToDatetime("2025-02-03T10:15:30+07:00"))  // "2025-02-03 10:15:30"
		fmt.Println(FormatISOToDatetime("2025-02-03T10:15:30z"))       // "2025-02-03 10:15:30"
		fmt.Println(FormatISOToDatetime("0001-01-01T00:00:00Z"))       // "0000-00-00 00:00:00"
	*/
}

// ดึงข้อมูล วันที่ตามตำแหน่งที่กำหนด
func FormatDateTimeByPosition(dateTime, position string) string {
	if dateTime != "" {
		// แยกวันที่และเวลา
		t := strings.Split(dateTime, " ")

		// ตรวจสอบว่า string มีวันที่และเวลาหรือไม่
		if len(t) < 2 {
			return dateTime // ถ้าไม่มีเวลาให้คืนค่าตามเดิม
		}

		// กรณีที่ตำแหน่งเป็น "0" คืนแค่วันที่
		if position == "0" {
			// ตรวจสอบหากเป็นวันที่ 0001-01-01 ให้แทนที่เป็น 0000-00-00
			if t[0] == "0001-01-01" {
				t[0] = "0000-00-00"
			}
			return t[0] // คืนแค่วันที่
		} else {
			// คืนค่าทั้งวันที่และเวลา
			return t[0] + " " + t[1]
		}
	} else {
		// ถ้าข้อมูลว่าง ให้คืนค่าว่าง
		return ""
	}

	/*
		Ex
		fmt.Println(FormatDateTimeByPosition("2025-02-03 10:15:30", "0"))       // ผลลัพธ์: 2025-02-03
		fmt.Println(FormatDateTimeByPosition("2025-02-03 10:15:30", "1"))       // ผลลัพธ์: 2025-02-03 10:15:30
		fmt.Println(FormatDateTimeByPosition("0001-01-01 00:00:00", "0"))       // ผลลัพธ์: 0000-00-00
		fmt.Println(FormatDateTimeByPosition("2025-02-03", "0"))                // ผลลัพธ์: 2025-02-03
	*/
}

// คำนวณจำนวนวันระหว่างวันที่สอง (ระหว่าง a และ b)
func DaysBetween(a, b time.Time) int {
	days := -a.YearDay()
	for year := a.Year(); year < b.Year(); year++ {
		days += time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).YearDay()
	}
	days += b.YearDay()

	return days

	/*
		Ex.
		fmt.Println(DaysBetween(a, b)) // ผลลัพธ์: 757
	*/
}

// Date แปลงสตริงที่มีรูปแบบตามที่กำหนดให้เป็นค่าเวลา (time.Time)
// ฟังก์ชันนี้ใช้สำหรับการแปลงวันที่ในรูปแบบที่กำหนด (ตามตัวแปร `dateLayout`) ให้เป็นประเภท time.Time
// ถ้าแปลงสำเร็จ จะคืนค่าผลลัพธ์เป็นเวลา
// ในกรณีที่แปลงไม่สำเร็จ (เกิดข้อผิดพลาด) จะไม่คืนค่าผลลัพธ์ที่ถูกต้องเพราะเราละเว้นการจัดการข้อผิดพลาด
// dateLayout คืนค่าเฉพาะวันที่
func Date(s string) time.Time {
	d, _ := time.Parse(dateLayout, s)
	return d
}

// ตรวจสอบว่าเวลา check อยู่ในช่วงระหว่าง start และ end หรือไม่
func InTimeSpan(start, end, check time.Time) bool {
	return (check.After(start) && check.Before(end)) || (check.Equal(start) || check.Equal(end))
}

// datetimeLayout คืนค่า วันที่และเวลา
func DateTime(s string) time.Time {
	d, _ := time.Parse(datetimeLayout, s)
	return d
}

// นับวัน ระหว่างช่วงวันที่
func CountDays(startDateStr, endDateStr string) (int, error) {
	layout := dateLayout // รูปแบบของวันที่ใน Go

	// แปลงวันที่จาก string เป็น time.Time
	startDate, err := time.Parse(layout, startDateStr)
	if err != nil {
		return 0, err
	}

	endDate, err := time.Parse(layout, endDateStr)
	if err != nil {
		return 0, err
	}

	// ใช้ for loop ในการนับจำนวนวัน
	var daysCount int
	for current := startDate; current.Before(endDate) || current.Equal(endDate); current = current.AddDate(0, 0, 1) {
		daysCount++
	}

	return daysCount, nil
}

// แปลงเดือนเป็นเดือนไทย
func ToThaiMonth(month int) (string, error) {
	if month < 1 || month > 12 {
		return "", errors.New("invalid month")
	}
	return thaiMonths[month-1], nil
}

// วันเวลา ปัจจุบัน ประเทศไทย แบบ string
func DateTimeNow() string {
	//location := time.FixedZone("Asia/Bangkok", 7*60*60) // 7 ชั่วโมง
	location := loadLocation()             //ตั้งโซนเวลาของประเทศไทย (Asia/Bangkok)
	currentTime := time.Now().In(location) // ใช้ Time Zone ที่กำหนด
	formattedTime := currentTime.Format(datetimeLayout)
	return formattedTime
}

// วันที่ปัจจุบัน ประเทศไทย แบบ time.Time
func TimeTimeNow() time.Time {
	dateFormat := datetimeLayout
	tt := DateTimeNow()
	parsedTime, err := time.Parse(dateFormat, tt)
	if err != nil {
		return time.Time{}
	}
	return parsedTime
}

// Struct หลักที่เก็บเวลา
type TimeTime struct {
	time time.Time
}

// ฟังก์ชันหลักสำหรับดึงวันเวลาปัจจุบัน
func TimeNow() *TimeTime {
	dateFormat := datetimeLayout
	tt := DateTimeNow()
	parsedTime, err := time.Parse(dateFormat, tt)
	if err != nil {
		return &TimeTime{time.Time{}}
	}
	return &TimeTime{time: parsedTime}
}

// เมธอดสำหรับดึงเฉพาะวันที่
func (d *TimeTime) DateOnly() string {
	return d.time.Format(dateLayout)
}

// เมธอดสำหรับดึงเฉพาะเวลา
func (d *TimeTime) TimeOnly() string {
	return d.time.Format(timeLayout)
}

// เมธอดสำหรับดึงวันเวลาเต็มรูปแบบ
func (d *TimeTime) DateTime() string {
	return d.time.Format(datetimeLayout)
}
