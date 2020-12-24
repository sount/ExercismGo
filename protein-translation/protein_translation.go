package protein

import "errors"

var (
	ErrStop =errors.New("Stop")
	ErrInvalidBase = errors.New("Invalid")
)

//FromCodon translates Codon to Protein
func FromCodon(Codon string )(string,error){

	switch Codon {
	case "AUG":
		return "Methionine",nil
	case "UUU","UUC":
		return "Phenylalanine",nil
	case "UUA","UUG":
		return "Leucine",nil
	case "UCU", "UCC", "UCA", "UCG":
		return  "Serine",nil
	case "UAU","UAC":
		return "Tyrosine",nil
	case "UGU","UGC":
		return "Cysteine",nil
	case "UGG":
		return "Tryptophan",nil
	case "UAA", "UAG", "UGA":
		return "",ErrStop
	}
	return "", ErrInvalidBase
}
//FromRNA gets an RNA string and uses FromCodon to translate it to its proteins
func FromRNA(RNA string) ([]string, error){
	var Proteins []string
	var error error
	for i:=0; i<len(RNA)-2; i+=3{
		Protein,err :=FromCodon(RNA[i:i+3])
		if err != nil{
			if err == ErrInvalidBase {
				error = err
			}
			break
		}
		Proteins = append(Proteins, Protein)

	}
	return Proteins, error
}



