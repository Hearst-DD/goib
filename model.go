package goib

import "errors"

// ItemType is the type of content encapsulated by the object
type ItemType string

const (
	// ArticleType item type
	ArticleType ItemType = "ARTICLE"
	// VideoType item type
	VideoType ItemType = "VIDEO"
	// CollectionType item type
	CollectionType ItemType = "COLLECTION"
	// SearchType item type
	SearchType ItemType = "SEARCH"
	// ImageType item type
	ImageType ItemType = "IMAGE"
	// GalleryType item type
	GalleryType ItemType = "GALLERY"
	// MapType item type
	MapType ItemType = "MAP"
	// LivevideoType item type
	LivevideoType = "LIVEVIDEO"
	// AudioType item type
	AudioType = "AUDIO"
	// ExternalContentType external content item type
	ExternalContentType = "EXTERNAL_CONTENT"
	// ExternalLinkType external content item type
	ExternalLinkType = "EXTERNAL_LINK"
	// ClosingsType item type
	ClosingsType = "CLOSINGS"
	// HTMLType HTML item type
	HTMLType = "HTML"
	// PersonType Person (AKA Author) item type
	PersonType = "PERSON"
	// TeaserType teasy teasy tease.
	TeaserType = "TEASER"
	// SettingsType is someone's idiot idea of a joke
	SettingsType = ""
	// DownloadFileType is a download file.
	DownloadFileType = "DOWNLOAD_FILE"
	// UnsupportedType is unsupported.
	UnsupportedType = "UNSUPPORTED"
)

type ClosingsFilter string

const (
	ClosingsAll    ClosingsFilter = "all"
	ClosingsClosed ClosingsFilter = "closed"
	ClosingsCount  ClosingsFilter = "count"
	ClosingsInst   ClosingsFilter = "institution"
)

var errUnsupportedType = errors.New("unsupported type")
var errTeaserMissingTarget = errors.New("teaser missing target")

// Receiver captures a type-agnostic representation of an API response as a
// step in processing a response. Its fields are a superset of all content fields,
// so any type can be captured and derived from this struct.
// This object should not be used outside of the IB API classes. It is exposed
// only to facilitate JSON unmarshalling.
type Receiver struct {
	Type                    ItemType            `json:"type"`
	ContentID               int                 `json:"content_id"`
	ContentName             string              `json:"content_name"`
	CollectionName          string              `json:"collection_name"`
	Items                   []Receiver          `json:"items"`
	TeaserTitle             string              `json:"teaser_title"`
	TeaserText              string              `json:"teaser_text"`
	TeaserImage             string              `json:"teaser_image"`
	PublicationDate         int64               `json:"publication_date"`
	Title                   string              `json:"title"`
	Subheadline             string              `json:"subheadline"`
	Text                    string              `json:"article_text"`
	Author                  string              `json:"author"`
	Flavors                 []VideoFlavor       `json:"flavors"`
	StartIndex              int                 `json:"start_index"`
	TotalCount              int                 `json:"total_count"`
	Keywords                string              `json:"keywords"`
	AltText                 string              `json:"alt_text"`
	Caption                 string              `json:"caption"`
	URLs                    []ImageURL          `json:"urls"`
	Media                   []Receiver          `json:"media"`
	RelatedMedia            []Receiver          `json:"related_media"`
	Authors                 []Person            `json:"author_objects"`
	Settings                []map[string]string `json:"settings"`
	Copyright               string              `json:"copyright"`
	CopyrightObjects        []CopyrightObject   `json:"copyright_objects"`
	ExternalContent         string              `json:"external_content"`
	Code                    string              `json:"code"`
	CanonicalURL            string              `json:"canonical_url"`
	URL                     string              `json:"url"`
	StaticMap               string              `json:"static_map"`
	InteractiveMap          string              `json:"interactive_map"`
	Email                   string              `json:"email"`
	Bio                     string              `json:"biography"`
	FullName                string              `json:"full_name"`
	Struct                  []interface{}       `json:"struct"`
	Photo                   []Image             `json:"photo"`
	Stream                  string              `json:"m3u8"`
	NavContext              []string            `json:"navigation_context"`
	AnalyticsCategory       string              `json:"analytics_category"`
	AdvertisingCategory     string              `json:"advertising_category"`
	AdvertisingCategoryPath string              `json:"advertising_category_path"`
	Dateline                string              `json:"author_location"`
	ExternalID              string              `json:"external_id"`
	ShowAds                 bool                `json:"show_ads"`
	Target                  *Receiver           `json:"target"`
	Captions                map[string]string   `json:"captions"` // not from IB, but needed for UnmarshalReceiver()
	LinkText                string              `json:"link_text"`
}

