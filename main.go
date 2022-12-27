package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"path"
	"syscall"

	"github.com/tidwall/gjson"
	"github.com/yanzay/tbot/v2"
)

var (
	bot *tbot.Server
	client *tbot.Client
)

func calcUsedDiskVolume(path string) uint64 {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		panic(err)
	}

	all := fs.Blocks * uint64(fs.Bsize)
	free := fs.Bfree * uint64(fs.Bsize)
	used := all - free
	return used*100/all
}

func main() {
	const (
		txtAfterClean string = "Объём занимаемого места: %d%%.\nПосле очистки: %d%%"
		txtNotClean string = "Объём занимаемого места: %d%%.\nОчистка не проводилась"
		folderDoesntExist string = "Папка '%s' не существует. Пропускается..."
		defaultVolume uint64 = 50
	)
	var configPath string

	flag.StringVar(&configPath, "config", "./config.json", "Path to a config file")
	flag.Parse()

	content, err := ioutil.ReadFile(configPath) 
	if err != nil {
		log.Println(err)
	}
	
	token := gjson.Get(string(content), "telegram-bot.token").Str
	channel := gjson.Get(string(content), "telegram-bot.channel").Str
	folders := gjson.Get(string(content), "folders").Array()
	maxVolume := gjson.Get(string(content), "maxVolume").Uint()
	fmt.Println(maxVolume)
	if maxVolume == 0 {
		maxVolume = defaultVolume
	}
	fmt.Println(maxVolume)

	bot = tbot.New(token)
	client = bot.Client()
	
	usedDiskBefore := calcUsedDiskVolume(string(filepath.Separator))
	fmt.Println(usedDiskBefore)
	if maxVolume < usedDiskBefore {
		for _, folder := range folders {
			_, err := os.Stat(folder.Str)
			if os.IsNotExist(err) {
				log.Println(fmt.Sprintf(folderDoesntExist, folder.Str))
				client.SendMessage(channel, fmt.Sprintf(folderDoesntExist, folder.Str))
			} else {
				dir, _ := ioutil.ReadDir(folder.Str)
				for _, d := range dir {
					os.RemoveAll(path.Join([]string{folder.Str, d.Name()}...))
				}
				usedDiskAfter := calcUsedDiskVolume(string(filepath.Separator))
				client.SendMessage(channel, fmt.Sprintf(txtAfterClean, usedDiskBefore, usedDiskAfter))
			}
		}
	} else {
		client.SendMessage(channel, fmt.Sprintf(txtNotClean, usedDiskBefore))
	}
}