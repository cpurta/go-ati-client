package util

import (
	"net/http"

	"github.com/cpurta/go-ati-client/config"
)

type BaseClient struct {
	Config     *config.Config
	HTTPClient *http.Client
}
