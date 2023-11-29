// helpers/queries.go

package helpers

// SelectAllTasksQuery returns the SQL query to select all tasks from a table
func SelectAllTasksQuery() string {
	return "SELECT * FROM posts"
}

// InsertTaskQuery returns the SQL query to insert a task into a table
func InsertTaskQuery() string {
	return "INSERT INTO posts (id, title, body) VALUES (?,?,?)"
}

// Add more queries as needed
///
