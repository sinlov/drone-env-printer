package env_printer_plugin

const (
	EnvPluginResultShareHost = "PLUGIN_RESULT_SHARE_HOST"

	msgTypeText        = "text"
	msgTypePost        = "post"
	msgTypeInteractive = "interactive"
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
