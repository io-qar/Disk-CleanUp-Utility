package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	
	"github.com/tidwall/gjson"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.json", "Path to a config file")
	flag.Parse()

	content, err := ioutil.ReadFile(configPath) 
	if err != nil {
    fmt.Println(err)
  }

	folders := gjson.Get(string(content), "folders")
	// fmt.Print(folders)

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
}