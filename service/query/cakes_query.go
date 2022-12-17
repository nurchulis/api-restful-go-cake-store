package cakes_query

import (
	"context"
	"fmt"
	"api-restful-cake-store/config"
	"api-restful-cake-store/models"
	"log"
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
			&cake.Images,
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

	queryText := fmt.Sprintf("INSERT INTO %v (title, description, rating, images, created_at, updated_at) values('%v','%v',%v,'%v','%v','%v')", table,
		cke.Title,
		cke.Description,
		cke.Rating,
		cke.Images,
		time.Now().Format(layoutDateTime),
		time.Now().Format(layoutDateTime))

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		log.Fatal("Failed Insert To Table ", err)
		return err
	}
	return nil
}

func GetDetail(ctx context.Context, id) ([]models.Cake, error) {
	log.Fatal("Cant connect to MySQL", id)
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
			&cake.Images,
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