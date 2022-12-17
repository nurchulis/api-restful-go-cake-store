package cakes_query

import (
	"context"
	"fmt"
	"api-restful-cake-store/config"
	"api-restful-cake-store/models"
	"log"
	"database/sql"
	"errors"
	"time"
)

const (
	table          = "cakes"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll
func GetAll(ctx context.Context) ([]models.Cake, error) {

	var cakes []models.Cake

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By title, rating DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var cake models.Cake
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&cake.ID,
			&cake.Title,
			&cake.Description,
			&cake.Rating,
			&cake.Image,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		cake.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		cake.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		cakes = append(cakes, cake)
	}

	return cakes, nil
}


func Insert(ctx context.Context, cke models.Cake) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (title, description, rating, image, created_at, updated_at) values('%v','%v',%v,'%v','%v','%v')", table,
		cke.Title,
		cke.Description,
		cke.Rating,
		cke.Image,
		time.Now().Format(layoutDateTime),
		time.Now().Format(layoutDateTime))

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		log.Fatal("Failed Insert To Table ", err)
		return err
	}
	return nil
}

func GetDetail(ctx context.Context, cke models.Cake) ([]models.Cake, error) {

	var cakes []models.Cake

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}
	
	queryText := fmt.Sprintf("SELECT * FROM %v where id = '%d'", table, cke.ID)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var cake models.Cake
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&cake.ID,
			&cake.Title,
			&cake.Description,
			&cake.Rating,
			&cake.Image,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		cake.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		cake.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		cakes = append(cakes, cake)
	}

	return cakes, nil
}

func Delete(ctx context.Context, cke models.Cake) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = '%d'", table, cke.ID)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada ")
	}

	return nil
}

// Update
func Update(ctx context.Context, cke models.Cake) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set title = '%s', description ='%s', rating = %d, image = '%s', updated_at = '%v' where id = '%d'",
		table,
		cke.Title,
		cke.Description,
		cke.Rating,
		cke.Image,
		time.Now().Format(layoutDateTime),
		cke.ID,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}