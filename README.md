# faker
generate fake words for you.

# languages
-  fa-IR
-  en-US

## generate
- Name
- Address
- Animal
- Car
- Color
- Country
- Email
- Gender
- Image
- Job
- Username
- URL
- Title
- Text
- RandomHash
- Phone
- Password
- Longitude
- Latitude


## examples:
```Go

import (
	"fmt"

	"github.com/emadghaffari/seeder/seeder"
)

// default language is [FA]
func main() {
    seeder.SetLang("fa") // for fa-IR
    seeder.SetLang("en") // for en-US

	fmt.Println("Name: ", seeder.Name())
	fmt.Println("Address: ", seeder.Address())
	fmt.Println("Animal: ", seeder.Animal())
	fmt.Println("Car: ", seeder.Car())
	fmt.Println("Color: ", seeder.Color())
	fmt.Println("Country: ", seeder.Country())
	fmt.Println("Email: ", seeder.Email())
	fmt.Println("Gender: ", seeder.Gender())
	fmt.Println("Image: ", seeder.Image())
	fmt.Println("Job: ", seeder.Job())
	fmt.Println("Username: ", seeder.Username())
	fmt.Println("URL: ", seeder.URL())
	fmt.Println("Title: ", seeder.Title())
	fmt.Println("Text: ", seeder.Text())
	fmt.Println("RandomHash: ", seeder.RandomHash(15))
	fmt.Println("Phone: ", seeder.Phone())
	fmt.Println("Password: ", seeder.Password())
	fmt.Println("Longitude: ", seeder.Longitude())
	fmt.Println("Latitude: ", seeder.Latitude())
}

```