package repositories

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type Repository[T any] interface {
	GetAll() ([]T, error)
	GetByID(id uint) (T, error)
	Create(item T) (T, error)
	Update(id uint, item T) (T, error)
	Delete(id uint) error
}

type GenericRepository[T any] struct {
	db        *sql.DB
	tableName string
	fields    []string
}

func NewGenericRepository[T any](db *sql.DB, tableName string, fields []string) *GenericRepository[T] {
	return &GenericRepository[T]{db: db, tableName: tableName, fields: fields}
}

func (r *GenericRepository[T]) GetAll() ([]T, error) {
	query := fmt.Sprintf("SELECT * FROM %s", r.tableName)
	rows, err := r.db.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var items []T
	for rows.Next() {
		var item T
		v := reflect.ValueOf(&item).Elem()
		dest := make([]interface{}, v.NumField())
		for i := 0; i < v.NumField(); i++ {
			dest[i] = v.Field(i).Addr().Interface()
		}

		err := rows.Scan(dest...)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *GenericRepository[T]) GetByID(id uint) (T, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE Id = ?", r.tableName)
	row := r.db.QueryRow(query, id)
	var item T
	v := reflect.ValueOf(&item).Elem()
		dest := make([]interface{}, v.NumField())
		for i := 0; i < v.NumField(); i++ {
			dest[i] = v.Field(i).Addr().Interface()
		}

		err := row.Scan(dest...)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return item, err
		}
	return item, nil
}

func (r *GenericRepository[T]) Create(item T) (T, error) {
	// Prepare placeholders and values for the SQL query
	placeholders := make([]string, len(r.fields))
	values := make([]interface{}, len(r.fields))
	v := reflect.ValueOf(item)
	// Ensure item is a struct
	if v.Kind() != reflect.Struct {
		return item, fmt.Errorf("item must be a struct")
	}

	for i, field := range r.fields {
		fieldValue := v.FieldByName(field)
		if !fieldValue.IsValid() {
			return item, fmt.Errorf("invalid field: %s", field)
		}
		placeholders[i] = "?"
		values[i] = fieldValue.Interface()
	}

// Add created_at and updated_at fields
created_at := time.Now()
updated_at := created_at
placeholders = append(placeholders, "?", "?")
	values = append(values, created_at, updated_at)
	// Create the SQL INSERT query
	query := fmt.Sprintf("INSERT INTO %s (%s, created_at, updated_at) VALUES (%s)", r.tableName, strings.Join(r.fields, ", "), strings.Join(placeholders, ", "))

	// Execute the SQL query
	result, err := r.db.Exec(query, values...)
	if err != nil {
		return item, err
	}

	// Retrieve the last inserted ID if needed
	lastID, err := result.LastInsertId()
	if err == nil {
		// Find and set the ID field if it exists
		idField := v.FieldByName("Id")
		if idField.IsValid() && idField.CanSet() {
			idField.Set(reflect.ValueOf(uint(lastID)))
		}
	}

	return item, nil
}

func (r *GenericRepository[T]) Update(id uint, item T) (T, error) {
	setClause := make([]string, len(r.fields))
	values := make([]interface{}, len(r.fields)+1)
	v := reflect.ValueOf(item)
	for i, field := range r.fields {
		setClause[i] = fmt.Sprintf("%s = ?", field)
		values[i] = v.FieldByName(field).Interface()
	}
	values[len(r.fields)] = id

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = ?", r.tableName, strings.Join(setClause, ", "))
	_, err := r.db.Exec(query, values...)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (r *GenericRepository[T]) Delete(id uint) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", r.tableName)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
