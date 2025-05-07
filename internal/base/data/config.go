package data

type Database struct {
	Driver          string `json:"driver" yaml:"driver"`
	Connection      string `json:"connection" yaml:"connection"`
	ConnMaxLifeTime int    `json:"conn_max_life_time" yaml:"conn_max_life_time"`
	MaxOpenConn     int    `json:"max_open_conn" yaml:"max_open_conn"`
	MaxIdleConn     int    `json:"max_idle_conn" yaml:"max_idle_conn"`
}
