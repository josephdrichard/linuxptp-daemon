package mapping

import (
	"github.com/openshift/linuxptp-daemon/addons/example"
	"github.com/openshift/linuxptp-daemon/addons/intel"
	"github.com/openshift/linuxptp-daemon/pkg/plugin"
)

var PluginMapping = map[string]plugin.New{
	"reference": example.Reference,
	"e810":      intel.E810,
}
