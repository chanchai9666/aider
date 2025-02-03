package aider

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtConfig struct {
	ExpirationTime int64  //เวลาหมดอายุของ JWT (แบบ Unix)
	Audience       string //ผู้รับที่ JWT นี้ถูกออกให้ ใคร ?
	Issuer         string //ผู้สร้าง JWT นี้
}

// ฟังก์ชันสำหรับสร้าง JWT Token รองรับ struct ใดๆ
// GenerateJWT สร้าง JWT Token
func GenerateJWT[T any](jwtKey []byte, conFig JwtConfig, claimsStruct T) (string, error) {
	claimsMap := StructToMapInterface(claimsStruct)
	// กำหนด Expiration ของ token (เช่น 15 นาที)
	// expirationTime := time.Now().Add(15 * time.Minute).Unix()

	// กำหนด IssuedAt, Expiration, Audience, และ Issuer
	claimsMap["iat"] = TimeTimeNow().Unix()  //เวลาที่สร้าง
	claimsMap["exp"] = conFig.ExpirationTime //เวลาที่หมดอายุ
	claimsMap["aud"] = conFig.Audience       // ผู้รับที่ JWT นี้ถูกออกให้ ใคร ?
	claimsMap["iss"] = conFig.Issuer         // ผู้สร้าง JWT นี้

	// สร้าง token พร้อมกับ claims และเซ็นด้วย secret key
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims(claimsMap))
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
		// ตรวจสอบว่า algorithm ตรงกับที่คาดหวัง
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	// ตรวจสอบว่าการยืนยัน token สำเร็จหรือไม่
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// ตรวจสอบ Expiration และ Issuer
	if claims.ExpiresAt.Before(time.Now()) {
		return nil, fmt.Errorf("token expired")
	}

	if claims.Issuer != "your-issuer222" {
		return nil, fmt.Errorf("invalid issuer")
	}

	if len(claims.Audience) == 0 {
		return nil, fmt.Errorf("audience not found")
	}

	if claims.Audience[0] != "your-audience111" {
		return nil, fmt.Errorf("invalid audience")
	}

	return claims, nil
}
