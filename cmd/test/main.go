package main

import (
	"fmt"
	"github.com/slink-go/awx-client-go/awx"
	"github.com/slink-go/awx-client-go/awx/api"
	"os"
	"strings"
	"time"
)

const size = 100

func main() {

	os.Setenv("GO_ENV", "dev")
	os.Setenv("LOGGING_LEVEL_ROOT", "INFO")

	connection, err := api.NewAwxClientBuilder().
		URL("http://awx.infra.loc/api").
		Username("overmind").
		Password("rj95cusE?Du[.*O").
		//Proxy(proxy).
		//CAFile(caFile).
		Insecure(true).
		Build()
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	//listJobTemplates(connection)
	//listWorkflowTemplates(connection)
	//getJobTemplate(connection, 337)
	//getWorkflowTemplate(connection, 479)
	start := time.Now()
	listAll(connection)
	stop := time.Now()
	taken := stop.Sub(start)
	fmt.Printf("\n\ntime taken: %v\n\n", taken)
}

func listAll(ac *api.Awx) {

	fetcher := awx.NewTemplateFetcher(
		ac,
		100*time.Millisecond,
		awx.NewAnyLabelFilter("CR_AUTO"),
		//awx.NewNegatedTemplateIdFilter(336, 405),
	)
	res, err := fetcher.Fetch()
	if err != nil {
		panic(err)
	}
	var idx = 1
	for _, p := range res {
		var lll []string
		for l := range p.Labels().Iter() {
			lll = append(lll, l)
		}
		fmt.Printf("[%02d] Template #%d: %d %s %s\n", idx, p.Id(), p.Kind(), p.Name(), labels(lll))
		for _, v := range p.Variables() {
			fmt.Printf("     %s (var=%s; kind=%s; default=%s)\n", v.Question(), v.Name(), v.Kind(), v.Default())
		}
		idx++
	}
}

func listJobTemplates(ac *api.Awx) {
	fmt.Println("\nList job templates")
	var idx = 1
	var page = 1
	var resource = ac.JobTemplates()
	for {
		// Send the request to retrieve the project:
		response, err := resource.Page(page, size).Send()
		if err != nil {
			panic(err)
		}
		// Print results
		for _, p := range response.Results() {
			var lll []string
			for _, l := range p.Labels() {
				lll = append(lll, l)
			}
			fmt.Printf("[%02d] Job template #%d: %s %s\n", idx, p.Id(), p.Name(), labels(lll))
			idx++
		}
		if !response.HasNext() {
			break
		}
		page++
	}
}
func listWorkflowTemplates(ac *api.Awx) {
	fmt.Println("\nList workflow templates")
	var idx = 1
	var page = 1
	var resource = ac.WorkflowJobTemplates()
	for {

		// Send the request to retrieve the project:
		response, err := resource.Page(page, size).Send()
		if err != nil {
			panic(err)
		}

		// Print results
		for _, p := range response.Results() {
			var lll []string
			for _, l := range p.Labels() {
				lll = append(lll, l)
			}
			fmt.Printf("[%02d] Workflow template #%d: %s %s\n", idx, p.Id(), p.Name(), labels(lll))
			idx++
		}
		if !response.HasNext() {
			break
		}
		page++
	}
}
func getJobTemplate(ac *api.Awx, id int) {
	fmt.Printf("\nGet job template #%d\n", id)
	var resource = ac.JobTemplates().Id(id)
	response, err := resource.Get().Send()
	if err != nil {
		panic(err)
	}

	// Print results
	fmt.Printf(
		"Job template #%d: %s %s %s\n",
		response.Result().Id(),
		response.Result().Name(),
		response.Result().Description(),
		labels(response.Result().Labels()),
	)

}
func getWorkflowTemplate(ac *api.Awx, id int) {
	fmt.Printf("\nGet workflow template #%d\n", id)
	var resource = ac.WorkflowJobTemplates().Id(id)
	response, err := resource.Get().Send()
	if err != nil {
		panic(err)
	}

	// Print results
	fmt.Printf(
		"Workflow template #%d: %s %s %s\n",
		response.Result().Id(),
		response.Result().Name(),
		response.Result().Description(),
		labels(response.Result().Labels()),
	)

}

func labels(input []string) string {
	if input == nil || len(input) == 0 {
		return ""
	}
	return fmt.Sprintf("(%s)", strings.Join(input, ", "))
}
