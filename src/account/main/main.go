package main
import (
	"account/utils"
	"fmt"
)
func main() {
	fmt.Println("面向对象的方式来完成.....")
	utils.NewMyFamilyAccount().MainMenu()
}