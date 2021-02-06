package seeder

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

// Name func
func Name() string {
	l := Random(0, len(name))
	return name[l]
}

// Email func
func Email() string {
	l := Random(0, len(email))
	return email[l]
}

// Gender func
func Gender() string {
	l := Random(0, len(gender))
	return gender[l]
}

// Address func
func Address() string {
	l := Random(0, len(address))
	return address[l]
}

// Animal func
func Animal() string {
	l := Random(0, len(animal))
	return animal[l]
}

// Car func
func Car() string {
	l := Random(0, len(car))
	return car[l]
}

// Color func
func Color() string {
	l := Random(0, len(color))
	return color[l]
}

// Country func
func Country() string {
	l := Random(0, len(country))
	return country[l]
}

// Image func
func Image() string {
	l := Random(0, len(image))
	return image[l]
}

// Phone func
func Phone() string {
	l := Random(0, len(phone))
	return phone[l]
}

// URL func
func URL() string {
	l := Random(0, len(url))
	return url[l]
}

// Password func
func Password() string {
	l := Random(0, len(password))
	return password[l]
}

// Latitude func
func Latitude() float64 {
	l := Random(0, len(latitude))
	return latitude[l]
}

// Longitude func
func Longitude() float64 {
	l := Random(0, len(longitude))
	return longitude[l]
}

// Username func
func Username() string {
	l := Random(0, len(username))
	return username[l]
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
	l := Random(0, len(job))
	return job[l]
}

// Title func
func Title() string {
	l := Random(0, len(title))
	return title[l]
}

// Text func
func Text() string {
	l := Random(0, len(text))
	return text[l]
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
