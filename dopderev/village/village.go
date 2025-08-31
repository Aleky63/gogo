package village

type VillageElement interface {
	Update()
	FlushInfo() string
}

type Village struct {
	Elements []VillageElement
}

func NewVillage() *Village {
	return &Village{
		Elements: []VillageElement{},
	}
}

func (v *Village) AddElement(e VillageElement) {
	v.Elements = append(v.Elements, e)
}

func (v *Village) UpdateAll() {
	for _, e := range v.Elements {
		e.Update()
	}
}

func (v Village) ShowAllInfo() string {
	info := ""
	for _, e := range v.Elements {
		info += e.FlushInfo()
	}
	return info
}
