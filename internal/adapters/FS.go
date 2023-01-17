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

func (f FS) DiskInfo() (entity.Info, error) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(string(filepath.Separator), &fs)
	if err != nil {
		return entity.Info{0, 0, 0, ""}, err
	}

	total := fs.Blocks * uint64(fs.Bsize)
	free := fs.Bfree * uint64(fs.Bsize)
	used := (total - free)*100/total

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
		
		if os.IsNotExist(err) {
			logs.Errors = append(logs.Errors, fmt.Sprintf(entity.ErrFolderDoesntExist, folder))
			continue
		}

		dir, err := ioutil.ReadDir(folder)
		if err != nil {
			logs.Errors = append(logs.Errors, fmt.Sprintf(entity.ErrFolderDoesntExist, folder))
		} else {
			for _, d := range dir {
				os.RemoveAll(path.Join([]string{folder, d.Name()}...))
			}
			logs.Info = append(logs.Info, fmt.Sprintf(entity.FolderDeleted, folder))
		}
	}

	return logs
}
