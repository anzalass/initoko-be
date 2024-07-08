package main

import (
	"fmt"
	"initoko/config"
	"initoko/routes"
	"initoko/utils/database"
	"initoko/utils/midtrans"
	"net/http"

	hAlamat "initoko/module/feature/alamat/handler"
	rAlamat "initoko/module/feature/alamat/repository"
	sAlamat "initoko/module/feature/alamat/service"
	hCheckout "initoko/module/feature/checkout/handler"
	rCheckout "initoko/module/feature/checkout/repository"
	sCheckout "initoko/module/feature/checkout/service"
	hJumbotron "initoko/module/feature/jumbotron/handler"
	rJumbotron "initoko/module/feature/jumbotron/repository"
	sJumbotron "initoko/module/feature/jumbotron/service"
	hKategori "initoko/module/feature/kategori/handler"
	rKategori "initoko/module/feature/kategori/repository"
	sKategori "initoko/module/feature/kategori/service"
	hProduct "initoko/module/feature/product/handler"
	rProduct "initoko/module/feature/product/repository"
	sProduct "initoko/module/feature/product/service"
	hUser "initoko/module/feature/users/handler"
	rUser "initoko/module/feature/users/repository"
	sUser "initoko/module/feature/users/service"
	hWishlist "initoko/module/feature/wishlist/handler"
	rWishlist "initoko/module/feature/wishlist/repository"
	sWishlist "initoko/module/feature/wishlist/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("hello world")
	config.InitConfig()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello gays")
	})

	corsConfig := middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,
	}
	e.Use(middleware.CORSWithConfig(corsConfig))

	var db = database.InitDatabase()
	database.MigrateDatabase(db)

	midtranss := midtrans.NewMidtrans()

	alamatrepo := rAlamat.NewAlamatRepository(db)
	alamatservice := sAlamat.NewAlamatService(alamatrepo)
	alamathandler := hAlamat.NewAlamatHandler(alamatservice)
	routes.RouteAlamat(e, alamathandler)

	jumbotronrepo := rJumbotron.NewJumbotronRepository(db)
	jumbotornservice := sJumbotron.NewJumbotronService(jumbotronrepo)
	jumbotronhandler := hJumbotron.NewJumbotronHandler(jumbotornservice)
	routes.RouteJumbotron(e, jumbotronhandler)

	kategorirepo := rKategori.NewKategoriRepository(db)
	kategoriservice := sKategori.NewKategoriService(kategorirepo)
	kategorihandler := hKategori.NewKategoriHandler(kategoriservice)
	routes.RouteKategori(e, kategorihandler)

	productrepo := rProduct.NewProductRepository(db)
	productservice := sProduct.NewProductService(productrepo, jumbotronrepo, kategorirepo)
	producthandler := hProduct.NewHandlerProduct(productservice)
	routes.RouteProduct(e, producthandler)

	userrepo := rUser.NewUserRepository(db)
	userservice := sUser.NewUsersService(userrepo)
	userhandler := hUser.NewUsersHandler(userservice)
	routes.RouteUser(e, userhandler)

	wishlistuser := rWishlist.NewReposioryWishlist(db)
	wishlistservice := sWishlist.NewWishlistService(wishlistuser)
	wishlisthandler := hWishlist.NewWishlistHandler(wishlistservice)
	routes.RouteWishlist(e, wishlisthandler)

	checkoutrepo := rCheckout.NewCheckoutRepository(db)
	checkoutservice := sCheckout.NewCheckoutService(checkoutrepo, midtranss)
	checkouthandler := hCheckout.NewCheckoutHandler(checkoutservice, midtranss)
	routes.RouteCheckout(e, checkouthandler)

	e.Logger.Fatal(e.Start(":8000"))

}
