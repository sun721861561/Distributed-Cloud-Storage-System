package config

const (
	// MySQLSource : The database source to connect to.
	// test:test is the username:password.
	// 127.0.0.1:3306 is the ip and port.
	// fileserver is the database name; 
	// charset=utf8 Data is transmitted in utf8 character encoding
	MySQLSource = "test:test@tcp(127.0.0.1:3306)/fileserver?charset=utf8"
)