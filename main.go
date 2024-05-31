package main

import (
	"fmt"
	"log"

	"THN-ex1/api"
)

//	@title			THN-ex1
//	@version		0.1
//	@description	app description
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	https://TODO.com
//	@contact.email	TODO@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	app, err := api.NewApp()
	if err != nil {
		log.Fatalf("app layer error: %v", err)
	}

	router := api.InitRouter(app)
	err = router.Run(app.Port())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server is running on port: ", app.Port())
}
