package configuration

import (
	"github.com/howeyc/fsnotify"
	"log"
)

type FileEventHandlerConstructor interface {    //�ļ��¼��������߽ӿ�
	new(filePath string) FileEventHandler
}

type FileEventHandler interface {               //�ļ��¼�����ӿ�
	Handle(event *fsnotify.FileEvent)
}
//�������ļ��������߽ṹ��
type BlacklistFileHandlerConstructor struct{}

func (bhc BlacklistFileHandlerConstructor) new(filePath string) FileEventHandler {
	return &BlacklistHandler{
		filePath: filePath,
	}
}

type BlacklistHandler struct {                 //����������ṹ��
	filePath string
}

func (h BlacklistHandler) Handle(event *fsnotify.FileEvent) {
	log.Println(event)                     //��־��¼�ļ��¼�
	log.Println(h.filePath)                //��־��¼�ļ�·��
	if event == nil {
		log.Print("blacklisthandler nil.")
	} else if event.Name == h.filePath && event.IsModify() {
		log.Println("blacklisthandler....")
		initBlacklistConfig()   //��ʼ��
	}
}

func (h BlacklistHandler) String() string {        //�����ļ�·��
	return "black\t" + h.filePath
}
//�������ļ�����   ����ͬ��
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
