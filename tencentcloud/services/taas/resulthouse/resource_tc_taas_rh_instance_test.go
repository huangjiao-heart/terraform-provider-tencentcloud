package taas_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	tcacctest "github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/acctest"
	"testing"
)

func TestAccTencentCloudTaasRHInstanceResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			tcacctest.AccPreCheck(t)
		},
		Providers: tcacctest.AccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTaasRHInstance,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("tencentcloud_taas_rh_instance.rh_instance", "id")),
			},
			{
				ResourceName:      "tencentcloud_taas_rh_instance.rh_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccTaasRHInstance = `

resource "tencentcloud_taas_rh_instance" "rh_instance" {
    instance_type = &lt;nil&gt;
  instance_name = &lt;nil&gt;
  volume = &lt;nil&gt;
    renew_flag = &lt;nil&gt;
  charge_type = &lt;nil&gt;
        tags {
		tag_key = &lt;nil&gt;
		tag_value = &lt;nil&gt;

  }
  period = &lt;nil&gt;
  count = &lt;nil&gt;
  auto_voucher = &lt;nil&gt;
  deal_name = &lt;nil&gt;
  root_password = &lt;nil&gt;
}

`
