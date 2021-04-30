package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Conf struct {
	User       string
	Pass       string
	Host       string
	RemotePath string
	LocalPath  string
}

func NewConf() Conf {
	file, err := os.Open("goscp.json")
	if err != nil {
		fmt.Println("goscp.json文件打開失敗", err)
	}
	defer file.Close()
	dec := json.NewDecoder(file)

	var conf Conf
	for dec.More() {
		err = dec.Decode(&conf)
		if err != nil {
			fmt.Println(err)
		}
	}
	return conf
}

func main() {
	conf := NewConf()

	f := flag.String("p", "", "pull or push")
	flag.Parse()
	flag.Usage()

	switch *f {
	case "pull": //遠程到本地
		fmt.Println("遠程到本地")
		downloadRemoteDirToLocal(conf)
	case "push": //本地到遠程
		fmt.Println("本地到遠程")
		uploadLocalDirToRemote(conf)
	}
}
