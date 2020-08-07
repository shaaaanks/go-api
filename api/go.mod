module github.com/shaaaanks/go-api/api

go 1.14

replace github.com/shaaaanks/go-api/kibisis => ./../kibisis

require (
	github.com/gorilla/mux v1.7.4
	github.com/rs/zerolog v1.19.0 // indirect
	github.com/shaaaanks/go-api/kibisis v0.0.0-20200803232955-1596882d891e
	k8s.io/apimachinery v0.18.6 // indirect
)
