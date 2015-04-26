package bucket

import (
	"sort"
)

type ReleaseList []*Release

func NewReleaseList() *ReleaseList {
	return &ReleaseList{}
}

func (list *ReleaseList) Append(release *Release) *ReleaseList {
	newList := append(*list, release)
	return &newList
}

func (list *ReleaseList) Len() int {
	return len(*list)
}

func (list *ReleaseList) Index(i int) *Release {
	return (*list)[i]
}

func (list *ReleaseList) Swap(i, j int) {
	l := *list
	l[i], l[j] = l[j], l[i]
}

func (list *ReleaseList) Less(i, j int) bool {
	vi, vj := list.Index(i).Version, list.Index(j).Version
	if vi.Major == vj.Major {
		return vi.Minor < vj.Minor
	} else {
		return vi.Major < vj.Major
	}
}

func (list *ReleaseList) Sort() {
	sort.Sort(list)
}

func (list *ReleaseList) FilterByPlatform(platform string) *ReleaseList {
	newList := NewReleaseList()
	for _, release := range *list {
		if release.Platform == platform {
			newList = newList.Append(release)
		}
	}
	return newList
}

func (list *ReleaseList) Latest() *Release {
	list.Sort()
	length := list.Len()
	if length > 0 {
		return list.Index(length - 1)
	} else {
		return nil
	}
}
