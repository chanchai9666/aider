package aider

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	rand2 "math/rand"
	"os"
	"reflect"

	"golang.org/x/crypto/bcrypt"
)

// ฟังก์ชันที่สร้าง map จาก slice ของ struct ใด ๆ โดยกำหนด key และ value ตามที่เลือก
/*
อธิบายโค้ด:

	1.	ฟังก์ชัน CreateMap: ใช้ generics เพื่อให้สามารถสร้าง map จาก slice ของ struct ใดๆ โดยมี key และ value ที่กำหนดได้ตามต้องการ:
	•	K: ประเภทของ key ที่ต้องเป็น comparable.
	•	V: ประเภทของ value.
	•	T: ประเภทของ struct.
	2.	ตัวอย่างโครงสร้างข้อมูล UserLogin และ Product: สร้าง struct UserLogin และ Product เพื่อใช้เป็นตัวอย่างข้อมูล.
	3.	การใช้งานใน main function:
	•	สร้าง slice ของ UserLogin และ Product.
	•	ใช้ CreateMap เพื่อสร้าง map โดยกำหนด key และ value ตามที่ต้องการ:
	•	สร้าง map ที่มี UserId เป็น key และ UserName เป็น value (userIdToUserName).
	•	สร้าง map ที่มี UserName เป็น key และ ProfileName เป็น value (userNameToProfileName).
	•	สร้าง map ที่มี ProductId เป็น key และ ProductName เป็น value (productIdToProductName).
	•	สร้าง map ที่มี Category เป็น key และ Price เป็น value (categoryToPrice).
	4.	แสดงผลลัพธ์: พิมพ์ผลลัพธ์ของ map ที่สร้างขึ้นเพื่อแสดงข้อมูลที่ถูกจัดกลุ่มและจัดเก็บใน map.

โค้ดนี้จะช่วยให้คุณสามารถสร้าง map จาก slice ของ struct ใดๆ ได้โดยกำหนด key และ value ตามที่ต้องการ, ทำให้มีความยืดหยุ่นในการใช้งาน.
*/

func CreateMap[K comparable, V any, T any](items []T, keyFunc func(T) K, valueFunc func(T) V) map[K]V {

	result := make(map[K]V)
	for _, item := range items {
		key := keyFunc(item)
		value := valueFunc(item)
		result[key] = value
	}
	return result

	/*
		EX ตัวอย่างการใช้งาน
			users := []UserLogin{
				{UserId: 1, UserName: "john_doe", ProfileName: "John Doe", Level: 5, ProfileId: 101},
				{UserId: 2, UserName: "jane_doe", ProfileName: "Jane Doe", Level: 10, ProfileId: 102},
				{UserId: 3, UserName: "sam_smith", ProfileName: "Sam Smith", Level: 15, ProfileId: 103},
			}

			userIdToUserName := CreateMap(users, func(u UserLogin) uint {
				return u.UserId
			}, func(u UserLogin) string {
				return u.UserName
			})
	*/
}

