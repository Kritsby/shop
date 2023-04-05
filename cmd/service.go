package main

import (
	"log"
	_ "shop/docs"
	"shop/internal/app"
	"shop/internal/config"
)

var cfg config.Config

func init() {
	err := cfg.InitCfg()
	if err != nil {
		panic(err)
	}
}

// @title SHOP API
// @version 1.0
// @description API Server for SHOP
// @BasePath /
func main() {
	err := app.Run(cfg)
	if err != nil {
		log.Println(err)
	}
}
