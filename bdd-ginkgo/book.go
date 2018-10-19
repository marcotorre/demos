package books

const novel = "NOVEL"
const shortStory = "SHORT STORY"

// Book represents any type of book
type Book struct {
	Title  string
	Author string
	Pages  int
}

// CategoryByLength returns the category of book determined by page length
func (b *Book) CategoryByLength() string {
	if b.Pages < 300 {
		return shortStory
	} else {
		return novel
	}
}
