package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

//application config //map template render
type AppConfig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
    InfoLog  *log.Logger
	InProduction bool
	Session *scs.SessionManager
}


