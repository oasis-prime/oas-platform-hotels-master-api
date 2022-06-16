package configs

type app struct {
	Debug bool   `mapstructure:"debug"`
	Env   string `mapstructure:"env"`
	Port  string `mapstructure:"port"`
}

type jwt struct {
	Secret          string `mapstructure:"secret"`
	Expirationdelta int    `mapstructure:"expirationdelta"`
}

type database struct {
	Name string `mapstructure:"name"`
	Read struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"read"`
	Write struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"write"`
}

type redis struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type Config struct {
	App      app      `mapstructure:"app"`
	Database database `mapstructure:"database"`
	Redis    redis    `mapstructure:"redis"`
	Jwt      jwt      `mapstructure:"jwt"`
}
