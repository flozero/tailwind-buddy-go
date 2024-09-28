package components

import (
	"tailwindbuddy/internal"
)

type ButtonProps struct {
	Color string
	Size  string
	Class string
}

const (
	ColorVariant internal.VariantKey = "Color"
	SizeVariant  internal.VariantKey = "Size"
)

type ButtonResult struct {
	Root func(ButtonProps) string
}

func Button() ButtonResult {
	globalDefaultVariants := internal.DefaultVariants[ButtonProps]{
		ColorVariant: "primary",
		SizeVariant:  "small",
	}

	root := internal.Compose(internal.ComposeOptions[ButtonProps]{
		Variants: internal.Variants[ButtonProps]{
			ColorVariant: {
				"primary":   /** @tw */ "bg-blue-500 text-white",
				"secondary": /** @tw */ "bg-gray-200 text-gray-800",
			},
			SizeVariant: {
				"small": /** @tw */ "text-sm",
				"large": /** @tw */"text-lg",
			},
		},
		CompoundVariants: []internal.CompoundVariant[ButtonProps]{
			{
				Conditions: map[internal.VariantKey]string{
					ColorVariant: "primary", 
					SizeVariant: "large",
			},
				Class: /** @tw */ "font-bold",
			},
		},
		DefaultVariants: globalDefaultVariants,
	}, /** @tw */ "px-4 py-2 rounded-xl shadow-lg")

	return ButtonResult{
		Root: root,
	}
}