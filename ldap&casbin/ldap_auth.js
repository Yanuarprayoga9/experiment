// package main

// import (
// 	// "fmt"
// 	"fmt"
// 	"log"

// 	"github.com/go-ldap/ldap/v3"
// )
// func authenticateLDAP(username, password string) bool {
// 	ldapServer := "ldap://localhost:389"
// 	baseDN := "dc=example,dc=org"

// 	conn, err := ldap.DialURL(ldapServer)
// 	if err != nil {
// 		log.Println("LDAP connection error:", err)
// 		return false
// 	}
// 	defer conn.Close()

// 	// Bind dulu sebagai admin untuk bisa cari user
// 	err = conn.Bind("cn=admin,dc=example,dc=org", "admin")
// 	if err != nil {
// 		log.Println("Admin bind failed:", err)
// 		return false
// 	}

// 	// Cari DN user berdasar username
// 	searchRequest := ldap.NewSearchRequest(
// 		baseDN,
// 		ldap.ScopeWholeSubtree,
// 		ldap.NeverDerefAliases,
// 		0,
// 		0,
// 		false,
// 		fmt.Sprintf("(uid=%s)", username),
// 		[]string{"dn"},
// 		nil,
// 	)

// 	sr, err := conn.Search(searchRequest)
// 	if err != nil || len(sr.Entries) == 0 {
// 		log.Println("User not found or search failed:", err)
// 		return false
// 	}

// 	userDN := sr.Entries[0].DN

// 	// Bind ulang pakai user DN dan password input untuk verifikasi password
// 	err = conn.Bind(userDN, password)
// 	if err != nil {
// 		log.Println("User bind failed:", err)
// 		return false
// 	}

// 	return true
// }
