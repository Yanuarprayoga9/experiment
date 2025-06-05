package main
import (
	"fmt"
	"log"

	"github.com/go-ldap/ldap/v3"
)
func authenticateLDAP(username, password string) bool {
	// LDAP server config
	ldapServer := "ldap://localhost:389"
	baseDN := "dc=example,dc=org"
	bindUsername := fmt.Sprintf("uid=%s,ou=users,%s", username, baseDN)

	conn, err := ldap.DialURL(ldapServer)
	if err != nil {
		log.Println("LDAP connection error:", err)
		return false
	}
	defer conn.Close()

	// Try to bind as the user
	err = conn.Bind(bindUsername, password)
	if err != nil {
		log.Println("LDAP authentication failed:", err)
		return false
	}
	return true
}
