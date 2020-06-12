// Package mfy converts meters, feet, and yards for a numeric argument
package main

import "fmt"

type Meters float64
type Feet float64

func (m Meters) String() string { return fmt.Sprintf("%g meters", m) }
func (f Feet) String() string   { return fmt.Sprintf("%g feet", f) }

// MToF converts Meters to Feet
func MToF(m Meters) Feet { return Feet(m / 0.3048) }

// FToM converts Feet to Meters
func FToM(f Feet) Meters { return Meters(f * 0.3048) }
