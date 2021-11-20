package models

//TemplateData holds datasent from handlers to template
type TemplateData struct {
	StringMap map[string]string
	IntMaop   map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