// Item is the base type of all items. It is not used outside the IB package, as
// we return full objects, partially populated
type Item interface {
	GetType() ItemType
	GetContentID() int
	GetTeaserTitle() string
	GetTeaserText() string
	GetPublicationDate() int64
}

// Collection represents a collection of IB Items and metadata about those items
type Collection struct {
	Type                    ItemType            `json:"type"`
	ContentID               int                 `json:"content_id"`
	TeaserTitle             string              `json:"teaser_title"`
	TeaserText              string              `json:"teaser_text"`
	TeaserImage             string              `json:"teaser_image"`
	CollectionName          string              `json:"collection_name"`
	ContentName             string              `json:"content_name"`
	TotalCount              int                 `json:"total_count"`
	StartIndex              int                 `json:"start_index"`
	Keywords                string              `json:"keywords"` // populated only in search results
	Items                   []Item              `json:"items"`
	Media                   []Item              `json:"media"`
	Settings                []map[string]string `json:"settings"`
	NavContext              []string            `json:"navigation_context"`
	AnalyticsCategory       string              `json:"analytics_category"`
	AdvertisingCategory     string              `json:"advertising_category"`
	AdvertisingCategoryPath string              `json:"advertising_category_path"`
}

func (c *Collection) GetType() ItemType {
	return CollectionType
}

func (c *Collection) GetContentID() int {
	return c.ContentID
}

func (c *Collection) GetTeaserTitle() string {
	return c.TeaserTitle
}

func (c *Collection) GetTeaserText() string {
	return c.TeaserText
}

func (c *Collection) GetPublicationDate() int64 {
	return 0 // collections do not have pub dates
}

// Article represents an IB article
type Article struct {
	Type                    ItemType `json:"type"`
	ContentID               int      `json:"content_id"`
	TeaserTitle             string   `json:"teaser_title"`
	TeaserText              string   `json:"teaser_text"`
	TeaserImage             string   `json:"teaser_image"`
	PublicationDate         int64    `json:"publication_date"`
	Title                   string   `json:"title"`
	Subheadline             string   `json:"subheadline"`
	Text                    string   `json:"article_text"`
	Authors                 []Person `json:"author_objects"`
	Media                   []Item   `json:"media"`
	RelatedMedia            []Item   `json:"related_media"`
	CanonicalURL            string   `json:"canonical_url"`
	URL                     string   `json:"url"`
	NavContext              []string `json:"navigation_context"`
	AnalyticsCategory       string   `json:"analytics_category"`
	AdvertisingCategory     string   `json:"advertising_category"`
	AdvertisingCategoryPath string   `json:"advertising_category_path"`
	Dateline                string   `json:"author_location"`
}

func (a *Article) GetType() ItemType {
	return ArticleType
}

func (a *Article) GetContentID() int {
	return a.ContentID
}

func (a *Article) GetTeaserTitle() string {
	return a.TeaserTitle
}

func (a *Article) GetTeaserText() string {
	return a.TeaserText
}

func (a *Article) GetPublicationDate() int64 {
	return a.PublicationDate
}

