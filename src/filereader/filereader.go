package filereader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type Response struct {
	Code int `json:"code"`
	Params string `json:"params"`
	Msg string `json:"msg"`
	Data []interface{} `json:"data"`
}

type JsonConfig struct {
	Api string `json:"api"`
	Response Response `json:"response"`
}

var Res []JsonConfig
var file_locker sync.Mutex

func New() []JsonConfig{
	path := GetAppPath()
	conf,status :=LoadConfig(path+"/api/api.json")
	if status == false {
		fmt.Println("配置文件读取失败")
		time.Sleep(2*time.Second)
		os.Exit(1)
	}
	Res = conf
	return Res
}

func GetAppPath() string{
	fmt.Println("正在获取当前路径...")
	file, _ := exec.LookPath(os.Args[0])

	path, _ := filepath.Abs(file)

	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}

func LoadConfig(filename string) ([]JsonConfig, bool) {
	fmt.Println("正在读取接口信息...")
	var conf []JsonConfig
	file_locker.Lock()
	data, err := ioutil.ReadFile(filename) //read config file
	file_locker.Unlock()
	if err != nil {
		fmt.Println("read json file error")
		return conf, false
	}
	datajson := []byte(data)

	err = json.Unmarshal(datajson, &conf)

	if err != nil {
		fmt.Println("unmarshal json file error")
		return conf, false
	}
	fmt.Println("接口信息获取完毕...")
	return conf, true
}

