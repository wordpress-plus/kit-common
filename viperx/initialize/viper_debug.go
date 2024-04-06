package initialize

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func ViperDebug(conf any, path ...string) *viper.Viper {
	v := viper.NewWithOptions()
	v.SetConfigFile(getConfigFile(path...))
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		unmarshalConfig(v, conf)
	})

	unmarshalConfig(v, conf)

	return v
}

func getConfigFileDebug(path ...string) string {

	var config string
	if len(path) != 0 {
		config = path[0]
	}
	// parse from cmd first
	if config == "" {
		fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
		fs.StringVar(&config, "c", "", "choose config file.")
		_ = flag.CommandLine.Parse(os.Args[1:])
	}

	if len(config) == 0 {
		if con := os.Getenv("config"); len(con) == 0 {
			config = "./config-local.yaml"
		} else {
			config = con
		}
	}

	// pickup default value
	if len(config) == 0 {
		panic("config file is not configured")
	}

	fmt.Printf("using viper config: %s\n", config)

	return config
}
