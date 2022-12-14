package hotbed

// API Objects
// https://docs.battlesnake.com/api

type move string

const (
	UP    move = "up"
	DOWN  move = "down"
	LEFT  move = "left"
	RIGHT move = "right"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Snake struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Health         int            `json:"health"`
	Body           []Point        `json:"body"`
	Head           Point          `json:"head"`
	Length         int            `json:"length"`
	Latency        string         `json:"latency"`
	Shout          string         `json:"shout"`
	Customizations Customizations `json:"customizations"`
}

type Customizations struct {
	Color string `json:"color"`
	Head  string `json:"head"`
	Tail  string `json:"tail"`
}

type Board struct {
	Height  int     `json:"height"`
	Width   int     `json:"width"`
	Food    []Point `json:"food"`
	Hazards []Point `json:"hazards"`
	Snakes  []Snake `json:"snakes"`
}

type GameState struct {
	Game  Game  `json:"game"`
	Turn  int   `json:"turn"`
	Board Board `json:"board"`
	You   Snake `json:"you"`
}

type Game struct {
	ID      string  `json:"id"`
	Ruleset Ruleset `json:"ruleset"`
	Map     string  `json:"map"`
	Source  string  `json:"source"`
	Timeout int     `json:"timeout"`
}

type Ruleset struct {
	Name     string          `json:"name"`
	Version  string          `json:"version"`
	Settings RulesetSettings `json:"settings"`
}

type RulesetSettings struct {
	FoodSpawnChance     int `json:"foodSpawnChance"`
	MinimumFood         int `json:"minimumFood"`
	HazardDamagePerTurn int `json:"hazardDamagePerTurn"`
}

type SnakeInfoResponse struct {
	APIVersion string `json:"apiversion"`
	Author     string `json:"author"`
	Color      string `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
}

type SnakeMoveResponse struct {
	Move  string `json:"move"`
	Shout string `json:"shout"`
}
