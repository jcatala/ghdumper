package config

type Config struct {
	Port	int64
	Ssl		bool
	Pemfile	string
	Key		string
	Verbose	bool
	ResponseFile	string
	AutoCL	bool
}