// usage: go run main.go 1f 1ft 1lb
package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/bravmi/gopl/chap2/tempconv"
)

type Feet float64
type Meters float64

func FtToM(ft Feet) Meters { return Meters(ft) * 0.3048 }
func MToFt(m Meters) Feet  { return Feet(m) / 0.3048 }

func (ft Feet) String() string  { return fmt.Sprintf("%.2fft", ft) }
func (m Meters) String() string { return fmt.Sprintf("%.2fm", m) }

type Pounds float64
type Kilograms float64

func LbToKg(lb Pounds) Kilograms { return Kilograms(lb) * 0.454 }
func KgToLb(kg Kilograms) Pounds { return Pounds(kg) / 0.454 }

func (lb Pounds) String() string    { return fmt.Sprintf("%.2flb", lb) }
func (kg Kilograms) String() string { return fmt.Sprintf("%.2fkg", kg) }

type Measure struct {
	n    float64
	unit string
}

func parseArg(s string) (Measure, error) {
	re := regexp.MustCompile(`([\d.]+)(\w+)`)
	match := re.FindStringSubmatch(s)
	if match == nil {
		return Measure{}, fmt.Errorf("expected <number><unit>, got %q", s)
	}
	n, err := strconv.ParseFloat(match[1], 64)
	if err != nil {
		return Measure{}, fmt.Errorf("%v isn't a number", match[1])
	}
	unit := strings.ToLower(match[2])
	return Measure{n, unit}, nil
}

func printMeasure(measure Measure) error {
	switch measure.unit {
	case "f":
		f := tempconv.Fahrenheit(measure.n)
		c := tempconv.FToC(f)
		fmt.Printf("%s = %s\n", f, c)
	case "ft":
		ft := Feet(measure.n)
		m := FtToM(ft)
		fmt.Printf("%s = %s\n", ft, m)
	case "lb":
		lb := Pounds(measure.n)
		kg := LbToKg(lb)
		fmt.Printf("%s = %s\n", lb, kg)
	default:
		return fmt.Errorf("unexpected unit %v", measure.unit)
	}
	return nil
}

func main() {
	for _, arg := range os.Args[1:] {
		measure, err := parseArg(arg)
		if err != nil {
			log.Fatal(err.Error())
		}
		err = printMeasure(measure)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
