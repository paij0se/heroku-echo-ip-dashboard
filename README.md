#heroku-echo-ip-dashboard

<h1>Steps</h1>

- Create the heroku database

```sh
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
        HerokuEchoIpDashboard "github.com/paij0se/ip/src"
        // .....
    )
   
    func main(){
        //...
        // Your echo app
        //...
        port, ok := os.LookupEnv("PORT")

        if !ok {
            port = "5000"
        }
        HerokuEchoIpDashboard.HerokuEchoIpDashboard(port) // <---- Put the port here
        fmt.Printf("server on port: %s", port)
        e.Logger.Fatal(e.Start(":" + port)) 
    }
    

```

<h1>Test everything with `go test`</h1>
