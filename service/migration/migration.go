package migration

import (
	"context"
	"api-restful-cake-store/config"
	"log"
	"io/ioutil"
)

// CreateTable Migration Not use ORM, just load migration and run query file
func CreatTable(ctx context.Context) error {
	files, err := ioutil.ReadDir("migration/command")
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
		content, err := ioutil.ReadFile("migration/command/"+file.Name())
		if err != nil {
			log.Fatal(err)
		}
		
		db, err := config.MySQL()
	
		if err != nil {
			log.Fatal("Can't connect to MySQL", err)
		}
	
		sql := string(content)
		_, err = db.ExecContext(ctx, sql)
	
		if err != nil {
			log.Fatal("Failed Migrate Table ", err)
			return err
		}
    }
	return nil
}