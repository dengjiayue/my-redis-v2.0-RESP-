package src

import (
	"fmt"
	"os"
	"sync"
	"time"
)

//my-redis的数据持久化

//aof持久化

const ( // 保存策略
	AofFsyncAlways   = iota // 每次都执行
	AofFsyncEverySec        // 每秒执行一次
	AofFsyncNo              // 不主动执行
)

type AofPersister struct {
	//aof文件
	AofFile *os.File
	//锁
	AofLock sync.Mutex
	//保存策略
	AofFsync int
	//数据缓冲区(每秒执行一次时使用)
	AofBuf *AofBuf

	// isRewriting 重写状态
	IsRewriting bool
}

type AofBuf struct {
	Lock sync.RWMutex
	Data []byte
}

// 保存数据
func (p *AofPersister) SaveData(cmdLine []byte) {
	//lock
	p.AofLock.Lock()
	defer p.AofLock.Unlock()
	if p.AofFsync == AofFsyncEverySec || p.IsRewriting { //重写的时候也写缓冲区
		p.AofBuf.Lock.Lock()
		defer p.AofBuf.Lock.Unlock()
		p.AofBuf.Data = append(p.AofBuf.Data, cmdLine...)
	} else if p.AofFsync == AofFsyncAlways {
		_, err := p.AofFile.Write(cmdLine)
		if err != nil {
			panic(err)
		}
	}
}

// 定时刷盘
func (p *AofPersister) Flush() {
	if p.AofFsync == AofFsyncEverySec {
		go func() {
			//定时器
			timer := time.NewTicker(time.Second)
			for {
				<-timer.C
				p.AofBuf.Lock.Lock()
				_, err := p.AofFile.Write(p.AofBuf.Data)
				if err != nil {
					panic(err)
				}
				p.AofBuf.Data = p.AofBuf.Data[:0]
				p.AofBuf.Lock.Unlock()
			}

		}()
	}
}

// aof重写
// 使用数据生成命令写到一个新的文件(临时)中,最后使用临时文件替换原来的文件

// 重写
func (p *AofPersister) Rewrite(s *Server) {
	p.IsRewriting = true
	defer func() {
		p.IsRewriting = false
	}()
	// 2. 准备重写
	// 2.1 关闭文件
	_ = p.AofFile.Close()
	// 2.2 打开文件
	aofFile, err := os.OpenFile(p.AofFile.Name(), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	//TODO 向文件写入数据
	s.WriteAof(aofFile)
	// 2.3 关闭文件
	_ = aofFile.Close()
	// 3. 重写完成
	// 3.1 重命名文件
	_ = os.Rename(p.AofFile.Name(), p.AofFile.Name()+".tmp")
	// 3.2 打开文件
	p.AofFile, err = os.OpenFile(p.AofFile.Name()+".tmp", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	// 替换aof
	p.AofFile = aofFile
}

// 读取数据并写入文件
func (s *Server) WriteAof(file *os.File) {
	for k, v := range s.M {
		//拼接成命令的形式
		cmdLine := fmt.Sprintf("*%d\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n", 3, len("set"), "set", len(k), k, len(v), v)
		//写入文件
		_, err := file.Write([]byte(cmdLine))
		if err != nil {
			panic(err)
		}
	}
}
