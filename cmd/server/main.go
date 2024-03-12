package main

import (
	"html/template"
	"net/http"
)

func main() {
	tmpl := template.Must(template.ParseFiles("internal/templates/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := []string{"Hello", "World"}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)

	http.HandleFunc("/convertColor", func(w http.ResponseWriter, r *http.Request) {
		// 	color := r.FormValue("color")

		// 	red, green, blue, err := HexToRGB(color)
		// 	if err != nil {
		// 		http.Error(w, err.Error(), http.StatusBadRequest)
		// 		return
		// 	}

		// 	html := fmt.Sprintf(`
		//     <div>RGB: %d, %d, %d</div>
		//     <div>HEX: %s</div>
		//     <div>HSL: %s</div>
		// `, red, green, blue, RGBToHex(red, green, blue), RGBToHSL(red, green, blue))

		// 	w.Write([]byte(html))
	})

	http.ListenAndServe(":8080", nil)
}
