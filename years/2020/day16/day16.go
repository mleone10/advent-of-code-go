package day16

import (
	"strings"

	"github.com/mleone10/advent-of-code-go/internal/set"
	"github.com/mleone10/advent-of-code-go/internal/slice"
)

type field struct {
	name       string
	rngA, rngB span
}

type span struct {
	min, max int
}

type ticket []ticketVal

type ticketVal struct {
	val                 int
	candidateFieldNames set.Set[string]
}

type Day16 struct {
	fields        []field
	myTicket      ticket
	nearbyTickets []ticket
}

func (d *Day16) AddField(name string, min1, max1, min2, max2 int) {
	if d.fields == nil {
		d.fields = []field{}
	}

	d.fields = append(d.fields, field{
		name,
		span{min1, max1},
		span{min2, max2},
	})
}

func (d *Day16) AddMyTicket(fields ...int) {
	d.myTicket = initTicket(fields...)
}

func initTicket(fields ...int) []ticketVal {
	ticket := []ticketVal{}
	for _, f := range fields {
		ticket = append(ticket, ticketVal{val: f})
	}
	return ticket
}

func (d *Day16) AddNearbyTicket(fields ...int) {
	if d.nearbyTickets == nil {
		d.nearbyTickets = []ticket{}
	}
	d.nearbyTickets = append(d.nearbyTickets, initTicket(fields...))
}

func (d Day16) CalcErrorRate() int {
	errRate := 0
	for _, nt := range d.nearbyTickets {
		errRate += slice.Sum(findInvalidValues(d, nt))
	}
	return errRate
}

func findInvalidValues(d Day16, t ticket) []int {
	vs := []int{}
	for _, f := range t {
		if len(findValidFieldNames(d, f.val)) == 0 {
			vs = append(vs, f.val)
		}
	}
	return vs
}

func findValidFieldNames(d Day16, f int) []string {
	fns := []string{}
	for _, field := range d.fields {
		if (f >= field.rngA.min && f <= field.rngA.max) || (f >= field.rngB.min && f <= field.rngB.max) {
			fns = append(fns, field.name)
		}
	}
	return fns
}

func (d Day16) CalcDepartureProduct() int {
	d.populateFieldCandidates()
	d.discardInvalidTickets()
	fm := buildFieldMap(d)

	// With the completed field map in hand, accumulate the product of the six "departure" fields on my ticket.
	departureProduct := 1
	for i, name := range fm {
		if strings.Contains(name, "departure") {
			departureProduct *= d.myTicket[i].val
		}
	}

	return departureProduct
}

func (d Day16) populateFieldCandidates() {
	for _, nt := range d.nearbyTickets {
		populateCandidateFieldNames(nt, d)
	}
}

func populateCandidateFieldNames(t ticket, d Day16) {
	for i, f := range t {
		t[i].candidateFieldNames = set.From(findValidFieldNames(d, f.val)...)
	}
}

func (d *Day16) discardInvalidTickets() {
	nts := []ticket{}
	for _, nt := range d.nearbyTickets {
		valid := true
		for _, f := range nt {
			if len(f.candidateFieldNames) == 0 {
				valid = false
			}
		}
		if valid {
			nts = append(nts, nt)
		}
	}
	d.nearbyTickets = nts
}

func buildFieldMap(d Day16) map[int]string {
	cfs := buildCandidateFieldSets(d)
	return simplifyFieldSets(cfs)
}

func buildCandidateFieldSets(d Day16) map[int]set.Set[string] {
	fm := map[int]set.Set[string]{}
	for i := range len(d.fields) {
		ss := []set.Set[string]{}
		for _, nt := range d.nearbyTickets {
			ss = append(ss, nt[i].candidateFieldNames)
		}
		fm[i] = set.Intersection(ss...)
	}
	return fm
}

func simplifyFieldSets(fs map[int]set.Set[string]) map[int]string {
	found := map[int]string{}

	if len(fs) == 0 {
		return found
	}

	// Find all field sets with only one element.
	for i, f := range fs {
		if f.Size() == 1 {
			found[i] = f.Slice()[0]
			delete(fs, i)
		}
	}

	// Remove all completed fields from the remaining field sets.
	for _, e := range found {
		for _, f := range fs {
			f.Remove(e)
		}
	}

	// Recurse through the remaining field sets.
	for i, f := range simplifyFieldSets(fs) {
		found[i] = f
	}

	return found
}
