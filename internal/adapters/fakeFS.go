package adapters

import (
	"clean-utility/internal/entity"
	"clean-utility/internal/interfaces"
	"errors"
	"fmt"
	"os"
	"path"
)

func (fs FakeFS) ErrorCheckingFS(e error, ls *entity.EventsLog, f string) {
	if os.IsNotExist(e) {
		ls.Errors = append(ls.Errors, fmt.Sprintf(entity.FolderDsntExistError, f))
	} else if os.IsPermission(e) {
		ls.Errors = append(ls.Errors, fmt.Sprintf(entity.PermissionError, f))
	} else {
		ls.Errors = append(ls.Errors, fmt.Sprintf("Ошибка во время очистки папок: %v\n", e))
	}
}

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

	for _, folder := range folders {
		_, err := os.Stat(folder)

		if err != nil {
			f.ErrorCheckingFS(err, &logs, folder)
			continue
		}

		dir, err := os.ReadDir(folder)
		f.ErrorCheckingFS(errors.New("timeout"), &logs, folder)
		if err != nil {
			logs.Errors = append(logs.Errors, fmt.Sprintf(entity.FolderDsntExistError, folder))
			continue
		}

		for _, d := range dir {
			err := os.RemoveAll(path.Join([]string{folder, d.Name()}...))
			if err != nil {
				f.ErrorCheckingFS(err, &logs, folder)
				continue
			}
			logs.Info = append(logs.Info, fmt.Sprintf(entity.FolderDeleted, folder))
		}
	}

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
