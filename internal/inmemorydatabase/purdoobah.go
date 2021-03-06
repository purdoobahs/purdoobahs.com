package inmemorydatabase

import (
	"fmt"
	"sort"

	"github.com/purdoobahs/purdoobahs.com/internal/purdoobahs"
)

// `All` returns every single Purdoobah.
func (ps *PurdoobahService) All() ([]*purdoobahs.Purdoobah, error) {
	allPurdoobahs := make([]*purdoobahs.Purdoobah, 0, len(ps.purdoobahs))

	for _, v := range ps.purdoobahs {
		allPurdoobahs = append(allPurdoobahs, v)
	}

	sort.Sort(purdoobahs.ByName(allPurdoobahs))
	return allPurdoobahs, nil
}

// `ByName` returns a single Purdoobah by their nickname.
func (ps *PurdoobahService) ByName(name string) (*purdoobahs.Purdoobah, error) {
	if purdoobah, ok := ps.purdoobahs[name]; ok {
		return purdoobah, nil
	}

	return &purdoobahs.Purdoobah{}, fmt.Errorf("no Purdoobah exists with that name")
}
