package purdoobahs

import "strings"

// Year defines a Purdoobahs year in school
type Year string

const (
	Alumni      Year = "alumni"
	SuperSenior      = "super-senior"
	Senior           = "senior"
	Junior           = "junior"
	Sophomore        = "sophomore"
	Freshman         = "freshman"
)

// ByName sorts Purdoobahs by name
type ByName []*Purdoobah

func (p ByName) Len() int {
	return len(p)
}

func (p ByName) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ByName) Less(i, j int) bool {
	return strings.Compare(p[i].Name, p[j].Name) < 0
}

type Purdoobah struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	BirthCertificateName struct {
		First  string `json:"first"`
		Middle string `json:"middle,omitempty"`
		Last   string `json:"last"`
	} `json:"birth_certificate_name"`
	Emoji string `json:"emoji"`

	Marching struct {
		YearsMarched []int  `json:"years_marched"`
		Shoutout     string `json:"shoutout,omitempty"`
	} `json:"marching"`

	Education struct {
		Major string `json:"major"`
		Minor string `json:"minor,omitempty"`
		Year  Year   `json:"year"`
	} `json:"education"`

	Hometown struct {
		City  string `json:"city"`
		State string `json:"state"`
	} `json:"hometown"`

	Alumni struct {
		Job string `json:"job,omitempty"`
	} `json:"alumni,omitempty"`

	Personal struct {
		Hobbies []string `json:"hobbies,omitempty"`
		Socials struct {
			Facebook  string `json:"facebook,omitempty"`
			Instagram string `json:"instagram,omitempty"`
			LinkedIn  string `json:"linkedin,omitempty"`
		} `json:"socials,omitempty"`
	} `json:"personal,omitempty"`

	Achievements struct {
		StudentLeader         []int `json:"student_leader,omitempty"`
		BottomFeederCommittee bool  `json:"bottom_feeder_committee,omitempty"`
		SpoonsassinsVictories []int `json:"spoonsassins_victories,omitempty"`
		KappaKappaPsi         bool  `json:"kappa_kappa_psi,omitempty"`
		TauBetaSigma          bool  `json:"tau_beta_sigma,omitempty"`
	} `json:"achievements,omitempty"`

	Metadata struct {
		Image struct {
			File string `json:"file"`
			Alt  string `json:"alt"`
		} `json:"image"`
	} `json:"metadata"`
}

func (p *Purdoobah) MarchedDuringYear(targetYear int) bool {
	for _, year := range p.Marching.YearsMarched {
		if year == targetYear {
			return true
		}
	}
	return false
}

func (p *Purdoobah) IsStudentLeader() bool {
	return len(p.Achievements.StudentLeader) > 0
}

func (p *Purdoobah) IsStudentLeaderInYear(targetYear int) bool {
	for _, year := range p.Achievements.StudentLeader {
		if year == targetYear {
			return true
		}
	}
	return false
}

func (p *Purdoobah) IsYear(targetYear Year) bool {
	return p.Education.Year == targetYear
}
