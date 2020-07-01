package structure

// Config Structure of config.json File
type Config struct {
	Host      string
	Username  string
	Password  string
	Port      int
	DBType    string
	Schema    string
	JwtSecret string
}
