package excuses

import (
	"fmt"
	"math/rand"
	"strings"
)

func Pet() string {
	sb := strings.Builder{}
	w := rand.Intn(len(pwho))
	p := rand.Intn(len(pet))
	d := rand.Intn(len(pdid))
	t := rand.Intn(len(pwhat))

	sb.WriteString(fmt.Sprintf("My %s's %s %s my %s.", pwho[w], pet[p], pdid[d], pwhat[t]))
	return sb.String()
}

func PossiblePets() int {
	return len(pwho) * len(pet) * len(pdid) * len(pwhat)
}
