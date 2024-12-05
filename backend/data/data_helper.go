package data

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Datacrud struct {
	SiteName             string
	Password_e           string
	LastUpdatedTimeStamp string
	Username             string
}

func (d *Datacrud) AddPassword() (uint, error) {
	exists, err := d.siteExists()
	if exists {
		log.Printf("ADDPASS: site %s already exists for user %s", d.SiteName, d.Username)
		return http.StatusBadRequest, errors.New("password for site already exists")
	}
	if err != nil {
		log.Printf("ADDPASS: error in checking if site %s exists in database: %s", d.SiteName, err.Error())
		return http.StatusInternalServerError, errors.New("internal server error, try again later")
	}

	pgsql_query := fmt.Sprintf("INSERT INTO passwords (username, sitename, encrypted_password) VALUES ('%s', '%s', '%s')",
		d.Username,
		d.SiteName,
		d.Password_e,
	)
	_, err = g_pgsql_conn.Exec(pgsql_query)
	if err != nil {
		log.Printf("ADDPASS: user %s adding password for site %s failed: %s", d.Username, d.SiteName, err.Error())
		return http.StatusInternalServerError, errors.New("internal server error, try again later")
	}

	return http.StatusOK, nil
}

func (d *Datacrud) DelPassword() (uint, error) {
	exists, err := d.siteExists()
	if !exists {
		log.Printf("DELPASS: site %s does not exists for user %s", d.SiteName, d.Username)
		return http.StatusBadRequest, errors.New("password for requested site does not exists")
	}
	if err != nil {
		log.Printf("DELPASS: error in checking if site %s exists in database: %s", d.SiteName, err.Error())
		return http.StatusInternalServerError, errors.New("internal server error, try again later")
	}

	pgsql_query := fmt.Sprintf("DELETE FROM passwords WHERE username='%s' AND sitename='%s'",
		d.Username,
		d.SiteName,
	)
	_, err = g_pgsql_conn.Exec(pgsql_query)
	if err != nil {
		log.Printf("DELPASS: user %s deleting password for site %s failed: %s", d.Username, d.SiteName, err.Error())
		return http.StatusInternalServerError, errors.New("internal server error, try again later")
	}

	return http.StatusOK, nil
}

func (d *Datacrud) siteExists() (bool, error) {
	var exists bool
	pgsql_query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM passwords WHERE username='%s' AND sitename='%s')", d.Username, d.SiteName)
	err := g_pgsql_conn.QueryRow(pgsql_query).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
