package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strings"
    "time"

    "github.com/casbin/casbin/v2"
    "github.com/go-ldap/ldap/v3"
    "github.com/golang-jwt/jwt/v4"
)

var (
    ldapServer   = "localhost:389"
    baseDN       = "dc=example,dc=org"
    adminDN      = "cn=admin,dc=example,dc=org"
    adminPass    = "admin"
    jwtKey       = []byte("secret_key") // ganti dengan secret mu sendiri
    casbinEnforcer *casbin.Enforcer
)

// Login request format
type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// JWT Claims custom (gunakan RegisteredClaims supaya ada expiry)
type Claims struct {
    Username string `json:"username"`
    jwt.RegisteredClaims
}

func main() {
    // Load Casbin model dan policy
    var err error
    casbinEnforcer, err = casbin.NewEnforcer("rbac_model.conf", "policy.csv")
    if err != nil {
        log.Fatalf("Failed to create Casbin enforcer: %v", err)
    }
    err = casbinEnforcer.LoadPolicy()
    if err != nil {
        log.Fatalf("Failed to load Casbin policy: %v", err)
    }

    http.HandleFunc("/login", loginHandler)
    http.Handle("/public", authMiddleware(http.HandlerFunc(publicHandler)))
    http.Handle("/private", authMiddleware(http.HandlerFunc(privateHandler)))
    http.Handle("/admin", authMiddleware(http.HandlerFunc(adminHandler)))

    fmt.Println("Server running at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
        return
    }

    var creds Credentials
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // LDAP authenticate
    ok, err := ldapAuthenticate(creds.Username, creds.Password)
    if err != nil {
        http.Error(w, "LDAP error: "+err.Error(), http.StatusInternalServerError)
        return
    }
    if !ok {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    // Generate JWT token
    expirationTime := time.Now().Add(1 * time.Hour)
    claims := &Claims{
        Username: creds.Username,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        http.Error(w, "Could not generate token", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "token": tokenString,
    })
}

func ldapAuthenticate(username, password string) (bool, error) {
    conn, err := ldap.DialURL("ldap://" + ldapServer)
    if err != nil {
        return false, err
    }
    defer conn.Close()

    // Bind as admin to search user DN
    err = conn.Bind(adminDN, adminPass)
    if err != nil {
        return false, fmt.Errorf("admin bind failed: %w", err)
    }

    searchRequest := ldap.NewSearchRequest(
        baseDN,
        ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
        fmt.Sprintf("(uid=%s)", ldap.EscapeFilter(username)),
        []string{"dn"},
        nil,
    )

    sr, err := conn.Search(searchRequest)
    if err != nil {
        return false, fmt.Errorf("search failed: %w", err)
    }
    if len(sr.Entries) != 1 {
        return false, fmt.Errorf("user does not exist or too many entries returned")
    }

    userDN := sr.Entries[0].DN

    // Bind as user to verify password
    err = conn.Bind(userDN, password)
    if err != nil {
        return false, nil // false = invalid credentials, no error
    }

    return true, nil
}

func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        claims := &Claims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })
        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        username := claims.Username
        if username == "" {
            http.Error(w, "Invalid token claims", http.StatusUnauthorized)
            return
        }

        action := methodToCasbinAction(r.Method)
        obj := r.URL.Path

        allowed, err := casbinEnforcer.Enforce(username, obj, action)
        if err != nil {
            http.Error(w, "Authorization error", http.StatusInternalServerError)
            return
        }
        if !allowed {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }

        next.ServeHTTP(w, r)
    })
}

func methodToCasbinAction(method string) string {
    switch method {
    case "GET":
        return "read"
    case "POST", "PUT", "DELETE":
        return "write"
    default:
        return "read"
    }
}

func publicHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello Public! Only users with 'user' role can access this"))
}

func privateHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello Private! Only 'admin' role can access this"))
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello Admin! Only 'admin' role can write here"))
}
