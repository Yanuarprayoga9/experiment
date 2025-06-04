package services

import (
	"fmt"
	"os"

	"github.com/go-ldap/ldap/v3"
)

func AuthenticateLDAPByEmail(email string, password string) (bool, error) {

	LDAP_URL := os.Getenv("LDAP_URL")
	LDAP_BIND_DN := os.Getenv("LDAP_BIND_ADMIN_DN")
	LDAP_BIND_PASSWORD := os.Getenv("LDAP_BIND_ADMIN_PASSWORD")
	LDAP_BASE_DN := os.Getenv("LDAP_BASE_DN")

	if LDAP_URL == "" {
		return false, fmt.Errorf("LDAP_URL is not set")
	}
	if LDAP_BIND_DN == "" {
		return false, fmt.Errorf("LDAP_BIND_DN is not set")
	}
	if LDAP_BIND_PASSWORD == "" {
		return false, fmt.Errorf("LDAP_BIND_PASSWORD is not set")
	}
	if LDAP_BASE_DN == "" {
		return false, fmt.Errorf("LDAP_BASE_DN is not set")
	}

	l, err := ldap.DialURL(LDAP_URL)
	if err != nil {
		return false, fmt.Errorf("failed to connect to LDAP: %v", err)
	}

	defer l.Close()

	err = l.Bind(LDAP_BIND_DN, LDAP_BIND_PASSWORD)
	if err != nil {
		fmt.Println("Binding with admin credentials", err)
		return false, fmt.Errorf("failed to bind: %v", err)
	}
	searchRequest := ldap.NewSearchRequest(
		LDAP_BASE_DN,                  // Base DN
		ldap.ScopeWholeSubtree,        // Search scope
		ldap.NeverDerefAliases,        // Dereference aliases
		0,                             // Size limit
		0,                             // Time limit
		false,                         // TypesOnly
		"(objectClass=inetOrgPerson)", // Search filter
		[]string{"mail"},              // Attributes to retrieve
		nil,                           // Controls
	)
	result, err := l.Search(searchRequest)
	if err != nil {
		return false, fmt.Errorf("search error: %v", err)
	}
	var currentUserDN *ldap.Entry

	for _, entry := range result.Entries {
		emailEntry := entry.GetAttributeValue("mail")
		if email == emailEntry {
			currentUserDN = entry
			break
		}
	}

	if currentUserDN == nil {
		return false, fmt.Errorf("user not found")
	}

	// Try to bind with the user's DN and password to verify credentials
	err = l.Bind(currentUserDN.DN, password)
	if err != nil {
		return false, fmt.Errorf("authentication failed: invalid credentials")
	}

	return true, nil
}
