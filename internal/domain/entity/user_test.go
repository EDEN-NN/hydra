package entity

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	tests := []struct {
		username string
		password string
		email    string
		name     string
		wantErr  bool
	}{
		{"osvaldinho", "osvaldo123", "osvaldo@gmail.com", "Osvaldo Nunes", false},
		{"os", "osvaldo123", "osvaldo@gmail.com", "Osvaldo Nunes", true},
		{"osvaldinho", "senha", "osvaldo@gmail.com", "Osvaldo Nunes", true},
		{"osvaldinho", "osvaldo123", "osvaldomail.com", "Osvaldo Nunes", true},
		{"osvaldinho", "osvaldo123", "osvaldo@gmail.com", "Os", true},
	}

	for _, tt := range tests {
		_, err := CreateUser(tt.username, tt.password, tt.email, tt.name)
		if (err != nil) != tt.wantErr {
			t.Errorf("CreateUser(%v) error = %v, wantErr %v", tt, err, tt.wantErr)
		}
	}
}
