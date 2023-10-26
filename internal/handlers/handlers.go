package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/shunmasa/bookings/internal/config"
	"github.com/shunmasa/bookings/internal/models"
	"github.com/shunmasa/bookings/internal/render"

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

	render.RenderTemplate(w, r,"home.page.gohtml",&models.TemplateData{})

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
   
  render.RenderTemplate(w,r, "about.page.gohtml", &models.TemplateData{
   StringMap: stringMap,
  })
  
  }


  func (m *Repository) Reservation(w http.ResponseWriter,r *http.Request) {
  
   render.RenderTemplate(w,r,"make-reservation.page.gohtml",&models.TemplateData{})
   }

   func (m *Repository) Generals(w http.ResponseWriter,r *http.Request) {
  
    render.RenderTemplate(w,r,"generals.page.gohtml",&models.TemplateData{})
    }
 
    func (m *Repository) Majors(w http.ResponseWriter,r *http.Request) {
  
      render.RenderTemplate(w,r,"majors.page.gohtml",&models.TemplateData{})
      }
   

      func (m *Repository) Availability(w http.ResponseWriter,r *http.Request) {
  
        render.RenderTemplate(w,r,"search-availability.page.gohtml",&models.TemplateData{})
        }

        func (m *Repository) PostAvailability(w http.ResponseWriter,r *http.Request) {
          start := r.Form.Get("start")
          end := r.Form.Get("end")
        
          w.Write([]byte(fmt.Sprintf("start date is %s and end is %s", start, end)))
         }
        func (m *Repository) Contact(w http.ResponseWriter,r *http.Request) {
  
          render.RenderTemplate(w,r,"contact.page.gohtml",&models.TemplateData{})
          }


type jsonResponse struct{
  Ok bool `json:"ok"`
  Message string  `json:"message"`
}
//AvailabilityJSON handles request for availability and send JSON response
    func (m *Repository) AvailabilityJSON(w http.ResponseWriter,r *http.Request) {
       resp := jsonResponse{
         Ok: true,
         Message: "Available",
       }

       out, err := json.MarshalIndent(resp,"","     ")
       if err != nil {
        log.Println(err)
       }
       log.Println(string(out))
      w.Header().Set("Content-Type","application/json")
      w.Write(out)
  
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
   
