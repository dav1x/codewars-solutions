package kata
import "fmt"


func SeriesSum(n int) string {
 
   var sum float64 = 0;
   
   for i:=0; i<=n-1; i++ {
    sum += float64(1) / float64(1 + i * 3)
  }
  
  return fmt.Sprintf("%.2f", sum)
}

