package db

var DB DataBases

type DataBases struct {
	MySQLDB mysqlDB
}

func init() {
	initMySQLDB()
}