// ฟังก์ชันที่สร้าง slice จาก struct ใด ๆ โดยกำหนดค่า value ที่จะนำมาใช้ใน slice นั้นได้
/*
อธิบายโค้ด:

	1.	ฟังก์ชัน CreateSlice: ฟังก์ชันนี้ใช้ generics เพื่อสร้าง slice จาก struct ใดๆ โดยกำหนดค่า value ที่จะนำมาใช้ใน slice:
	•	V: ประเภทของ value.
	•	T: ประเภทของ struct.
	2.	ตัวอย่างโครงสร้างข้อมูล UserLogin และ Product: สร้าง struct UserLogin และ Product เพื่อใช้เป็นตัวอย่างข้อมูล.
	3.	การใช้งานใน main function:
	•	สร้าง slice ของ UserLogin และ Product.
	•	ใช้ CreateSlice เพื่อสร้าง slice โดยกำหนดค่า value ที่จะนำมาใช้ใน slice:
	•	สร้าง slice ที่มี UserName จาก UserLogin (userNames).
	•	สร้าง slice ที่มี ProfileId จาก UserLogin (profileIds).
	•	สร้าง slice ที่มี ProductName จาก Product (productNames).
	•	สร้าง slice ที่มี Price จาก Product (productPrices).
	4.	แสดงผลลัพธ์: พิมพ์ผลลัพธ์ของ slice ที่สร้างขึ้นเพื่อแสดงข้อมูลที่ถูกจัดเก็บใน slice.

โค้ดนี้ทำให้คุณสามารถสร้าง slice จาก struct ใดๆ ได้โดยกำหนดค่า value ตามที่ต้องการ, ทำให้มีความยืดหยุ่นในการใช้งาน.
*/
func CreateSlice[V any, T any](items []T, valueFunc func(T) V) []V {
	var result []V
	for _, item := range items {
		value := valueFunc(item)
		result = append(result, value)
	}
	return result

	/*
		EX ตัวอย่างการใช้งาน
		users := []UserLogin{
			{UserId: 1, UserName: "john_doe", ProfileName: "John Doe", Level: 5, ProfileId: 101},
			{UserId: 2, UserName: "jane_doe", ProfileName: "Jane Doe", Level: 10, ProfileId: 102},
			{UserId: 3, UserName: "sam_smith", ProfileName: "Sam Smith", Level: 15, ProfileId: 103},
		}

		userNames := CreateSlice(users, func(u UserLogin) string {
			return u.UserName
		})
	*/
}

// ToMap แปลง slice ของ struct ใด ๆ ที่ใช้ key เป็นประเภทเดียวกันกับค่าที่ส่งเข้ามา
// K comparable: แทนที่ด้วย type ใด ๆ ก็ได้ที่สามารถใช้เป็น key ใน map (เช่น int, string, etc.)
// T any: แทนที่ด้วย type ใด ๆ ก็ได้ (Generics)
// items []T: slice ของ struct ที่ต้องการแปลง
// keyFunc func(T) K: ฟังก์ชันที่ใช้ในการดึง key จาก struct
/*
อธิบายการทำงานของโค้ด:

	1.	การประกาศฟังก์ชัน ToMap:
	•	ฟังก์ชันนี้ใช้ Generics เพื่อให้สามารถทำงานกับ key และ value ที่เป็น type ใดก็ได้
	•	K comparable: กำหนดให้ K เป็นประเภทที่สามารถใช้เป็น key ใน map ได้ (ต้องสามารถเปรียบเทียบได้ เช่น int, string เป็นต้น)
	•	T any: กำหนดให้ T เป็นประเภทใดก็ได้
	•	items []T: รับ slice ของ struct ที่ต้องการแปลง
	•	keyFunc func(T) K: รับฟังก์ชันที่ใช้ในการดึง key จาก struct โดยคืนค่าเป็นประเภท K
	2.	การสร้าง map:
	•	สร้าง map ที่ชื่อ result โดยมี key เป็นประเภท K และ value เป็นประเภท T
	3.	การวน loop ผ่านแต่ละ item ของ slice:
	•	ใช้ for loop เพื่อวนผ่านแต่ละ item ใน slice items
	•	ใช้ keyFunc เพื่อดึง key จาก item
	•	นำ item ไปเก็บใน map result โดยใช้ key ที่ดึงมาได้
	4.	การคืนค่า:
	•	คืนค่า map ที่สร้างขึ้น

การใช้งาน:

ในฟังก์ชัน main:

	•	สร้าง instance ของ struct User และกำหนดค่า field ต่าง ๆ
	•	เรียกฟังก์ชัน ToMap เพื่อแปลง struct เป็น map โดยใช้ ID (ประเภท uint) เป็น key
	•	เรียกฟังก์ชัน ToMap เพื่อแปลง struct เป็น map โดยใช้ Username (ประเภท string) เป็น key
	•	พิมพ์ผลลัพธ์ที่ได้ออกมา

ผลลัพธ์จะเป็น map ที่มี key เป็นประเภทที่กำหนด (เช่น uint, string) และค่าของ map เป็น struct User ซึ่งสามารถใช้งานต่อไปได้ตามต้องการ
*/
func ToMap[K comparable, T any](items []T, keyFunc func(T) K) map[K]T {
	// สร้าง map ที่จะเก็บผลลัพธ์ โดยมี key เป็นประเภท K และ value เป็น struct ของ type T
	result := make(map[K]T)

	// วน loop ผ่านแต่ละ item ใน slice
	for _, item := range items {
		// ใช้ keyFunc เพื่อดึง key จาก item
		key := keyFunc(item)
		// เก็บ item ลงใน map โดยใช้ key ที่ดึงมาได้
		result[key] = item
	}

	// คืนค่า map ที่แปลงเรียบร้อยแล้ว
	return result

	/*
		EX ตัวอย่างการใช้งาน
				myUsers := []User{
				{ID: 1, Username: "john_doe", Email: "john@example.com", Age: 30, Active: true},
				{ID: 2, Username: "jane_doe", Email: "jane@example.com", Age: 25, Active: false},
				{ID: 3, Username: "sam_smith", Email: "sam@example.com", Age: 40, Active: true},
			}

			userMapByID := ToMap(myUsers, func(u User) uint {
				return u.ID
			})
	*/
}

