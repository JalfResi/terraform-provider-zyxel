package main

import "github.com/hashicorp/terraform/helper/schema"

type MyClient struct {

}

func (c *MyClient) Get(id string) (Record, bool)  {
	return Record{}, true
}


type Record struct {
	Address string
}

func updateAddress(d *schema.ResourceData, m interface{}) error {
	return nil
}