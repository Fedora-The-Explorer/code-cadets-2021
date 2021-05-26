package config

import "github.com/kelseyhightower/envconfig"

// Cfg is the single instance of configuration that gets automatically populated from the
// environment variables once the  module loads.
var Cfg Config

// Config contains all the configuration needed for service to work.
type Config struct {
	Rabbit         rabbitConfig `split_words:"true"`
	SqliteDatabase string       `split_words:"true"`
}

type rabbitConfig struct {
	ConsumerBetFromControllerQueue string `split_words:"true" required:"true"`
	ConsumerEventUpdateQueue       string `split_words:"true" required:"true"`
	ConsumerBetFromControllerName  string `split_words:"true" default:"calculatorbet"`
	ConsumerEventUpdateName        string `split_words:"true" default:"calculatorevent"`
	ConsumerAutoAck                bool   `split_words:"true" default:"true"`
	ConsumerExclusive              bool   `split_words:"true" default:"false"`
	ConsumerNoLocal                bool   `split_words:"true" default:"false"`
	ConsumerNoWait                 bool   `split_words:"true" default:"false"`
	PublisherBetCalculatedQueue    string `split_words:"true" required:"true"`
	PublisherDeclareDurable        bool   `split_words:"true" default:"true"`
	PublisherDeclareAutoDelete     bool   `split_words:"true" default:"false"`
	PublisherDeclareExclusive      bool   `split_words:"true" default:"false"`
	PublisherDeclareNoWait         bool   `split_words:"true" default:"false"`
	PublisherExchange              string `split_words:"true" default:""`
	PublisherMandatory             bool   `split_words:"true" default:"false"`
	PublisherImmediate             bool   `split_words:"true" default:"false"`
	DeclareDurable                 bool   `split_words:"true" default:"true"`
	DeclareAutoDelete              bool   `split_words:"true" default:"false"`
	DeclareExclusive               bool   `split_words:"true" default:"false"`
	DeclareNoWait                  bool   `split_words:"true" default:"false"`
}

// Load loads the configuration on bootstrap, this avoid injecting the same config object
// everywhere.
func Load() {
	err := envconfig.Process("", &Cfg)
	if err != nil {
		panic(err)
	}
}
