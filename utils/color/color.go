package color

//Color is a main color schem
type Color struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Hex         string `json:"hex"`
}

//GetColors for export colors
func GetColors() []Color {
	var Colors = []Color{
		Color{"black", "Чёрный", "#000"},
		Color{"red", "красный", "#f00"},
		Color{"yellow", "жёлтый", "#ff0"},
	}
	return Colors
}
