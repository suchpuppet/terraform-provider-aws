package aws

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceAwsApiGatewayRestApi(t *testing.T) {
	rName := acctest.RandString(8)
	dataSourceName := "data.aws_api_gateway_rest_api.test"
	resourceName := "aws_api_gateway_rest_api.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAwsApiGatewayRestApiConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "arn", resourceName, "arn"),
					resource.TestCheckResourceAttrPair(dataSourceName, "root_resource_id", resourceName, "root_resource_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "tags", resourceName, "tags"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "policy", resourceName, "policy"),
					resource.TestCheckResourceAttrPair(dataSourceName, "api_key_source", resourceName, "api_key_source"),
					resource.TestCheckResourceAttrPair(dataSourceName, "minimum_compression_size", resourceName, "minimum_compression_size"),
					resource.TestCheckResourceAttrPair(dataSourceName, "binary_media_types", resourceName, "binary_media_types"),
					resource.TestCheckResourceAttrPair(dataSourceName, "endpoint_configuration", resourceName, "endpoint_configuration"),
					resource.TestCheckResourceAttrPair(dataSourceName, "execution_arn", resourceName, "execution_arn"),
				),
			},
		},
	})
}

func testAccDataSourceAwsApiGatewayRestApiConfig(r string) string {
	return fmt.Sprintf(`
resource "aws_api_gateway_rest_api" "test" {
  name = "%[1]s"
}

data "aws_api_gateway_rest_api" "test" {
  name = "${aws_api_gateway_rest_api.test.name}"
}
`, r)
}
