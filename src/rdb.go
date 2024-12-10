package src

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
)

//实现RDB持久化

// rdb持久化的结构体
type RdbPersister struct {
	//rdb文件
	RdbFile *os.File
	//锁
	RdbLock sync.Mutex
	//数据缓冲区
	RdbBuf *map[string]string // 保存键值对,持久化完成加载到主存中
	//是否正在持久化
	IsRewriting bool
}

// 保存数据
func (p *RdbPersister) SaveData(m map[string]string) {
	//lock
	p.RdbLock.Lock()
	defer p.RdbLock.Unlock()
	//二进制打包
	data, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	p.RdbBuf = &m
	//写入文件
	_, err = p.RdbFile.Write(data)
	if err != nil {
		panic(err)
	}
}

// 重新加载数据
func (p *RdbPersister) LoadData() map[string]string {
	//lock
	p.RdbLock.Lock()
	defer p.RdbLock.Unlock()
	//读取文件
	data, err := ioutil.ReadFile(p.RdbFile.Name())
	if err != nil {
		panic(err)
	}
	//二进制解包
	var m map[string]string
	err = json.Unmarshal(data, &m)
	if err != nil {
		panic(err)
	}
	return m
}
