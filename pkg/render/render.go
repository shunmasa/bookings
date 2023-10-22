package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/shunmasa/bookings/pkg/config"
	"github.com/shunmasa/bookings/pkg/models"
)

func AddDefaultData(td *models.TemplateData)*models.TemplateData{
	return td
}

// var functions = template.FuncMap{

// }

var app *config.AppConfig
//new Templates sets the config for the template package
func NewTemplates(a *config.AppConfig){
	app = a
}



//needs data sending to template
//RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, gohtml string,td *models.TemplateData){
	var tc map[string]*template.Template 
	//if cache is true 
	if app.UseCache{
	   //get the template cache from the app 
		tc = app.TemplateCache
	}else{
		//create a template chache 
		tc,_= CreateTemplateCache()
	}
    

    // t,err := CreateTemplateCache()
	// if err != nil{
	// 	log.Fatal(err)
	// }
	//get requested template from cache
    t,ok := tc[gohtml]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}
    //execute as buffer directly
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
	_ = t.Execute(buf,td)//pass data ,td

	_,err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writting template to browser",err)
	}

	// err =t.Execute(buf,nil)
	// if err != nil{
	// 	log.Println(err)
	// }

	//render the template
	
	//every thingle request parse files ???? ineficient 
	// parsedTemplate, _ := template.ParseFiles("./templates/" + gohtml, "./templates/base.layout.gohtml")
	// err:= parsedTemplate.Execute(w , nil)
	// if err != nil {
	// 	fmt.Println("error parsing temolate:", err)
	// 	return 
	// }
}


//return map type and error

func CreateTemplateCache()(map[string]*template.Template,error){
	// myCache:= make(map[string]*template.Template)
	//Parse renders
	myCache := make(map[string]*template.Template)
 
	//get all of the files named *.page.temp from ./templates
	pages,err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil{
		return myCache,err
	}

	//range through all files ending with *.page.tmpl
	for _,page := range pages{
		//file path file name itself
		name :=filepath.Base(page)
		log.Println("this is base", name)
		//this is base home.page.gohtml
		//this is base about.page.gohtml

		//call template by name, name and page 
		ts, err :=template.New(name).ParseFiles(page)
		log.Println("this is page:", page)//page: templates/home.page.gohtml 
		//this is page: templates/home.page.gohtml
		log.Println("this is ts:", ts)
		//2023/10/21 17:10:41 this is ts: &{<nil> 0x140002a0b80 0x140002c10e0 0x140002b07e0}
		//2023/10/21 17:10:41 this is ts: &{<nil> 0x140002a0080 0x140002c0000 0x140002b0060}
		if err != nil{
			return myCache,err
		}
      ///single path ,layout only
		matches,err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil{
			return myCache,err
		}
       //match length 
		if len(matches) > 0 {
			ts,err = ts.ParseGlob("./templates/*.layout.gohtml")

			if err!= nil {
				return myCache,err
			}
		} 
		//home.page.gohtml || about.page.gohtml
		//&{<nil> 0x140002a0080 0x140002c0000 0x140002b0060 cache 
		myCache[name] = ts
	}

	return myCache,nil
 }
 




















// //Pointing value template
// var tc = make(map[string] *template.Template)

// func RenderTemplate(w http.ResponseWriter,t string){
//      //Data structure 
// 	 var gohtml *template.Template
// 	 var err error

// 	 //check to see if we already the template in the cache
// 	 _,inMap := tc[t]
//      if  !inMap {

//     log.Println("using tempoalte and adding to cache")
//      // need to create template
// 	 err = createTemplateCache(t)
// 	 if err != nil {
// 		log.Println(err)
// 	 }
// 	 }else {
//      // we have the template in the cache
// 	 log.Println("using cached template")
	 
// 	}

// 	gohtml = tc[t]

// 	err = gohtml.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	 }

// }


// func createTemplateCache(t string) error{
//      templates := []string {
// 		fmt.Sprintf("./templates/%s",t),
// 		"./templates/base.layout.gohtml",

// 	 }

// 	gohtml, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}
// 	//add template to cache (map)
// 	tc[t] = gohtml

// 	return nil
// }


