package smtp

type Config struct {
	Enable   bool
	Host     string
	Port     int
	Username string
	Password string
	From     string
	To       []string
}
