package cli

import (
	"context"
	"net/http"
	"time"

	"github.com/rxtreme8/gluetun/internal/constants"
	"github.com/rxtreme8/gluetun/internal/healthcheck"
)

func (c *cli) HealthCheck(ctx context.Context) error {
	const timeout = 10 * time.Second
	httpClient := &http.Client{Timeout: timeout}
	healthchecker := healthcheck.NewChecker(httpClient)
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	const url = "http://" + constants.HealthcheckAddress
	return healthchecker.Check(ctx, url)
}
