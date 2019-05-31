package development

const host = "localhost"
const dbname = "database_test"
const user = "juniorbarros"
const source = "host="+host+" port=5432 user="+user+" dbname="+dbname+" sslmode=disable"

func GetSource() string {
	return source
}