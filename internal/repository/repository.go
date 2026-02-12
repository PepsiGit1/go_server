package repository

type Repository struct {
	// Add database connection here when needed
	// db *sql.DB
}

func NewRepository() *Repository {
	return &Repository{}
}
