package vcenter

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"liyongcool.nat300.top/iaas/vsphere-go-sdk/common"
)

/*
* VM Operations
 */
type VM struct {
	con *common.Connector
	uri string
}

type ListVMResult struct {
	Vm         string `json:"vm"`
	Name       string `json:"name"`
	PowerState string `json:"power_state"`
	CpuCount   int    `json:"cpu_count"`
	MemSizeMiB int    `json:"memory_size_MiB"`
}

type ValueOfListVMsResult struct {
	Value []ListVMResult `json:"value"`
}

func (o *VM) List() ([]ListVMResult, error) {
	header := make(map[string]string)
	header["vmware-api-session-id"] = o.con.Sid
	resp, err := o.con.Invoker.SendRequest(o.uri, header, nil, "GET")
	if err != nil {
		log.Error("ListVMs Error", err)
		return nil, err
	}

	vms := ValueOfListVMsResult{}
	err = json.Unmarshal(resp.Data, &vms)
	if err != nil {
		log.Error("ListVMs Response Data", string(resp.Data))
		return nil, err
	}
	return vms.Value, err
}

type CDROMInfo struct {
	Key   string `json:"key,omitempty"`
	Value struct {
		StartConnected bool `json:"start_connected,omitempty"`
		Backing        struct {
			DeviceAccessType string `json:"device_access_type,omitempty"`
			HostDevice       string `json:"host_device,omitempty"`
			AutoDetect       bool   `json:"auto_detect,omitempty"`
			ISOFile          string `json:"iso_file,omitempty"`
			Type             string `json:"type,omitempty"`
		} `json:"backing,omitempty"`
		AllowGuestControl bool   `json:"allow_guest_control,omitempty"`
		Label             string `json:"label,omitempty"`
		State             string `json:"state,omitempty"`
		Type              string `json:"type,omitempty"`
		Sata              struct {
			Bus  int `json:"bus,omitempty"`
			Unit int `json:"unit,omitempty"`
		} `json:"sata,omitempty"`
		Ide struct {
			Master  bool `json:"master,omitempty"`
			Primary bool `json:"primary,omitempty"`
		} `json:"ide,omitempty"`
	} `json:"value,omitempty"`
}

type MemoryInfo struct {
	SizeMiB                int  `json:"size_MiB,omitempty"`
	HotAddEnabled          bool `json:"hot_add_enabled,omitempty"`
	HotAddIncrementSizeMiB int  `json:"hot_add_increment_size_MiB,omitempty"`
	HotAddLimitMiB         int  `json:"hot_add_limit_MiB,omitempty"`
}

type ParallelPortInfo struct {
	Key   string `json:"key,omitempty"`
	Value struct {
		Label   string `json:"label,omitempty"`
		Backing struct {
			AutoDetect bool   `json:"auto_detect,omitempty"`
			File       string `json:"file,omitempty"`
			Type       string `json:"type,omitempty"`
			HostDevice string `json:"host_device,omitempty"`
		} `json:"backing,omitempty"`
		AllowGuestControl bool   `json:"allow_guest_control,omitempty"`
		State             string `json:"state,omitempty"`
		StartConnected    bool   `json:"start_connected,omitempty"`
	} `json:"value,omitempty"`
}

type DiskInfo struct {
	Key   string `json:"key,omitempty"`
	Value struct {
		SATA struct {
			Bus  int `json:"bus,omitempty"`
			Unit int `json:"unit,omitempty"`
		} `json:"sata,omitempty"`
		SCSI struct {
			Bus  int `json:"bus,omitempty"`
			Unit int `json:"unit,omitempty"`
		} `json:"scsi,omitempty"`
		Ide struct {
			Master  bool `json:"master,omitempty"`
			Primary bool `json:"primary,omitempty"`
		} `json:"ide,omitempty"`
		Backing struct {
			VMDKFile string `json:"vmdk_file,omitempty"`
			Type     string `json:"type,omitempty"`
		} `json:"backing,omitempty"`
		Label    string `json:"label,omitempty"`
		Type     string `json:"type,omitempty"`
		Capacity int    `json:"capacity,omitempty"`
	} `json:"value,omitempty"`
}

type SataAdapterInfo struct {
	Key   string `json:"key,omitempty"`
	Value struct {
		Bus           int    `json:"bus,omitempty"`
		PCISlotNumber int    `json:"pci_slot_number,omitempty"`
		Label         string `json:"label,omitempty"`
		Type          string `json:"type,omitempty"`
	} `json:"value,omitempty"`
}

type CPUInfo struct {
	HotRemoveEnabled bool `json:"hot_remove_enabled,omitempty"`
	Count            int  `json:"count,omitempty"`
	HotAddEnabled    bool `json:"hot_add_enabled,omitempty"`
	CoresPerSocket   int  `json:"cores_per_socket,omitempty"`
}

