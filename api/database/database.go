package database

import (
	"github.com/gorilla/mux"
	kibisis "github.com/shaaaanks/go-api/kibisis"
)

type DB struct {
	Driver         kibisis.Database
	GenerateRouter mux.Router
}

func (d *DB) Conn() error {
	return d.Driver.Conn()
}

func (d *DB) Init() error {
	return d.Driver.Init()
}

func (d *DB) Create(item interface{}) error {
	return d.Driver.Create(item)
}

func (d *DB) Update(id string, item interface{}) error {
	return d.Driver.Update(id, item)
}

func (d *DB) Delete(id string) error {
	return d.Driver.Delete(id)
}

func (d *DB) Find(id string) (interface{}, error) {
	return d.Driver.Find(id)
}

func (d *DB) FindAll() ([]interface{}, error) {
	return d.Driver.FindAll()
}
