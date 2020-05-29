package ex1

type Kelvin float64

func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

func KToC(k Kelvin) Celsius { return Celsius(k) + AbsoluteZeroC }
