package auth

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"passm/data"

	_ "github.com/lib/pq"
)

// secret key_ "github.com/lib/pq"
var SECRET_KEY = "a1b2c3d4e5f6g7h8"

type authlogin struct {
	Username string
	Password string
	Emailid  string
}

func (a authlogin) verifyLogin() (uint16, error) {

	pgsql_userlogin := a.getInfoFromDB()
	password_d, err := decryptPassword(pgsql_userlogin.Password)
	if err != nil {
		log.Printf("LOGIN: error in decrypting passworf for user %s: %s", a.Username, err.Error())
		return http.StatusInternalServerError, errors.New("internal server error, try again later")
	}
	if a.Password != password_d {
		log.Printf("LOGIN: user:%s, received password does not match with database, incorrect password", a.Username)
		return http.StatusUnauthorized, errors.New("incorrect password")
	}
	return http.StatusOK, nil
}

func (a authlogin) getInfoFromDB() authlogin {
	var pgsql_userlogin authlogin
	pgsql_query := fmt.Sprintf("select username, password_enc from users where username='%s'", a.Username)
	r := g_pgsql_conn.QueryRowContext(context.Background(), pgsql_query)

	r.Scan(&pgsql_userlogin.Username, &pgsql_userlogin.Password)
	return pgsql_userlogin
}

func (a authlogin) register() (uint, error) {
	u_exists, e_exists, err := a.userExists()
	if err != nil {
		log.Printf("REGISTER: error in checking if user %s already exists: %s", a.Username, err.Error())
		return http.StatusInternalServerError, errors.New("internal server error, try again later")
	}
	if u_exists {
		log.Printf("REGISTER: user %s already exists, register failed", a.Username)
		return http.StatusBadRequest, errors.New("user already exists")
	} else if e_exists {
		log.Printf("REGISTER: email id %s already used, register failed", a.Username)
		return http.StatusBadRequest, errors.New("email id is already used")
	} else {
		password_e := encryptPassword(a.Password)
		pgsql_query := fmt.Sprintf("INSERT INTO users (username, email, password_enc) VALUES ('%s', '%s', '%s')",
			a.Username,
			a.Emailid,
			password_e,
		)

		_, err = g_pgsql_conn.Exec(pgsql_query)
		if err != nil {
			log.Printf("REGISTER: error in inserting new user %s data in database: %s", a.Username, err.Error())
			return http.StatusInternalServerError, errors.New("internal server error, try again later")
		}
		return http.StatusOK, nil
	}
}

func (a authlogin) userExists() (bool, bool, error) {
	var u_exists, e_exists bool

	pgsql_query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM users WHERE username='%s')", a.Username)
	err := g_pgsql_conn.QueryRow(pgsql_query).Scan(&u_exists)
	if err != nil {
		return false, false, err
	}

	pgsql_query = fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM users WHERE email='%s')", a.Emailid)
	err = g_pgsql_conn.QueryRow(pgsql_query).Scan(&e_exists)
	if err != nil {
		return false, false, err
	}
	return u_exists, e_exists, nil
}

func (a authlogin) fetchPasswordList() (uint, []data.PasswordDataOnLogin, error) {
	pgsql_query := fmt.Sprintf("SELECT sitename, encrypted_password FROM passwords where username='%s'", a.Username)
	r, err := g_pgsql_conn.Query(pgsql_query)
	if err != nil {
		log.Printf("LOGIN: user %s fetching password list failed %s", a.Username, err.Error())
		return http.StatusInternalServerError, nil, errors.New("internal server error, try again later")
	}

	data_arr := make([]data.PasswordDataOnLogin, 0)
	for r.Next() {
		var data data.PasswordDataOnLogin
		err := r.Scan(&data.SiteName, &data.Password_e)
		if err != nil {
			log.Printf("LOGIN: user %s scanning sitename failed %s", a.Username, err.Error())
			return http.StatusInternalServerError, nil, errors.New("internal server error, try again later")
		}
		data_arr = append(data_arr, data)
	}
	return http.StatusOK, data_arr, nil
}

func decryptPassword(password_e string) (string, error) {
	password_d, err := base64.StdEncoding.DecodeString(password_e)
	return string(password_d), err
	// secret_key_bytes := []byte(SECRET_KEY)

	// ciphertext, err := hex.DecodeString(password_e)
	// if err != nil {
	// 	return "", err
	// }

	// block, err := aes.NewCipher(secret_key_bytes)
	// if err != nil {
	// 	return "", err
	// }

	// plaintext := make([]byte, len(ciphertext))
	// block.Decrypt(plaintext, ciphertext)
	// return string(plaintext), nil
}

func encryptPassword(password_d string) string {
	plaintext_bytes := []byte(password_d)
	return base64.StdEncoding.EncodeToString(plaintext_bytes)

	// secret_key_bytes := []byte(SECRET_KEY)
	// plaintext_bytes := []byte(password_d)

	// block, err := aes.NewCipher(secret_key_bytes)
	// if err != nil {
	// 	return "", err
	// }

	// ciphertext := make([]byte, len(plaintext_bytes))
	// block.Encrypt(ciphertext, plaintext_bytes)
	// return hex.EncodeToString(ciphertext), nil
}
