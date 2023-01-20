package app

import (
	"clean-utility/internal/adapters"
	"testing"
)

func TestApplication_Run(t *testing.T) {
	tests := []struct {
		name    string
		a       Application
		wantErr bool
	}{
		{
			name: "ошибка получения информации о диске",
			a: Application{
				FSService: adapters.NewBadFakeFS(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Application.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
