package test

import (
    "testing"
    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/stretchr/testify/assert"
)

func TestTerraformModule(t *testing.T) {
    t.Parallel()

    opts := &terraform.Options{
        TerraformDir: "/home/circleci/project/terratest/infra-module",
        TerraformBinary : "terragrunt",
    }

    defer terraform.TgDestroyAll(t, opts)
    terraform.TgApplyAll(t, opts)
    httpInstances := terraform.OutputList(t, opts, "http_ip")
	dbInstances := terraform.OutputList(t, opts, "db_ip")

	assert.Equal(t, 2, len(httpInstances), "Expected 2 HTTP instances to be created")
	assert.Equal(t, 3, len(dbInstances), "Expected 3 DB instances to be created")


    vpcCidr := terraform.Output(t, opts, "vpc_cidr")
    assert.Equal(t, vpcCidr, "192.168.0.0/16", "Expected VPC CIDR block to be 192.168.0.0/16")

}
