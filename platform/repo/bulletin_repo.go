package repo

import (
	"go-rest-gin-k8s/platform/bulletin"
	"log"
	"time"
)

func GetBulletins() ([]bulletin.Bulletin, error) {
	const q = `SELECT author, content, created_at FROM bulletins ORDER BY created_at DESC LIMIT 100`
	db, _ := GetConexion()
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	results := make([]bulletin.Bulletin, 0)

	for rows.Next() {
		var author string
		var content string
		var createdAt time.Time
		err = rows.Scan(&author, &content, &createdAt)
		if err != nil {
			return nil, err
		}
		results = append(results, bulletin.Bulletin{author, content, createdAt})
	}
	return results, nil
}
func AddBulletin(b bulletin.Bulletin) error {
	const q = `INSERT INTO bulletins(author, content, created_at) VALUES ($1, $2, $3)`
	db, _ := GetConexion()
	_, err := db.Exec(q, b.Author, b.Content, b.CreatedAt)
	return err
}
func Migrations() {
	var err error
	db, err := GetConexion()
	//db, err = sql.Open("postgres", dbInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	_, err = db.Query(Migration)
	if err != nil {
		log.Println("Failed run Migrations", err.Error())
		panic(err)
		return
	}
	log.Println("Running Migrations....")
}
