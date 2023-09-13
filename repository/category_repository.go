package repository

import (
	"database/sql"
	"final-project-enigma-clean/model"
)

type CategoryRepository interface {
	Save(category model.Category) error
	FindById(id string) (model.Category, error)
	FindAll() ([]model.Category, error)
	Update(category model.Category) error
	Delete(id string) error
}

type categoryRepository struct {
	db *sql.DB
}

// FindById implements categoryRepository.
func (c *categoryRepository) FindById(id string) (model.Category, error) {
	row := c.db.QueryRow("SELECT id, name FROM category WHERE id = $1", id)
	var category model.Category
	err := row.Scan(&category.Id, &category.Name)
	if err != nil {
		return model.Category{}, err
	}

	return category, nil

}

// Delete implements categoryRepository.
func (c *categoryRepository) Delete(id string) error {
	_, err := c.db.Exec("DELETE FROM category WHERE id= $1", id)
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements categoryRepository.
func (c *categoryRepository) FindAll() ([]model.Category, error) {
	rows, err := c.db.Query("SELECT id, name FROM category")
	if err != nil {
		return nil, err
	}
	var categories []model.Category
	for rows.Next() {
		var category model.Category
		rows.Scan(&category.Id, &category.Name)
		categories = append(categories, category)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return categories, nil
}

// Save implements categoryRepository.
func (c *categoryRepository) Save(category model.Category) error {
	_, err := c.db.Exec("INSERT INTO category VALUES ($1,$2)", category.Id, category.Name)
	if err != nil {
		return err
	}
	return nil
}

// Update implements categoryRepository.
func (c *categoryRepository) Update(category model.Category) error {
	_, err := c.db.Exec("UPDATE category SET name=$2 WHERE id=$1", category.Id, category.Name)
	if err != nil {
		return err
	}
	return nil
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}
