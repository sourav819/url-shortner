package models

var migrationModels = []interface{}{
	//table name here
	&UrlInfo{},
	&User{},
}

func GetMigrationModels() []interface{} {
	return migrationModels
}
