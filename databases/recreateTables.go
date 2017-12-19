package databases

//RecreateTables Elimina y Crea las tablas de la base de datos.
func RecreateTables() {
	db := GetConnectionDB()
	defer db.Close()

	DropTables()
	CreateTables()
}
