package seabotserver

// структуры для обмена данными между ботами и сервером.

// FROM BOT = FB
type FBBvb struct {
	Place int   `json:"place"`
	Ships []int `json:"ships"`
}

type FBTurn struct {
	Shot [2]int `json:"shot"`
}

type FBProfile struct {
	Info int `json:"info"`
}

type FromBot struct {
	Auth    string     `json:"auth"`
	Exit    bool       `json:"exit"`
	Bvb     *FBBvb     `json:"bvb"`
	Bvr     *FBBvb     `json:"bvr"`
	Turn    *FBTurn    `json:"turn"`
	Profile *FBProfile `json:"profile"`
}

// TO BOT = TB

type ToBot struct {
	Auth    *TBAuth    `json:"auth,omitempty"`
	Bvb     *TBBvb     `json:"bvb,omitempty"`
	Turn    *TBTurn    `json:"turn,omitempty"`
	End     *TBEnd     `json:"end,omitempty"`
	Error   *TBError   `json:"error,omitempty"`
	Profile *TBProfile `json:"profile,omitempty"`
}

type TBProfile struct {
	Gnum int `json:"gnum"`
}

type TBEnd struct {
	Winner int64     `json:"winner"`
	Ships  *[100]int `json:"opponent,omitempty"`
}

type TBTurn struct {
	ID       int64           `json:"id,omitempty"`
	Result   int             `json:"result,omitempty"`
	Opponent *TBOpponentTurn `json:"opponent,omitempty"`
}

type TBOpponentTurn struct {
	Shot   [2]int `json:"shot"`
	Result int    `json:"result,omitempty"`
}

type TBBvb struct {
	Wait  int       `json:"wait,omitempty"`
	ID    int64     `json:"id,omitempty"`
	Name  string    `json:"name,omitempty"`
	Ships *[100]int `json:"ships,omitempty"`
}

type TBAuth struct {
	OK    bool   `json:"ok"`
	Error string `json:"error"`
	ID    int64  `json:"id"`
	User  int64  `json:"userid"`
}

type TBError struct {
	Error string `json:"error"`
}
