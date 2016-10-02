package hmgo

import (
	"github.com/iteny/hmgo/context"
)

// Config is the main struct for BConfig
type Config struct {
	AppName             string //Application name
	RunMode             string //Running Mode: dev | prod
	RouterCaseSensitive bool
	ServerName          string
	RecoverPanic        bool
	RecoverFunc         func(*context.Context)
	CopyRequestBody     bool
	EnableGzip          bool
	MaxMemory           int64
	EnableErrorsShow    bool
	Listen              Listen
	WebConfig           WebConfig
	Log                 LogConfig
}

// Listen holds for http and https related config
type Listen struct {
	Graceful      bool // Graceful means use graceful module to start the server
	ServerTimeOut int64
	ListenTCP4    bool
	EnableHTTP    bool
	HttpAddr      string
	HttpPort      int
	EnableHTTPS   bool
	HTTPSAddr     string
	HTTPSPort     int
	HTTPSCertFile string
	HTTPSKeyFile  string
	EnableAdmin   bool
	AdminAddr     string
	AdminPort     int
	EnableFcgi    bool
	EnableStdIo   bool // EnableStdIo works with EnableFcgi Use FCGI via standard I/O
}

// WebConfig holds web related config
type WebConfig struct {
	AutoRender             bool
	EnableDocs             bool
	FlashName              string
	FlashSeparator         string
	DirectoryIndex         bool
	StaticDir              map[string]string
	StaticExtensionsToGzip []string
	TemplateLeft           string
	TemplateRight          string
	ViewsPath              string
	EnableXSRF             bool
	XSRFKey                string
	XSRFExpire             int
	Session                SessionConfig
}

// SessionConfig holds session related config
type SessionConfig struct {
	SessionOn               bool
	SessionProvider         string
	SessionName             string
	SessionGCMaxLifetime    int64
	SessionProviderConfig   string
	SessionCookieLifeTime   int
	SessionAutoSetCookie    bool
	SessionDomain           string
	EnableSidInHttpHeader   bool //	enable store/get the sessionId into/from http headers
	SessionNameInHttpHeader string
	EnableSidInUrlQuery     bool //	enable get the sessionId from Url Query params
}

// LogConfig holds Log related config
type LogConfig struct {
	AccessLogs  bool
	FileLineNum bool
	Outputs     map[string]string // Store Adaptor : config
}

var (
	// BConfig is the default config for Application
	HmConfig *Config
	// AppConfig is the instance of Config, store the config information from file
	// AppConfig *beegoAppConfig
	// AppPath is the absolute path to the app
	AppPath string
	// GlobalSessions is the instance for the session manager
	// GlobalSessions *session.Manager

	// appConfigPath is the path to the config files
	appConfigPath string
	// appConfigProvider is the provider for the config, default is ini
	appConfigProvider = "ini"
)
