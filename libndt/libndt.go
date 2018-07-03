// Part of Measurement Kit <https://measurement-kit.github.io/>.
// Measurement Kit is free software under the BSD license. See AUTHORS
// and LICENSE for more information on the copying conditions.

package libndt

import (
	"github.com/measurement-kit/libndt-go/swig"
	"log"
)

// Version contains a version number
type Version uint

// VersionMajor is libndt's major version number
var VersionMajor = Version(swig.GetVersionMajor())

// VersionMinor is libndt's minor version number
var VersionMinor = Version(swig.GetVersionMinor())

// VersionPatch is libndt's patch version number
var VersionPatch = Version(swig.GetVersionPatch())

// NettestFlags contains flags for selecting subtests
type NettestFlags uint8

// NettestFlagUpload runs the upload subtest
var NettestFlagUpload = NettestFlags(swig.GetNettestFlagUpload())

// NettestFlagDownload runs the download subtest
var NettestFlagDownload = NettestFlags(swig.GetNettestFlagDownload())

// NettestFlagDownloadExt runs the multi-stream download subtest
var NettestFlagDownloadExt = NettestFlags(swig.GetNettestFlagDownloadExt())

// Verbosity indicates libndt verbosity
type Verbosity = uint

// VerbosityQuiet indicates that the library is quiet
var VerbosityQuiet = Verbosity(swig.GetVerbosityQuiet())

// VerbosityWarning indicates that the library only emits warnings
var VerbosityWarning = Verbosity(swig.GetVerbosityWarning())

// VerbosityInfo indicates that the library does not emit debug messages
var VerbosityInfo = Verbosity(swig.GetVerbosityInfo())

// VerbosityDebug indicates that the library emits all log messages
var VerbosityDebug = Verbosity(swig.GetVerbosityDebug())

// ProtocolFlags contains flags enabling protocol features
type ProtocolFlags = uint

// ProtocolFlagJSON enable wrapping NDT messages in JSON objects
var ProtocolFlagJSON = ProtocolFlags(swig.GetProtocolFlagJson())

// ProtocolFlagTLS enables using TLS
var ProtocolFlagTLS = ProtocolFlags(swig.GetProtocolFlagTls())

// ProtocolFlagWebSockets uses WebSockets rather than the legacy NDT protocol
var ProtocolFlagWebSockets = ProtocolFlags(swig.GetProtocolFlagWebsockets())

// Timeout is a timeout in seconds
type Timeout = uint

// Settings contains NDT client settings. When a field is not set, the default
// value configured insider of libndt will be used.
type Settings struct {
	// MlabnsURL contains the mlab-ns URL to use.
	MlabnsURL     string
	// Timeout indicates the timeout for I/O, in seconds.
	Timeout       Timeout
	// Hostname is the host to connect to.
	Hostname      string
	// Port is the port to use.
	Port          string
	// Metadata contains additional metadata to be sent to the server.
	Metadata      map[string]string
	// NettestFlags contains flags selecting specific sub tests.
	NettestFlags  NettestFlags
	// Verbosity controls the test verbosity.
	Verbosity     Verbosity
	// ProtocolFlags selects specific protocl features.
	ProtocolFlags ProtocolFlags
	// MaxRuntime is the maximum amount of seconds for which a subtest can run.
	MaxRuntime    Timeout
	// SOCKS5hPort is the port to eventually be used for SOCKSv5.
	SOCKS5hPort   string
}

// NewSettings creates an instance of settings
func NewSettings() Settings {
	settings := Settings{}
	settings.Metadata = make(map[string]string)
	return settings
}

// Client is a NDT client.
type Client struct {
	settings Settings
}

// NewClient creates a new client with default settings.
func NewClient() Client {
	return Client{}
}

// NewClientWithSettings creates a new client with specific settings.
func NewClientWithSettings(s Settings) Client {
	return Client{settings: s}
}

