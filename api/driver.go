package main

import kibisis "github.com/shaaaanks/go-api"

type driver struct {
	driver kibisis.Database
}

func (d *driver) conn() error {
	return d.driver.Conn()
}

func (d *driver) init() error {
	return d.driver.Init()
}

func (d *driver) create(item interface{}) error {
	return d.driver.Create(item)
}

func (d *driver) update(id string, item interface{}) error {
	return d.driver.Update(id, item)
}

func (d *driver) delete(id string) error {
	return d.driver.Delete(id)
}

func (d *driver) find(id string) (interface{}, error) {
	return d.driver.Find(id)
}

func (d *driver) findAll() ([]interface{}, error) {
	return d.driver.FindAll()
}
