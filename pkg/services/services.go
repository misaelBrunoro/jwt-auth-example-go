package services

import (
	"fmt"
	"net/http"

	"github.com/jwt-auth/pkg/config"
)

func InitServices(cfg *config.Config) error {
	http.HandleFunc("/login", Login)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/refresh", Refresh)

	connection := fmt.Sprintf(":%d", cfg.HttpPort)

	fmt.Printf("Http running in %s", connection)

	return http.ListenAndServe(connection, nil)
}
