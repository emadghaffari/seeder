package seeder

import (
	"math/rand"
	"reflect"
	"time"
)

// Name func
func Name() string {
	l := random(0, len(name))
	return name[l]
}

// Email func
func Email() string {
	l := random(0, len(email))
	return email[l]
}

// Gender func
func Gender() string {
	l := random(0, len(gender))
	return gender[l]
}

// Address func
func Address() string {
	l := random(0, len(address))
	return address[l]
}

// Animal func
func Animal() string {
	l := random(0, len(animal))
	return animal[l]
}

// Car func
func Car() string {
	l := random(0, len(car))
	return car[l]
}

// Color func
func Color() string {
	l := random(0, len(color))
	return color[l]
}

// Country func
func Country() string {
	l := random(0, len(country))
	return country[l]
}

// Image func
func Image() string {
	l := random(0, len(image))
	return image[l]
}

// Phone func
func Phone() string {
	l := random(0, len(phone))
	return phone[l]
}

// URL func
func URL() string {
	l := random(0, len(url))
	return url[l]
}

// Password func
func Password() string {
	l := random(0, len(password))
	return password[l]
}

// Latitude func
func Latitude() float64 {
	l := random(0, len(latitude))
	return latitude[l]
}

// Longitude func
func Longitude() float64 {
	l := random(0, len(longitude))
	return longitude[l]
}

// Username func
func Username() string {
	l := random(0, len(username))
	return username[l]
}

// Avatar func
func Avatar() string {
	l := random(0, len(avatar))
	return avatar[l]
}

// Job func
func Job() string {
	l := random(0, len(job))
	return job[l]
}

// RandomArray from elements
func RandomArray(items interface{}) interface{} {
	switch reflect.TypeOf(items).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(items)
		l := random(0, s.Len())
		return s.Index(l)
	default:
		return nil
	}
}

// random between numbers
func random(min, max int) int {
	if min > max {
		return 0
	}
	rand.Seed(time.Now().UnixNano())
	random := (rand.Intn((max - min)) + min)
	return random
}
