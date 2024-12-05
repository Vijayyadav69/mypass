package auth

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"passm/data"

	_ "github.com/lib/pq"
)

type loginRequest struct {
	Username   string `json:"username"`
	Password_d string `json:"password_d"`
	Email_id   string `json:"email_id"`
}

var g_pgsql_conn *sql.DB

func SetConnectionForAuth(pgsql_conn *sql.DB) {
	g_pgsql_conn = pgsql_conn
}

func Register(w http.ResponseWriter, r *http.Request) {
	var lr loginRequest
	err := json.NewDecoder(r.Body).Decode(&lr)
	if err != nil {
		log.Printf("REGISTER: error in /register received data %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid data"))
		return
	}

	log.Printf("REGISTER: register request received for user %s, emailid %s", lr.Username, lr.Email_id)

	if lr.Username == "" || lr.Email_id == "" || lr.Password_d == "" {
		log.Println("REGISTER: blank field received")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("all fields are required"))
		return
	}

	ul := authlogin{Username: lr.Username, Password: lr.Password_d, Emailid: lr.Email_id}
	if code, err := ul.register(); err != nil {
		log.Printf("REGISTER: register failed for %s", ul.Username)
		w.WriteHeader(int(code))
		w.Write([]byte(err.Error()))
		return
	}

	log.Printf("REGISTER: register request processed for user %s, emailid %s", lr.Username, lr.Email_id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user created"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	var lr loginRequest
	err := json.NewDecoder(r.Body).Decode(&lr)
	if err != nil {
		log.Printf("LOGIN: error in /login received data %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid data"))
		return
	}

	log.Printf("LOGIN: login request received for user %s, emailid %s", lr.Username, lr.Email_id)

	if lr.Username == "" || lr.Password_d == "" {
		log.Println("LOGIN: blank field received")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("both fields username and password are required"))
		return
	}

	ul := authlogin{Username: lr.Username, Password: lr.Password_d}
	if code, err := ul.verifyLogin(); err != nil {
		log.Printf("LOGIN: verify login failed for user %s", ul.Username)
		w.WriteHeader(int(code))
		w.Write([]byte(err.Error()))
		return
	}

	code, data_arr, err := ul.fetchPasswordList()
	if err != nil {
		log.Printf("LOGIN: fetch password list failed for user %s", ul.Username)
		w.WriteHeader(int(code))
		w.Write([]byte(err.Error()))
		return
	}

	log.Printf("LOGIN: login request processed for user %s", lr.Username)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data_arr)
}

func AddPass(w http.ResponseWriter, r *http.Request) {
	var pa data.PasswordAdd
	err := json.NewDecoder(r.Body).Decode(&pa)
	if err != nil {
		log.Printf("ADDPASS: error in /addpass received data %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid data"))
		return
	}

	log.Printf("ADDPASS: add pass request received for user %s, site %s", pa.Username, pa.SiteName)

	if pa.Username == "" || pa.SiteName == "" || pa.Password_e == "" {
		log.Println("ADDPASS: blank field received")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("all fields are required"))
		return
	}

	dc := data.Datacrud{Username: pa.Username, SiteName: pa.SiteName, Password_e: pa.Password_e}
	code, err := dc.AddPassword()
	if err != nil {
		log.Println("ADDPASS: addpassword failed for user", dc.Username)
		w.WriteHeader(int(code))
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("password added"))
	log.Printf("ADDPASS: user %s, addpass processed for site %s", dc.Username, dc.SiteName)
}

func DelPass(w http.ResponseWriter, r *http.Request) {
	var pd data.PasswordDel
	err := json.NewDecoder(r.Body).Decode(&pd)
	if err != nil {
		log.Printf("DELPASS: error in /delpass received data %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid data"))
		return
	}

	log.Printf("DELPASS: del pass request received for user %s, site %s", pd.Username, pd.SiteName)

	if pd.Username == "" || pd.SiteName == "" {
		log.Println("DELPASS: blank field received")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("all fields are required"))
		return
	}

	dc := data.Datacrud{Username: pd.Username, SiteName: pd.SiteName}
	code, err := dc.DelPassword()
	if err != nil {
		log.Println("DELPASS: del password failed for user", dc.Username)
		w.WriteHeader(int(code))
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("password deleted"))
	log.Printf("DELPASS: user %s, delpass processed for site %s", dc.Username, dc.SiteName)
}
