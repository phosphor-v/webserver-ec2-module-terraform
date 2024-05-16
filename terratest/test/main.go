package test

import (
    "testing"
    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/stretchr/testify/assert"
)

func TestTerraformModule(t *testing.T) {
    t.Parallel()

    opts := &terraform.Options{
        TerraformDir: "../infra-module",
        TerraformBinary : "terragrunt",
    }

    defer terraform.Destroy(t, opts)
    terraform.InitAndApply(t, opts)

    ec2Instances := terraform.OutputList(t, opts, "ec2_instance_ids")
    assert.Equal(t, len(ec2Instances), 2, "Expected two EC2 instances to be created")

    vpcCidr := terraform.Output(t, opts, "vpc_cidr")
    assert.Equal(t, vpcCidr, "192.168.0.0/16", "Expected VPC CIDR block to be 192.168.0.0/16")

    dbAccessible := terraform.OutputBool(t, opts, "db_accessible_from_internet")
    assert.False(t, dbAccessible, "Database should not be accessible from the internet")
}
