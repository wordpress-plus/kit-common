package initialize

import (
	"flag"
	"fmt"
	jasypt "github.com/alice52/jasypt-go"
	jasyptv "github.com/alice52/jasypt-go/viper"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func Viper(conf any, path ...string) *viper.Viper {
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

func getConfigFile(path ...string) string {

	var config string
	// parse from cmd first

	fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	fs.StringVar(&config, "c", "", "choose config file.")
	_ = flag.CommandLine.Parse(os.Args[1:])

	// parse from code
	if len(config) == 0 && len(path) != 0 {
		config = path[0]
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

// unmarshalConfig
//  1. parse config
//  2. do decrypt by jasypt
func unmarshalConfig(v *viper.Viper, conf any) {
	if err := jasyptv.Unmarshal(v, jasypt.New(), &conf); err != nil {
		panic(err)
	}
}