// golangClient implements the swig.NdtClient interface.
type golangClient struct {
	// TODO(bassosimone): add here room for the callbacks to override
}

// Swigcptr() is part of the swig.NdtClient interface.
func (*golangClient) Swigcptr() uintptr {
	return 0
}

// SwigIsNdtClient() is part of the swig.NdtClient interface.
func (*golangClient) SwigIsNdtClient() {
}

// DirectorInterface() is part of the swig.NdtClient interface.
func (*golangClient) DirectorInterface() interface{} {
	return nil
}

// Run() is part of the swig.NdtClient interface.
func (*golangClient) Run() bool {
	return false
}

// OnWarning() is part of the swig.NdtClient interface.
func (*golangClient) OnWarning(msg string) {
	log.Printf("WARNING: %s\n", msg)
}

// OnInfo() is part of the swig.NdtClient interface.
func (*golangClient) OnInfo(msg string) {
	log.Printf("INFO: %s\n", msg)
}

// OnDebug() is part of the swig.NdtClient interface.
func (*golangClient) OnDebug(msg string) {
	log.Printf("DEBUG: %s\n", msg)
}

// OnPerformance() is part of the swig.NdtClient interface.
func (*golangClient) OnPerformance(nettestID byte, numFlows byte, measuredBytes float64, measurementInterval float64, elapsed float64, maxRuntime float64) {
	log.Printf("PERFORMANCE\n")
}

// OnResult() is part of the swig.NdtClient interface.
func (*golangClient) OnResult(scope string, name string, value string) {
	log.Printf("RESULT: %s.%s=%s\n", scope, name, value)
}

// OnServerBusy() is part of the swig.NdtClient interface.
func (*golangClient) OnServerBusy(reason string) {
	log.Printf("BUSY: %s\n", reason)
}

// SwigIsGolangClient() is used to distinguish this type in the type system. See
// also https://github.com/swig/swig/issues/418.
func (*golangClient) SwigIsGolangClient() {
}

// newDirectorNdtClient() provide type safety checks to swig.NewDirectoryNdtClient()
func newDirectorNdtClient(clnt swig.NdtClient, settings swig.NdtSettings) swig.NdtClient {
	return swig.NewDirectorNdtClient(clnt, settings)
}

// Run() runs the test.
func (clnt Client) Run() bool {
	cxxSettings := swig.NewNdtSettings()
	if clnt.settings.MlabnsURL != "" {
		cxxSettings.SetMlabnsUrl(clnt.settings.MlabnsURL)
	}
	if clnt.settings.Timeout > 0 {
		cxxSettings.SetTimeout(clnt.settings.Timeout)
	}
	if clnt.settings.Hostname != "" {
		cxxSettings.SetHostname(clnt.settings.Hostname)
	}
	if clnt.settings.Port != "" {
		cxxSettings.SetPort(clnt.settings.Port)
	}
	for k, v := range clnt.settings.Metadata {
		cxxSettings.AddMetadata(k, v)
	}
	if clnt.settings.Verbosity != 0 {
		cxxSettings.SetVerbosity(uint(clnt.settings.Verbosity))
	}
	if clnt.settings.ProtocolFlags != 0 {
		cxxSettings.SetProtocolFlags(clnt.settings.ProtocolFlags)
	}
	if clnt.settings.MaxRuntime > 0 {
		cxxSettings.SetMaxRuntime(clnt.settings.MaxRuntime)
	}
	if clnt.settings.SOCKS5hPort != "" {
		cxxSettings.SetSocks5hPort(clnt.settings.SOCKS5hPort)
	}

	// Here we pass the client with overriden methods as underlying
	// implementation to the director class, such that the overriden
	// methods of the Go class are called by the C++ code.
	overriden := &golangClient{}
	director := newDirectorNdtClient(overriden, cxxSettings)
	defer swig.DeleteDirectorNdtClient(director)
	return director.Run()
}
