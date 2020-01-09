package main

import (
	"html/template"
	"net/http"

	"github.com/abdimussa87/Intern-Seek-Version-1/delivery/http/handler"
	"github.com/abdimussa87/Intern-Seek-Version-1/user/repository"
	userRep "github.com/abdimussa87/Intern-Seek-Version-1/user/repository"
	"github.com/abdimussa87/Intern-Seek-Version-1/user/service"
	userServ "github.com/abdimussa87/Intern-Seek-Version-1/user/service"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"

	_ "github.com/lib/pq"
)

var tmpl = template.Must(template.ParseGlob("../../ui/templates/*"))

func indexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "index.html", nil)

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "login.html", nil)
}

func main() {

	dbconn, err := gorm.Open("postgres", "user=postgres dbname=gorminterndb password='P@$$wOrDd' sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()
	// dbconn.DropTableIfExists(&entity.CompanyDetail{}, &entity.User{})
	// errs := dbconn.CreateTable(&entity.User{}, &entity.CompanyDetail{}).GetErrors()

	// if len(errs) > 0 {
	// 	panic(errs)
	// }

	userRepo := userRep.NewUserGormRepoImpl(dbconn)
	userServi := userServ.NewUserServiceImpl(userRepo)

	compRepo := repository.NewCompanyGormRepoImpl(dbconn)
	compServ := service.NewCompanyService(compRepo)

	//userHandler := handler.NewUserHandler(userServi)

	compHandler := handler.NewCompanyHandler(compServ, userServi)

	router := httprouter.New()

	router.GET("/v1/company", compHandler.GetCompanies)
	router.GET("/v1/company/:id", compHandler.GetSingleCompany)
	router.POST("/v1/company/:id", compHandler.PostCompany)
	router.PUT("/v1/company/update/:id", compHandler.PutCompany)
	router.DELETE("/v1/company/delete/:id", compHandler.DeleteCompany)

	http.ListenAndServe(":8181", router)

	// mux := http.NewServeMux()
	// fs := http.FileServer(http.Dir("../../ui/assets"))
	// mux.Handle("/assets/", http.StripPrefix("/assets", fs))
	// mux.HandleFunc("/", indexHandler)
	// mux.HandleFunc("/login", loginHandler)
	// mux.HandleFunc("/signup", userHandler.SignUp)
	// http.ListenAndServe(":8080", mux)
}
