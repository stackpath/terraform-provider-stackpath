package stackpath

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestStackpathConvertContainerSecurityContext(t *testing.T) {
	// Note that this is different from runtime container security context

	// Keep in mind that within the schema security context
	// is a single-element array
	r := resourceComputeWorkloadSecurityContextSchema().Elem.(*schema.Resource)
	s := r.Schema

	rawData := map[string]interface{}{
		"allow_privilege_escalation": true,
		"read_only_root_filesystem":  true,
		"run_as_user":                "999",
		"run_as_group":               "991",
		"run_as_non_root":            true,
		"capabilities": []interface{}{
			map[string]interface{}{
				"add": []interface{}{
					"NET_ADMIN",
				},
				"drop": []interface{}{
					"NET_BROADCAST",
					"FOO",
				},
			},
		},
	}
	testData := schema.TestResourceDataRaw(t, s, rawData)

	secContext := convertComputeWorkloadSecurityContext("", testData)

	assert.NotNil(t, secContext)
	assert.True(t, secContext.AllowPrivilegeEscalation)
	assert.True(t, secContext.ReadOnlyRootFilesystem)
	assert.Equal(t, "999", secContext.RunAsUser)
	assert.Equal(t, "991", secContext.RunAsGroup)
	assert.NotNil(t, secContext.Capabilities)

	assert.Len(t, secContext.Capabilities.Add, 1)
	assert.Len(t, secContext.Capabilities.Drop, 2)
}

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
