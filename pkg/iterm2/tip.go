package iterm2

// Tip is a direct translation of the iTerm2 internal representation of a Tip of
// the Day.
type Tip struct {
	ID    int
	Title string
	Body  string
	URL   string
}

// Tips facilitates clean and unified interaction with clusters of iTerm2 Tips.
type Tips []Tip

func (tip *Tip) Empty() bool {
	return tip == nil || *tip == Tip{}
}
