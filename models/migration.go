package models

var migrationModels = []interface{}{
	//table name here
	&UrlInfo{},
}

func GetMigrationModels() []interface{} {
	return migrationModels
}
