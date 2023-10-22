package handlers

import (
	"net/http"

	"github.com/shunmasa/bookings/pkg/config"
	"github.com/shunmasa/bookings/pkg/models"
	"github.com/shunmasa/bookings/pkg/render"
)




var Repo *Repository
//Repository type
type Repository struct {
  App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig)*Repository{
  return &Repository{
   App:a,

  }
}
//New Handlers sets the repository for the handlers
func Newhandlers(r *Repository){
  Repo = r
}


//Home is the home page handler
  func (m *Repository)Home(w http.ResponseWriter,r *http.Request){
  remoteIP := r.RemoteAddr
  //first user visits website, register remoteIP 
  m.App.Session.Put(r.Context(),"remote_ip",remoteIP)

	render.RenderTemplate(w, "home.page.gohtml",&models.TemplateData{})

  }
  
  
  
  //About is the about page handler
  func (m *Repository) About(w http.ResponseWriter,r *http.Request) {
   


   stringMap := make(map[string]string)
   stringMap["test"] = "Hello, again."
   //send the data to the template
  //  m.App.Session
  remoteIP := m.App.Session.GetString(r.Context(),"remote_ip")
  // m.App.Session.
  stringMap["remote_ip"] = remoteIP
   
  render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{
   StringMap: stringMap,
  })
  
  }
















  
//    func addValues(x,y int) int{
// 	   return x + y 
//    }
  
  
  
//    func Divide(w http.ResponseWriter,r *http.Request){
// 	   f, err := divideValues(100.0, 10.0)
// 	   if err != nil {
// 		   fmt.Fprintf(w,"cannot divide by 0 ")
// 		   return //stop executing function 
// 	   }
  
// 	   fmt.Fprintf(w, "%f divided by %f is %f", 100.0, 10.0, f)
  
  
//    }
  
//    func divideValues(x, y float32) (float32,error){
// 	   if y <= 0 {
// 		   err:= errors.New("cannot divide by zero")
// 		   return 0 ,err
// 	   }
  
// 	   result := x / y
// 		return result, nil
  
//    }




   
//    func addValues(x,y int) int{
// 	   return x + y 
//    }
   
   
   
//    func Divide(w http.ResponseWriter,r *http.Request){
// 	   f, err := divideValues(100.0, 10.0)
// 	   if err != nil {
// 		   fmt.Fprintf(w,"cannot divide by 0 ")
// 		   return //stop executing function 
// 	   }
   
// 	   fmt.Fprintf(w, "%f divided by %f is %f", 100.0, 10.0, f)
   
   
//    }
   
//    func divideValues(x, y float32) (float32,error){
// 	   if y <= 0 {
// 		   err:= errors.New("cannot divide by zero")
// 		   return 0 ,err
// 	   }
   
// 	   result := x / y
// 		return result, nil
   
//    }
   