// ToTripleNestedMap แปลง slice ของ struct ใด ๆ ที่ใช้ key เป็นประเภทเดียวกันกับค่าที่ส่งเข้ามาเป็น map ซ้อนกันสามชั้น
// K1, K2, K3 comparable: แทนที่ด้วย type ใด ๆ ก็ได้ที่สามารถใช้เป็น key ใน map (เช่น int, string, etc.)
// T any: แทนที่ด้วย type ใด ๆ ก็ได้ (Generics)
// items []T: slice ของ struct ที่ต้องการแปลง
// keyFunc1, keyFunc2, keyFunc3 func(T) K: ฟังก์ชันที่ใช้ในการดึง key จาก struct
func ToTripleNestedMap[K1 comparable, K2 comparable, K3 comparable, T any](items []T, keyFunc1 func(T) K1, keyFunc2 func(T) K2, keyFunc3 func(T) K3) map[K1]map[K2]map[K3]T {
	// สร้าง map ที่จะเก็บผลลัพธ์ โดยมี key เป็นประเภท K1 และ value เป็น map[K2]map[K3]T
	result := make(map[K1]map[K2]map[K3]T)

	// วน loop ผ่านแต่ละ item ใน slice
	for _, item := range items {
		// ใช้ keyFunc1, keyFunc2 และ keyFunc3 เพื่อดึง key จาก item
		key1 := keyFunc1(item)
		key2 := keyFunc2(item)
		key3 := keyFunc3(item)

		// ตรวจสอบว่า key1 มีอยู่ใน map result หรือไม่ ถ้าไม่มีให้สร้าง map ใหม่
		if _, exists := result[key1]; !exists {
			result[key1] = make(map[K2]map[K3]T)
		}

		// ตรวจสอบว่า key2 มีอยู่ใน map result[key1] หรือไม่ ถ้าไม่มีให้สร้าง map ใหม่
		if _, exists := result[key1][key2]; !exists {
			result[key1][key2] = make(map[K3]T)
		}

		// เก็บ item ลงใน map ซ้อนกันโดยใช้ key1, key2 และ key3 ที่ดึงมาได้
		result[key1][key2][key3] = item
	}

	// คืนค่า map ที่แปลงเรียบร้อยแล้ว
	return result

	/*
			อธิบายการทำงานของโค้ด:

			1.	การประกาศฟังก์ชัน ToTripleNestedMap:
			•	ฟังก์ชันนี้ใช้ Generics เพื่อให้สามารถทำงานกับ key และ value ที่เป็น type ใดก็ได้
			•	K1, K2, K3 comparable: กำหนดให้ K1, K2 และ K3 เป็นประเภทที่สามารถใช้เป็น key ใน map ได้ (ต้องสามารถเปรียบเทียบได้ เช่น int, string เป็นต้น)
			•	T any: กำหนดให้ T เป็นประเภทใดก็ได้
			•	items []T: รับ slice ของ struct ที่ต้องการแปลง
			•	keyFunc1, keyFunc2, keyFunc3 func(T) K: รับฟังก์ชันที่ใช้ในการดึง key จาก struct โดยคืนค่าเป็นประเภท K1, K2 และ K3
			2.	การสร้าง map:
			•	สร้าง map ที่ชื่อ result โดยมี key เป็นประเภท K1 และ value เป็น map ที่มี key เป็นประเภท K2 และ value เป็น map ที่มี key เป็นประเภท K3 และ value เป็นประเภท T
			3.	การวน loop ผ่านแต่ละ item ของ slice:
			•	ใช้ for loop เพื่อวนผ่านแต่ละ item ใน slice items
			•	ใช้ keyFunc1, keyFunc2 และ keyFunc3 เพื่อดึง key จาก item
			•	ตรวจสอบว่า key1 มีอยู่ใน map result หรือไม่ ถ้าไม่มีให้สร้าง map ใหม่สำหรับ key1
			•	ตรวจสอบว่า key2 มีอยู่ใน map result[key1] หรือไม่ ถ้าไม่มีให้สร้าง map ใหม่สำหรับ key2
			•	นำ item ไปเก็บใน map ซ้อนกันโดยใช้ key1, key2 และ key3 ที่ดึงมาได้
			4.	การคืนค่า:
			•	คืนค่า map ที่สร้างขึ้น

		การใช้งาน:

		ในฟังก์ชัน main:

			•	สร้าง instance ของ struct User และกำหนดค่า field ต่าง ๆ
			•	เรียกฟังก์ชัน ToTripleNestedMap เพื่อแปลง struct เป็น triple nested map โดยใช้ Username (ประเภท string) เป็น key ชั้นแรก, Email (ประเภท string) เป็น key ชั้นสอง และ ID (ประเภท uint) เป็น key ชั้นสาม
			•	เรียกฟังก์ชัน ToTripleNestedMap เพื่อแปลง struct เป็น triple nested map โดยใช้ Age (ประเภท int) เป็น key ชั้นแรก, Active (ประเภท bool) เป็น key ชั้นสอง และ ID (ประเภท uint) เป็น key ชั้นสาม
			•	พิมพ์ผลลัพธ์ที่ได้ออกมา

		ผลลัพธ์จะเป็น triple nested map ที่มี key ซ้อนกันตามที่กำหนด (เช่น Username -> Email -> ID, Age -> Active -> ID) และค่าของ map เป็น struct User ซึ่งสามารถใช้งานต่อไปได้ตามต้องการ

	*/
}

