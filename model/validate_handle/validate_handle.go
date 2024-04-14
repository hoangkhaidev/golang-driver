package validate_handle

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func errorTagFunc[T interface{}](obj interface{}, snp string, fieldName, typeErr string) error {
    o := obj.(T)

    if !strings.Contains(snp, fieldName) {
        return nil
    }

    fieldArr := strings.Split(snp, ".")
    rsf := reflect.TypeOf(o)

    for i := 1; i < len(fieldArr); i++ {
        field, found := rsf.FieldByName(fieldArr[i])
        if found {
            if fieldArr[i] == fieldName {
				name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
				tagCustom := "msg_" + typeErr
                customMessage := field.Tag.Get(tagCustom)
                if customMessage != "" {
                    return fmt.Errorf("%s: %s", name, customMessage)
                }
                return nil
            } else {
                if field.Type.Kind() == reflect.Ptr {
                    // If the field type is a pointer, dereference it
                    rsf = field.Type.Elem()
                } else {
                    rsf = field.Type
                }
            }
        }
    }
    return nil
}

func ValidateFunc[T interface{}](obj interface{}, validate *validator.Validate) (errs error) {
    o := obj.(T)

    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in Validate:", r)
            errs = fmt.Errorf("can't validate %+v", r)
        }
    }()

    if err := validate.Struct(o); err != nil {
        errorValid := err.(validator.ValidationErrors)
        for _, e := range errorValid {
            // snp  X.Y.Z
			
            snp := e.StructNamespace()

            errmgs := errorTagFunc[T](obj, snp, e.Field(), e.ActualTag())
            if errmgs != nil {
                errs = errors.Join(errs, fmt.Errorf("%w", errmgs))
            } else {
                errs = errors.Join(errs, fmt.Errorf("%w", e))
            }
        }
    }

    if errs != nil {
        return errs
    }

    return nil
}