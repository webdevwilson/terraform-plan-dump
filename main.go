package main

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/plans/planfile"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Please provide path to Terraform plan file")
	}

	path := os.Args[1]

	planFile, err := planfile.Open(path)
	if err != nil {
		panic(err)
	}

	plan, err := planFile.ReadPlan()
	if err != nil {
		panic(err)
	}

	dump := make(map[string] string)
	for _, res := range plan.Changes.Resources {
		dump[res.Addr.String()] = res.Action.String()
	}

	b, err := json.Marshal(dump)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(b))
}
