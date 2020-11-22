package main

import (
	"context"
	"time"

	http "github.com/gobench-io/gobench/clients/http"
	"github.com/gobench-io/gobench/executor/scenario"
)

// export scenarios
func export() scenario.Vus {
	return scenario.Vus{
		scenario.Vu{
			Nu:   30, // 1 vu
			Rate: 1000,
			Fu:   f,
		},
	}
}

func f(ctx context.Context, vui int) {
	var err error

	headers := map[string]string{
		"Content-Type": "application/json",
	}
	client, err := http.NewHttpClient(ctx, "Nginx_Welcome_Page")

	period := time.Duration(30) * time.Second
	elapsed := time.Now().Add(period)
	for {
		if elapsed.Before(time.Now()) {
			return
		}
		// create new user
		if _, err = client.Get(ctx, "http://nginx", headers); err != nil {
			return
		}
	}
}
