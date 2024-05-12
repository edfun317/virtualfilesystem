package domain

import "testing"

func TestContainsSpace(t *testing.T) {

	tests := []struct {
		name   string
		text   string
		wantIt bool
	}{
		{"valid", "testUser", false},
		{"contains spaces", "test user", true},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			isContain := containsSpace(tt.text)
			if isContain != tt.wantIt {
				t.Errorf("containsSpace() for %s, wantIt %v, got: %v", tt.text, tt.wantIt, isContain)
			}
		})

	}
}

func TestValidateName(t *testing.T) {
	tests := []struct {
		name     string
		username string
		wantErr  bool
	}{
		{"valid", "testUser", false},
		{"too short", "tu", true},
		{"too long", "thisiswaytoolongforausername", true},
		{"invalid chars", "user!@#", true},
		{"contains spaces", "test user", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateName(tt.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateUsername() for %s, wantErr %v, error: %v", tt.username, tt.wantErr, err)
			}
		})
	}
}
