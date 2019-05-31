package production

const host = "localhost"
const dbname = "database_test"
const user = "juniorbarros"
const password = "pass"
const source = "host="+host+" port=5432 user="+user+" dbname="+dbname+" password="+password+" sslmode=disable"

func GetSource() string {
	return source
}

