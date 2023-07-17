package models

import "time"

// Repo is a database model for a repository.
type Repo struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`
	ProjectName string    `db:"project_name"`
	Description string    `db:"description"`
	Private     bool      `db:"private"`
	Mirror      bool      `db:"mirror"`
	Hidden      bool      `db:"hidden"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
