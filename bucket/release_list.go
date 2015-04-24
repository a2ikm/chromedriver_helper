package bucket

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
