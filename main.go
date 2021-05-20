package main

import (
	"os"
	xyzgw "q-change-assignment/src/xyz/gateways"
	xyzsv "q-change-assignment/src/xyz/services"

	cgw "q-change-assignment/src/cashier/gateways"
	csv "q-change-assignment/src/cashier/services"

	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	defer e.Close()

	// ASSIGNMENT #1
	sv0 := xyzsv.NewXYZService()
	g0 := e.Group("xyz")
	xyzgw.NewHTTPGateway(g0, sv0)

	// // ASSIGNMENT #2
	sv1 := csv.NewCashierService()
	g1 := e.Group("cashier")
	cgw.NewHTTPGateway(g1, sv1)

	// entry point
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
