package listprops

// ContainsNRow liefert true, falls die Liste l
// den String x n mal hintereinander enthält.
func ContainsNRow(l []string, x string, n int) bool {
	// TODO
	count :=0 
	for i := 0; i < len(l); i++ {
		
		if l[i] == x {
			count++
			if count > n{
				return true
			}
			
		}
	}
	return false
}
