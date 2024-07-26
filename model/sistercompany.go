package model

import "github.com/google/uuid"

type sistercompany struct {
	sistercompanyid      uuid.UUID
	parentid             uuid.UUID
	sistercompanyname    string
	sistercompanyinitial string
	description          string
	createdon            string
	createdby            string
	modifiedon           string
	modifiedby           string
	status               int
}
