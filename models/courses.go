// class for courses implementing content interface
package models

type Courses struct {
	AbstractContent
	Miniature          string
	VideosTitles       []string
	VideosDescriptions []string
	VideosURL          []string
	// Otros campos que puedas necesitar
}
