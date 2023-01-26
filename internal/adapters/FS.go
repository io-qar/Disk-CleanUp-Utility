package adapters

import (
	"clean-utility/internal/entity"
	"clean-utility/internal/interfaces"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"syscall"
)

type FS struct{}

func NewFS() interfaces.FS {
	return FS{}
}

func ErrorCheckingFS(e error, ls entity.EventsLog, f string) []string {
	if os.IsNotExist(e) {
		ls.Errors = append(ls.Errors, fmt.Sprintf(entity.FolderDsntExistError, f))
	} else if os.IsPermission(e) {
		ls.Errors = append(ls.Errors, fmt.Sprintf(entity.PermissionError, f))
	} else {
		ls.Errors = append(ls.Errors, fmt.Sprintf("Ошибка во время очистки папок: %v\n", e))
	}

	return ls.Errors
}

func (f FS) DiskInfo() (entity.Info, error) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(string(filepath.Separator), &fs)
	if err != nil {
		return entity.Info{0, 0, 0, ""}, err
	}

	total := fs.Blocks * uint64(fs.Bsize)
	free := fs.Bfree * uint64(fs.Bsize)
	used := (total - free) * 100 / total

	return entity.Info{
		Total:  total,
		Free:   free,
		Used:   used,
		FSType: "Linux",
	}, nil
}

func (f FS) ClearedFolders(folders []string) entity.EventsLog {
	logs := entity.EventsLog{}

	for _, folder := range folders {
		_, err := os.Stat(folder)
		
		if err != nil {
			logs.Errors = append(ErrorCheckingFS(err, logs, folder))
			continue
		}

		dir, err := ioutil.ReadDir(folder)
		if err != nil {
			logs.Errors = append(logs.Errors, fmt.Sprintf(entity.FolderDsntExistError, folder))
			continue
		}
		
		for _, d := range dir {
			err := os.RemoveAll(path.Join([]string{folder, d.Name()}...))
			if err != nil {
				logs.Errors = append(ErrorCheckingFS(err, logs, folder))
				continue
			}
			logs.Info = append(logs.Info, fmt.Sprintf(entity.FolderDeleted, folder))
		}
	}

	return logs
}
