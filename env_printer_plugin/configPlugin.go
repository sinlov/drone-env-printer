package env_printer_plugin

const (
	EnvPluginResultShareHost = "PLUGIN_RESULT_SHARE_HOST"

	NamePluginDebug   = "config.debug"
	EnvPluginTimeOut  = "PLUGIN_TIMEOUT_SECOND"
	NamePluginTimeOut = "config.timeout_second"

	EnvPrinterPaddingLeftMax  = "PLUGIN_ENV_PRINTER_PADDING_LEFT_MAX"
	NamePrinterPaddingLeftMax = "config.env_printer_padding_left_max"

	EnvPrinterPrintKeys  = "PLUGIN_ENV_PRINTER_PRINT_KEYS"
	NamePrinterPrintKeys = "config.env_printer_print_keys"
)

var (
	// supportMsgType
	//supportMsgType = []string{
	//	msgTypeText,
	//	msgTypePost,
	//	msgTypeInteractive,
	//}

	cleanResultEnvList = []string{
		EnvPluginResultShareHost,
	}
)

type (
	// Config env_printer_plugin private config
	Config struct {
		EnvPrintKeys   []string
		PaddingLeftMax int

		Debug bool

		TimeoutSecond uint
	}
)
