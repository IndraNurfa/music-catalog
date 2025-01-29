package configs

type (
	Config struct {
		Service       Service       `maspstructure:"service"`
		Database      Database      `mapstructure:"database"`
		SpotifyConfig SpotifyConfig `mapstructure:"spotifyConfig"`
	}

	Service struct {
		Port      string `mapstructure:"port"`
		SecretJWT string `mapstructure:"secretJWT"`
	}

	Database struct {
		DataSourceName string `mapstructure:"dataSourcename"`
	}

	SpotifyConfig struct {
		ClientID     string `mapstructure:"clientID"`
		ClientSecret string `mapstructure:"clientSecret"`
	}
)
