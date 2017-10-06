package datacentred

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/datacentred/datacentred-go"
)

func TestAccUser_basic(t *testing.T) {
	var user = datacentred.User{}
	var email = fmt.Sprintf("test.%s@test.com", acctest.RandString(5))
	var password = acctest.RandString(10)
	var firstName = "Test"
	var lastName = "User"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDatacentredUserDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDatacentredUser_basic(email, password),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatacentredUserExists("datacentred_user.user_1", &user),
					resource.TestCheckResourceAttr(
						"datacentred_user.user_1", "email", email),
					resource.TestCheckResourceAttr(
						"datacentred_user.user_1", "first_name", firstName),
					resource.TestCheckResourceAttr(
						"datacentred_user.user_1", "last_name", lastName),
				),
			},
			resource.TestStep{
				Config: testAccDatacentredUser_update(email),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatacentredUserExists("datacentred_user.user_1", &user),
					resource.TestCheckResourceAttr(
						"datacentred_user.user_1", "email", email),
					resource.TestCheckResourceAttr(
						"datacentred_user.user_1", "first_name", "Changed"),
					resource.TestCheckResourceAttr(
						"datacentred_user.user_1", "last_name", "Changed"),
				),
			},
		},
	})
}

func testAccCheckDatacentredUserDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "datacentred_user" {
			continue
		}

		_, err := datacentred.FindUser(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("User still exists")
		}
	}

	return nil
}

func testAccCheckDatacentredUserExists(n string, user *datacentred.User) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		found, err := datacentred.FindUser(rs.Primary.ID)
		if err == nil {
			return err
		}

		if found.Id != rs.Primary.ID {
			return fmt.Errorf("User not found")
		}

		*user = *found

		return nil
	}
}

func testAccDatacentredUser_basic(email, password string) string {
	return fmt.Sprintf(`
    resource "datacentred_user" "user_1" {
      email = "%s"
      password = "%s"
      first_name = "Test"
      last_name = "User"
    }
  `, email, password)
}

func testAccDatacentredUser_update(email string) string {
	return fmt.Sprintf(`
    resource "datacentred_user" "user_1" {
      email = "%s"
      first_name = "Changed"
      last_name = "Changed"
      password = "Changed123!"
    }
  `, email)
}
