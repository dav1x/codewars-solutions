package kata
import "strings"

func DNAStrand(dna string) string {
  var output string;
  
  if dna != "" {
      for _, letter := range strings.ToUpper(dna){
        switch letter {
          case 'A':
            output += "T"
          case 'T':
            output += "A"
          case 'C':
            output += "G"
          case 'G':
            output += "C"
        }
        
      }
  }
  return string(output)
}
