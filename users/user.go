package users

import (
	"fmt"
	"time"
	"go/src/models"

)

func AltaUsuario()  {
	u := new(models.User)
	u.AddUser(10,"pablo", time.Now(), true)

	fmt.Println(u)
}