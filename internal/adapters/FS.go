package adapters

import (
	"clean-utility/internal/entity"
	"clean-utility/internal/interfaces"
)

type FS struct{}

func NewFS() interfaces.FS {
	return FS{}
}

func (f FS) DiskInfo() (entity.Info, error) {
	return entity.Info{
		Total:  100500,
		Free:   5000,
		Used:   95500,
		FSType: "REAL",
	}, nil
}

func (f FS) ClearedFolders(folders []string) entity.EventsLog {
	logs := entity.EventsLog{}
	logs.Info = append(logs.Info, "реальная информация")
	return logs
}
