package adapters

import (
	"clean-utility/internal/entity"
	"clean-utility/internal/interfaces"
)

type FakeFS struct{}

func NewFaleFS() interfaces.FS {
	return FakeFS{}
}

func (f FakeFS) DiskInfo() (entity.Info, error) {
	return entity.Info{
		Total:  100500,
		Free:   5000,
		Used:   95500,
		FSType: "FAKE",
	}, nil
}

func (f FakeFS) ClearedFolders(folders []string) entity.EventsLog {
	logs := entity.EventsLog{}
	logs.Info = append(logs.Info, "фейковая информация")
	return logs
}
