package setting

import (
	"flag"
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	Title    string
	AllowIps string
	Users    []string
	Currency []string
	Expense  []string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Trello struct {
	AppKey     string
	Token      string
	BoardID    string
	ListNumber int
}

var TrelloSetting = &Trello{}

var cfg *ini.File

func Setup() {
	var err error

	file := confFile()

	cfg, err = ini.Load(file)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("trello", TrelloSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second

}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo err: %v", err)
	}
}

func confFile() string {

	confPtr := flag.String("conf", "", "set private app.ini file")
	flag.Parse()

	if *confPtr != "" {
		return *confPtr
	} else {
		return "conf/app.ini"
	}
}
