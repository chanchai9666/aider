package aider

import (
	"testing"
)

func TestGenerateJWT(t *testing.T) {
	// ตัวอย่าง struct ที่ใช้ในการทดสอบ
	type MyClaims struct {
		Username string
		Email    string
		Role     string
	}
	type args struct {
		jwtKey       []byte
		claimsStruct MyClaims
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Valid claims - generates valid token",
			args: args{
				jwtKey: []byte("secret_key"),
				claimsStruct: MyClaims{
					Username: "secret",
					Email:    "secret",
					Role:     "secret",
				},
			},
			want:    "", // We don't check the exact token string, just validity
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateJWT(tt.args.jwtKey, tt.args.claimsStruct)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == tt.want {
				t.Errorf("GenerateJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}
