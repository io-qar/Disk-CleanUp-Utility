package entity

import (
	"bytes"
	"log"
)

const (
	TxtAfterClean        string = "Объём занимаемого места: %d%%.\nПосле очистки: %d%%\n\n"
	TxtNotClean          string = "Объём занимаемого места: %d%%.\nОчистка не проводилась\n\n"
	FolderDeleted        string = "Папка '%s' была очищена\n"
	FolderDsntExistError string = "Папка '%s' не была найдена и была пропущена\n"
	PermissionError      string = "Доступ к '%s' не был получен\n"
	CreationError        string = "Ошибка при создании приложения: %v\n"
	RunError             string = "Ошибка выполнения: %v\n"
)

var (
	Buf bytes.Buffer
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)