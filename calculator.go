package bodyfat

// Person provides information about the person for which
// we will calculate the body-fat percentage.
type Person struct {
	Gender GenderEnum
	Age    int
}

// Measurements provides the body-fat caliper measurements
// in millimeters.
type Measurements struct {
	FrontArm float64
	BackArm  float64
	Waist    float64
	Back     float64
}

// Calculator defines the abstract calculator interface
type Calculator interface {
	Name() string
	Calculate(p *Person, m *Measurements) float64
}
