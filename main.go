package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"vsphere-go-sdk/cis"
	"vsphere-go-sdk/common"
	"vsphere-go-sdk/content"
	"vsphere-go-sdk/vcenter"
)

func common_test() {

	url := "http://localhost:8088/v1/user/login"

	client := common.NewClient(url)
	//headers:=make(map[string]string)
	//headers["session"]="test"
	method := "GET"
	//testBody:="gdhsdsdsghd"
	//body,_:=json.Marshal(testBody)
	resBody, _ := client.SendRequest(url, nil, nil, method)
	log.Debug(resBody)

}

func content_test() {
	client := common.NewClient("https://128.179.0.241/rest/")
	log.Debug(client)
	sess, err := cis.NewCIS(client).GetSessionHandle().CreateSession(cis.CodeBase64("root@vsphere.local", "Root@2021"))
	c := content.NewContent(client, sess)
	l := c.NewLibrary()
	log.Debug(*c)
	strs, err := l.ListLibraries()
	log.Debug("GetLibraryList: ", strs, err)

	i := l.NewItem()
	strss, err := i.GetItemByLibraryID(strs[0])
	log.Debug("GetItemByLibraryId: ", strss, err)
	err = cis.NewCIS(client).GetSessionHandle().DeleteSession(sess)
}

func cis_test() {
	code := cis.CodeBase64("root@vsphere.local", "Root@2021")
	log.Debug("base64:", code)

	client := common.NewClient("https://128.179.0.241/rest/")
	//client := common.NewClient("http://127.0.0.1/")
	log.Debug(client)

	c := cis.NewCIS(client)
	log.Debug(*c)

	shandle := c.GetSessionHandle()
	log.Debug(shandle)

	sess, err := shandle.CreateSession(code)
	log.Debug("CreateSession: ", sess, err)

	err = shandle.DeleteSession(sess)
	log.Debug("DeleteSession: ", err)
}

func vcenter_test() {
	client := common.NewClient("https://128.179.0.241/rest/")
	sess, err := cis.NewCIS(client).GetSessionHandle().CreateSession(cis.CodeBase64("root@vsphere.local", "Root@2021"))

	vc := vcenter.NewVCenter(client, sess)
	vm := vc.NewVM()
	vms, err := vm.ListVMs()
	log.Debug("VMs: ", vms, err)
	log.Debug("================================================")

	vmi, err := vm.GetVMInfo(vms[0].Vm)
	log.Debug("VMI: ", vmi, err)
	log.Debug("================================================")

	f := vc.NewFolder()
	folders, err := f.List()
	log.Debug("Folders: ", folders, err)
	log.Debug("================================================")

	err = cis.NewCIS(client).GetSessionHandle().DeleteSession(sess)
}

func main() {
	TestModule := ""
	help := false
	flag.StringVar(&TestModule, "t", "", "Test Module. eg: cis , vcenter , content")
	flag.BoolVar(&help, "h", false, "Show Usage.")
	flag.Parse()

	log.SetOutput(os.Stdout)     //设置日志的输出为标准输出
	log.SetLevel(log.DebugLevel) //设置日志的显示级别，这一级别以及更高级别的日志信息将会输出
	log.SetReportCaller(true)    //设置日志的调用文件，调用函数

	if help {
		flag.Usage()
		return
	}

	switch TestModule {
	case "cis":
		// test cis module
		cis_test()
	case "vcenter":
		// test vcenter module
		vcenter_test()
	case "content":
		//test content moudle
		content_test()
	default:
		flag.Usage()
	}
}