// ตรวจสอบค่า ว่ามีใน Slice ไหม
func InSlice[T comparable](value T, array []T) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
	/*
		Ex.
		ตัวอย่างการใช้งานกับ slice ของ int
		intSlice := []int{1, 2, 3, 4, 5}
		fmt.Println(InSlice(3, intSlice)) // true
		fmt.Println(InSlice(6, intSlice)) // false

		ตัวอย่างการใช้งานกับ slice ของ string
		stringSlice := []string{"apple", "banana", "cherry"}
		fmt.Println(InSlice("banana", stringSlice)) // true
		fmt.Println(InSlice("grape", stringSlice))  // false

		ตัวอย่างการใช้งานกับ slice ของ custom type
		type Person struct {
			Name string
			Age  int
		}

		personSlice := []Person{
			{"Alice", 30},
			{"Bob", 25},
			{"Charlie", 35},
		}

		fmt.Println(InSlice(Person{"Bob", 25}, personSlice))       // true
		fmt.Println(InSlice(Person{"Dave", 40}, personSlice))      // false
	*/
}

func DD(values ...interface{}) {
	for k, v := range values {
		fmt.Println("==x0xx0xx0xx0xx0xx0xx0xx0xx0x==", k, "==x0xx0xx0xx0xx0xx0xx0xx0xx0x==")
		b, err := json.MarshalIndent(v, "", "  ")
		if err == nil {
			fmt.Println(string(b))
		}
	}

	fmt.Println("DD Print Exit")
	os.Exit(0)
}

