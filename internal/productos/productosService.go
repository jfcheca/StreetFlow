package productos

import (

// Acá se importan todas las librerías que vamos a estar usando. Los de github son paquetes personalizados de mi proyecto
	"github.com/jfcheca/StreetFlow/internal/domain"
)


//Acá se define una nueva interfaz llamada Service.
//Una interfaz en Go define un conjunto de métodos que un tipo debe implementar. Si se quiere implementar la interfaz Repository,
//se deben implementar todos los métodos que contiene esa interfaz. 
type Service interface {

//Se declaran los métodos que cualquier tipo que implemente la interfaz Repository debe definir.
//El método CrearProducto toma un parámetro "p" de tipo domain.Producto y devuelve un valor de tipo error.
	CrearProducto(p domain.Producto) (domain.Producto, error)

}


//Define una estructura llamada service que contiene un campo r de tipo Repository
type service struct {
	r Repository
//Declara un campo en la estructura service llamado r.
//Este campo es de tipo Repository, que es una interfaz.

//Propósito
//El propósito de esta estructura service es encapsular una dependencia del tipo Repository. 
//Esto permite que la lógica del servicio interactúe con la capa de repositorio a través de la interfaz Repository, 
//promoviendo un diseño desacoplado y facilitando la inyección de dependencias.
}

//NewService es una función constructora para crear una nueva instancia del servicio que implementa la interfaz Service
//La función toma un parámetro r de tipo Repository, que es una interfaz.
//La función devuelve un valor de tipo Service, que también es una interfaz.
func NewService(r Repository) Service {

//La función crea una nueva instancia de la estructura service, inicializando su campo r con el valor del parámetro r.
//La expresión &service{r} devuelve un puntero a la nueva instancia de service.
//Esta instancia de service es devuelta como un valor de tipo Service porque service implementa la interfaz Service	
	return &service{r}
}
//Propósito
//El propósito de la función NewService es encapsular la lógica de creación de la instancia de service y 
//asegurarse de que la instancia devuelta cumple con la interfaz Service. Esto permite usar la interfaz Service para trabajar 
//con el servicio en lugar de la implementación concreta, promoviendo un diseño desacoplado y facilitando la inyección de dependencias.

//Ventajas
//Desacoplamiento: Al depender de una interfaz en lugar de una implementación concreta, el código del servicio es más flexible y fácil de cambiar.
//Inyección de Dependencias: Facilita la inyección de dependencias, permitiendo pasar cualquier implementación de Repository.
//Modularidad: Promueve la separación de responsabilidades y modularidad en el código, haciendo que sea más fácil de mantener y probar.


// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR UN NUEVO PRODUCTO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

// CrearProducto crea un nuevo producto utilizando el repositorio y devuelve el producto creado
func (s *service) CrearProducto(p domain.Producto) (domain.Producto, error) {
//(s *service) indica que esta función es un método que pertenece a la estructura service. s es un receptor de método 
//que es un puntero a la estructura service.
//CrearProducto(p domain.Producto) es el nombre del método y sus parámetros. El método toma un parámetro p de tipo domain.Producto.
//(domain.Producto, error) indica que el método devuelve dos valores: un domain.Producto y un error.

//El método llama al método CrearProducto de la interfaz Repository (accesible a través del campo r de la estructura service).
//Pasa el producto p al método del repositorio.
//El resultado de esta llamada se asigna a dos variables: productoCreado y err.

////Llamada a un Método de una Interfaz:
//s.r accede al campo r de la estructura service. Este campo es de tipo Repository, que es una interfaz.   
//llama al método CrearProducto definido por la interfaz Repository. 
//productoCreado, err := es una asignación múltiple en Go.
	productoCreado, err := s.r.CrearProducto(p)
	
//La llamada s.r.CrearProducto(p) devuelve dos valores:
//productoCreado: El producto creado (de tipo domain.Producto).
//err: Un error (de tipo error), que será nil si la operación fue exitosa, o contendrá información sobre el error si ocurrió un problema.	
    if err != nil {
//Si la llamada al método del repositorio devuelve un error (err no es nil), se retorna un domain.Producto vacío y el error.
        return domain.Producto{}, err
		
    }
    return productoCreado, nil
}