
package proverb

// Proverb returns an array of strings and takes different words as input
func Proverb(rhyme []string) []string {

	k := []string{}
	for i:=0; i<len(rhyme);i++{
		if i == len(rhyme)-1{
			k= append(k,"And all for the want of a " + rhyme[0] + ".")
		}else{
		    k= append(k, "For want of a "+ rhyme[i]+" the "+ rhyme[i+1] + " was lost.")
		}
	}
	return k
}
