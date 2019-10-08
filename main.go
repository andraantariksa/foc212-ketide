package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/labstack/echo"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func findAndParseTemplates(rootDir string, funcMap template.FuncMap) (*template.Template, error) {
	cleanRoot := filepath.Clean(rootDir)
	pfx := len(cleanRoot) + 1
	root := template.New("")

	err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".html") {
			if e1 != nil {
				return e1
			}

			b, e2 := ioutil.ReadFile(path)
			if e2 != nil {
				return e2
			}

			name := path[pfx:]
			t := root.New(name).Funcs(funcMap)
			t, e2 = t.Parse(string(b))
			if e2 != nil {
				return e2
			}
		}

		return nil
	})

	return root, err
}

type Users struct {
	ID        int64     `xorm:"pk not null autoincr"`
	Username  string    `xorm:"unique not null"`
	Email     string    `xorm:"not null"`
	Password  string    `xorm:"not null"`
	CreatedAt time.Time `xorm:"created not null"`
	UpdatedAt time.Time `xorm:"updated not null"`
}

type Code struct {
	ID        int64  `xorm:"pk not null autoincr"`
	Language  string `xorm:"not null"`
	Code      string `xorm:"not null"`
	Stdin     string
	Stdout    string
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

func migrateTable(db *xorm.Engine) {
	if err := db.Sync(new(Users)); err != nil {
		panic("Table users migrate error")
	}

	if err := db.Sync(new(Code)); err != nil {
		panic("Table code migrate error")
	}
}

// func dropTable(db *xorm.Engine) {
// 	if err := db.DropTables(new(Users)); err != nil {
// 		panic("Table users drop error")
// 	}

// 	if err := db.DropTables(new(Code)); err != nil {
// 		panic("Table code drop error")
// 	}
// }

func dbConnect() *xorm.Engine {
	dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	dbEngine, err := xorm.NewEngine("postgres", dbURL)

	if err != nil {
		panic("Database engine creation error")
	}

	if err := dbEngine.Ping(); err != nil {
		panic("Database ping error")
	}

	return dbEngine
}

func main() {
	db := dbConnect()
	migrateTable(db)

	e := echo.New()

	e.Static("/static", "static")

	// renderer := &TemplateRenderer{
	// 	 templates: template.Must(template.ParseGlob("templates/*.html")),
	// }
	renderer := &TemplateRenderer{
		templates: template.Must(findAndParseTemplates("templates", template.FuncMap{})),
	}
	// renderer, err := findAndParseTemplates("templates", a)
	e.Renderer = renderer

	// Main

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "home.html", map[string]interface{}{})
	})

	e.GET("/help", func(c echo.Context) error {
		return c.Render(http.StatusOK, "help.html", map[string]interface{}{})
	})

	// User

	e.GET("/signup", func(c echo.Context) error {
		return c.Render(http.StatusOK, "user/signup.html", map[string]interface{}{})
	})

	e.POST("/signup", func(c echo.Context) error {
		users := new(Users)
		if err := c.Bind(users); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		if _, err := db.InsertOne(users); err != nil {
			return c.String(http.StatusOK, fmt.Sprintf("%s", err))
		}

		return c.String(http.StatusOK, "Ok")
	})

	e.GET("/signin", func(c echo.Context) error {
		return c.Render(http.StatusOK, "user/signin.html", map[string]interface{}{})
	})

	e.GET("/settings", func(c echo.Context) error {
		return c.Render(http.StatusOK, "user/settings.html", map[string]interface{}{})
	})

	e.GET("/recent", func(c echo.Context) error {
		return c.Render(http.StatusOK, "code.html", map[string]interface{}{})
	})

	// Code

	e.GET("/:id", func(c echo.Context) error {
		return c.Render(http.StatusOK, "single-code.html", map[string]interface{}{})
	})

	e.Logger.Debug(e.Start(":1324"))
}
