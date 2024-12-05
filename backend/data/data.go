package data

import "database/sql"

type PasswordDataOnLogin struct {
	SiteName             string `json:"sitename"`
	Password_e           string `json:"password_e"`    // blank until requested
	LastUpdatedTimeStamp string `json:"lastupdatedts"` //
}

type PasswordRequest struct {
	Username string
	SiteName string
}

type PasswordResponse struct {
	SiteName   string
	Password_e string
}

type PasswordAdd struct {
	Username   string `json:"username"`
	SiteName   string `json:"sitename"`
	Password_e string `json:"password_e"`
}

type PasswordDel struct {
	Username string
	SiteName string
}

var g_pgsql_conn *sql.DB

func SetConnectionForData(pgsql_conn *sql.DB) {
	g_pgsql_conn = pgsql_conn
}
