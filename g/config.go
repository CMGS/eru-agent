package g

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/projecteru/eru-agent/common"
	"github.com/projecteru/eru-agent/defines"
	"github.com/projecteru/eru-agent/logs"
	"gopkg.in/yaml.v2"
)

var Config = defines.AgentConfig{}

func LoadConfig() {
	var configPath string
	var version bool
	flag.BoolVar(&logs.Mode, "DEBUG", false, "enable debug")
	flag.StringVar(&configPath, "c", "agent.yaml", "config file")
	flag.BoolVar(&version, "v", false, "show version")
	flag.Parse()
	if version {
		logs.Info("Version", common.VERSION)
		os.Exit(0)
	}
	load(configPath)
}

func load(configPath string) {
	if _, err := os.Stat(configPath); err != nil {
		logs.Assert(err, "config file invaild")
	}

	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		logs.Assert(err, "Read config file failed")
	}

	if err := yaml.Unmarshal(b, &Config); err != nil {
		logs.Assert(err, "Load config file failed")
	}

	if Config.HostName, err = os.Hostname(); err != nil {
		logs.Assert(err, "Load hostname failed")
	}
	logs.Debug("Configure:", Config)
}
