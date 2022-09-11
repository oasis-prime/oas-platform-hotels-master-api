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
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	PWD      string `mapstructure:"pwd"`
	Instance struct {
		Connection struct {
			Name string `mapstructure:"name"`
		} `mapstructure:"connection"`
	} `mapstructure:"instance"`
	Socket struct {
		Dir bool `mapstructure:"dir"`
	} `mapstructure:"socket"`
}

type google struct {
	Places    string `mapstructure:"places"`
	Pubsubkey string `mapstructure:"pubsubkey"`
	Projectid string `mapstructure:"projectid"`
}

type redis struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type hotelbeds struct {
	Endpoint       string `mapstructure:"endpoint"`
	SecureEndopint string `mapstructure:"secureendopint"`
	Key            string `mapstructure:"key"`
	Secret         string `mapstructure:"secret"`
	Format         string `mapstructure:"format"`
}

type Config struct {
	App       app       `mapstructure:"app"`
	Database  database  `mapstructure:"database"`
	Redis     redis     `mapstructure:"redis"`
	Jwt       jwt       `mapstructure:"jwt"`
	Hotelbeds hotelbeds `mapstructure:"hotelbeds"`
	Google    google    `mapstructure:"google"`
}
