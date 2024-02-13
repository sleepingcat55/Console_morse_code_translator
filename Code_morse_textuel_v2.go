package main

import (
	"fmt"
	"strings"
	"regexp"
	"os"
	"bufio"
)

var (
	alphabet = [36]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	at_morse = [36]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--..", ".----", "..---", "...--", "....-", ".....", "-....", "--...", "---..", "----.", "-----"}
)
var inputReader *bufio.Reader
var inputt string
var err error

func codage_morse(mot_a_coder string) string {
	var resu string
	for i:=0; i<len(mot_a_coder); i++ {
		for j:=0; j<len(alphabet); j++ {
			if string(mot_a_coder[i]) == alphabet[j] {
			//	fmt.Printf("%d", j)
				resu += at_morse[j]	// lettre trouvé, on concatene dans resu toutes les lettres les unes a la suite des autres
				// on ajoute / pour separer chaque lettres d'un mot sinon impossible de faire une partie decodage vu qu'il n'y a pas de meme quantité de .-
				resu += "/"	
			} else if string(mot_a_coder[i]) == " " {
				resu += "_"	// permet de simuler un espace
				// on va essayer d'eviter les 26 espaces
				previous_lettre := len(resu)-1
				if string(resu[previous_lettre]) == "_" {
					break
				}
			}
			
		}
	}
	//fmt.Println()
	return resu
	// fmt.Printf("%s\n", mot_coder)
}

func decodage_morse(mot_a_decoder string) string {
	var resu string
	var tempo string
	var entmpo string
	for i:=0; i<len(mot_a_decoder); i++ {
		// fmt.Printf("%c\n", mot_a_decoder[i])
		tempo += string(mot_a_decoder[i])
		
		if string(mot_a_decoder[i]) == "/" {
			re, _ := regexp.Compile("/")
			entmpo = re.ReplaceAllString(tempo, "")
			// fmt.Printf("valeur de entmpo dans le if : %s\n", entmpo)

			for j:=0; j<len(at_morse); j++ {
				if entmpo == at_morse[j] {
					//	fmt.Printf("%d", j)
					resu += alphabet[j]	// lettre trouvé, on concatene dans resu toutes les lettres les unes a la suite des autres
					entmpo = ""
					tempo = ""
				}
			}
		} else if string(mot_a_decoder[i]) == "_" {
			// fmt.Printf("Dans la 2eme condition : %c\n", mot_a_decoder[i])
			tempo = ""
			resu += " "
		} 
	}
	return resu
}

func main() {
	// entree := "abcd zsx Remi 12 Gillou 34 90"	// marche
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Entrez la phrase que vous voulez coder en Morse : ")
	inputt, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("Entree : %s\n", inputt)
	}

	// mise en forme
	mot_to_coder := strings.ToLower(inputt)

	//partie codage fonctionnelle
	codage := codage_morse(mot_to_coder)
	fmt.Printf("Codage du messsage : %s\n", codage)

	// partie decodage fonctionnelle
	decodage := decodage_morse(codage)
	fmt.Printf("Decodage du message : %s\n", decodage)
	

	


}