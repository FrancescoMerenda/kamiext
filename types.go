package kamiext

type Response struct {
	Data  any    `json:"data"`
	Error string `json:"error"`
}
type Extension struct {
	Name            string   `json:"name"`
	ExtensionId     string   `json:"id"`
	Author          string   `json:"author"`
	ExtensionSource string   `json:"source"`
	Sources         []Source `json:"sources"`
	Logo            string   `json:"logo"`
	Executable      string   `json:"executable"`
	UpdatesUrl      string   `json:"updatesurl"`
	Updater         string   `json:"updater"`
	Version         string   `json:"version"`
}
type Source struct {
	Name            string `json:"name"`
	Methods         []string `json:"methods"`
	Id              string `json:"id"`
	Logo            string   `json:"logo"`
	Domain          string `json:"domain"`
	ExtensionId     string `json:"extensionid"`
}
type Comic struct {
	Title         string    `json:"title"`
	Id            string    `json:"id"`
	Url           string    `json:"url"`
	SourceId      string    `json:"sourceid"`
	SourceName    string    `json:"sourcename"`
	ComicType     string    `json:"comictype"`
	ReaderMode    string    `json:"readermode"`
	ReleaseStatus string    `json:"releasestatus"`
	Summary       string    `json:"summary"`
	Cover         string    `json:"cover"`
	Authors       []string  `json:"authors"`
	Artists       []string  `json:"artists"`
	Altnames      []string  `json:"altnames"`
	Genres        []string  `json:"genres"`
	Year          string    `json:"year"`
	Chapters      []Chapter `json:"chapters"`
}

type SearchComic struct {
	Title  string `json:"title"`
	Id     string `json:"id"`
	Handle string `json:"handle"`
	Cover  string `json:"cover"`
}
type Chapter struct {
	Name      string   `json:"name"`
	Number    string   `json:"number"`
	Timestamp int      `json:"timestamp"`
	Handle    string   `json:"handle"`
	Pages     []string `json:"pages"`
}
