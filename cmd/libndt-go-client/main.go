// Part of Measurement Kit <https://measurement-kit.github.io/>.
// Measurement Kit is free software under the BSD license. See AUTHORS
// and LICENSE for more information on the copying conditions.

package main

import (
	"github.com/measurement-kit/libndt-go/libndt"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	// TODO(bassosimone): correctly deal with signals in libndt
	signal.Ignore(syscall.SIGPIPE)
	settings := libndt.NewSettings()
	settings.NettestFlags |= libndt.NettestFlagDownload | libndt.NettestFlagUpload
	settings.Verbosity = libndt.VerbosityDebug
	settings.ProtocolFlags |= libndt.ProtocolFlagJSON
	settings.Metadata["client.application"] = "measurement-kit/libndt-go"
	settings.MlabnsURL = "https://mlab-ns.appspot.com/ndt?policy=random"
	settings.OnWarningCallbacks = append(settings.OnWarningCallbacks, func(s string) {
		log.Printf("<warn> %s\n", s)
	})
	settings.OnInfoCallbacks = append(settings.OnInfoCallbacks, func(s string) {
		log.Printf("<info> %s\n", s)
	})
	settings.OnDebugCallbacks = append(settings.OnDebugCallbacks, func(s string) {
		log.Printf("<debug> %s\n", s)
	})
	settings.OnPerformanceCallbacks = append(settings.OnPerformanceCallbacks, func(dir libndt.NettestFlags, nconn byte, numBytes float64, interval float64, elapsed float64, maxRuntime float64) {
		log.Printf("<performance> %d - %d flows - %f kbit/s - %f completion\n", uint8(dir), nconn, numBytes*8.0/interval/1000.0, elapsed/maxRuntime)
	})
	settings.OnResultCallbacks = append(settings.OnResultCallbacks, func(scope string, name string, value string) {
		log.Printf("<result> %s.%s = %s\n", scope, name, value)
	})
	settings.OnServerBusyCallbacks = append(settings.OnServerBusyCallbacks, func(m string) {
		log.Printf("<server-busy> %s\n", m)
	})
	client := libndt.NewClientWithSettings(settings)
	client.Run()
}
