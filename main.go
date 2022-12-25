package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
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
	var configPath string

	flag.StringVar(&configPath, "config", "./config.json", "Path to a config file")
	flag.Parse()

	content, err := ioutil.ReadFile(configPath) 
	if err != nil {
		fmt.Println(err)
	}
	
	token := gjson.Get(string(content), "telegram-bot.token").Str
	channel := gjson.Get(string(content), "telegram-bot.channel").Str
	folders := gjson.Get(string(content), "folders").Array()
	maxVolume := gjson.Get(string(content), "maxVolume").Uint()
	bot = tbot.New(token)
	client = bot.Client()
	
	usedDiskBefore := calcUsedDiskVolume("/")
	
	if maxVolume < usedDiskBefore {
		for _, folder := range folders {
			err = os.RemoveAll(folder.String())
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(folder, "was removed")

			err = os.Mkdir(folder.String(), 0755)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(folder, "was created again")
		}
		usedDiskAfter := calcUsedDiskVolume("/")

		client.SendMessage(channel, fmt.Sprintf("Объём занимаемого места: %d%%.\nПосле очистки: %d%%", usedDiskBefore, usedDiskAfter))
	} else {
		// fmt.Println("!=")
		client.SendMessage(channel, fmt.Sprintf("Объём занимаемого места: %d%%.\nОчистка не проводилась", usedDiskBefore))
	}
	err = bot.Start()
	if err != nil {
		panic(err)
	}
}