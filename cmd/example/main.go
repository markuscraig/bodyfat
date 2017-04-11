package main

import (
	"fmt"

	"github.com/fitquick/bodyfat"
)

func main() {
	// create a 'durnin and wormersly' body-fat calculator
	calc := bodyfat.NewDWCalculator()

	// define the person's attributes
	p := &bodyfat.Person{
		Gender: bodyfat.GenderMale,
		Age:    42,
	}

	// define the body measurements
	m := &bodyfat.Measurements{
		FrontArm: 4.0,
		BackArm:  6.0,
		Waist:    6.0,
		Back:     16.0,
	}

	// calculate the body-fat percentage
	bf := calc.Calculate(p, m)

	// dump the results
	fmt.Printf("BODY-FAT %% = %f\n", bf)
}
