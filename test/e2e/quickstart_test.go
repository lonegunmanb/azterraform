package e2e

import (
	"fmt"
	helper "github.com/Azure/terraform-module-test-helper"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"os"
	"strings"
	"testing"
)

func Test_ChangedQuickstarts(t *testing.T) {
	input := os.Getenv("CHANGED_FOLDERS")
	folders := strings.Split(input, ",")
	for _, f := range folders {
		fmt.Printf("Testing %s", f)
		t.Run(f, func(t *testing.T) {

			helper.RunE2ETest(t, fmt.Sprintf("../../%s", f), "", terraform.Options{
				Vars:     nil,
				VarFiles: nil,
				//RetryableTerraformErrors: nil,
				//MaxRetries:               0,
				//TimeBetweenRetries:       0,
				Logger:  logger.Default,
				Upgrade: true,
				NoColor: true,
			}, func(t *testing.T, output helper.TerraformOutput) {

			})
		})
	}
}
