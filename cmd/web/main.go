package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/shunmasa/bookings/pkg/config"
	"github.com/shunmasa/bookings/pkg/handlers"
	"github.com/shunmasa/bookings/pkg/render"
)


const portNumber = ":8080" 

var app config.AppConfig
var session *scs.SessionManager

func main(){

 //Map tempalte only once and cache 
//change this to true when in production  
session = scs.New()
session.Lifetime = 24 * time.Hour
session.Cookie.Persist = true 
session.Cookie.SameSite = http.SameSiteLaxMode
session.Cookie.Secure = app.InProduction //dev or production bool 

app.Session = session



//call render
tc, err := render.CreateTemplateCache()
if err != nil{
  log.Fatal("can not create tempalte cache")
}

app.TemplateCache = tc
app.UseCache = false //kind of global state 


repo:= handlers.NewRepo(&app)
handlers.Newhandlers(repo)


//access to the conf variable 
render.NewTemplates(&app)

//request pointer allocate process 
//go run *.go

  // http.HandleFunc("/",handlers.Repo.Home)//Repo add bolean value 
  // http.HandleFunc("/about",handlers.Repo.About)

//   http.HandleFunc("/divide",Divide)

//   http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
// 	n, err := fmt.Fprintf(w,"Hello,world")
//     if err != nil {
// 		fmt.Println(err)
// 	})

// 	fmt.Println(fmt.Sprintf("NUmber of Bytes Written:%d",n))// differents type 
//   }
   

  //start web server
  fmt.Printf("Starting application on port %s\n", portNumber)

  srv:= &http.Server{
  Addr:portNumber,
  Handler:routes(&app),
 }
 err = srv.ListenAndServe()
 log.Fatal(err)

}













