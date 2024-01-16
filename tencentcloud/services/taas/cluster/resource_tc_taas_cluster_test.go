package cluster_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	tcacctest "github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/acctest"
	"testing"
)

func TestAccTencentCloudTaasClusterResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			tcacctest.AccPreCheck(t)
		},
		Providers: tcacctest.AccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTaasCluster,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("tencentcloud_taas_cluster.cluster", "id")),
			},
			{
				ResourceName:      "tencentcloud_taas_cluster.cluster",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccTaasCluster = `

resource "tencentcloud_taas_cluster" "cluster" {
    cluster_name = &lt;nil&gt;
  description = &lt;nil&gt;
  public_access = &lt;nil&gt;
  tags {
		tag_key = &lt;nil&gt;
		tag_value = &lt;nil&gt;

  }
    associated_rh_instance {
		region = &lt;nil&gt;
		instance_id = &lt;nil&gt;

  }
  authorized_tcr_registry_infos {
		service_account_id = &lt;nil&gt;
		region = &lt;nil&gt;
		registry_id = &lt;nil&gt;
		namespaces = &lt;nil&gt;

  }
  third_party_service_accounts {
		service_account_id = &lt;nil&gt;
		domain = &lt;nil&gt;
		username = &lt;nil&gt;
		password = &lt;nil&gt;

  }
  vpc_config {
		vpc_id = &lt;nil&gt;
		subnet_id = &lt;nil&gt;

  }
    }

`