// Video represents an IB video
type Video struct {
	Type                    ItemType      `json:"type"`
	ContentID               int           `json:"content_id"`
	TeaserTitle             string        `json:"teaser_title"`
	TeaserText              string        `json:"teaser_text"`
	TeaserImage             string        `json:"teaser_image"`
	PublicationDate         int64         `json:"publication_date"`
	Authors                 []Person      `json:"author_objects"`
	Title                   string        `json:"title"`
	Subheadline             string        `json:"subheadline"`
	Flavors                 []VideoFlavor `json:"flavors"`
	Media                   []Item        `json:"media"`
	CanonicalURL            string        `json:"canonical_url"`
	URL                     string        `json:"url"`
	NavContext              []string      `json:"navigation_context"`
	AnalyticsCategory       string        `json:"analytics_category"`
	AdvertisingCategory     string        `json:"advertising_category"`
	AdvertisingCategoryPath string        `json:"advertising_category_path"`
	ShowAds                 bool          `json:"show_ads"`
	Stream                  string        `json:"m3u8"`
}

func (v *Video) GetType() ItemType {
	return VideoType
}

func (v *Video) GetContentID() int {
	return v.ContentID
}

func (v *Video) GetTeaserTitle() string {
	return v.TeaserTitle
}

func (v *Video) GetTeaserText() string {
	return v.TeaserText
}

func (v *Video) GetPublicationDate() int64 {
	return v.PublicationDate
}

