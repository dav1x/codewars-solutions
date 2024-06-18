package kata
// https://www.codewars.com/kata/55cbd4ba903825f7970000f5/train/go

func GetGrade(a,b,c int) rune {
  
    switch avg := (a + b + c) / 3;  {
      case avg >= 90:
        return 'A'
      case avg >= 80 && avg < 90:
        return 'B'
      case avg >= 70 && avg < 80:
        return 'C'
      case avg >= 60 && avg < 70:
        return 'D'
      default:
        return 'F'
    }
}
