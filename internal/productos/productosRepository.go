package productos

// Acá se importan todas las librerías que vamos a estar usando. Los de github son paquetes personalizados de mi proyecto
import (
	"fmt"
	"log"
	"github.com/jfcheca/StreetFlow/internal/domain"
	"github.com/jfcheca/StreetFlow/pkg/store"
)

//Acá se define una nueva interfaz llamada Repository.
//Una interfaz en Go define un conjunto de métodos que un tipo debe implementar. Si se quiere implementar la interfaz Repository,
//se deben implementar todos los métodos que contiene esa interfaz. 
type Repository interface{

//Se declaran los métodos que cualquier tipo que implemente la interfaz Repository debe definir.
//El método CrearProducto toma un parámetro "producto" de tipo domain.Producto y devuelve un Producto y un valor de tipo error.
CrearProducto(p domain.Producto) (domain.Producto, error)

}
//El propósito de esta estructura repository es encapsular la implementación de la lógica de almacenamiento utilizando
// la interfaz StoreInterfaceProducto. Esto permite a la estructura "repository" depender de una abstracción (la interfaz) en lugar
// de una implementación específica, facilitando la inyección de dependencias y el cambio de implementaciones de almacenamiento
// sin modificar el código del repositorio.

//Define un nuevo tipo llamado repository que es una estructura (struct).
type repository struct {
//Declara un campo en la estructura repository llamado storage.
//Este campo es de tipo store.StoreInterfaceProducto, que es una interfaz definida en tu paquete store.
	storage store.StoreInterfaceProducto
}

// NewRepository crea un nuevo repositorio
//La función toma un parámetro llamado storage de tipo store.StoreInterfaceProducto.
//La función devuelve un valor de tipo Repository, que es una interfaz.

//La función crea una nueva instancia de la estructura repository, inicializando su campo storage con el valor del parámetro storage.
//La expresión &repository{storage: storage} devuelve un puntero a la nueva instancia de repository.
//Esta instancia de repository es devuelta como un valor de tipo Repository porque repository implementa la interfaz Repository.

//El propósito de la función NewRepository es encapsular la lógica de creación de la instancia de repository y asegurarse
// de que la instancia devuelta cumple con la interfaz Repository. Esto permite usar la interfaz Repository para trabajar
//  con el repositorio en lugar de la implementación concreta, promoviendo un diseño desacoplado y más fácil de mantener.

//Ventajas de Usar una Función Constructora
//Encapsulación: La función NewRepository encapsula la lógica de creación de la instancia, asegurando que siempre se inicialice correctamente.
//Inyección de Dependencias: Facilita la inyección de dependencias, permitiendo pasar cualquier implementación de store.StoreInterfaceProducto.
//Desacoplamiento: Promueve el uso de la interfaz Repository en lugar de la implementación concreta, haciendo que el código sea más flexible y fácil de cambiar o probar.
func NewRepository(storage store.StoreInterfaceProducto) Repository {
    return &repository{storage: storage}
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR PRODUCTO >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

//Esto define un método con un receptor de tipo *repository. El receptor r es un puntero a la estructura repository,
// lo que significa que el método puede modificar el receptor y acceder a sus campos.

//Este método se llama CrearProducto y toma un parámetro p de tipo domain.Producto.
//Devuelve dos valores: un domain.Producto y un error.

func (r *repository) CrearProducto(p domain.Producto) (domain.Producto, error) {
//Llama al método CrearProducto del campo storage, que es de tipo store.StoreInterfaceProducto.
//Si CrearProducto de storage devuelve un error, err contendrá ese error.	

//En resumen crea una variable "err" y la iguala a "r" que es el receptor de la funcion, storage que es el campo pasado como parámetro en
//la función NewRepository y llama al metodo CrearProducto de la interfaz store.StoreInterfaceProducto y le pasa el parametro p de la función
	err := r.storage.CrearProducto(p)

//Si err no es nil, lo que indica que hubo un error al intentar crear el producto, se ejecuta el bloque de manejo de errores.
//log.Printf registra un mensaje de error, incluyendo detalles del producto p y el error err.
//Se devuelve un valor vacío de domain.Producto y un error envuelto con fmt.Errorf. El %w en fmt.Errorf permite que el error original 
//se incluya en el error devuelto, lo que facilita el desempaquetado y la inspección de errores más adelante.	
	if err != nil {
		log.Printf("Error al crear el producto %v: %v\n", p, err)
		return domain.Producto{}, fmt.Errorf("error creando producto: %w", err)
	}
//Si no hubo errores, se devuelve el producto p y un valor nil para el error, indicando que la operación fue exitosa.	
	return p, nil
}

//Propósito del Método
//El método CrearProducto en la estructura repository tiene el propósito de delegar la creación de un producto a la implementación 
//de almacenamiento (storage), manejar cualquier error que ocurra durante este proceso y devolver el producto creado o un error si 
//la operación falla.