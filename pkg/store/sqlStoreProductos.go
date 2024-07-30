package store

// Acá se importan todas las librerías que vamos a estar usando. Los de github son paquetes personalizados de mi proyecto
import (
	"database/sql"
	"fmt"
	"github.com/jfcheca/StreetFlow/internal/domain"

)

//Define una estructura que representa la interacción con una base de datos SQL
//Define una nueva estructura llamada sqlStoreProductos.
//La estructura contiene un único campo db, que es un puntero a una instancia de sql.DB.

//La estructura sqlStoreProductos está diseñada para interactuar con una base de datos SQL. El campo db es un puntero 
//a sql.DB, que es una estructura proporcionada por el paquete database/sql para manejar conexiones a bases de datos SQL.
type sqlStoreProductos struct {
	db *sql.DB
}

//NewSqlStoreProductos es una función constructora para crear una instancia de sqlStoreProductos y devolverla como 
//una implementación de la interfaz StoreInterfaceProducto

//db *sql.DB es un parámetro de la función, que es un puntero a una instancia de sql.DB. sql.DB es la estructura 
//del paquete database/sql que representa la conexión a una base de datos SQL.
//StoreInterfaceProducto es el tipo de valor que la función devolverá. Es una interfaz que debe ser implementada por sqlStoreProductos.
func NewSqlStoreProductos(db *sql.DB) StoreInterfaceProducto {
//La función retorna una nueva instancia de sqlStoreProductos.
//&sqlStoreProductos{ db:db }crea una nueva instancia de la estructura sqlStoreProductos, inicializando su campo db con el valor del parámetro db.
//El operador & obtiene la dirección de la instancia recién creada, lo que da un puntero a sqlStoreProductos	
	return &sqlStoreProductos{
		db: db,
	}
}
//Propósito de la Función
//El propósito de NewSqlStoreProductos es crear e inicializar una nueva instancia de sqlStoreProductos y devolverla como un tipo de 
//StoreInterfaceProducto. Esto permite que el código que usa la función no tenga que preocuparse por la implementación concreta de 
//sqlStoreProductos y solo interactúe con la interfaz StoreInterfaceProducto


// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR UN NUEVO PRODUCTO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

//(s *sqlStoreProductos) indica que la función es un método de la estructura sqlStoreProductos. s es el receptor del método, 
//que es un puntero a sqlStoreProductos.
//CrearProducto(producto domain.Producto) error es el nombre del método y sus parámetros. El método toma un parámetro producto 
//de tipo domain.Producto y devuelve un valor de tipo error.
func (s *sqlStoreProductos) CrearProducto(producto domain.Producto) error {

//query es una cadena de texto que contiene la consulta SQL para insertar un nuevo producto en la tabla productos	
//Los signos de interrogación (?) son marcadores de posición para los valores que se proporcionarán más adelante. Esto previene inyecciones 
//de SQL y permite que el controlador de base de datos maneje el formato adecuado de los datos.
	query := "INSERT INTO productos (id_categoria, nombre, talle, descripcion, precio) VALUES (?, ?, ?, ?, ?);"

//s.db.Prepare(query) prepara la consulta SQL para su ejecución. Esto crea una declaración (stmt) que puede ser ejecutada 
//posteriormente con los valores proporcionados.	
	stmt, err := s.db.Prepare(query)
//Si hay un error al preparar la consulta, se devuelve un error formateado con fmt.Errorf.
	if err != nil {
		return fmt.Errorf("error preparing query: %w", err)
	}
//defer stmt.Close() asegura que la declaración preparada se cierre cuando la función termine, liberando los recursos asociados	
	defer stmt.Close()

//stmt.Exec(...) ejecuta la declaración preparada con los valores proporcionados para los marcadores de posición (?) en la consulta SQL.
//Los valores del producto (producto.Id_categoria, producto.Nombre, producto.Talle, producto.Descripcion, producto.Precio) se pasan a la consulta.	
	_, err = stmt.Exec(producto.Id_categoria, producto.Nombre, producto.Talle, producto.Descripcion, producto.Precio)

//Si hay un error al ejecutar la consulta, se devuelve un error formateado con fmt.Errorf.	
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}
//Si la ejecución de la consulta es exitosa y no hay errores, se retorna nil, indicando que no hubo errores.
	return nil
}