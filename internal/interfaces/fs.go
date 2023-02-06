package interfaces

import "clean-utility/internal/entity"

type FS interface {
	DiskInfo() (entity.Info, error)
	ClearedFolders([]string) entity.EventsLog
}
