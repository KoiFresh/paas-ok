package multipart

import (
	"errors"
	"fmt"
	"mime/multipart"
	"slices"
	"strconv"
	"strings"
)

type Form multipart.Form

func (form *Form) GetFiles(key string) []*multipart.FileHeader {
	if files, ok := form.File[key]; ok && len(files) > 0 {
		return files
	}

	return []*multipart.FileHeader{}
}

func (form *Form) GetStringOrDefault(key string, def string) string {
	if values, ok := form.Value[key]; ok && len(values) > 0 {
		return values[0]
	}

	return def
}

func (form *Form) GetStringEnumOrFirst(key string, values []string) (string, error) {
	if value, ok := form.Value[key]; ok && len(value) > 0 {
		if slices.Contains(values, value[0]) {
			return value[0], nil
		}

		return "", errors.New(fmt.Sprintf("Unsupported value. The %s hast to be one of [ %s ]", key, strings.Join(values, ", ")))
	}

	return values[0], nil
}

func (form *Form) GetIntOrDefault(key string, def int) (int, error) {
	if values, ok := form.Value[key]; ok && len(values) > 0 {
		value, err := strconv.ParseInt(values[0], 10, 64)
		if err != nil {
			return 0, err
		}

		return int(value), nil
	}

	return def, nil
}
