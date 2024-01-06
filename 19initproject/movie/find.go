package movie

func FindName(imdbID string) string {
	switch imdbID {
	case "tt4154796":
		return "Avengers: endgame"
	case "tt1825683":
		return "Black panther"
	}
	return "not found"
}