type SerialPortInfo struct {
	Key   string `json:"key,omitempty"`
	Value struct {
		Label   string `json:"label,omitempty"`
		Backing struct {
			File            string `json:"file,omitempty"`
			Pipe            string `json:"pipe,omitempty"`
			Type            string `json:"type,omitempty"`
			NetworkLocation string `json:"network_location,omitempty"`
			Proxy           string `json:"proxy,omitempty"`
			AutoDetect      bool   `json:"auto_detect,omitempty"`
			NORXLoss        bool   `json:"no_rx_loss,omitempty"`
			HostDevice      string `json:"host_device,omitempty"`
		} `json:"backing,omitempty"`
		StartConnected    bool   `json:"start_connected,omitempty"`
		YieldOnPoll       bool   `json:"yiel_on_poll,omitempty"`
		AllowGuestControl bool   `json:"allow_guest_control,omitempty"`
		State             string `json:"state,omitempty"`
	} `json:"value,omitempty"`
}

type SCSIAdapterInfo struct {
	Key   string `json:"key,omitempty"`
	Value struct {
		SCSI struct {
			Bus  int `json:"bus,omitempty"`
			Unit int `json:"unit,omitempty"`
		} `json:"scsi,omitempty"`
		PCISlotNumber int    `json:"pci_slot_number,omitempty"`
		Label         string `json:"label,omitempty"`
		Type          string `json:"type,omitempty"`
		Sharing       string `json:"sharing,omitempty"`
	} `json:"value,omitempty"`
}

type FloppyInfo struct {
	Key   string `json:"key,omitempty"`
	Value struct {
		Label   string `json:"label,omitempty"`
		Backing struct {
			AutoDetect bool   `json:"auto_detect,omitempty"`
			Type       string `json:"type,omitempty"`
			ImageFile  string `json:"image_file,omitempty"`
			HostDevice string `json:"host_device,omitempty"`
		} `json:"backing,omitempty"`
		AllowGuestControl bool   `json:"allow_guest_control,omitempty"`
		State             string `json:"state,omitempty"`
		StartConnected    bool   `json:"start_connected,omitempty"`
	} `json:"value,omitempty"`
}

type NicInfo struct {
	Key   string `json:"key,omitempty"`
	Value struct {
		StartConnected bool `json:"start_connected,omitempty"`
		PCISlotNumber  int  `json:"pci_slot_number,omitempty"`
		Backing        struct {
			DistributedSwitchUUID string `json:"distributed_switch_uuid,omitempty"`
			DistributedPort       string `json:"distributed_port,omitempty"`
			OpaqueNetworkID       string `json:"opaque_network_id,omitempty"`
			OpaqueNetworkType     string `json:"opaque_network_type,omitempty"`
			HostDevice            string `json:"host_device,omitempty"`
			ConnectionCookie      int    `json:"connection_cookie,omitempty"`
			NetworkName           string `json:"network_name,omitempty"`
			Type                  string `json:"type,omitempty"`
			Network               string `json:"network,omitempty"`
		} `json:"backing,omitempty"`
		MacAddress              string `json:"mac_address,omitempty"`
		MacType                 string `json:"mac_type,omitempty"`
		AllowGuestControl       bool   `json:"allow_guest_control,omitempty"`
		WakeOnLanEnabled        bool   `json:"wake_on_lan_enabled,omitempty"`
		Label                   string `json:"label,omitempty"`
		State                   string `json:"state,omitempty"`
		Type                    string `json:"type,omitempty"`
		UPTCompatibilityEnabled bool   `json:"upt_compatibility_enabled,omitempty"`
	} `json:"value,omitempty"`
}

type BootInfo struct {
	EFILegacyBoot   bool   `json:"efi_legacy_boot,omitempty"`
	Delay           int    `json:"delay,omitempty"`
	RetryDelay      int    `json:"retry_delay,omitempty"`
	EnterSetupMode  bool   `json:"enter_setup_mode,omitempty"`
	NetworkProtocol string `json:"network_protocol,omitempty"`
	Type            string `json:"type,omitempty"`
	Retry           bool   `json:"retry,omitempty"`
}

type BootDeviceInfo struct {
	Disks []string `json:"disks,omitempty"`
	Nic   string   `json:"nic,omitempty"`
	Type  string   `json:"type,omitempty"`
}

type HardwareInfo struct {
	UpgradeVersion string `json:"upgrade_version,omitempty"`
	UpgradeError   string `json:"upgrade_error,omitempty"`
	UpgradePolicy  string `json:"upgrade_policy,omitempty"`
	UpgradeStatus  string `json:"upgrade_status,omitempty"`
	Version        string `json:"version,omitempty"`
}

