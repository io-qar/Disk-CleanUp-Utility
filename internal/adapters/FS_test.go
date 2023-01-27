package adapters

import (
	"clean-utility/internal/interfaces"
	"testing"
)

func Test_ClearedFolders(t *testing.T) {
	tests := []struct{
		name string
		FSService interfaces.FS
		folders []string
		logs []string
	}{
		{
			name: "test folder doesn't exist",
			FSService: NewFakeFS(),
			folders: []string{"../../test 1"},
			logs: []string{""},
		},
		{
			name: "test folder permission",
			FSService: NewFakeFS(),
			folders: []string{"../../testpermission"},
			logs: []string{""},
		},
		{
			name: "test right folder",
			FSService: NewFakeFS(),
			folders: []string{"../../test"},
			logs: []string{""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if logs := tt.FSService.ClearedFolders(tt.folders); (logs.Errors == nil) /*!= tt.logs*/ {
				t.Errorf("ClearedFolders() error = %v, logs %v", logs, tt.logs)
			}
		})
	}
}