Basic GIN Project
===============

Provides essentials that most web applications need - MVC pattern, user authorisation, SQL db migration, admin dashboard, javascript form validation, and can be easily extended.

It consists of the following core components:

- GIN - A web microframework (with best performance atm) for Golang - https://github.com/gin-gonic/gin
- GIN middlewares [gin-csrf](https://github.com/utrack/gin-csrf), [gin/contrib/sessions](https://github.com/gin-gonic/contrib/tree/master/sessions)
- pq - Postgres driver for the Go language - https://github.com/lib/pq
- sqlx - Relational database access interface - https://github.com/jmoiron/sqlx
- sql-migrate - SQL schema migration tool - https://github.com/rubenv/sql-migrate
- go.rice - Asset packaging tool for easy deployment - https://github.com/GeertJohan/go.rice
- logrus - advanced Go logger - https://github.com/Sirupsen/logrus
- Twitter Bootstrap - popular HTML, CSS, JS framework for developing responsive, mobile first web projects - http://getbootstrap.com
- Parsley JS - form validation - http://parsleyjs.org
- Bootstrap Markdown Editor with image upload - https://github.com/inacho/bootstrap-markdown-editor
- blackfriday - markdown processor - https://github.com/russross/blackfriday 

# TODO
Server-side model validation with https://godoc.org/gopkg.in/bluesuncorp/validator.v5

i18n (probably)

Gorbac role-based access control (may be) - https://github.com/mikespook/gorbac

# Screenshots
## Home page
![](/public/images/screenshot_home.jpg)
## Dashboard
![](/public/images/screenshot_dashboard.jpg)
## Markdown editor
![](/public/images/screenshot_markdown.jpg)
## Fancy 404, 405, 500 error pages
![](/public/images/screenshot_error.jpg)

# Usage
```
git clone https://github.com/denisbakhtin/ginbasic.git
cd ginbasic
go get .
```
Copy sample config `cp config/config.json.example config/config.json`, create postgresql database, modify config/config.json accordingly.

Type `go run main.go -migrate=up` to create users and pages tables.

`go run main.go` to launch web server.

# Deployment
```
go generate && go build && rm ./*.rice-box.go
```
Upload `ginbasic` binary to your server. If you find `rice embed-go` is running slow on your system, consider using other [go.rice packing options](https://github.com/GeertJohan/go.rice#tool-usage) with `go generate` command.

# Project structure

`/config`

Contains application configuration file.

`/controllers`

All your controllers that serve defined routes.

`/helpers`

Helper functions.

`/migrations`

Database schema migrations

`/models`

You database models.

`/public`

It has all your static files

`/system`

Core functions and structs.

`/views`

Your views using standard `Go` template system.

`main.go`

This file starts your web application, contains routes definition & some custom middlewares.

# Make it your own

I assume you have followed installation instructions and you have `ginbasic` installed in your `GOPATH` location.

Let's say I want to create `Amazing Website`. I create new `GitHub` repository `https://github.com/denisbakhtin/amazingwebsite` (of course replace that with your own repository).

Now I have to prepare `ginbasic`. First thing is that I have to delete its `.git` directory.

I issue:

```
rm -rf src/github.com/denisbakhtin/ginbasic/.git
```

Then I want to replace all references from `github.com/denisbakhtin/ginbasic` to `github.com/denisbakhtin/amazingwebsite`:

```
grep -rl 'github.com/denisbakhtin/ginbasic' ./ | xargs sed -i 's/github.com\/denisbakhtin\/ginbasic/github.com\/denisbakhtin\/amazingwebsite/g'
```

Now I have to move all `ginbasic` files to the new location:

```
mv src/github.com/denisbakhtin/ginbasic/ src/github.com/denisbakhtin/amazingwebsite
```

And push it to my new repository at `GitHub`:

```
cd src/github.com/denisbakhtin/amazingwebsite
git init
git add --all .
git commit -m "Amazing Website First Commit"
git remote add origin https://github.com/denisbakhtin/amazingwebsite.git
git push -u origin master
```

You can now go back to your `GOPATH` and check if everything is ok:

```
go install github.com/denisbakhtin/amazingwebsite
```

And that's it. 

# Continuous Development

For Continuous Development I recommend using `Reflex` - https://github.com/cespare/reflex

You can install `Reflex` by issuing:

```
go get github.com/cespare/reflex
```

Then create a config file `reflex.conf` in your `GOPATH`:

```
# Restart server when .go, .html files change
-sr '(\.go|\.html)$' go run main.go
```

Now if you run:

```
reflex -c reflex.conf
```

Project should automatically rebuild itself when a change occurs. For more options read https://github.com/cespare/reflex#usage

