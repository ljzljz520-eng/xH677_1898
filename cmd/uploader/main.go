package main

import (
	"fmt"
	"os"
	"social-uploader/internal/app"
	"social-uploader/internal/domain"
	"social-uploader/internal/flow077"
	"social-uploader/internal/store"
)

func main() {
	path := "social.db"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	db, e := store.Open(path)
	if e != nil {
		panic(e)
	}
	defer db.Close()
	svc := app.New(db)
	p := flow077.New(svc)
	r, _ := domain.NewRecord("demo-1", "demo", "EMP-001", 1200)
	out, e := p.ImportBatch("demo", []domain.Record{r})
	if e != nil {
		panic(e)
	}
	fmt.Println(out.Text)
}
