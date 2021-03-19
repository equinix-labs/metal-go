package main

import (
	"context"
	"fmt"
	gometal "github.com/t0mk/gometal"
	"os"
)

func main() {
	include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
	exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)
	configuration := gometal.NewConfiguration()
	api_client := gometal.NewAPIClient(configuration)
	resp, r, err := api_client.PlansApi.FindPlans(context.Background()).Include(include).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.FindPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindPlans`: PlanList
	fmt.Fprintf(os.Stdout, "Response from `PlansApi.FindPlans`: %v\n", resp)
}