func DDD(v ...interface{}) (err error) {
	for k, v := range v {
		fmt.Println("==x0xx0xx0xx0xx0xx0xx0xx0xx0x==", k, "==x0xx0xx0xx0xx0xx0xx0xx0xx0x==")
		b, err := json.MarshalIndent(v, "", "  ")
		if err == nil {
			fmt.Println(string(b))
		}
	}
	return
}

// HashPassword เข้ารหัสรหัสผ่านโดยใช้ bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword ตรวจสอบว่ารหัสผ่านตรงกับรหัสผ่านที่เข้ารหัส
func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// สุ่ม string ตามจำนวนที่ต้องการ
func RandomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand2.Intn(len(letters))]
	}
	return string(b)
}

// เข้ารหัสข้อความ
func EncryptData(plaintext, key string) (string, error) {
	// keyProfile: เก็บค่าของ key (ไม่จำเป็นต้องมีการเปลี่ยนแปลง)
	keyProfile := key
	// newPlaintext: แปลงข้อความธรรมดา (plaintext) เป็น byte array
	newPlaintext := []byte(plaintext)
	// newKey: แปลง key เป็น byte array
	newKey := []byte(keyProfile)

	// สร้าง cipher block จาก key
	block, err := aes.NewCipher(newKey)
	if err != nil {
		return "", err
	}

	// สร้างตัว encryptor แบบ AES-GCM จาก cipher block
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// สร้างค่า nonce แบบสุ่ม
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// เข้ารหัสข้อความโดยใช้ AES-GCM พร้อมกับ nonce
	ciphered := gcm.Seal(nonce, nonce, newPlaintext, nil)

	// แปลงข้อมูลที่เข้ารหัสเป็น string base64
	return base64.StdEncoding.EncodeToString(ciphered), nil
}

// ถอดรหัสข้อตวาม
func DecryptData(cipheredStr, key string) (string, error) {
	keyProfile := key
	newKey := []byte(keyProfile)
	// แปลงข้อมูลที่เข้ารหัสกลับมาเป็น byte array
	ciphered, err := base64.StdEncoding.DecodeString(cipheredStr)
	if err != nil {
		return "", err
	}

	// สร้าง cipher block จาก key
	block, err := aes.NewCipher(newKey)
	if err != nil {
		return "", err
	}

	// สร้างตัว decryptor แบบ AES-GCM จาก cipher block
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// ตรวจสอบความยาวของข้อมูลที่เข้ารหัส
	nonceSize := gcm.NonceSize()
	if len(ciphered) < nonceSize {
		return "", fmt.Errorf("ciphered too short")
	}

	// แยก nonce ออกจากข้อมูลที่เข้ารหัส
	nonce, ciphered := ciphered[:nonceSize], ciphered[nonceSize:]

	// แก้รหัสข้อความโดยใช้ AES-GCM พร้อมกับ nonce
	plaintext, err := gcm.Open(nil, nonce, ciphered, nil)
	if err != nil {
		return "", err
	}

	// แปลงข้อความที่ถอดรหัสเป็น string
	return string(plaintext), nil
}

// ตรวจสอบ ความถูกต้องของรหัส
func IsEncrypt(key []byte, secure string) bool {
	cipherText, err := base64.RawStdEncoding.DecodeString(secure)
	if err != nil {
		// return secure, err
		log.Println("err cipherText", err)
		return false
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		// return secure, err
		log.Println("err block", err)
		return false
	}

	if len(cipherText) < aes.BlockSize {
		// return secure, err
		log.Println("err cipherText < aes.BlockSize", err)
		return false
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	// return string(cipherText), nil
	return true
}

// ฟังก์ชันแปลง struct ใดๆ ให้เป็น map[string]interface{}
func StructToMapInterface(data interface{}) map[string]interface{} {
	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)

	result := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		result[field.Name] = v.Field(i).Interface()
	}
	return result
}

// แปลง String เป็น Byte
func ToByte(value string) []byte {
	return []byte(value)
}
