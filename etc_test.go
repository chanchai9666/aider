package aider

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "รหัสผ่านที่ถูกต้อง",
			args:    args{password: "mySecretPassword"},
			wantErr: false, // ไม่ควรเกิดข้อผิดพลาด
		},
		{
			name:    "รหัสผ่านว่างเปล่า",
			args:    args{password: ""},
			wantErr: false, // bcrypt ยังสามารถสร้าง hash สำหรับรหัสผ่านว่างเปล่าได้
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashedPassword, err := HashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, ต้องการข้อผิดพลาด = %v", err, tt.wantErr)
				return
			}
			// ตรวจสอบว่ารหัสผ่านที่เข้ารหัสแล้ว ไม่เหมือนกับรหัสผ่านต้นฉบับ
			if tt.args.password == hashedPassword {
				t.Errorf("HashPassword() = %v, ต้องการแตกต่างจากรหัสผ่านต้นฉบับ", hashedPassword)
			}
			// ตรวจสอบว่ารหัสผ่านที่เข้ารหัสแล้ว ตรงกับรหัสผ่านต้นฉบับ โดยใช้ bcrypt
			if err == nil {
				if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(tt.args.password)); err != nil {
					t.Errorf("รหัสผ่านที่เข้ารหัสแล้ว ไม่ตรงกับรหัสผ่านต้นฉบับ")
				}
			}
		})
	}
}

func TestCheckPassword(t *testing.T) {
	// Set up the hashed password for the test cases
	originalPassword := "mySecretPassword"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(originalPassword), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	type args struct {
		password       string
		hashedPassword string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "เคสรหัสผ่านตรงกัน",
			args: args{
				password:       originalPassword,
				hashedPassword: string(hashedPassword),
			},
			want: true,
		},
		{
			name: "เคสรหัสผ่านไม่ตรงกัน",
			args: args{
				password:       "wrongPassword",
				hashedPassword: string(hashedPassword),
			},
			want: false,
		},
		{
			name: "เคสรหัสผ่านว่างเปล่า",
			args: args{
				password:       "",
				hashedPassword: string(hashedPassword),
			},
			want: false,
		},
		{
			name: "เคสรหัสผ่านมีรูปแบบต่างกัน (เช่น ตัวพิมพ์ใหญ่เล็ก)",
			args: args{
				password:       "Mysecretpassword", // Note the different case
				hashedPassword: string(hashedPassword),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPassword(tt.args.password, tt.args.hashedPassword); got != tt.want {
				t.Errorf("CheckPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncryptData(t *testing.T) {
	type args struct {
		plaintext string
		key       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Valid Data", args: args{plaintext: "Secret Message", key: "beee33dfe3640026d2da28b2c002cb9b"}, wantErr: false},
		{name: "Empty Plaintext", args: args{plaintext: "", key: "mySecretKey"}, wantErr: true},        // อาจจะปรับแต่งตามการออกแบบระบบของคุณ
		{name: "Invalid Key", args: args{plaintext: "Secret Message", key: "shortKey"}, wantErr: true}, // ปรับแต่งความยาวของ key ตามการออกแบบระบบของคุณ
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := EncryptData(tt.args.plaintext, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// We cannot check for the exact encrypted data due to randomness of nonce
			// But we can check if the length of the encrypted data is greater than 0
		})
	}
}
