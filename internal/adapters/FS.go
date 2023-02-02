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
	logger := NewLogger()

	if os.IsNotExist(e) {
		logger.Error("Папка %s не существует, пропускается", f)
		ls.Errors = append(ls.Errors, fmt.Sprintf(entity.FolderDsntExistError, f))
	} else if os.IsPermission(e) {
		logger.Error("Ошибка доступа к %s, пропускается", f)
		ls.Errors = append(ls.Errors, fmt.Sprintf(entity.PermissionError, f))
	} else {
		logger.Error("Другая ошибка во время очистки каталогов")
		ls.Errors = append(ls.Errors, fmt.Sprintf("Ошибка во время очистки папок: %v\n", e))
	}
}

func (f FS) DiskInfo() (entity.Info, error) {
	logger := NewLogger()
	
	logger.Info("Сбор информации об объёме диска")
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(string(filepath.Separator), &fs)
	if err != nil {
		logger.Error("Ошибка во время сбора информации о диске, параметры структуры диска обнулены")
		return entity.Info{0, 0, 0, ""}, err
	}

	total := fs.Blocks * uint64(fs.Bsize)
	free := fs.Bfree * uint64(fs.Bsize)
	used := (total - free) * 100 / total
	logger.Info("Сбор информации об объёме диска выполнен")

	return entity.Info{
		Total:  total,
		Free:   free,
		Used:   used,
		FSType: "Linux",
	}, nil
}

func (f FS) ClearedFolders(folders []string) entity.EventsLog {
	logs := entity.EventsLog{}
	logger := NewLogger()

	logger.Info("Начало очистки каталогов")
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
			logger.Info("Удаление %s", d.Name())
			err := os.RemoveAll(path.Join([]string{folder, d.Name()}...))
			if err != nil {
				f.ErrorCheckingFS(err, &logs, folder)
				continue
			}
			logger.Info("Удаление %s завершено", d.Name())
			logs.Info = append(logs.Info, fmt.Sprintf(entity.FolderDeleted, folder))
		}
	}
	logger.Info("Очистка каталогов завершена")

	return logs
}