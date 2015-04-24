package bucket

type ReleaseList []Release

func NewReleaseList() *ReleaseList {
	return &ReleaseList{}
}

func (list ReleaseList) add(release Release) ReleaseList {
	return append(list, release)
}
