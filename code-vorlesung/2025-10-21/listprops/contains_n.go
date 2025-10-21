package listprops

// ContainsN liefert true, falls die Liste l
// den String x mindestens n mal enthält.
func ContainsN(l []string, x string, n int) bool {
	count :=0 
	for i := 0; i < len(l); i++ {
		
		if l[i] == x {
			count++
			if count == n{
				return true
			}
			
		}
	}
	return false
}
