package inmemorydatabase

import (
	"sort"
	"time"

	"github.com/purdoobahs/purdoobahs.com/internal/purdoobahs"
)

// `CurrentSection` returns all the Purdoobahs that are marching this academic year.
func (ps *PurdoobahService) CurrentSection() (*purdoobahs.Section, error) {
	currentSection := &purdoobahs.Section{
		StudentLeaders: make([]*purdoobahs.Purdoobah, 0),
		SuperSeniors:   make([]*purdoobahs.Purdoobah, 0),
		Seniors:        make([]*purdoobahs.Purdoobah, 0),
		Juniors:        make([]*purdoobahs.Purdoobah, 0),
		Sophomores:     make([]*purdoobahs.Purdoobah, 0),
		Freshmen:       make([]*purdoobahs.Purdoobah, 0),
	}

	currentAcademicYear := ps.currentAcademicYear()

	for _, p := range ps.purdoobahs {
		if !p.MarchedDuringYear(currentAcademicYear) {
			continue
		}

		if p.IsStudentLeaderInYear(currentAcademicYear) {
			currentSection.StudentLeaders = append(currentSection.StudentLeaders, p)
		} else if p.IsYear(purdoobahs.SuperSenior) {
			currentSection.SuperSeniors = append(currentSection.SuperSeniors, p)
		} else if p.IsYear(purdoobahs.Senior) {
			currentSection.Seniors = append(currentSection.Seniors, p)
		} else if p.IsYear(purdoobahs.Junior) {
			currentSection.Juniors = append(currentSection.Juniors, p)
		} else if p.IsYear(purdoobahs.Sophomore) {
			currentSection.Sophomores = append(currentSection.Sophomores, p)
		} else if p.IsYear(purdoobahs.Freshman) {
			currentSection.Freshmen = append(currentSection.Freshmen, p)
		}
	}

	sort.Sort(purdoobahs.ByName(currentSection.StudentLeaders))
	sort.Sort(purdoobahs.ByName(currentSection.SuperSeniors))
	sort.Sort(purdoobahs.ByName(currentSection.Seniors))
	sort.Sort(purdoobahs.ByName(currentSection.Juniors))
	sort.Sort(purdoobahs.ByName(currentSection.Sophomores))
	sort.Sort(purdoobahs.ByName(currentSection.Freshmen))

	return currentSection, nil
}

// `SectionByYear` returns all the Purdoobahs that marched during the given year.
func (ps *PurdoobahService) SectionByYear(targetYear int) ([]*purdoobahs.Purdoobah, error) {
	sectionByYear := make([]*purdoobahs.Purdoobah, 0)

	for _, p := range ps.purdoobahs {
		if p.MarchedDuringYear(targetYear) {
			sectionByYear = append(sectionByYear, p)
		}
	}

	return sectionByYear, nil
}

// `currentAcademicYear` returns the value of the current academic year.
// e.g. if the academic year is "Fall 2020 -> Spring 2021", it will return "2020"
func (ps *PurdoobahService) currentAcademicYear() int {
	t := time.Now()
	switch t.Month() {
	case time.January,
		time.February,
		time.March,
		time.April,
		time.May,
		time.June,
		time.July,
		time.August:
		return t.Year() - 1
	default:
		return t.Year()
	}
}
