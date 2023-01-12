package e2e

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	helper "github.com/Azure/terraform-module-test-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
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
			path := fmt.Sprintf("../../%s", f)
			defer func() {
				if !t.Failed() {
					_ = helper.RecordVersionSnapshot(t, filepath.Join("..", ".."), f)
				}
			}()
			helper.RunE2ETest(t, path, "", terraform.Options{
				Upgrade: true,
			}, nil)
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
