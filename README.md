<img src="https://media.discordapp.net/attachments/950041049458438164/978468065501138964/Screenshot_from_2022-05-23_20-22-01.png?width=608&height=402"/>

<h1>Steps</h1>

- Create the heroku database

```rs
$ heroku addons:create heroku-postgresql:hobby-dev
Creating heroku-postgresql:hobby-dev on ⬢ go-getting-started... free
Created postgresql-curved-22223 as DATABASE_URL
Database has been created and is available
 ! This database is empty. If upgrading, you can transfer
 ! data from another database with pg:copy
Use heroku addons:docs heroku-postgresql to view documentation
```

<h1>Usage</h1>

```go
package main

import (
	"log"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	HerokuEchoIpDashboard "github.com/paij0se/heroku-echo-ip-dashboard/src"
	re "github.com/paij0se/heroku-echo-ip-dashboard/src/controllers"
)

func main() {
	e := echo.New()
	HerokuEchoIpDashboard.HerokuEchoIpDashboard(e) // init the dashboard

	e.GET("/", func(c echo.Context) error {
		re.Requester(c.Scheme() + "://" + c.Request().Host) // This is going to count all the visitors of "/"
		//return c.File("public/index.html") // Static file
        
		//return c.String(http.StatusOK, "Hello, World!") // a hello world
        
	})

	port := os.Getenv("PORT")

	if port == "" {
		log.Println("The port to use is not declared, using port 8080")

		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))

}

```

Note: The static files are serving in "herokudashboard/herokudashboard/public"

<h1>Testing</h1>

- Test everything with Linux/MacOS

```sh
❯ DATABASE_URL=postgres://xxxx go test # Note: You get the postgres url with `heroku config -a app_name`

```

<h1>Routes</h1>

```go
	e.Static("/dashboard", "src/public") // The fronted of dashboard
	e.POST("/ip/update", controllers.UpdateData) // the route where it post the ip
	e.GET("/ip", controllers.GetIp) // the route of where you get the current ip
	e.GET("/ip/all", controllers.ReturnIps) // all the data of the database
```

<h1>TODO</h1>

- [ ] Authentication

- [ ] A better fronted