type VMInfo struct {
	Cdroms        []CDROMInfo        `json:"cdroms,omitempty"`
	Floppies      []FloppyInfo       `json:"floppies,omitempty"`
	Memory        MemoryInfo         `json:"memory"`
	Disks         []DiskInfo         `json:"disks,omitempty"`
	SataAdapters  []SataAdapterInfo  `json:"sata_adapters,omitempty"`
	Cpu           CPUInfo            `json:"cpu"`
	ScsiAdapters  []SCSIAdapterInfo  `json:"scsi_adapters,omitempty"`
	PowerState    string             `json:"power_state"`
	Name          string             `json:"name,omitempty"`
	SerialPorts   []SerialPortInfo   `json:"serial_ports,omitempty"`
	Nics          []NicInfo          `json:"nics,omitempty"`
	Boot          BootInfo           `json:"boot"`
	BootDevices   []BootDeviceInfo   `json:"boot_devices,omitempty"`
	ParallelPorts []ParallelPortInfo `json:"parallel_ports,omitempty"`
	GuestOS       string             `json:"guest_OS"`
	Hard          HardwareInfo       `json:"hardware"`
}

type ValueOfVMInfo struct {
	Value VMInfo `json:"value"`
}

func (o *VM) Get(vm string) (*VMInfo, error) {
	header := make(map[string]string)
	header["vmware-api-session-id"] = o.con.Sid
	uri := o.uri + "/" + vm
	resp, err := o.con.Invoker.SendRequest(uri, header, nil, "GET")
	if err != nil {
		log.Error("GetVMInfo Error: ", err)
		return nil, err
	}

	vmi := ValueOfVMInfo{}
	err = json.Unmarshal(resp.Data, &vmi)
	if err != nil {
		log.Error("GetVMInfo Response Data", string(resp.Data))
		return nil, err
	}
	return &(vmi.Value), err
}

func (o *VM) Delete(vm string) error {
	header := make(map[string]string)
	header["vmware-api-session-id"] = o.con.Sid
	uri := o.uri + "/" + vm
	_, err := o.con.Invoker.SendRequest(uri, header, nil, "DELETE")
	if err != nil {
		log.Error("Delete SendRequest Error, ", err)
	}
	return err
}

func (o *VM) NewHardware(vm string) *Hardware {
	return &Hardware{
		con: o.con,
		uri: o.uri + "/" + vm + "/hardware",
	}
}

func (o *VM) NewPower(vm string) *Power {
	return &Power{
		con: o.con,
		uri: o.uri + "/" + vm + "/power",
	}
}

/*
* Hardware Operations
 */
type Hardware struct {
	con *common.Connector
	uri string
}

func (o *Hardware) NewDisk() *Disk {
	return &Disk{
		con: o.con,
		uri: o.uri + "/disk",
	}
}

/*
* Disk Operations
 */
type Disk struct {
	con *common.Connector
	uri string
}

/*
* Power Operations
 */
type Power struct {
	con *common.Connector
	uri string
}

type PowerInfo struct {
	CleanPowerOff bool   `json:"clean_power_off,omitempty"`
	State         string `json:"state,omitempty"`
}

type ValueOfPowerInfo struct {
	Value PowerInfo `json:"value"`
}

func (o *Power) Get() (*PowerInfo, error) {
	header := make(map[string]string)
	header["vmware-api-session-id"] = o.con.Sid
	resp, err := o.con.Invoker.SendRequest(o.uri, header, nil, "GET")
	if err != nil {
		log.Error("Power Get SendRequest Error, ", err)
		return nil, err
	}

	vpi := ValueOfPowerInfo{}
	err = json.Unmarshal(resp.Data, &vpi)
	if err != nil {
		log.Error("Power Get Unmarshal Error, ", err)
		return nil, err
	}
	return &(vpi.Value), err
}

type PowerOpType int

const (
	POWER_OP_RESET   PowerOpType = 1
	POWER_OP_START   PowerOpType = 2
	POWER_OP_STOP    PowerOpType = 3
	POWER_OP_SUSPEND PowerOpType = 4
)

func (o *Power) post(op PowerOpType) error {
	header := make(map[string]string)
	header["vmware-api-session-id"] = o.con.Sid

	op_uri := "/"
	switch op {
	case POWER_OP_RESET:
		op_uri += "reset"
	case POWER_OP_STOP:
		op_uri += "stop"
	case POWER_OP_START:
		op_uri += "start"
	case POWER_OP_SUSPEND:
		op_uri += "suspend"
	}

	_, err := o.con.Invoker.SendRequest(o.uri+op_uri, header, nil, "POST")
	if err != nil {
		log.Error("Power post SendRequest error, ", err)
	}
	return err
}

func (o *Power) Reset() error {
	return o.post(POWER_OP_RESET)
}

func (o *Power) Stop() error {
	return o.post(POWER_OP_STOP)
}

func (o *Power) Start() error {
	return o.post(POWER_OP_START)
}

func (o *Power) Suspend() error {
	return o.post(POWER_OP_SUSPEND)
}
