package routes

import (
	"initoko/module/feature/alamat"
	"initoko/module/feature/checkout"
	"initoko/module/feature/jumbotron"
	"initoko/module/feature/kategori"
	"initoko/module/feature/product"
	"initoko/module/feature/users"
	"initoko/module/feature/wishlist"

	"github.com/labstack/echo/v4"
)

func RouteAlamat(e *echo.Echo, h alamat.HandlerALamatInterface) {
	e.POST("/alamat/create", h.CreateAlamat())
	e.DELETE("/alamat/delete/:id", h.DeleteAlamatById())
	e.GET("/alamat/:id", h.GetAlamatById())
	e.PUT("/alamat/update/:id", h.UpdateAlamat())
}

func RouteJumbotron(e *echo.Echo, h jumbotron.JumbotronHandlerInterface) {
	e.POST("/jumbotron/create", h.CreateJumbotron())
	e.DELETE("/jumbotron/delete/:id", h.DeleteJumbotron())
	e.GET("/jumbotron/:id", h.GetJumbotronById())
	e.PUT("/jumbotron/update/:id", h.UpdateJumbotron())
	e.GET("/jumbotron/all", h.GetAllJumbotron())
}
func RouteKategori(e *echo.Echo, h kategori.KategoriHandlerInterface) {
	e.POST("/kategori/create", h.CreateKategori()) // done
	e.DELETE("/kategori/delete/:id", h.DeleteKategori())
	e.GET("/kategori/:id", h.GetKategoriByID())       // done
	e.PUT("/kategori/update/:id", h.UpdateKategori()) // done
	e.GET("/kategori/all", h.GetAllKategori())        // done
	e.GET("/kategori/name", h.GetAllKategoriName())   // done
}
func RouteProduct(e *echo.Echo, h product.ProductHandlerInterface) {
	e.POST("/product/create", h.CreateProduct())       //done
	e.DELETE("/product/delete/:id", h.DeleteProduct()) //done
	e.GET("/product/:id", h.GetProductByID())
	e.PUT("/product/update/:id", h.UpdateProduct()) // done
	e.GET("/product/all", h.GetAllProductByFilter())
	e.POST("/product/foto", h.UploadFotoProduct())
	e.GET("/product", h.GetAllProduct()) //done
	e.GET("/product/related/:kategori", h.GetProductByKategori())
	e.GET("/homepage", h.Homepage())
}
func RouteUser(e *echo.Echo, h users.HandlerUserInterface) {
	e.POST("/user/register", h.RegisterUser())
	e.POST("/user/registergoogle", h.RegisterUserGoogle())
	e.DELETE("/user/delete/:id", h.DeleteUser())
	e.POST("/user/login", h.LoginUser())
	e.PUT("/user/update/:id", h.UpdateUser())
	e.PUT("/user/avatar/:id", h.UpdateAvatar())
	e.PUT("/user/aktivasi/:kode/:email", h.AktivasiAkun())
	e.POST("/user/reqotp", h.SendOtp())
	e.GET("/user/profile/:email", h.ProfilePage())
	e.GET("/user/alamat/:email", h.AlamatPage())
}

func RouteWishlist(e *echo.Echo, h wishlist.HandlerWishlistInterface) {
	e.POST("/wishlist/create", h.AddWishlist())
	e.DELETE("/wishlist/delete/:id", h.DeleteWishListById())
	e.GET("/wishlist/:id", h.GetWishlistByIdUser())
}

func RouteCheckout(e *echo.Echo, h checkout.HandelerCheckout) {
	e.POST("/checkout/create", h.CreateTransaksi())
	e.PUT("/checkout/status/:id", h.ChangeStatus())
	e.POST("/checkout/status-pembayaran", h.ChangeStatusPembayaran())
	e.GET("/pesanan/:email", h.AllPesananByEmail())
	e.GET("/pesanan", h.AllPesanan())
	e.GET("/pesanan/detail/:id", h.DetailPesananById())
	e.PUT("/pesanan/inputresi/:id", h.InputResi())

}
