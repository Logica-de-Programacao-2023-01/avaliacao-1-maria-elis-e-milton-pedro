package q2
)
func ProblemsSolved(answers [][3]bool) int {
	contagem := 0
	for i := 0; i < len(answers); i++ {
		certezas := 0
		for j := 0; j < 3; j++ {
			if answers[i][j] {
				certezas++
			}
		}
		if certezas >= 2 {
			contagem++
		}
	}
	return contagem
}
