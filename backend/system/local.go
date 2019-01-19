package system

// LocalDatastoreSetup ...
func LocalDatastoreSetup() CloudSQLSetting {
	dbc := &CloudSQLSetting{
		ID:          1,
		User:        "testuser",
		Password:    "testpass",
		InstanceStr: "127.0.0.1:3306",
		DBName:      "testdb"}
	return *dbc
}
