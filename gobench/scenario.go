package scenario

import (
	"context"

	http "github.com/gobench-io/gobench/clients/http"
	"github.com/gobench-io/gobench/executor/scenario"
)

// export scenarios
func export() scenario.Vus {
	return scenario.Vus{
		{
			Nu:   1000, // 1 vu
			Rate: 200,
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

	for {
		// create new user
		if _, err = client.Get(ctx, "http://nginx", headers); err != nil {
			return
		}
	}
}
