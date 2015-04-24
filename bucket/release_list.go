package bucket

import (
	"sort"
)

type ReleaseList []Release

func NewReleaseList() *ReleaseList {
	return &ReleaseList{}
}

func (list *ReleaseList) Append(release *Release) *ReleaseList {
	newList := append(*list, *release)
	return &newList
}

func (list *ReleaseList) Len() int {
	return len(*list)
}

func (list *ReleaseList) Index(i int) *Release {
	return &(*list)[i]
}

func (list *ReleaseList) Swap(i, j int) {
	l := *list
	l[i], l[j] = l[j], l[i]
}

func (list *ReleaseList) Less(i, j int) bool {
	l := *list
	if l[i].Version.Major == l[j].Version.Major {
		return l[i].Version.Minor < l[j].Version.Minor
	} else {
		return l[i].Version.Major < l[j].Version.Major
	}
}

func (list *ReleaseList) Sort() {
	sort.Sort(list)
}
