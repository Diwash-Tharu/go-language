package main

import (
	
	"html/template"
	"os"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Port string
	DBSource string
}

//  initilizinge 
func loadConfig() Config {
	return Config{
		Port: getEnv("PORT", "8080"),
		DBSource: getEnv("DB_SOURCE", "./data/orders.db"),

	}
}

func getEnv(key, defaultValue string) string {
	if value :=os.Getenv(key);value != ""{
		return value 
	}
	return defaultValue
}

// fila addn the template so that go and gem can use it 

func loadTemplates(router *gin.Engine) error{
	functions :=template.FuncMap{
		"add": func(a, b int) int {
			return a + b },
	}
	tmpl, err :=template.New("").Funcs(functions).ParseGlob("templates/*tmpl")
	if err != nil {
		return err 
	}

	router.SetHTMLTemplate(tmpl)
	return nil 
}