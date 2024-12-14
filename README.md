# ozmap
api inteemedio ozmap en golag

# 1st step: dependencies meeded
go get -u github.com/gin-gonic/gin
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/gin-swagger/swaggerFiles
go install github.com/swaggo/swag/cmd/swag@latest


# 2nd step: init swagger definition
swag init
go generate

