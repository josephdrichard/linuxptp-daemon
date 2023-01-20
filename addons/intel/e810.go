package intel

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/openshift/linuxptp-daemon/pkg/plugin"
	ptpv1 "github.com/openshift/ptp-operator/api/v1"
	"os/exec"
)

type E810Opts struct {
	CmdLine []string `json:"cmdline"`
}

func OnPTPConfigChangeE810(nodeProfile *ptpv1.PtpProfile) error {
	glog.Error("calling onPTPConfigChange for E810")
	var e810Opts E810Opts
	var err error
	var e810OptsByteArray []byte
	var stdout []byte

	for name, opts := range (*nodeProfile).Plugins {
		glog.Infof("name=" + name)
		if name == "e810" {
			e810OptsByteArray, _ = json.Marshal(opts)
			//glog.Infof("Marshalled: " + string(e810OptsByteArray))
			err = json.Unmarshal(e810OptsByteArray, &e810Opts)
			if err != nil {
				glog.Error("e810 failed to unmarshal opts: " + err.Error())
			}

			for _, cmd := range e810Opts.CmdLine {
				glog.Info("e810 pluging executing: " + cmd)
				stdout, err = exec.Command("/usr/bin/bash", "-c", cmd).Output()
				if err != nil {
					glog.Error("failed to run cmd:" + err.Error())
				} else {
					glog.Info(string(stdout))
				}
			}

		}
	}
	return nil
}

func E810(name string) *plugin.Plugin {
	if name != "e810" {
		glog.Errorf("Plugin must be initialized as 'e810'")
		return nil
	}
	glog.Errorf("registering e810")
	return &plugin.Plugin{Name: "e810",
		OnPTPConfigChange: OnPTPConfigChangeE810,
		PopulateHwConfig:  PopulateHwConfigE810,
	}
}

func PopulateHwConfigE810(hwconfigs *[]ptpv1.HwConfig) error {
	glog.Info("Calling PopulateHwConfigE810")
	hwConfig := ptpv1.HwConfig{
		DeviceID: "e810",
		VendorID: "intel",
		Failed:   false,
		Status:   "done",
	}
	*hwconfigs = append(*hwconfigs, hwConfig)
	return nil
}
