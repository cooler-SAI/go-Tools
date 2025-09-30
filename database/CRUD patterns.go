package database

/*// Basic CRUD patterns to memorize:

// 1. CONNECTION
db, err := sql.Open("postgres", "your-dsn-here")
defer db.Close()
db.Ping()

// 2. CREATE (INSERT)
db.Exec("INSERT INTO table (col1, col2) VALUES ($1, $2)", val1, val2)

// 3. READ MULTIPLE
rows, err := db.Query("SELECT * FROM table")
defer rows.Close()
for rows.Next() {
rows.Scan(&var1, &var2)
}

// 4. READ SINGLE
db.QueryRow("SELECT ...").Scan(&var1)

// 5. UPDATE
db.Exec("UPDATE table SET col1 = $1 WHERE id = $2", newVal, id)

// 6. DELETE
db.Exec("DELETE FROM table WHERE id = $1", id)*/
