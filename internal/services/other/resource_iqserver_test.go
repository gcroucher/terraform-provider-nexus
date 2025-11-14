package other_test

import (
	"fmt"
	"testing"

	"github.com/datadrivers/terraform-provider-nexus/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceIQServer(t *testing.T) {
	resName := "nexus_iqserver.acceptance"
	iqURL := fmt.Sprintf("https://iq-test-%s.example.com", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acceptance.AccPreCheck(t) },
		Providers: acceptance.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIQServerConfig(iqURL, "testuser", "testpass123"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "id", "iqserver"),
					resource.TestCheckResourceAttr(resName, "enabled", "false"),
					resource.TestCheckResourceAttr(resName, "url", iqURL),
					resource.TestCheckResourceAttr(resName, "authentication_type", "USER"),
					resource.TestCheckResourceAttr(resName, "username", "testuser"),
					resource.TestCheckResourceAttr(resName, "timeout_seconds", "60"),
				),
			},
			{
				ResourceName:            resName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"}, // Password not returned by API
			},
		},
	})
}

func TestAccResourceIQServerUpdate(t *testing.T) {
	resName := "nexus_iqserver.acceptance"
	iqURL1 := fmt.Sprintf("https://iq-test-%s.example.com", acctest.RandString(10))
	iqURL2 := fmt.Sprintf("https://iq-updated-%s.example.com", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acceptance.AccPreCheck(t) },
		Providers: acceptance.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIQServerConfig(iqURL1, "testuser", "testpass123"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "url", iqURL1),
					resource.TestCheckResourceAttr(resName, "username", "testuser"),
					resource.TestCheckResourceAttr(resName, "timeout_seconds", "60"),
				),
			},
			{
				Config: testAccResourceIQServerConfig(iqURL2, "updateduser", "newpass456"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "url", iqURL2),
					resource.TestCheckResourceAttr(resName, "username", "updateduser"),
					resource.TestCheckResourceAttr(resName, "timeout_seconds", "60"),
				),
			},
		},
	})
}

func TestAccResourceIQServerWithOptionalFields(t *testing.T) {
	resName := "nexus_iqserver.acceptance"
	iqURL := fmt.Sprintf("https://iq-test-%s.example.com", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acceptance.AccPreCheck(t) },
		Providers: acceptance.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIQServerConfigWithOptionalFields(iqURL, "testuser", "testpass123"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "id", "iqserver"),
					resource.TestCheckResourceAttr(resName, "enabled", "false"),
					resource.TestCheckResourceAttr(resName, "show_link", "true"),
					resource.TestCheckResourceAttr(resName, "url", iqURL),
					resource.TestCheckResourceAttr(resName, "authentication_type", "USER"),
					resource.TestCheckResourceAttr(resName, "username", "testuser"),
					resource.TestCheckResourceAttr(resName, "timeout_seconds", "120"),
					resource.TestCheckResourceAttr(resName, "use_trust_store_for_url", "true"),
					resource.TestCheckResourceAttr(resName, "fail_open_mode_enabled", "true"),
				),
			},
		},
	})
}

func testAccResourceIQServerConfig(url, username, password string) string {
	return fmt.Sprintf(`
resource "nexus_iqserver" "acceptance" {
	enabled             = false
	url                 = "%s"
	authentication_type = "USER"
	username            = "%s"
	password            = "%s"
	timeout_seconds     = 60
}
`, url, username, password)
}

func testAccResourceIQServerConfigWithOptionalFields(url, username, password string) string {
	return fmt.Sprintf(`
resource "nexus_iqserver" "acceptance" {
	enabled                 = false
	show_link               = true
	url                     = "%s"
	authentication_type     = "USER"
	username                = "%s"
	password                = "%s"
	timeout_seconds         = 120
	use_trust_store_for_url = true
	fail_open_mode_enabled  = true
}
`, url, username, password)
}
