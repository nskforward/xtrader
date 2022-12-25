package config

var _instance *Config

func Get() *Config {
	if _instance == nil {
		_instance = newConfig()
	}
	return _instance
}
