package config

type Conf struct {
	AppName  string
	Database database
}

type database struct {
	HostName      string
	Port          int
	ConnectionMax int
}
