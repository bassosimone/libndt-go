// Part of Measurement Kit <https://measurement-kit.github.io/>.
// Measurement Kit is free software under the BSD license. See AUTHORS
// and LICENSE for more information on the copying conditions.

package libndt

import (
	"github.com/measurement-kit/libndt-go/swig"
	"log"
)

// TODO(bassosimone): add documentation

var VersionMajor uint32 = uint32(swig.GetVersionMajor())

var VersionMinor uint32 = uint32(swig.GetVersionMinor())

var VersionPatch uint32 = uint32(swig.GetVersionPatch())

var NettestFlagUpload uint8 = uint8(swig.GetNettestFlagUpload())

var NettestFlagDownload uint8 = uint8(swig.GetNettestFlagDownload())

var NettestFlagDownloadExt uint8 = uint8(swig.GetNettestFlagDownloadExt())

var VerbosityQuiet uint32 = uint32(swig.GetVerbosityQuiet())

var VerbosityWarning uint32 = uint32(swig.GetVerbosityWarning())

var VerbosityInfo uint32 = uint32(swig.GetVerbosityInfo())

var VerbosityDebug uint32 = uint32(swig.GetVerbosityDebug())

var ProtocolFlagJSON uint32 = uint32(swig.GetProtocolFlagJson())

var ProtocolFlagTLS uint32 = uint32(swig.GetProtocolFlagTls())

var ProtocolFlagWebSockets uint32 = uint32(swig.GetProtocolFlagWebsockets())

type Settings struct {
	MlabnsURL     string
	Timeout       uint16
	Hostname      string
	Port          string
	Metadata      map[string]string
	NettestFlags  uint8
	Verbosity     uint32
	ProtocolFlags uint32
	MaxRuntime    uint16
	SOCKS5hPort   string
}

func NewSettings() Settings {
	settings := Settings{}
	settings.Metadata = make(map[string]string)
	return settings
}

type Client struct {
	settings Settings
}

func NewClient() Client {
	return Client{}
}

func NewClientWithSettings(s Settings) Client {
	return Client{settings: s}
}

type golangClient struct {
	// TODO(bassosimone): add here room for the callbacks to override
}

func (*golangClient) Swigcptr() uintptr {
	return 0
}

func (*golangClient) SwigIsNdtClient() {
}

func (*golangClient) DirectorInterface() interface{} {
	return nil
}

func (*golangClient) Run() bool {
	return false
}

func (*golangClient) OnWarning(msg string) {
	log.Printf("WARNING: %s\n", msg)
}

func (*golangClient) OnInfo(msg string) {
	log.Printf("INFO: %s\n", msg)
}

func (*golangClient) OnDebug(msg string) {
	log.Printf("DEBUG: %s\n", msg)
}

func (*golangClient) OnPerformance(nettestID byte, numFlows byte, measuredBytes float64, measurementInterval float64, elapsed float64, maxRuntime float64) {
	log.Printf("PERFORMANCE\n")
}

func (*golangClient) OnResult(scope string, name string, value string) {
	log.Printf("RESULT: %s.%s=%s\n", scope, name, value)
}

func (*golangClient) OnServerBusy(reason string) {
	log.Printf("BUSY: %s\n", reason)
}

func (*golangClient) IsGolangClient() {
	// Needed by the type system to single out the golangClient.
}

func newDirectorNdtClient(clnt swig.NdtClient, settings swig.NdtSettings) swig.NdtClient {
	// Wrapper because swig.NewDirectorNdtClient does not perform type checks.
	return swig.NewDirectorNdtClient(clnt, settings)
}

func (clnt Client) Run() bool {
	cxxSettings := swig.NewNdtSettings()
	if clnt.settings.Verbosity != 0 {
		// TODO(bassosimone): here we currently use wide enough integers for
		// holding types but this makes passing stuff around in golang less
		// efficient; so, we should probably use `uint` directly.
		cxxSettings.SetVerbosity(uint(clnt.settings.Verbosity))
	}
	// TODO(bassosimone): wrap more settings variables here

	// TODO(bassosimone): explain what is happening here
	overriden := &golangClient{}
	director := newDirectorNdtClient(overriden, cxxSettings)
	defer swig.DeleteDirectorNdtClient(director)
	return director.Run()
}
