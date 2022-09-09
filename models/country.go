package models

type Country struct {
	Name     string `json:"name"`
	Capital  string `json:"capital"`
	Language string `json:"language"`
}

type PublicCountry struct {
	Name     string `json:"name"`
	Capital  string `json:"capital"`
	Language []struct {
		A         string `json:"iso639_1"`
		B         string `json:"iso639_2"`
		Name      string `json:"name"`
		NativName string `json:"nativeName"`
	} `json:"languages"`
}
