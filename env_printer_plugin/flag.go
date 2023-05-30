package env_printer_plugin

import (
	"github.com/sinlov/drone-info-tools/drone_info"
	"github.com/urfave/cli/v2"
	"log"
)

func BindCliFlag(c *cli.Context, cliVersion, cliName string, drone drone_info.Drone) Plugin {
	config := Config{
		EnvPrintKeys:   c.StringSlice("config.env_printer_print_keys"),
		PaddingLeftMax: c.Int("config.env_printer_padding_left_max"),

		Debug: c.Bool("config.debug"),

		TimeoutSecond: c.Uint("config.timeout_second"),
	}

	if config.Debug {
		log.Printf("config.timeout_second: %v", config.TimeoutSecond)
	}

	p := Plugin{
		Name:    cliName,
		Version: cliVersion,
		Drone:   drone,
		Config:  config,
	}
	return p
}

// Flag
// set flag at here
func Flag() []cli.Flag {
	return []cli.Flag{
		// env_printer_plugin start
		// new flag string template if no use, please replace this
		&cli.StringSliceFlag{
			Name:    "config.env_printer_print_keys,env_printer_print_keys",
			Usage:   "if use this args, will print env by keys",
			EnvVars: []string{"PLUGIN_ENV_PRINTER_PRINT_KEYS"},
		},
		&cli.IntFlag{
			Name:    "config.env_printer_padding_left_max,env_printer_padding_left_max",
			Usage:   "set env printer padding left max count, minimum 24, default 32",
			EnvVars: []string{"PLUGIN_ENV_PRINTER_PADDING_LEFT_MAX"},
			Value:   32,
		},
		// env_printer_plugin end
		//&cli.StringFlag{
		//	Name:    "config.new_arg,new_arg",
		//	Usage:   "",
		//	EnvVars: []string{"PLUGIN_new_arg"},
		//},
		// file_browser_plugin end
	}
}

// HideFlag
// set env_printer_plugin hide flag at here
func HideFlag() []cli.Flag {
	return []cli.Flag{
		//&cli.UintFlag{
		//	Name:    "config.timeout_second,timeout_second",
		//	Usage:   "do request timeout setting second.",
		//	Hidden:  true,
		//	Value:   10,
		//	EnvVars: []string{"PLUGIN_TIMEOUT_SECOND"},
		//},
	}
}

// CommonFlag
// Other modules also have flags
func CommonFlag() []cli.Flag {
	return []cli.Flag{
		&cli.UintFlag{
			Name:    "config.timeout_second,timeout_second",
			Usage:   "do request timeout setting second.",
			Hidden:  true,
			Value:   10,
			EnvVars: []string{"PLUGIN_TIMEOUT_SECOND"},
		},
		&cli.BoolFlag{
			Name:    "config.debug,debug",
			Usage:   "debug mode",
			Value:   false,
			EnvVars: []string{"PLUGIN_DEBUG"},
		},
	}
}
