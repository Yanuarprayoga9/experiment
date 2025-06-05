package main

import (
	"github.com/casbin/casbin/v2"
	"log"
)

var e *casbin.Enforcer

func init() {
	var err error
	e, err = casbin.NewEnforcer("rbac_model.conf", "policy.csv")
	if err != nil {
		log.Fatalf("Failed to create Casbin enforcer: %v", err)
	}
}

func checkPermission(username, obj, act string) bool {
	allowed, err := e.Enforce(username, obj, act)
	if err != nil {
		log.Println("Casbin error:", err)
		return false
	}
	return allowed
}
