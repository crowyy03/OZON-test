package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	_ "github.com/lib/pq"
)

const (
	connstr = "postgres://root:1234@localhost:5432/urls?sslmode=disable"
)

func main() {
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	migrationsPath := "./migrations"

	files, err := os.ReadDir(migrationsPath)
	if err != nil {
		panic(err)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			content, err := os.ReadFile(filepath.Join(migrationsPath, file.Name()))
			if err != nil {
				panic(err)
			}
			if _, err := db.Exec(string(content)); err != nil {
				panic(fmt.Errorf("ошибка выполнения миграции %s: %w", file.Name(), err))
			}
			fmt.Printf("Миграция %s применена успешно\n", file.Name())
		}
	}
}
