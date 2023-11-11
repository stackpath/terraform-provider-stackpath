package stackpath

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestStackpathConvertContainerRuntime_SecurityContext(t *testing.T) {

	r := resourceComputeContainerRuntimeEnvironment()
	s := r.Schema

	testData := schema.TestResourceDataRaw(t, s, map[string]interface{}{
		"termination_grace_period_seconds": 60,
		"share_process_namespace":          true,
		"security_context": []interface{}{
			map[string]interface{}{
				"run_as_group":        "101",
				"run_as_user":         "99",
				"run_as_non_root":     true,
				"supplemental_groups": []interface{}{"42"},
				"sysctl": map[string]interface{}{
					"foo.bar": "1",
					"foo.baz": "2",
				},
			},
		},
	})

	converted := convertComputeWorkloadRuntimeContainer("", testData)
	assert.NotNil(t, converted)
	assert.NotNil(t, converted.SecurityContext)
	assert.Len(t, converted.SecurityContext.Sysctls, 2)

	assert.Equal(t, "101", converted.SecurityContext.RunAsGroup)
	assert.Equal(t, "99", converted.SecurityContext.RunAsUser)
	assert.True(t, converted.SecurityContext.RunAsNonRoot)
	assert.Len(t, converted.SecurityContext.SupplementalGroups, 1)

	assert.True(t, converted.ShareProcessNamespace)
	assert.Equal(t, converted.TerminationGracePeriodSeconds, "60")

}

func TestStackpathConvertContainerRuntime_DNS(t *testing.T) {

	r := resourceComputeContainerRuntimeEnvironment()
	s := r.Schema

	testData := schema.TestResourceDataRaw(t, s, map[string]interface{}{
		"dns": []interface{}{
			map[string]interface{}{
				"host_aliases": []interface{}{
					map[string]interface{}{
						"address":   "192.168.3.4",
						"hostnames": []interface{}{"domain.com", "domain2.com"},
					},
				},
				"resolver_config": []interface{}{
					map[string]interface{}{
						"nameservers": []interface{}{"172.16.42.1"},
						"search":      []interface{}{"domain.com", "test.domain.com"},
						"options": map[string]interface{}{
							"timeout": "10",
						},
					},
				},
			},
		},
	})

	converted := convertComputeWorkloadRuntimeContainer("", testData)
	assert.NotNil(t, converted)
	assert.NotNil(t, converted.HostAliases)
	assert.Len(t, converted.HostAliases, 1)
	assert.Equal(t, "192.168.3.4", converted.HostAliases[0].IP)
	assert.Len(t, converted.HostAliases[0].Hostnames, 2)

	assert.NotNil(t, converted.DNSConfig)
	assert.Len(t, converted.DNSConfig.Nameservers, 1)
	assert.Len(t, converted.DNSConfig.Searches, 2)
	assert.Len(t, converted.DNSConfig.Options, 1)
	assert.Equal(t, "timeout", converted.DNSConfig.Options[0].Name)

}
