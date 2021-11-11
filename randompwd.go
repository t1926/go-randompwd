package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func lataaTiedosto() []byte {
	var ladattavaTiedosto string

	fmt.Printf("File to load (in the file every word should be on its own row)? ")
	fmt.Scanf("%s", &ladattavaTiedosto)

	tiedostonSisalto, virhe := ioutil.ReadFile(ladattavaTiedosto)

	if virhe != nil {
		log.Fatal(virhe)
	}

	return tiedostonSisalto
}

func arvoLuku(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	arvottuLuku := rand.Intn(max-min+1) + min
	return arvottuLuku
}

func main() {
	tiedostonSisalto := lataaTiedosto()
	kaikkiSanat := strings.Split(string(tiedostonSisalto), "\n")
	sanojaYhteensa := len(kaikkiSanat) - 1
	sanojenLukumaara := 1
	//fmt.Println(sanojaYhteensa)

	for sanojenLukumaara != 0 {
		var lause string
		var valiMerkki string
		var salasanalauseidenLukumaara int

		fmt.Printf("How many words (0 = quit)? ")
		fmt.Scanf("%d", &sanojenLukumaara)

		if sanojenLukumaara > 0 {
			fmt.Printf("Which character between words? ")
			fmt.Scanf("%s", &valiMerkki)
			fmt.Printf("How many passphareses? ")
			fmt.Scanf("%d", &salasanalauseidenLukumaara)
		}

		for a := 0; a < salasanalauseidenLukumaara; a++ {
			for b := 0; b < sanojenLukumaara; b++ {
				apuSana := ""
				arvottuSana := kaikkiSanat[arvoLuku(1, sanojaYhteensa)]
				isoKirjain := arvoLuku(1, len(arvottuSana))

				for c := 0; c < len(arvottuSana); c++ {
					if c == isoKirjain-1 {
						apuSana += strings.ToUpper(string(arvottuSana[c]))
					} else {
						apuSana += string(arvottuSana[c])
					}
				}

				lause += apuSana

				if b < sanojenLukumaara-1 {
					lause += valiMerkki // lisää annettu välimerkki, jos ei ole viimeinen sana
				} else {
					lause += strconv.Itoa(arvoLuku(0, 9)) + "\n" // jos on viimeinen sana, lisää viimeiseksi randomluku väliltä 0-9
				}
			}
		}

		fmt.Printf("\n")
		fmt.Println(lause)
		fmt.Printf("\n")
	}
}
