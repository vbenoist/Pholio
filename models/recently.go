package models

// Difference between Lastly / Lately is just from configured date -- TODO
type RecentlyRecords struct {
	Lastly []Record
	Lately []Record
}
