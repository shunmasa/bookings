
package models


//send data to tempalte,from handers to template
type TemplateData struct{
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	Data  map[string]interface{}
	//form
	CSRFToken string
	Flash string
	Warning string
	Error string
  }
  