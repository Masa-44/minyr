package yr

import (
    	"bufio"
    	"fmt"
    	"os"
	"log"
    	"strings"
    	"strconv"
    	"github.com/Masa-44/funtemps/conv"
)

//Denne pakken blir brukt for aa ha selve funksjonene i. I main.go vil vi kalle paa disse funksjonene.


func ConvTemperature() {

	//Setter input og output filnavn.

	//inputFilename := "kjevik-temp-celsius-20220318-20230318.csv"
	outputFilename := "kjevik-temp-fahr-20220318-20230318.csv"

	//Sjekker om kjevik fahr versjon av filen allerede eksisterer.

	if _, err := os.Stat(outputFilename); err == nil {
		//Om filen eksisterer spor den om vi vil genere filen paa nytt.
		fmt.Printf("Fil '%s' eksisterer allerede. Vil du generere filen paa nytt? (j/n): ", outputFilename)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			answer := scanner.Text()
			if strings.ToLower(answer) == "j" || strings.ToLower(answer) == "ja" {
				//Om brukeren vil genere filen aa nytt gaar den ut av loopen.
				break
			} else if strings.ToLower(answer) == "n" || strings.ToLower(answer) == "nei" {
				//Om brukeren ikke onsker aa genere filen paa nytt returner den og gaar ut av funksjonen.
				fmt.Println("Avslutter...")
				return
			} else {
				//Om brukeren ikke gir gyldig input i  Scanneren spor den paa nytt.
				fmt.Print("Invalid answer. Do you want to regenerate the file? (j/n): ")
			}
		}
	}


	inputFile := openInputFile()
	defer inputFile.Close()

	outputFile, err := createOutputFile()
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	//De neste to linjene er for aa skippe linje en i kjevik filen for loop begynner.
	if scanner.Scan() {
	}

	for scanner.Scan() {
    		line := scanner.Text()
    		fields := strings.Split(line, ";")
    		celsius, err := strconv.ParseFloat(fields[3], 64)

		if err != nil {
        	log.Fatal(err)
    		}

    		fahrenheit := conv.CelcsiusToFahrenheit(celsius)
    		fields[3] = fmt.Sprintf("%.2f", fahrenheit)
    		line = strings.Join(fields, ";")
    		fmt.Fprintln(outputFile, line)
		}

		if err := scanner.Err(); err != nil {
    		log.Fatal(err)
	}



}


func openInputFile() *os.File {
	file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
	log.Fatal(err)
	}
return file
}

func createOutputFile() (*os.File, error) {
	outputFilePath := "kjevik-temp-fahr-20220318-20230318.csv"
	if _, err := os.Stat(outputFilePath); err == nil {
		fmt.Printf("File %s already exists. Deleting...\n", outputFilePath)
		err := os.Remove(outputFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not delete file: %v", err)
	}
}

outputFile, err := os.Create(outputFilePath)
if err != nil {
	return nil, fmt.Errorf("could not create file: %v", err)
}
return outputFile, nil
}

