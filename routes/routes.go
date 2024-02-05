package routes

import (
	"net/http"

	"github.com/product-manager/controllers"
)

func LoadRoutes() {
	//"/" -> o caminho e index é a função
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
}
