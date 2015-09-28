package main

import (
	"log"
	"os"

	"github.com/abourget/rally"
)

func main() {
	rally.Debug = true
	r := rally.New(os.Getenv("RALLY_API_KEY"))
	//hr, err := r.GetHierarchicalRequirement("123")
	//hr, err := r.GetHierarchicalRequirement("32667874860")
	//hr, err := r.GetHierarchicalRequirement("32667874860")
	//hr, err := r.GetHierarchicalRequirement("11802")
	//_, err := r.QueryHierarchicalRequirement("17242")
	_, err := r.QueryPortfolioItemFeature("589")
	if err != nil {
		log.Fatalln("Error:", err)
	}

	//fmt.Printf("Found:\n Parent: %#v\n ParentRef: %#v\nIteration: %#v\nIterationRef: %#v\n", hr.Parent, hr.ParentRef, hr.Iteration, hr.IterationRef)
	//fmt.Printf("Found:\n CreationDate: %s\nIteration: %#v\n", hr.CreationDate.Format(time.RFC850), hr.Iteration)
}
