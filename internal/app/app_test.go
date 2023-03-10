package app

import (
	"clean-utility/internal/adapters"
	"testing"
)

func TestApplication_Run(t *testing.T) {
	tests := []struct{
		name    string
		a       Application
		wantErr bool
	}{
		{
			name: "ошибка получения информации о диске",
			a: Application{
				FSService: adapters.NewBadFakeFS(),
				Logger: adapters.NewLogger(),
			},
			wantErr: true,
		},
		{
			name: "ошибка отправки сообщения",
			a: Application{
				FSService: adapters.NewFakeFS(),
				NotificationService: adapters.NewFakeTgBot("5471768780:AAEAbreeE6DDECknHmMrlD2Mfvedb5GIQ-w"),
				MaxVolume: 7,
				Folders: []string{"test"},
				To: "",
				Logger: adapters.NewLogger(),
			},
			wantErr: true,
		},
		{
			name: "успешный вывод сообщения",
			a: Application{
				FSService: adapters.NewFakeFS(),
				NotificationService: adapters.NewFakeTgBot("5471768780:AAEAbreeE6DDECknHmMrlD2Mfvedb5GIQ-w"),
				MaxVolume: 5,
				Folders: []string{"test"},
				To: "1028417962",
				Logger: adapters.NewLogger(),
			},
			wantErr: false,
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
