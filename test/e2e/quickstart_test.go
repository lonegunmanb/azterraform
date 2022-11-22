package e2e

import (
	"fmt"
	helper "github.com/Azure/terraform-module-test-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"os"
	"strings"
	"testing"
)

func Test_ChangedQuickstarts(t *testing.T) {
	input := os.Getenv("CHANGED_FOLDERS")
	folders := strings.Split(input, ",")
	if input == "" {
		var err error
		folders, err = allExamples()
		if err != nil {
			t.Fatalf(err.Error())
		}
	}
	for _, f := range folders {
		f = strings.TrimSpace(f)
		t.Run(f, func(t *testing.T) {
			helper.RunE2ETest(t, fmt.Sprintf("../../%s", f), "", terraform.Options{
				Upgrade: true,
			}, func(t *testing.T, output helper.TerraformOutput) {

			})
		})
	}
}

func allExamples() ([]string, error) {
	examples, err := os.ReadDir("../../quickstart")
	if err != nil {
		return nil, err
	}
	var r []string
	for _, f := range examples {
		if !f.IsDir() {
			continue
		}
		r = append(r, f.Name())
	}
	return r, nil
}
