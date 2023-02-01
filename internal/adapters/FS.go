package adapters

import (
	"clean-utility/internal/entity"
	"clean-utility/internal/interfaces"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"syscall"
)

type FS struct{}

func NewFS() interfaces.FS {
	return FS{}
}

func (fs FS) ErrorCheckingFS(e error, ls *entity.EventsLog, f string) {
	if os.IsNotExist(e) {
		entity.ErrorLogger.Printf("Папка %s не существует, пропускается", f)
		ls.Errors = append(ls.Errors, fmt.Sprintf(entity.FolderDsntExistError, f))
	} else if os.IsPermission(e) {
		entity.ErrorLogger.Printf("Ошибка доступа к %s, пропускается", f)
		ls.Errors = append(ls.Errors, fmt.Sprintf(entity.PermissionError, f))
	} else {
		entity.ErrorLogger.Printf("Другая ошибка во время очистки каталогов")
		ls.Errors = append(ls.Errors, fmt.Sprintf("Ошибка во время очистки папок: %v\n", e))
	}
}

func (f FS) DiskInfo() (entity.Info, error) {
	entity.InfoLogger.Println("Сбор информации об объёме диска")
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(string(filepath.Separator), &fs)
	if err != nil {
		entity.ErrorLogger.Println("Ошибка во время сбора информации о диске, параметры структуры диска обнулены")
		return entity.Info{0, 0, 0, ""}, err
	}

	total := fs.Blocks * uint64(fs.Bsize)
	free := fs.Bfree * uint64(fs.Bsize)
	used := (total - free) * 100 / total
	entity.InfoLogger.Println("Сбор информации об объёме диска выполнен")

	return entity.Info{
		Total:  total,
		Free:   free,
		Used:   used,
		FSType: "Linux",
	}, nil
}

func (f FS) ClearedFolders(folders []string) entity.EventsLog {
	logs := entity.EventsLog{}

	entity.InfoLogger.Println("Начало очистки каталогов")
	for _, folder := range folders {
		_, err := os.Stat(folder)

		if err != nil {
			f.ErrorCheckingFS(err, &logs, folder)
			continue
		}

		dir, err := os.ReadDir(folder)
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
	entity.InfoLogger.Println("Очистка каталогов завершена")

	return logs
}