package system

import "Dotato-di-una-libreria/backend/util"

// CloudSQLSetting ...
type CloudSQLSetting struct {
	ID           int64 `goon:"id"`
	User         string
	Password     string
	InstanceStr  string
	DBName       string
	MaxIdleConns int
	MaxOpenConns int
}

// InstanceStrFormat ...
func InstanceStrFormat() string {
	if util.IsLocal() {
		return "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local"
	}
	return "%s:%s@unix(/cloudsql/%s)/%s?parseTime=True"
}
