package main

import (
	"fmt"

	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/utils"
)

func main() {
	token, _ := utils.CreateJWT(1)
	fmt.Print(utils.ValidateJWT(token))
}
