package space

type Planet string


func Age(seconds float64, planet Planet) float64{
	earth_age := seconds/31557600
	switch planet {
	case "Earth":
		return earth_age
	case "Mercury":
		return earth_age / 0.2408467
	case "Venus":
		return earth_age / 0.61519726
	case "Mars":
		return earth_age / 1.8808158
	case "Neptune":
		return earth_age / 164.79132
	case "Uranus":
		return earth_age / 84.016846
	case "Jupiter":
		return earth_age / 11.862615
	case "Saturn":
		return earth_age / 29.447498
	}
	return 0
}