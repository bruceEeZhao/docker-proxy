package configuration

import (
	"github.com/howeyc/fsnotify"
	"log"
)

type FileEventHandlerConstructor interface {    //文件事件处理建设者接口
	new(filePath string) FileEventHandler
}

type FileEventHandler interface {               //文件事件处理接口
	Handle(event *fsnotify.FileEvent)
}
//黑名单文件处理建设者结构体
type BlacklistFileHandlerConstructor struct{}

func (bhc BlacklistFileHandlerConstructor) new(filePath string) FileEventHandler {
	return &BlacklistHandler{
		filePath: filePath,
	}
}

type BlacklistHandler struct {                 //黑名单管理结构体
	filePath string
}

func (h BlacklistHandler) Handle(event *fsnotify.FileEvent) {
	log.Println(event)                     //日志记录文件事件
	log.Println(h.filePath)                //日志记录文件路径
	if event == nil {
		log.Print("blacklisthandler nil.")
	} else if event.Name == h.filePath && event.IsModify() {
		log.Println("blacklisthandler....")
		initBlacklistConfig()   //初始化
	}
}

func (h BlacklistHandler) String() string {        //返回文件路径
	return "black\t" + h.filePath
}
//白名单文件管理   以下同上
type WhitelistFileHandlerConstructor struct{}

func (whc WhitelistFileHandlerConstructor) new(filePath string) FileEventHandler {
	return &WhitelistHandler{
		filePath: filePath,
	}
}

type WhitelistHandler struct {
	filePath string
}

func (h WhitelistHandler) Handle(event *fsnotify.FileEvent) {
	if event == nil {
		log.Print("whitelisthandler nil.")
	} else if event.Name == h.filePath && event.IsModify() {
		initWhitelistConfig()
	}
}

func (h WhitelistHandler) String() string {
	return "white\t" + h.filePath
}
