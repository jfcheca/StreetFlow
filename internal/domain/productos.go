package domain

type Producto struct{
	ID                 int             `json:"id"`
	Nombre             string          `json:"nombre"`
	Talle			   string		   `json:"talle"`
	Descripcion        string          `json:"descripcion"`
	Precio             float64         `json:"precio"`
	Id_categoria	   string	       `json:"id_categoria"`
	Imagenes           []Imagen        `json:"imagenes"`


}