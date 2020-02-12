package space

//Planet is the name of the planet.
type Planet string

const (
	mercuryConst float64 = 0.2408467
	venusConst   float64 = 0.61519726
	earthConst   float64 = 1
	marsConst    float64 = 1.8808158
	jupiterConst float64 = 11.862615
	saturnConst  float64 = 29.447498
	uranusConst  float64 = 84.016846
	neptuneConst float64 = 164.79132
)

//Age calculates the age of a person on a planet given the number of seconds for which they have been alive and the planet they are on.
func Age(sec float64, planet Planet) float64 {
	sec = sec / 31557600
	switch planet {
	case "Mercury":
		return sec / mercuryConst
	case "Venus":
		return sec / venusConst
	case "Earth":
		return sec / earthConst
	case "Mars":
		return sec / marsConst
	case "Jupiter":
		return sec / jupiterConst
	case "Saturn":
		return sec / saturnConst
	case "Uranus":
		return sec / uranusConst
	case "Neptune":
		return sec / neptuneConst
	}
	panic("The input planet is undefined!")
}
