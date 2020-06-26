// Ex: create a tool that converts between values from std out and args.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"ex2/conv"
)

func main() {
	var input []string

	if len(os.Args) > 1 {
		input = os.Args[1:]
	} else {
		stdIn := bufio.NewScanner(os.Stdin)
		for stdIn.Scan() {
			input = append(input, stdIn.Text())
		}
	}
	for _, arg := range input {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		m := conv.Meters(t)
		f := conv.Feet(t)
		c := conv.Celsius(t)
		h := conv.Fahrenheit(t)
		k := conv.Kilograms(t)
		p := conv.Pounds(t)

		fmt.Printf("%s = %s, %s = %s\n",
			m, m.ToFeet(), f, f.ToMeters())
		fmt.Printf("%s = %s, %s = %s\n",
			c, c.ToFahrenheit(), h, h.ToCelsius())
		fmt.Printf("%s = %s, %s = %s\n",
			k, k.ToPounds(), p, p.ToKilograms())
	}
}

//!-
