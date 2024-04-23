package main

import (
	"html/template"
	"log"
	"net/http"
)

type RadioButton struct {
	Name       string
	Value      string
	IsDisabled bool
	IsChecked  bool
	Text       string
}

type PageVariables struct {
	PageTitle        string
	PageRadioButtons []RadioButton
	Answer           string
}

func main() {
	http.HandleFunc("/", DisplayRadioButtons)
	http.HandleFunc("/selected", UserSelected)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func DisplayRadioButtons(w http.ResponseWriter, r *http.Request) {
	Title := "Which do you prefer?"
	MyRadioButtons := []RadioButton{
		RadioButton{"animalselect", "cats", false, false, "Cats"},
		RadioButton{"animalselect", "dogs", false, false, "Dogs"},
	}

	MyPageVariables := PageVariables{
		PageTitle:        Title,
		PageRadioButtons: MyRadioButtons,
	}

	t, err := template.ParseFiles("select.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, MyPageVariables)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func UserSelected(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	youranimal := r.Form.Get("animalselect")

	Title := "Your preferred animal"
	MyPageVariables := PageVariables{
		PageTitle: Title,
		Answer:    youranimal,
	}

	t, err := template.ParseFiles("select.html")
	if err != nil {
		log.Print("template parsing error :", err)
	}

	err = t.Execute(w, MyPageVariables)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
}
