package adapters

import (
	"clean-utility/internal/entity"
	"clean-utility/internal/interfaces"
	"errors"
)

type FakeFS struct{}

func NewFakeFS() interfaces.FS {
	return FakeFS{}
}

func (f FakeFS) DiskInfo() (entity.Info, error) {
	return entity.Info{
		Total:  10,
		Free:   4,
		Used:   6,
		FSType: "Windows",
	}, nil
}

func (f FakeFS) ClearedFolders(folders []string) entity.EventsLog {
	logs := entity.EventsLog{}
	logs.Info = append(logs.Info, "фейковая информация")
	return logs
}

type BadFakeFS struct{}

func NewBadFakeFS() interfaces.FS {
	return BadFakeFS{}
}

func (bf BadFakeFS) DiskInfo() (entity.Info, error) {
	return entity.Info{}, errors.New("fake error")
}

func (bf BadFakeFS) ClearedFolders(folders []string) entity.EventsLog {
	return entity.EventsLog{}
}
