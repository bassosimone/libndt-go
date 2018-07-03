// Part of Measurement Kit <https://measurement-kit.github.io/>.
// Measurement Kit is free software under the BSD license. See AUTHORS
// and LICENSE for more information on the copying conditions.

package main

import (
	"github.com/measurement-kit/libndt-go/libndt"
	"os/signal"
	"syscall"
)

func main() {
	// TODO(bassosimone): correctly deal with signals is libndt
	signal.Ignore(syscall.SIGPIPE)
	settings := libndt.NewSettings()
	settings.NettestFlags |= libndt.NettestFlagDownload | libndt.NettestFlagUpload
	settings.Verbosity = libndt.VerbosityDebug
	settings.ProtocolFlags |= libndt.ProtocolFlagJSON
	settings.Metadata["client.application"] = "measurement-kit/libndt-go"
	settings.MlabnsURL = "https://mlab-ns.appspot.com/ndt?policy=random"
	client := libndt.NewClientWithSettings(settings)
	client.Run()
}
