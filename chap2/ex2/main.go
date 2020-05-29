package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bravmi/gopl/chap2/tempconv"
)

type Feet float64
type Meters float64

func FtToM(ft Feet) Meters { return Meters(ft) * 0.3048 }
func MToFt(m Meters) Feet  { return Feet(m) / 0.3048 }

func (ft Feet) String() string  { return fmt.Sprintf("%.2f ft", ft) }
func (m Meters) String() string { return fmt.Sprintf("%.2f m", m) }

type Pounds float64
type Kilograms float64

func LbToKg(lb Pounds) Kilograms { return Kilograms(lb) * 0.454 }
func KgToLb(kg Kilograms) Pounds { return Pounds(kg) / 0.454 }

func (lb Pounds) String() string    { return fmt.Sprintf("%.2f lbs", lb) }
func (kg Kilograms) String() string { return fmt.Sprintf("%.2f lbs", kg) }

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))

		ft := Feet(t)
		m := Meters(t)
		fmt.Printf("%s = %s, %s = %s\n",
			ft, FtToM(ft), m, MToFt(m))

		lb := Pounds(t)
		kg := Kilograms(t)
		fmt.Printf("%s = %s, %s = %s\n",
			lb, LbToKg(lb), kg, KgToLb(kg))

		fmt.Println()
	}
}
