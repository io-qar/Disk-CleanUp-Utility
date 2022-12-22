package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	unix "golang.org/x/sys/unix"
	
	"github.com/tidwall/gjson"
	"github.com/yanzay/tbot/v2"
)

var (
	bot *tbot.Server
	client *tbot.Client
)

func main() {
	var (
		configPath string
		stat unix.Statfs_t
	)

	flag.StringVar(&configPath, "config", "./config.json", "Path to a config file")
	flag.Parse()

	content, err := ioutil.ReadFile(configPath) 
	if err != nil {
    fmt.Println(err)
  }
	
	token := gjson.Get(string(content), "telegram-bot.token")
	bot = tbot.New(token.Str)
	client = bot.Client()

	wd, err := os.Getwd()
	unix.Statfs(wd, &stat)

	// Available blocks * size per block = available space in bytes
	fmt.Println(stat.Bavail * uint64(stat.Bsize))
	
	folders := gjson.Get(string(content), "folders")
	for _, folder := range folders.Array() {
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
	
	bot.HandleMessage("/start", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "hi")
	})
	err = bot.Start()
	if err != nil {
		panic(err)
	}
}