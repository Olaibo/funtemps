package main

import (
	"flag"
	"fmt"

	"github.com/RobertRegen/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var cels float64
var kelv float64
var out string
var funfacts string

var Svar float64

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	// Du må selv definere flag-variablene for "C" og "K"
	flag.Float64Var(&cels, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelv, "K", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala som skal brukes når funfacts skal vises

}

func main() {

	flag.Parse()

	/**
	    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
	    pakkene implementeres.

	    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
	    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
	    hvor mange flagg og argumenter er spesifisert på kommandolinje.

	        fmt.Println("len(flag.Args())", len(flag.Args()))
			    fmt.Println("flag.NFlag()", flag.NFlag())

	    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
	    brukes for å utelukke ugyldige kombinasjoner:
	    -F, -C, -K kan ikke brukes samtidig
	    disse tre kan brukes med -out, men ikke med -funfacts
	    -funfacts kan brukes kun med -t
	    ...
	    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
	    implementert i flag-pakken og at den vil skrive ut "Usage" med
	    beskrivelsene av flagg-variablene, som angitt i parameter fire til
	    funksjonene Float64Var og StringVar
	*/

	// Her er noen eksempler du kan bruke i den manuelle testingen
	// fmt.Println(fahr, out, funfacts)

	// fmt.Println("len(flag.Args())", len(flag.Args()))
	// fmt.Println("flag.NFlag()", flag.NFlag())

	// fmt.Println(isFlagPassed("out"))

	// Eksempel på enkel logikk

	erlik := "="
	f := "°F"
	c := "°C"
	k := "°K"

	// FahrenheitToCelsius
	if out == "C" && isFlagPassed("F") {
		// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
		// skal returnere °C
		//fmt.Println("0°F er -17.78°C")

		Svar = conv.FarhenheitToCelsius(fahr)

		fmt.Printf("%.9g %s %s", fahr, f, erlik)
		if Svar == float64(int(Svar)) {
			fmt.Printf("%d %s\n", int(Svar), c) //printer om Svar == heltall
		} else {
			fmt.Printf("%.2f %s\n", Svar, c) //printer om Svar == decimaltall
		}

	}

	//CelsiusToFahrenheit
	if out == "F" && isFlagPassed("C") {
		Svar = conv.CelsiusToFahrenheit(cels)

		fmt.Printf("%.9g %s %s", cels, c, erlik)
		if Svar == float64(int(Svar)) {
			fmt.Printf("%d %s\n ", int(Svar), f) //printer om Svar == heltall
		} else {
			fmt.Printf("%.2f %s\n ", Svar, f) //printer om Svar == decimaltall
		}
	}

	//FahrenheitToKelvin
	if out == "K" && isFlagPassed("F") {
		Svar = conv.FahrenheitToKelvin(fahr)

		fmt.Printf("%.9g %s %s", fahr, f, erlik)
		if Svar == float64(int(Svar)) {
			fmt.Printf("%d %s\n", int(Svar), k) //printer om Svar == heltall
		} else {
			fmt.Printf("%.2f %s\n", Svar, k) //printer om Svar == decimaltall
		}
	}

	//KelvinToFahrenheit
	if out == "F" && isFlagPassed("K") {
		Svar = conv.KelvinToFahrenheit(kelv)

		fmt.Printf("%.9g %s %s", kelv, k, erlik)
		if Svar == float64(int(Svar)) {
			fmt.Printf("%d %s\n", int(Svar), f) //printer om Svar == heltall
		} else {
			fmt.Printf("%.2f %s\n", Svar, f) //printer om Svar == decimaltall
		}
	}

	//CelsiusToKelvin
	if out == "K" && isFlagPassed("C") {
		Svar = conv.CelsiusToKelvin(cels)

		fmt.Printf("%.9g %s %s", cels, c, erlik)
		if Svar == float64(int(Svar)) {
			fmt.Printf("%d %s\n", int(Svar), k) //printer om Svar == heltall
		} else {
			fmt.Printf("%.2f %s\n", Svar, k) //printer om Svar == decimaltall
		}
	}

	//KelvinToCelsius
	if out == "C" && isFlagPassed("K") {
		Svar = conv.KelvinToCelsius(kelv)

		fmt.Printf("%.9g %s %s", kelv, k, erlik)
		if Svar == float64(int(Svar)) {
			fmt.Printf("%d %s\n", int(Svar), c) //printer om Svar == heltall
		} else {
			fmt.Printf("%.2f %s\n", Svar, c) //printer om Svar == decimaltall
		}
	}

}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
