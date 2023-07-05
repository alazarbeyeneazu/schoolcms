package utils

import (
	"bytes"
	"crypto/rand"
	"errors"
	"io"
	"mime/multipart"
	"os"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	SmallLetters   = "abcdefghijklmnopqrstuvwxyz"
	CapitalLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits         = "0987654321"
)

func GenerateCustomRandomString(set string, length int) string {
	randString := make([]byte, length)
	_, _ = io.ReadAtLeast(rand.Reader, randString, length) //nolint:errcheck // since length = len(randString)

	for i := 0; i < len(randString); i++ {
		randString[i] = set[int(randString[i])%len(set)]
	}

	return string(randString)
}

// PrepareMultipartFormFile creates a multipart/form-data body.
//
// NOTE: be sure to close the writer before sending the request
func PrepareMultipartFormFile(filePath, fieldName string) (*bytes.Buffer, *multipart.Writer, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(fieldName, file.Name())
	if err != nil {
		return nil, nil, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, nil, err
	}

	err = file.Close()
	if err != nil {
		return nil, nil, err
	}

	return body, writer, nil
}

func CheckForNullUUID(errorMessage string) validation.RuleFunc {
	return func(value interface{}) error {
		s, ok := value.(uuid.UUID)
		if !ok {
			return errors.New("faild to parse value to uuid")
		}
		nuuid := uuid.NullUUID{}
		if s == nuuid.UUID {
			return errors.New(errorMessage)
		}
		return nil
	}
}

func CheckForNullUUIDString(errorMessage string) validation.RuleFunc {
	return func(value interface{}) error {
		s, ok := value.(string)
		if !ok {
			return errors.New("faild to parse value to uuid string")
		}
		id, err := uuid.Parse(s)
		if err != nil {
			return errors.New("faild to parse value to uuid string")
		}
		nuuid := uuid.NullUUID{}
		if id == nuuid.UUID {
			return errors.New(errorMessage)
		}
		return nil
	}
}

// Calculate n percentage of x
func PercentageChange(n, x decimal.Decimal) (percentage decimal.Decimal) {
	if !x.IsZero() && !n.IsZero() {
		percentage = (n.Mul(x)).Div(decimal.NewFromInt(100))
		return percentage
	}
	return decimal.Zero
}

func Contains[T comparable](value T, array []T) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}

	return false
}
