# clean_architecture_blog
Attempt to implement ideas of [the Clean Architecure](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) in a simple project

## Setup and run

Control host/port via setting env vars in the `.env` file.

```bash
docker-compose up
```

## Used Stack

* Go
  * [Gin](https://pkg.go.dev/github.com/gin-gonic/gin)
  * [Gorm](https://pkg.go.dev/gorm.io/gorm)
* Typescript
  * [React](https://yarnpkg.com/package/react)

Client side implemented as [SPA](https://en.wikipedia.org/wiki/Single-page_application) website, and the backend gives you the bundler on request to any path except the `api/` one!
