package interfaces

type Persona struct{
	Id string `json:"id"`
	Rut string `json:"rut"`
	Nombre string `json:"nombre"`
	Apellido string `json:"apellido"`
}