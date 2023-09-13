package main

import (
	"fmt"
	"log"

	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
)

func main() {
	var option string
	var secretkey string
	var appname string
	var accountname string

	fmt.Println("Seleccione una opcion:")
	fmt.Println("1 - Generar secret key")
	fmt.Println("2 - Validar un OTP")

	fmt.Scanln(&option)

	switch option {
	case "1":
		fmt.Println("Ingrese el nombre de la aplicacion:")
		fmt.Scanln(&appname)

		fmt.Println("Ingrese el nombre de la cuenta:")
		fmt.Scanln(&accountname)

		// Crear una nueva clave secreta OTP
		key, err := totp.Generate(totp.GenerateOpts{
			Issuer:      appname,
			AccountName: accountname,
		})
		if err != nil {
			log.Fatal("Error al generar la clave secreta OTP:", err)
		}

		// Generar y mostrar el código QR
		err = generateQRCode(key.URL())
		if err != nil {
			log.Fatal("Error al generar el código QR:", err)
		}
		fmt.Println("Escanee el QR utilizando Google Authenticator")
		secretkey = key.Secret()
		fmt.Println("Anote la siguiente secret key generada:")
		fmt.Println(secretkey)

		return

	case "2":
		fmt.Println("Ingrese la secret key")
		//TODO Guardar la entrada de texto
		//en la variable secretkey
		//utilizando la funcion Scanln()

		// Validar un OTP
		fmt.Print("Ingrese el OTP que generó Google Authenticator: ")
		var inputOTP string
		//TODO Guardar la entrada de texto
		//en la variable inputOTP
		//utilizando la funcion Scanln()

		valid := totp.Validate(inputOTP, secretkey)

		//TODO Si valid es verdadero entonces
		//imprimir "Acceso permitodo", en caso contrario
		//imprimir "Acceso denegado"
	}

}

func generateQRCode(text string) error {
	qrcode, err := qrcode.New(text, qrcode.Medium)
	if err != nil {
		return err
	}

	return qrcode.WriteFile(256, "qrcode.png")
}