// VideoFlavor represents a flavor (i.e. resolution) of an IB Video
type VideoFlavor struct {
	Type     string `json:"video_type"`
	URL      string `json:"url"`
	Bitrate  int    `json:"bitrate"`
	Duration int    `json:"duration"`
	FileSize int    `json:"file_size"`
	Codec    string `json:"codec"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
}

// Image represents an IB image content piece
type Image struct {
	Type                    ItemType          `json:"type"`
	ContentID               int               `json:"content_id"`
	TeaserTitle             string            `json:"teaser_title"`
	TeaserText              string            `json:"teaser_text"`
	TeaserImage             string            `json:"teaser_image"`
	PublicationDate         int64             `json:"publication_date"`
	AltText                 string            `json:"alt_text"`
	Caption                 string            `json:"caption"`
	Author                  string            `json:"author"`
	Authors                 []Person          `json:"author_objects"`
	Title                   string            `json:"title"`
	Subheadline             string            `json:"subheadline"`
	Keywords                string            `json:"keywords"`
	URLs                    []ImageURL        `json:"urls"`
	Copyright               string            `json:"copyright"`
	CopyrightObjects        []CopyrightObject `json:"copyright_objects"`
	URL                     string            `json:"url"`
	CanonicalURL            string            `json:"canonical_url"`
	NavContext              []string          `json:"navigation_context"`
	AnalyticsCategory       string            `json:"analytics_category"`
	AdvertisingCategory     string            `json:"advertising_category"`
	AdvertisingCategoryPath string            `json:"advertising_category_path"`
}

func (i *Image) GetType() ItemType {
	return ImageType
}

func (i *Image) GetContentID() int {
	return i.ContentID
}

func (i *Image) GetTeaserTitle() string {
	return i.TeaserTitle
}

func (i *Image) GetTeaserText() string {
	return i.TeaserText
}

func (i *Image) GetPublicationDate() int64 {
	return i.PublicationDate
}

// ImageURL is a URL flavor for an image
type ImageURL struct {
	Version string `json:"version"`
	Height  int    `json:"height"`
	Width   int    `json:"width"`
	URL     string `json:"url"`
	Mime    string `json:"mime"`
}

// Gallery represents an image gallery
type Gallery struct {
	Type                    ItemType          `json:"type"`
	ContentID               int               `json:"content_id"`
	TeaserTitle             string            `json:"teaser_title"`
	TeaserText              string            `json:"teaser_text"`
	TeaserImage             string            `json:"teaser_image"`
	PublicationDate         int64             `json:"publication_date"`
	Authors                 []Person          `json:"author_objects"`
	Keywords                string            `json:"keywords"`
	Title                   string            `json:"title"`
	Subheadline             string            `json:"subheadline"`
	Media                   []Item            `json:"media"`
	Items                   []Item            `json:"items"`
	Captions                map[string]string `json:"captions"`
	CanonicalURL            string            `json:"canonical_url"`
	URL                     string            `json:"url"`
	NavContext              []string          `json:"navigation_context"`
	AnalyticsCategory       string            `json:"analytics_category"`
	AdvertisingCategory     string            `json:"advertising_category"`
	AdvertisingCategoryPath string            `json:"advertising_category_path"`
}

func (g *Gallery) GetType() ItemType {
	return GalleryType
}

func (g *Gallery) GetContentID() int {
	return g.ContentID
}

func (g *Gallery) GetTeaserTitle() string {
	return g.TeaserTitle
}

func (g *Gallery) GetTeaserText() string {
	return g.TeaserText
}

func (g *Gallery) GetPublicationDate() int64 {
	return g.PublicationDate
}

// Audio represents an audio clip
type Audio struct {
	Type                    ItemType `json:"type"`
	ContentID               int      `json:"content_id"`
	Title                   string   `json:"title"`
	Subheadline             string   `json:"subheadline"`
	TeaserTitle             string   `json:"teaser_title"`
	TeaserText              string   `json:"teaser_text"`
	Authors                 []Person `json:"author_objects"`
	CanonicalURL            string   `json:"canonical_url"`
	URL                     string   `json:"url"`
	Media                   []Item   `json:"media"`
	Stream                  string   `json:"stream"`
	PublicationDate         int64    `json:"publication_date"`
	NavContext              []string `json:"navigation_context"`
	AnalyticsCategory       string   `json:"analytics_category"`
	AdvertisingCategory     string   `json:"advertising_category"`
	AdvertisingCategoryPath string   `json:"advertising_category_path"`
}

func (a *Audio) GetType() ItemType {
	return AudioType
}

func (a *Audio) GetContentID() int {
	return a.ContentID
}

func (a *Audio) GetTeaserTitle() string {
	return a.TeaserTitle
}

func (a *Audio) GetTeaserText() string {
	return ""
}

func (a *Audio) GetPublicationDate() int64 {
	return a.PublicationDate
}

// Livevideo represents a live stream
type Livevideo struct {
	Type                    ItemType `json:"type"`
	ContentID               int      `json:"content_id"`
	Title                   string   `json:"title"`
	Subheadline             string   `json:"subheadline"`
	TeaserTitle             string   `json:"teaser_title"`
	TeaserText              string   `json:"teaser_text"`
	PublicationDate         int64    `json:"publication_date"`
	Authors                 []Person `json:"author_objects"`
	CanonicalURL            string   `json:"canonical_url"`
	URL                     string   `json:"url"`
	Media                   []Item   `json:"media"`
	Stream                  string   `json:"stream"`
	ExternalID              string   `json:"external_id"`
	NavContext              []string `json:"navigation_context"`
	AnalyticsCategory       string   `json:"analytics_category"`
	AdvertisingCategory     string   `json:"advertising_category"`
	AdvertisingCategoryPath string   `json:"advertising_category_path"`
	ShowAds                 bool     `json:"show_ads"`
}

func (l *Livevideo) GetType() ItemType {
	return LivevideoType
}

func (l *Livevideo) GetContentID() int {
	return l.ContentID
}

func (l *Livevideo) GetTeaserTitle() string {
	return l.TeaserTitle
}

func (l *Livevideo) GetTeaserText() string {
	return l.TeaserText
}

func (l *Livevideo) GetPublicationDate() int64 {
	return l.PublicationDate
}

// Map represents a map
type Map struct {
	Type                    ItemType `json:"type"`
	ContentID               int      `json:"content_id"`
	PublicationDate         int64    `json:"publication_date"`
	TeaserTitle             string   `json:"teaser_title"`
	TeaserText              string   `json:"teaser_text"`
	Title                   string   `json:"title"`
	Subheadline             string   `json:"subheadline"`
	StaticMap               string   `json:"static_map"`
	InteractiveMap          string   `json:"interactive_map"`
	CanonicalURL            string   `json:"canonical_url"`
	URL                     string   `json:"url"`
	NavContext              []string `json:"navigation_context"`
	AnalyticsCategory       string   `json:"analytics_category"`
	AdvertisingCategory     string   `json:"advertising_category"`
	AdvertisingCategoryPath string   `json:"advertising_category_path"`
}

func (m *Map) GetType() ItemType {
	return MapType
}

func (m *Map) GetContentID() int {
	return m.ContentID
}

func (m *Map) GetTeaserTitle() string {
	return m.TeaserTitle
}

func (m *Map) GetTeaserText() string {
	return ""
}

func (m *Map) GetPublicationDate() int64 {
	return m.PublicationDate
}

// ExternalContent represents an external content object
type ExternalContent struct {
	Type            ItemType      `json:"type"`
	ContentID       int           `json:"content_id"`
	PublicationDate int64         `json:"publication_date"`
	TeaserTitle     string        `json:"teaser_title"`
	ExternalContent string        `json:"external_content"`
	Struct          []interface{} `json:"struct"`
}

func (e *ExternalContent) GetType() ItemType {
	return ExternalContentType
}

func (e *ExternalContent) GetContentID() int {
	return e.ContentID
}

func (e *ExternalContent) GetTeaserTitle() string {
	return e.TeaserTitle
}

func (e *ExternalContent) GetTeaserText() string {
	return ""
}

func (e *ExternalContent) GetPublicationDate() int64 {
	return e.PublicationDate
}

// ExternalLink represents an external link object
type ExternalLink struct {
	Type            ItemType `json:"type"`
	ContentID       int      `json:"content_id"`
	PublicationDate int64    `json:"publication_date"`
	TeaserTitle     string   `json:"teaser_title"`
	TeaserText      string   `json:"teaser_text"`
	CanonicalURL    string   `json:"canonical_url"`
	URL             string   `json:"url"`
	Media           []Item   `json:"media"`
}

func (e *ExternalLink) GetType() ItemType {
	return ExternalLinkType
}

func (e *ExternalLink) GetContentID() int {
	return e.ContentID
}

func (e *ExternalLink) GetTeaserTitle() string {
	return e.TeaserTitle
}

func (e *ExternalLink) GetTeaserText() string {
	return e.TeaserText
}

func (e *ExternalLink) GetPublicationDate() int64 {
	return e.PublicationDate
}

// HTMLContent represents a content object that contains a raw HTML payload
type HTMLContent struct {
	Type                    ItemType `json:"type"`
	ContentID               int      `json:"content_id"`
	PublicationDate         int64    `json:"publication_date"`
	Code                    string   `json:"code"`
	URL                     string   `json:"url"`
	TeaserTitle             string   `json:"teaser_title"`
	NavContext              []string `json:"navigation_context"`
	AnalyticsCategory       string   `json:"analytics_category"`
	AdvertisingCategory     string   `json:"advertising_category"`
	AdvertisingCategoryPath string   `json:"advertising_category_path"`
}

func (h *HTMLContent) GetType() ItemType {
	return HTMLType
}

func (h *HTMLContent) GetContentID() int {
	return h.ContentID
}

func (h *HTMLContent) GetTeaserTitle() string {
	return h.TeaserTitle
}

func (h *HTMLContent) GetTeaserText() string {
	return ""
}

func (h *HTMLContent) GetPublicationDate() int64 {
	return h.PublicationDate
}

// Person represents an IB person
type Person struct {
	Type                    ItemType `json:"type"`
	ContentID               int      `json:"content_id"`
	Blurb                   string   `json:"teaser_text"`
	FullName                string   `json:"full_name"`
	Title                   string   `json:"title"`
	TeaserImage             string   `json:"teaser_image"`
	PublicationDate         int64    `json:"publication_date"`
	Bio                     string   `json:"biography"`
	Photo                   []Image  `json:"photo,omitempty"`
	Email                   string   `json:"email"`
	FacebookUsername        string   `json:"facebook_username"`
	FacebookUID             string   `json:"facebook_uid"`
	TwitterUsername         string   `json:"twitter_username"`
	GPlusUID                string   `json:"gplus_uid"`
	StoriesCOID             int      `json:"recent_stories"`
	CanonicalURL            string   `json:"canonical_url"`
	URL                     string   `json:"url"`
	NavContext              []string `json:"navigation_context"`
	AnalyticsCategory       string   `json:"analytics_category"`
	AdvertisingCategory     string   `json:"advertising_category"`
	AdvertisingCategoryPath string   `json:"advertising_category_path"`
}

func (p *Person) GetType() ItemType {
	return PersonType
}

func (p *Person) GetContentID() int {
	return p.ContentID
}

func (p *Person) GetTeaserTitle() string {
	return p.FullName
}

func (p *Person) GetTeaserText() string {
	return p.Blurb
}

func (p *Person) GetPublicationDate() int64 {
	return p.PublicationDate
}

type CopyrightObject struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

type ClosingsResponse struct {
	Count              ClsCount                    `json:"count"`
	Institutions       map[string][]ClsInstitution `json:"institutions,omitempty"`
	ClosedInstitutions map[string][]ClsInstitution `json:"closed_institutions,omitempty"`
	Institution        ClsInstitution              `json:"institution"`
}

type ClsCount struct {
	Total           int   `json:"total"`
	PublicationDate int64 `json:"publication_date"`
}

type ClsInstitution struct {
	Name            string `json:"name"`
	PublicationDate int64  `json:"publication_date"`
	City            string `json:"city"`
	County          string `json:"county"`
	State           string `json:"state"`
	ProviderID      string `json:"provider_id"`
	Status          string `json:"status"`
}

// Settings represents a collection of settings
type Settings struct {
	ContentID int               `json:"content_id"`
	Settings  map[string]string `json:"settings"`
}

func (c *Settings) GetType() ItemType {
	return SettingsType
}

func (c *Settings) GetContentID() int {
	return c.ContentID
}

func (c *Settings) GetTeaserTitle() string {
	return ""
}

func (c *Settings) GetTeaserText() string {
	return ""
}

func (c *Settings) GetPublicationDate() int64 {
	return 0 // collections do not have pub dates
}

// Teaser represents ... something
type Teaser struct {
	Type                    ItemType `json:"type"`
	ContentID               int      `json:"content_id"`
	Title                   string   `json:"title"`
	TeaserTitle             string   `json:"teaser_title"`
	TeaserText              string   `json:"teaser_text"`
	PublicationDate         int64    `json:"publication_date"`
	Authors                 []Person `json:"author_objects"`
	Media                   []Item   `json:"media"`
	NavContext              []string `json:"navigation_context"`
	AnalyticsCategory       string   `json:"analytics_category"`
	AdvertisingCategory     string   `json:"advertising_category"`
	AdvertisingCategoryPath string   `json:"advertising_category_path"`
	Target                  Item     `json:"target"`
}

func (t *Teaser) GetType() ItemType {
	return TeaserType
}

func (t *Teaser) GetContentID() int {
	return t.ContentID
}

func (t *Teaser) GetTeaserTitle() string {
	return t.TeaserTitle
}

func (t *Teaser) GetTeaserText() string {
	return t.TeaserText
}

func (t *Teaser) GetPublicationDate() int64 {
	return t.PublicationDate
}

// DownloadFile represents a file download object
type DownloadFile struct {
	Type            ItemType `json:"type"`
	ContentID       int      `json:"content_id"`
	PublicationDate int64    `json:"publication_date"`
	TeaserTitle     string   `json:"teaser_title"`
	TeaserText      string   `json:"teaser_text"`
	LinkText        string   `json:"link_text"`
	URL             string   `json:"url"`
}

func (d *DownloadFile) GetType() ItemType {
	return DownloadFileType
}

func (d *DownloadFile) GetContentID() int {
	return d.ContentID
}

func (d *DownloadFile) GetTeaserTitle() string {
	return d.TeaserTitle
}

func (d *DownloadFile) GetTeaserText() string {
	return d.TeaserText
}

func (d *DownloadFile) GetPublicationDate() int64 {
	return d.PublicationDate
}
