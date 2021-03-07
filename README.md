# terraform-plan-dump

Dumps the resource id's and actions in a plan file into json structure like:

```
{
    "resource.id": "Create",
    "another.id": "Delete"
}
```

## Setup and Usage

**Note: Terraform version in go.mod should match version of Terraform used to create the plan**

Setup: `go get`

To use, `go run main.go someplat.tfplan > plan.json`