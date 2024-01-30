package pulse

type H2HResponse struct {
	Stats       map[string][]*StatsElement `json:"stats"`
	HeadToHeads []*Fixture                 `json:"headToHeads"`
	Fixtures    map[string][]*Fixture      `json:"fixtures"`
}

type StatsElement struct {
	Name           string                  `json:"name"`
	Value          float64                 `json:"value"`
	Description    string                  `json:"description"`
	AdditionalInfo map[string]*interface{} `json:"additionalInfo"`
}

type Gameweek struct {
	Id         float32     `json:"id"`
	CompSeason *CompSeason `json:"compSeason"`
	Gameweek   float32     `json:"gameweek"`
}

type CompSeason struct {
	Label       string       `json:"label"`
	Competition *Competition `json:"competition"`
	Id          float32      `json:"id"`
}

type Competition struct {
	Abbreviation string            `json:"abbreviation"`
	Description  string            `json:"description"`
	Level        string            `json:"level"`
	Source       string            `json:"source"`
	Id           float32           `json:"id"`
	AltIds       map[string]string `json:"altIds"`
}

type Fixture struct {
	Gameweek           *Gameweek         `json:"gameweek"`
	Kickoff            *Kickoff          `json:"kickoff"`
	ProvisionalKickoff *Kickoff          `json:"provisionalKickoff"`
	Teams              []*TeamElement    `json:"teams"`
	Replay             bool              `json:"replay"`
	Ground             *Ground           `json:"ground"`
	NeutralGround      bool              `json:"neutralGround"`
	Status             string            `json:"status"`
	Phase              string            `json:"phase"`
	Outcome            string            `json:"outcome"`
	Attendance         float32           `json:"attendance"`
	Clock              *Clock            `json:"clock"`
	FixtureType        string            `json:"fixtureType"`
	ExtraTime          bool              `json:"extraTime"`
	Shootout           bool              `json:"shootout"`
	BehindClosedDoors  bool              `json:"behindClosedDoors"`
	Goals              []*Goal           `json:"goals"`
	Id                 float32           `json:"id"`
	AltIds             map[string]string `json:"altIds"`
	PenaltyShootouts   []*interface{}    `json:"penaltyShootouts"`
}

type FixtureTeam struct {
	Team  *Team   `json:"team"`
	Score float32 `json:"score"`
}

type TeamElement struct {
	Team  *Team   `json:"team"`
	Score float32 `json:"score"`
}

type Team struct {
	Name      string            `json:"name"`
	Club      *Club             `json:"club"`
	TeamType  string            `json:"teamType"`
	ShortName string            `json:"shortName"`
	Id        float32           `json:"id"`
	AltIds    map[string]string `json:"altIds"`
}

type Club struct {
	Name      string  `json:"name"`
	ShortName string  `json:"shortName"`
	Abbr      string  `json:"abbr"`
	Id        float32 `json:"id"`
}

type Ground struct {
	Name   string  `json:"name"`
	City   string  `json:"city"`
	Source string  `json:"source"`
	Id     float32 `json:"id"`
}

type Kickoff struct {
	Completeness float32 `json:"completeness"`
	Millis       float32 `json:"millis"`
	Label        string  `json:"label"`
}

type Clock struct {
	Secs  float32 `json:"secs"`
	Label string  `json:"label"`
}

type Goal struct {
	PersonId    float32 `json:"personId"`
	Clock       *Clock  `json:"clock"`
	Phase       string  `json:"phase"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
}
