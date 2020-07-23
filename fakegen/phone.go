package fakegen

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
)

var phone Phoner

// GetPhoner serves as a constructor for Phoner interface
func GetPhoner() Phoner {
	mu.Lock()
	defer mu.Unlock()

	if phone == nil {
		phone = &Phone{}
	}
	return phone
}

// SetPhoner sets custom Phoner
func SetPhoner(p Phoner) {
	phone = p
}

// Phoner serves overall tele-phonic contact generator
type Phoner interface {
	PhoneNumber(ctx context.Context, v reflect.Value) (interface{}, error)
	TollFreePhoneNumber(ctx context.Context, v reflect.Value) (interface{}, error)
	E164PhoneNumber(ctx context.Context, v reflect.Value) (interface{}, error)
}

// Phone struct
type Phone struct {
}

func (p Phone) phonenumber() string {
	randInt, _ := RandomInt(1, 10)
	str := strings.Join(IntToString(randInt), "")
	return fmt.Sprintf("%s-%s-%s", str[:3], str[3:6], str[6:10])
}

// PhoneNumber generates phone numbers of type: "201-886-0269"
func (p Phone) PhoneNumber(ctx context.Context, v reflect.Value) (interface{}, error) {
	return p.phonenumber(), nil
}

// Phonenumber get fake phone number
func Phonenumber() string {
	p := Phone{}
	return p.phonenumber()
}

func (p Phone) tollfreephonenumber() string {
	out := ""
	boxDigitsStart := []string{"777", "888"}

	ints, _ := RandomInt(1, 9)
	for index, v := range IntToString(ints) {
		if index == 3 {
			out += "-"
		}
		out += v
	}
	return fmt.Sprintf("(%s) %s", boxDigitsStart[rand.Intn(1)], out)
}

// TollFreePhoneNumber generates phone numbers of type: "(888) 937-7238"
func (p Phone) TollFreePhoneNumber(ctx context.Context, v reflect.Value) (interface{}, error) {
	return p.tollfreephonenumber(), nil
}

// TollFreePhoneNumber get fake TollFreePhoneNumber
func TollFreePhoneNumber() string {
	p := Phone{}
	return p.tollfreephonenumber()
}

func (p Phone) e164PhoneNumber() string {
	out := ""
	boxDigitsStart := []string{"7", "8"}
	ints, _ := RandomInt(1, 10)

	for _, v := range IntToString(ints) {
		out += v
	}
	return fmt.Sprintf("+%s%s", boxDigitsStart[rand.Intn(1)], strings.Join(IntToString(ints), ""))
}

// E164PhoneNumber generates phone numbers of type: "+27113456789"
func (p Phone) E164PhoneNumber(ctx context.Context, v reflect.Value) (interface{}, error) {
	return p.e164PhoneNumber(), nil
}

// E164PhoneNumber get fake E164PhoneNumber
func E164PhoneNumber() string {
	p := Phone{}
	return p.e164PhoneNumber()
}
