package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"liyongcool.nat300.top/iaas/vsphere-go-sdk/appliance"
	"liyongcool.nat300.top/iaas/vsphere-go-sdk/cis"
	"liyongcool.nat300.top/iaas/vsphere-go-sdk/common"
	"liyongcool.nat300.top/iaas/vsphere-go-sdk/content"
	"liyongcool.nat300.top/iaas/vsphere-go-sdk/vcenter"
	"os"
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
	log.Info(resBody)

}

func content_test() {
	client := common.NewClient("https://128.179.0.241/rest/")
	log.Info(client)
	sess, err := cis.NewCIS(client).GetSessionHandle().CreateSession(cis.CodeBase64("root@vsphere.local", "Root@iaas12321"))
	c := content.NewContent(client, sess)
	l := c.NewLibrary()
	log.Info(*c)
	strs, err := l.ListLibraries()
	log.Info("GetLibraryList: ", strs, err)

	i := l.NewItem()
	strss, err := i.GetItemByLibraryID(strs[0])
	log.Info("GetItemByLibraryId: ", strss, err)
	info, err := i.GetItemInfoByItemID(strss[0])
	log.Info("GetItemInfoByItemID: ", info, err)
	err = cis.NewCIS(client).GetSessionHandle().DeleteSession(sess)
}

func cis_test() {
	code := cis.CodeBase64("root@vsphere.local", "Root@2021")
	log.Info("base64:", code)

	client := common.NewClient("https://128.179.0.241/rest/")
	//client := common.NewClient("http://127.0.0.1/")
	log.Info(client)

	c := cis.NewCIS(client)
	log.Info(*c)

	shandle := c.GetSessionHandle()
	log.Info(shandle)

	sess, err := shandle.CreateSession(code)
	log.Info("CreateSession: ", sess, err)

	info, err := shandle.Update(sess)
	log.Info("UpdateSession: ", info, err)

	err = shandle.DeleteSession(sess)
	log.Info("DeleteSession: ", err)
}

func vcenter_test() {
	client := common.NewClient("https://128.179.0.241/rest/")
	sess, err := cis.NewCIS(client).GetSessionHandle().CreateSession(cis.CodeBase64("root@vsphere.local", "Root@iaas12321"))
	vc := vcenter.NewVCenter(client, sess)

	nw := vc.NewNetwork()
	nws, err := nw.List()
	log.Info("Networks: ", nws, err)

	ds := vc.NewDataStore()
	dss, err := ds.List()
	log.Info("DataStores: ", dss, err)
	log.Info("================================================")

	host := vc.NewHost()
	hosts, err := host.List()
	log.Info("Hosts: ", hosts, err)
	log.Info("================================================")

	vm := vc.NewVM()
	vms, err := vm.List()
	log.Info("VMs: ", vms, err)
	log.Info("================================================")

	vmi, err := vm.Get(vms[0].Vm)
	log.Info("VMI: ", vmi, err)
	log.Info("================================================")

	f := vc.NewFolder()
	folders, err := f.List()
	log.Info("Folders: ", folders, err)
	log.Info("================================================")

	c := vc.NewCluster()
	cs, err := c.List()
	log.Info("Clusters: ", cs, err)
	log.Info("================================================")

	r := vc.NewResourcePool()
	rs, err := r.List()
	log.Info("ResourcePools: ", rs, err)
	log.Info("================================================")

	//vt := vc.NewVMTemplate().NewLibraryItems().NewItem("10574872-f28b-4f1e-b1a2-aae3a79905d4")
	vt := vc.NewVMTemplate().NewLibraryItems().NewItem("9893ad34-32d2-4a22-87c4-2e31806abadd")
	req := &vcenter.VMTemplateDeployReqeust{}
	req.Spec.Name = "LucyFly"
	req.Spec.Description = "I am Lucy"
	req.Spec.PoweredOn = true
	req.Spec.Placement.ClusterID = "domain-c7"
	req.Spec.Placement.FolderID = "group-v3"
	req.Spec.Placement.Host = "host-12"
	req.Spec.VMHomeStorage = &vcenter.VMTemplateDeployHomeStorage{
		DataStore: "datastore-60",
	}
	/*
		req.Spec.HardwareCustom = &vcenter.VMTemplateHDCustom{
			NICs: []vcenter.VMTemplateHDCustomNIC{
				vcenter.VMTemplateHDCustomNIC{Key: "4003"},
			},
		}
		req.Spec.HardwareCustom.NICs[0].Value.Network = "network-19"
	*/
	//req.Spec.VMHomeStorage = nil
	req.Spec.HardwareCustom = nil
	vmid, err := vt.Deploy(req)
	log.Info("vm: ", vmid, err)
	log.Info("================================================")

	eth := vc.NewVM().NewHardware(vmid).NewEthernet()
	nics, err := eth.List()
	log.Info("Ethernets: ", nics, err)
	for _, nic := range nics {
		ni, err := eth.Get(nic.Nic)
		log.Info("NIC Info: ", ni, " error: ", err)
	}
	log.Info("================================================")

	p := vm.NewPower(vmid)
	pi, err := p.Get()
	log.Info("PowerInfo: ", pi, err)
	log.Info("================================================")
	if pi.State == "POWERED_ON" {
		err = p.Stop()
		log.Info("Power OFF ", err)
		log.Info("================================================")
	}

	err = vm.Delete(vmid)
	log.Info("delete vm: ", vmid, err)
	log.Info("================================================")

	err = cis.NewCIS(client).GetSessionHandle().DeleteSession(sess)
}

func appliance_test() {
	client := common.NewClient("https://128.179.0.241/rest/")
	log.Info(client)
	sess, err := cis.NewCIS(client).GetSessionHandle().CreateSession(cis.CodeBase64("root@vsphere.local", "Root@2021"))
	if err != nil {
		log.Info("GetSession err: ", err)
	}
	c := appliance.NewAppliance(client, sess)
	l := c.NewNetworking()
	log.Info(*c)
	strs, SDKErr := l.Get()
	log.Info("GetNetworking: ", strs, SDKErr)
	err = cis.NewCIS(client).GetSessionHandle().DeleteSession(sess)
}

func main() {
	TestModule := ""
	help := false
	flag.StringVar(&TestModule, "t", "", "Test Module. eg: cis , vcenter , content, appliance")
	flag.BoolVar(&help, "h", false, "Show Usage.")
	flag.Parse()

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)    //设置日志的输出为标准输出
	log.SetLevel(log.InfoLevel) //设置日志的显示级别，这一级别以及更高级别的日志信息将会输出
	log.SetReportCaller(true)   //设置日志的调用文件，调用函数
	log.SetFormatter(&log.JSONFormatter{})

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
	case "appliance":
		//test appliance moudle
		appliance_test()
	default:
		flag.Usage()
	}
}
