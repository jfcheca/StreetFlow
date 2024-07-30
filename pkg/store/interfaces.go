package store

import (
	"github.com/jfcheca/StreetFlow/internal/domain"
)
//*************************************************************PRODUCTOS******************************************************
type StoreInterfaceProducto interface {

	CrearProducto(producto domain.Producto) error //comienza con una letra mayúscula, lo que indica que la función es exportada y,
//	 por lo tanto, accesible desde otros paquetes. "producto" es el parametro que va a recibir la funcion, "domain.Producto" es el tipo de dato
//  de este parámetro. Producto es una estructura (struct) definida en el paquete domain. "error" es el tipo de dato que la función devuelve


//	BuscarProducto(id int) (domain.Producto, error)
//	BuscarTodosLosProductos() ([]domain.Producto, error)
//	UpdateProducto(id int, p domain.Producto) error
//	Patch(id int, updatedFields map[string]interface{}) (domain.Producto, error)
//	DeleteProducto(id int) error
//	ExistsByID(id int) bool
//	ObtenerNombreCategoria(id int) (string, error) // Añadir este método
}