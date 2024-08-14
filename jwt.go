package aider

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ฟังก์ชันสำหรับสร้าง JWT Token รองรับ struct ใดๆ
func GenerateJWT(jwtKey []byte, claimsStruct interface{}) (string, error) {
	// แปลง struct เป็น map เพื่อใช้เป็น claims
	claimsMap := StructToMapInterface(claimsStruct)

	// กำหนด Expiration ของ token
	claimsMap["iat"] = TimeNowLocationTH().Unix()
	claimsMap["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// สร้าง token พร้อมกับ claims และเซ็นด้วย secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claimsMap))
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ฟังก์ชันสำหรับตรวจสอบ JWT Token
func VerifyJWT(jwtKey []byte, tokenString string) (*jwt.RegisteredClaims, error) {
	claims := &jwt.RegisteredClaims{}

	// ตรวจสอบและแปลง token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	// ตรวจสอบว่าการยืนยัน token สำเร็จหรือไม่
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
