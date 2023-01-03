package config

type Mysql struct {
	Host     string //服务器ip
	port     string //服务器端口
	Username string
	Password string
	Dbname   string //连接的数据库名
}

func (m *Mysql) Dsn() string {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.port + ")/" + m.Dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
}
