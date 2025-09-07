package conf

import (
	"errors"
	"sync"

	"github.com/nextadm/pkg/runner"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

var (
	conf *viper.Viper
	once sync.Once
)

func InitConf(options *runner.ServerRunOptions) {
	once.Do(func() {
		config, err := NewConf(options)
		if err != nil {
			panic(err)
		}
		conf = config
	})
}

func NewConf(options *runner.ServerRunOptions) (*viper.Viper, error) {
	if options.Endpoint == "" || options.Path == "" {
		return nil, errors.New("endpoint or path is empty")
	}

	v := viper.New()
	err := v.AddRemoteProvider("etcd3", options.Endpoint, options.Path)
	if err != nil {
		return nil, err
	}

	v.SetConfigType("yaml")

	err = v.ReadRemoteConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}
