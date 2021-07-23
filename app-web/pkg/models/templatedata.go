package models

// Contiene datos enviados desde el controlador a la plantilla
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMao  map[string]float32
	Data      map[string]interface{}
	CRFToken  string
	Flash     string
	Warning   string
	Error     string
}
