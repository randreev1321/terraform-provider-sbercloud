package eps

import (
	"testing"

	"github.com/sbercloud-terraform/terraform-provider-sbercloud/sbercloud/acceptance"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccEnterpriseProjectDataSource_basic(t *testing.T) {
	dataSourceName := "data.sbercloud_enterprise_project.test"
	dc := acceptance.InitDataSourceCheck(dataSourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccEnterpriseProjectDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					dc.CheckResourceExists(),
					resource.TestCheckResourceAttr(dataSourceName, "name", "default"),
					resource.TestCheckResourceAttr(dataSourceName, "id", "0"),
				),
			},
		},
	})
}

const testAccEnterpriseProjectDataSource_basic = `
data "sbercloud_enterprise_project" "test" {
  name = "default"
}
`
