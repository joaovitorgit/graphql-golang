package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name, description, CategoryID string) (*Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO course(id, name, description, category_id) VALUES($1,$2,$3,$4)", id, name, description, CategoryID)
	if err != nil {
		return nil, err
	}
	return &Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryID:  CategoryID,
	}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM course")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []Course{}
	for rows.Next() {
		var id, name, description, categoryID string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		courses = append(courses, Course{ID: id, Name: name, Description: description, CategoryID: categoryID})
	}
	return courses, nil
}

func (c *Course) FindByCategoryId(CategoryID string) ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM course WHERE category_id=$1", CategoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []Course{}
	for rows.Next() {
		var id, name, description, categoryID string
		if err := rows.Scan(&id, &name, &description, &categoryID); err != nil {
			return nil, err
		}
		courses = append(courses, Course{ID: id, Description: description, Name: name, CategoryID: categoryID})
	}
	return courses, nil
}
