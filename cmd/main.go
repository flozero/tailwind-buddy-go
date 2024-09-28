package main

import (
	"html/template"
	"net/http"
	"tailwindbuddy/components"
)

// Handler for the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	button := components.Button()

	// Create template for the home page
	tmpl := template.Must(template.New("home").Parse(`
	<!DOCTYPE html>
	<html lang="en" class="h-full">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Home Page</title>
		<link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css" rel="stylesheet">
	</head>
	<body class="flex justify-center items-center h-full">
		<button type="button" class="{{.buttonClasses}}">
			<h1>Welcome to the Homepage</h1>
			<p>This is a simple page using a templating system.</p>
		</button>
	</body>
	</html>
	`))

	// Execute the template, passing dynamic classes
	tmpl.Execute(w, map[string]string{
		"buttonClasses": button.Root(components.ButtonProps{
			Color: "primary",
			Size:  "large",
			Class: "bg-purple-500",
		}),
	})
}

func main() {
	// Handle the "/" route
	http.HandleFunc("/", homeHandler)
	// Start the server
	http.ListenAndServe(":8080", nil)
}