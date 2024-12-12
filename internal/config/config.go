package config

import (
	"flag"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Logger   Logger   `mapstructure:",squash"`
	Postgres Postgres `mapstructure:",squash"`
	Telegram Telegram `mapstructure:",squash"`
	Server   Server   `mapstructure:",squash"`
}

type Logger struct {
	OutputPaths []string
}

type Postgres struct {
	User     string `mapstructure:"PG_USER"`
	DBName   string `mapstructure:"PG_DBNAME"`
	Password string `mapstructure:"PG_PASSWORD"`
	Host     string `mapstructure:"PG_HOST"`
	Port     string `mapstructure:"PG_PORT"`
}

type Telegram struct {
	BotToken       string `mapstructure:"TG_BOT_TOKEN"`
	AdminChatID    int64  `mapstructure:"ADMIN_CHAT_ID"`
	PollingTimeout time.Duration
	IsSecondRoster bool
}

type Server struct {
	Port          string `mapstructure:"APP_PORT"`
	APIPrefix     string
	CORS          CORS
	IsAPIDisabled bool
}

type CORS struct {
	AllowedHeaders     []string
	AllowedOrigins     []string
	AllowedMethods     []string
	AllowedCredentials bool
}

var (
	Default Config = Config{
		Logger: Logger{
			OutputPaths: []string{
				"stderr",
			},
		},
		Telegram: Telegram{
			PollingTimeout: 2 * time.Second,
		},
		Server: Server{
			APIPrefix: "/api/v1",
			CORS: CORS{
				AllowedHeaders:     DefaultAllowedHeaders,
				AllowedOrigins:     DefaultAllowedOrigins,
				AllowedMethods:     DefaultAllowedMethods,
				AllowedCredentials: true,
			},
		},
	}
)

func parseFeatureFlags() {
	isSecondRoster := flag.Bool("sr", false, "sr flag defines if bot instance is for processing second roster applications")
	isAPIDisabled := flag.Bool("noapi", false, "noapi flag defines if application REST API server is disabled")
	flag.Parse()

	Default.Telegram.IsSecondRoster = *isSecondRoster
	Default.Server.IsAPIDisabled = *isAPIDisabled
}

func NewConfig() (cfg *Config, err error) {
	parseFeatureFlags()
	v := viper.New()

	v.AutomaticEnv()
	// without v.BindEnv call Viper can't load env variables from runtime
	// empty string argument can be replaced with any string as long as BindEnv is called and no error is returned
	// if no argument is provided (e.g. v.BindEnv()) then "missing key to bind to" error is returned
	err = v.BindEnv("")
	if err != nil {
		return nil, err
	}

	err = v.Unmarshal(&Default)
	if err != nil {
		return nil, err
	}

	return &Default, nil
}
