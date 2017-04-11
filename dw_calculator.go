package bodyfat

import (
	"math"
)

// DWCalculator defines the 'Durnin and Womersly'
// body-fat calculator type
type DWCalculator struct{}

// DensityCoeff defines the density coefficients used to
// calculate the body-fat percentage for a particular
// gender and age-range.
type DensityCoeff struct {
	A float64
	B float64
}

// NewDWCalculator returns a new 'Durnin and Womersly'
// body-fat calculator
func NewDWCalculator() *DWCalculator {
	return &DWCalculator{}
}

// Name returns the name of this body-fat calculator
func (c *DWCalculator) Name() string {
	return "Durnin and Womersley"
}

// Calculate returns the body-fat percentage based
// on the given caliper measurements.
func (c *DWCalculator) Calculate(p *Person, m *Measurements) float64 {
	// get the density coefficients for the given person attributes
	coeff := getDensityCoeff(p)

	// sum the measurements (in millimeters)
	sum := m.Waist + m.FrontArm + m.BackArm + m.Back

	// calculate the body-fat percentage
	d := coeff.A - (coeff.B * math.Log10(sum))
	bf := (495 / d) - 450

	// return the body-fat percentage
	return bf
}

// getDensityCoeff returns the 'Durnin and Womersly' density
// coefficients for the given person attributes
func getDensityCoeff(p *Person) DensityCoeff {
	var a, b float64

	// determine the density coefficients
	age := p.Age
	if p.Gender == GenderMale {
		if age < 17 {
			a = 1.1533
			b = 0.0643
		} else if (age >= 17) && (age <= 19) {
			a = 1.1620
			b = 0.0630
		} else if (age >= 20) && (age <= 29) {
			a = 1.1631
			b = 0.0632
		} else if (age >= 30) && (age <= 39) {
			a = 1.1422
			b = 0.0544
		} else if (age >= 40) && (age <= 49) {
			a = 1.1620
			b = 0.0700
		} else {
			a = 1.1715
			b = 0.0779
		}
	} else {
		if age < 17 {
			a = 1.1369
			b = 0.0598
		} else if (age >= 17) && (age <= 19) {
			a = 1.1549
			b = 0.0678
		} else if (age >= 20) && (age <= 29) {
			a = 1.1599
			b = 0.0717
		} else if (age >= 30) && (age <= 39) {
			a = 1.1423
			b = 0.0632
		} else if (age >= 40) && (age <= 49) {
			a = 1.1333
			b = 0.0612
		} else {
			a = 1.1339
			b = 0.0645
		}
	}

	// return the coefficients
	return DensityCoeff{
		A: a,
		B: b,
	}
}
