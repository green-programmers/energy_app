package structs

type User struct {
	Username string
	Email    string
	Password string
}

type Db struct {
	Host     string
	Port     int
	DBname   string
	Password string
	Name     string
}
