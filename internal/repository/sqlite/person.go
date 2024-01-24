package sqlite

import (
	"database/sql"
	"fmt"
	"task_peopler/internal/models"
)

type PersonRepo struct {
	db *sql.DB
}

func NewPersonRepo(db *sql.DB) *PersonRepo {
	return &PersonRepo{db: db}
}

func (s *PersonRepo) Create(p *models.Person) (*models.Person, error) {
	query := `INSERT INTO People (first_name, last_name, patronymic, age, gender, nationality) VALUES (?, ?, ?, ?, ?, ?)`

	result, err := s.db.Exec(query, p.FirstName, p.FirstName, p.Patronymic, p.Age, p.Gender, p.Nationality)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert id: %v", err)
	}

	p.ID = id

	return p, nil
}

func (s *PersonRepo) Update(p *models.Person) (*models.Person, error) {
	query := `
		UPDATE People
		SET first_name=?, last_name=?, patronymic=?, age=?, gender=?, nationality=?
		WHERE id=?
	`

	result, err := s.db.Exec(query, p.FirstName, p.LastName, p.Patronymic, p.Age, p.Gender, p.Nationality, p.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("no rows updated, record with ID %d not found", p.ID)
	}

	return p, nil
}

func (s *PersonRepo) Delete(id int64) error {
	query := `DELETE FROM People WHERE id=?`

	result, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows deleted, record with ID %d not found", id)
	}

	return nil
}

func (s *PersonRepo) GetByID(id int64) (*models.Person, error) {
	query := `SELECT id, first_name, last_name, patronymic, age, gender, nationality FROM People WHERE id=?`

	row := s.db.QueryRow(query, id)

	var person models.Person
	err := row.Scan(&person.ID, &person.FirstName, &person.LastName, &person.Patronymic, &person.Age, &person.Gender, &person.Nationality)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no person found with ID %d", id)
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	return &person, nil
}

func (s *PersonRepo) GetPeople(filters map[string]interface{}, page, pageSize int) ([]*models.Person, error) {
	query := `SELECT id, first_name, last_name, patronymic, age, gender, nationality FROM People WHERE 1`

	// filters to the query
	args := make([]interface{}, 0)
	for key, value := range filters {
		query += fmt.Sprintf(" AND %s=?", key)
		args = append(args, value)
	}

	query += " LIMIT ? OFFSET ?"
	args = append(args, pageSize, (page-1)*pageSize)

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	people := make([]*models.Person, 0)
	for rows.Next() {
		var person models.Person
		err := rows.Scan(&person.ID, &person.FirstName, &person.LastName, &person.Patronymic, &person.Age, &person.Gender, &person.Nationality)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		people = append(people, &person)
	}

	return people, nil
}
