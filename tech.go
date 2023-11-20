package excuses

import (
	"fmt"
	"math/rand"
	"strings"
)

func Tech() string {
	sb := strings.Builder{}
	d := rand.Intn(len(tdid))

	sb.WriteString(fmt.Sprintf("The %s %s %s the %s %s.", TechWHo(), TechWhat(), tdid[d], TechWHo(), TechWhat()))
	return sb.String()
}

func PossibleTech() int {
	return (len(twho) * 2) * len(tdid) * (len(twhat) * 2)
}

func TechWHo() string {
	w := rand.Intn(len(twho))
	return fmt.Sprintf("%s's", twho[w])
}

func TechWhat() string {
	l := rand.Intn(len(twhat))
	return twhat[l]
}
