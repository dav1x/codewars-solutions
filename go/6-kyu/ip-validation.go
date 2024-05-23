package kata
import (
  "regexp"
)

func Is_valid_ip(ip string) bool {

  // ^((25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)(\.|$)){4}\b
  ipv4_regex := `^(((25[0-5]|2[0-4][0-9]|[1]?[0-9][0-9]?)(\.|$)){4})`  
  match, _ := regexp.MatchString(ipv4_regex, ip)

  
  if ! match { return false } 	  
  return true 
  
  /* Submitted solution was WAY easier

  func Is_valid_ip(ip string) bool {
	res := net.ParseIP(ip)
	if res == nil {
		return false
	}
	return true
  }
  */
}
