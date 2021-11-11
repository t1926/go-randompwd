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

/*
        1. Rakenna sama kokonaisuus funktioista
                1.1 Tiedoston lataus
                1.2 Kuinka monta sanaa
                1.3 Mikä välimerkki (vai arvotaanko !#-)
                1.4 Kuinka monta salasanalausetta
                1.5 Arvo ison kirjaimen paikka jokaisessa sanassa
                1.6 Arvo numeron paikka koko lauseessa
        2. Arvo jokaiseen sanaan ison kirjaimen paikka
        3. Arvo numero johonkin kohtaan koko lausetta --> huomioi välimerkki
                eli jos välimerkki, niin numero ennen tai jälkeen
        4. Pitäisikö suolata eli tehdä sanoihin kirjoitusvirheitä vaihtamalla
                kahden kirjaimen paikkaa
*/


func main() {
        var ladattavaTiedosto string
        fmt.Printf("File to load (in the file every word should be on its own row)? ")
        fmt.Scanf("%s", &ladattavaTiedosto)
        tiedostonSisalto, virhe := ioutil.ReadFile(ladattavaTiedosto)
        if virhe != nil {
                log.Fatal(virhe)
        }

        kaikkiSanat := strings.Split(string(tiedostonSisalto), "\n")
        sanojaYhteensa := len(kaikkiSanat) - 1
        fmt.Println(sanojaYhteensa)

        sanojenLukumaara := 1
        valiMerkki := ""
        salasanojenLukumaara := 0

        for sanojenLukumaara != 0 {
                fmt.Printf("How many words (0 = quit)? ")
                fmt.Scanf("%d", &sanojenLukumaara) // bufio.scanner?

                if sanojenLukumaara > 0 {
                        fmt.Printf("Which character between words? ")
                        fmt.Scanf("%s", &valiMerkki)

                        fmt.Printf("How many passphareses? ")
                        fmt.Scanf("%d", &salasanojenLukumaara)
                }

                var lause string

                for a := 0; a < salasanojenLukumaara; a++ {

                        for i := 0; i < sanojenLukumaara; i++ {
                                rand.Seed(time.Now().UnixNano())
                                min := 0
                                max := sanojaYhteensa
                                arvottuLuku := rand.Intn(max-min+1) + min

                                if i%2 != 0 {
                                        lause += strings.Title(kaikkiSanat[arvottuLuku])
                                } else {
                                        var apusana string
                                        sana := kaikkiSanat[arvottuLuku]
                                        sananPituus := len(sana)

                                        for i := 0; i < sananPituus; i++ {
                                                if i < sananPituus-1 {
                                                        apusana += string(sana[i])
                                                } else {
                                                        apusana += strings.Title(string(sana[i]))
                                                }
                                        }
                                        lause += apusana
                                }

                                if i < sanojenLukumaara-1 {
                                        lause += valiMerkki // "-"
                                } else {
                                        rand.Seed(time.Now().UnixNano())
                                        min := 1
                                        max := 9
                                        lause += strconv.Itoa(rand.Intn(max-min+1)+min) + "\n"
                                }
                        }
                }

                fmt.Printf("\n")
                fmt.Println(lause)
                fmt.Printf("\n")
        }
}