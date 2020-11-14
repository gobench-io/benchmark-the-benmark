package main

import (
	"context"
	"fmt"

	http "github.com/gobench-io/gobench/clients/http"
	"github.com/gobench-io/gobench/executor/scenario"
)

// export scenarios
func export() scenario.Vus {
	return scenario.Vus{
		{
			Nu:   1, // 1 vu
			Rate: 0.5,
			Fu:   f,
		},
	}
}

func f(ctx context.Context, vui int) {
	var err error

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	fmt.Printf("sss%+v\n", ctx)
	client, err := http.NewHttpClient(ctx, "Nginx_Welcome_Page")

	// create new user
	if _, err = client.Get(ctx, "http://nginx", headers); err != nil {
		return
	}

	return
}
func main() {
	ex := export()
	ex[0].Fu(context.TODO(), 0)
}
