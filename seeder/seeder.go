package seeder

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"

	"github.com/google/uuid"
)

var (
	lang = "fa"
)

func SetLang(str string) error {

	if str == "fa" {
		lang = "fa"
		return nil
	}

	if str == "en" {
		lang = "en"
		return nil
	}

	return fmt.Errorf("invalid lang[fa,en]")

}

// Name func
func Name() string {
	l := Random(0, len(name[lang]))
	return name[lang][l]
}

// Email func
func Email() string {
	l := Random(0, len(email[lang]))
	return email[lang][l]
}

// Gender func
func Gender() string {
	l := Random(0, len(gender[lang]))
	return gender[lang][l]
}

// Address func
func Address() string {
	l := Random(0, len(address[lang]))
	return address[lang][l]
}

// Animal func
func Animal() string {
	l := Random(0, len(animal[lang]))
	return animal[lang][l]
}

// Car func
func Car() string {
	l := Random(0, len(car[lang]))
	return car[lang][l]
}

// Color func
func Color() string {
	l := Random(0, len(color[lang]))
	return color[lang][l]
}

// Country func
func Country() string {
	l := Random(0, len(country[lang]))
	return country[lang][l]
}

// Image func
func Image() string {
	l := Random(0, len(image[lang]))
	return image[lang][l]
}

// Phone func
func Phone() string {
	l := Random(0, len(phone[lang]))
	return phone[lang][l]
}

// URL func
func URL() string {
	l := Random(0, len(url[lang]))
	return url[lang][l]
}

// Password func
func Password() string {
	l := Random(0, len(password[lang]))
	return password[lang][l]
}

// Latitude func
func Latitude() float64 {
	l := Random(0, len(latitude[lang]))
	return latitude[lang][l]
}

// Longitude func
func Longitude() float64 {
	l := Random(0, len(longitude[lang]))
	return longitude[lang][l]
}

// Username func
func Username() string {
	l := Random(0, len(username[lang]))
	return username[lang][l]
}

// Avatar func
func Avatar() string {

	to := Random(0, len(topTypeOptions))
	ac := Random(0, len(accessoriesTypeOptions))
	fa := Random(0, len(facialHairTypeOptions))
	fas := Random(0, len(facialHairColorOptions))
	cl := Random(0, len(clotheTypeOptions))
	ey := Random(0, len(eyeTypeOptions))
	eys := Random(0, len(eyebrowTypeOptions))
	mo := Random(0, len(mouthTypeOptions))
	sk := Random(0, len(skinColorOptions))
	ha := Random(0, len(hairColorTypes))
	cls := Random(0, len(clotheColorOptions))

	return fmt.Sprintf("https://avataaars.io/?accessoriesType=%s&avatarStyle=Circle&clotheColor=%s&clotheType=%s&eyeType=%s&eyebrowType=%s&facialHairColor=%s&facialHairType=%s&hairColor=%s&mouthType=%s&skinColor=%s&topType=%s", accessoriesTypeOptions[ac], clotheColorOptions[cls], clotheTypeOptions[cl], eyeTypeOptions[ey], eyebrowTypeOptions[eys], facialHairColorOptions[fas], facialHairTypeOptions[fa], hairColorTypes[ha], mouthTypeOptions[mo], skinColorOptions[sk], topTypeOptions[to])
}

// Job func
func Job() string {
	l := Random(0, len(job[lang]))
	return job[lang][l]
}

// Title func
func Title() string {
	l := Random(0, len(title[lang]))
	return title[lang][l]
}

// Text func
func Text() string {
	l := Random(0, len(text[lang]))
	return text[lang][l]
}

// RandomArray from elements
func RandomArray(items interface{}) interface{} {
	switch reflect.TypeOf(items).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(items)
		l := Random(0, s.Len())
		return s.Index(l)
	default:
		return nil
	}
}

// RandomHasher generate a random of string
func RandomHash(lenght int) string {
	letters := []int32("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789-&()_")
	rand.Seed(time.Now().UnixNano())
	b := make([]int32, lenght)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Random between numbers
func Random(min, max int) int {
	if min > max {
		return 0
	}
	rand.Seed(time.Now().UnixNano())
	Random := (rand.Intn((max - min)) + min)
	return Random
}

func NewUUID() string {
	return uuid.Must(uuid.NewRandom()).String()
}
