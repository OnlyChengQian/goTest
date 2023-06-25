package command

import (
	facade2 "advt/internal/facade"
	"log"
	"os"
	"sync"
)

var facade facade2.InterfaceFacade

// VAT费用系数
var vat = map[string]float64{
	"UK": 1.2,
	"DE": 1.19,
	"FR": 1.2,
	"IT": 1.22,
	"ES": 1.21,
	"PL": 1.23,
	"RO": 1.19,
	"NL": 1.2,
}

// CBT费用系数
var cbt = map[string]float64{
	"US": 0.013,
	"CA": 0.013,
	"UK": 0.0145,
	"AU": 0.014,
	"DE": 0.0175,
	"ES": 0.0175,
	"IT": 0.0175,
	"FR": 0.0175,
}

// 站点成交费率
var siteFvfRate = map[string]float64{
	"US": 0.1255,
	"CA": 0.1255,
	"UK": 0.128,
	"AU": 0.134,
	"FR": 0.12,
	"DE": 0.12,
	"IT": 0.12,
	"ES": 0.12,
}

var accountCFG = []string{"athleticsshop2018", "outdoormarket", "highdealcrafts", "idealsgarden", "stgreatdeals", "amart2010", //uk
	"craft_mall", "diyhomedecor_uk", "qimpservices", "cldepot", "suntekstore_uk", "worldwidelectro", "homedecorlover", "healthbeauty20",
	"lucky3celectronics", "trandetree007uk", "sportsmod", "sfcautoparts", "protoyszoom", "szsfc-sp", "hobbiewoo", "runningtolife",
	"aclmart", "jollyojolly", "motorsonline2020", "enjoybuyhome", "injoyjo", "ecdepot", "homedecoronline_20", "idealhome2016", "petsloveit",
	"tooysgoing2016", "hobbiesgogo", "deardolity", "flyforevermall", "carhouseuk", "hobby2023", "motohomeau"}

func init() {
	var once sync.Once
	once.Do(func() {
		facade = facade2.NewFacade()
	})
}

type ScriptProvider interface {
	handle() error
}

func Run() {
	args := os.Args
	command := args[1]
	if command == "" {
		log.Fatalln("请输入需要执行的脚本")
	}
	script := scriptList[command]
	if script == nil {
		log.Fatalln("/command/scriptList 里没有记录有脚本:<<" + command + ">> 请确认脚本名称输入是否正确")
	}
	err := script.handle()
	if err != nil {
		log.Fatalln(err)
	}
}
