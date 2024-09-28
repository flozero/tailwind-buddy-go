package internal

import (
	"reflect"
	"strings"
)

type VariantKey string

type Variants[T any] map[VariantKey]map[string]string
type DefaultVariants[T any] map[VariantKey]string
type CompoundVariant[T any] struct {
	Conditions map[VariantKey]string
	Class      string
}

type ComposeOptions[T any] struct {
	Variants         Variants[T]
	CompoundVariants []CompoundVariant[T]
	DefaultVariants  DefaultVariants[T]
}

func uniquifyClasses(classes []string) string {
	uniqueMap := make(map[string]bool)
	for _, class := range classes {
		if class != "" {
			uniqueMap[class] = true
		}
	}
	
	uniqueClasses := make([]string, 0, len(uniqueMap))
	for class := range uniqueMap {
		uniqueClasses = append(uniqueClasses, class)
	}
	
	return strings.Join(uniqueClasses, " ")
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func Compose[T any](options ComposeOptions[T], baseClass string) func(T) string {
	return func(props T) string {
		classSet := []string{baseClass}

		propsValue := reflect.ValueOf(props)
		propsType := propsValue.Type()

		defaultVariantsToIgnore := []string{}

		for i := 0; i < propsValue.NumField(); i++ {
			field := propsType.Field(i)
			value := propsValue.Field(i)
			if (field.Name != "Class") {
				if (value.String() != "") {
					defaultVariantsToIgnore = append(defaultVariantsToIgnore, field.Name)
				}
			}
		}

		// Apply default variants
		for variantKey, defaultValue := range options.DefaultVariants {
			if !contains(defaultVariantsToIgnore, string(variantKey)) {
				if variantClasses, ok := options.Variants[variantKey][defaultValue]; ok {
					classSet = append(classSet, variantClasses)
				}
			} else {
				if variantClasses, ok := options.Variants[variantKey][propsValue.FieldByName(string(variantKey)).String()]; ok {
					classSet = append(classSet, variantClasses)
				}
			}
		}

		// Apply compound variants
		for _, cv := range options.CompoundVariants {
			isMatch := true
			for condKey, condValue := range cv.Conditions {
				field, ok := propsType.FieldByName(string(condKey))
				if !ok {
					isMatch = false
					break
				}
				propValue := propsValue.FieldByName(field.Name).String()
				if propValue == "" {
					propValue = options.DefaultVariants[condKey]
				}
				if propValue != condValue {
					isMatch = false
					break
				}
			}
			if isMatch {
				classSet = append(classSet, cv.Class)
			}
		}

		// Add custom class if provided
		if classField, ok := propsType.FieldByName("Class"); ok {
			customClass := propsValue.FieldByName(classField.Name).String()
			if customClass != "" {
				classSet = append(classSet, customClass)
			}
		}

		return uniquifyClasses(classSet)
	}
}