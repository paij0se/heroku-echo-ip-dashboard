
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
    import (
        HerokuEchoIpDashboard "github.com/paij0se/heroku-echo-ip-dashboard/src"
        // .....
    )
   
    func main(){
        //...
        // Your echo app
        //...
        e := echo.New()
        HerokuEchoIpDashboard.HerokuEchoIpDashboard(e) // initialize the dashboard with the e type.

        e.GET("/", func(c echo.Context) error {
            go HerokuEchoIpDashboard.NumberOfTimesToCount(c) // Put this on your main route.
            // this is going to count each request.
            // Note: You can used it without a go routine.
            return c.String(http.StatusOK, "Hello, World!")
        })
        port := os.Getenv("PORT")

        if port == "" {
            log.Println("The port to use is not declared, using port 8080")

            port = "8080"
        }
        e.Logger.Fatal(e.Start(":" + port))
    }
    

```

<h1>Testing</h1>

- Test everything with Linux/MacOS

```sh
❯ DATABASE_URL=postgres://xxxx go test # Note: You get the postgres url with `heroku config -a app_name`

```

<h1>TODO</h1>

- [ ] Authentication

- [ ] A better fronted

