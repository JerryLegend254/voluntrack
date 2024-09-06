package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/JerryLegend254/voluntrack/pkg/models/mysql"
)

type application struct {
	templateCache map[string]*template.Template
	infoLog       *log.Logger
	errorLog      *log.Logger
}

func main() {

	addr := flag.String("addr", ":8080", "Server port")
	dsn := flag.String("dsn", "root:pass254@tcp(localhost:3307)/voluntrack?parseTime=true", "MYSQL DSN")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	tc, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	db, err := openDB(*dsn)

	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	err = mysql.SetUpDB(db)
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{templateCache: tc, infoLog: infoLog, errorLog: errorLog}

	srv := http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	infoLog.Println("Server started on port,", *addr)
	errorLog.Fatal(srv.ListenAndServe())

}
