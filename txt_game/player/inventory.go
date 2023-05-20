package character

type Inventory struct {
	items []Item
}

func NewInventory() Menu {
	return &Inventory{
		items:make([]Item, 0),
	}
}

func (inv Inventory) Display() string {
	horisontal 	:= "--------------------------------------------------------------\n"
	vertical 	:= "||                                                          ||\n"

	out := horisontal
	out += vertical
	for _, item := range inv.items {
		namelen := len(item.Name())
		desclen := len(item.Desc())
		name_startindx := len(vertical)/2 - namelen/2
		desc_startindx := len(vertical)/2 - desclen/2
		out += vertical[:name_startindx] + item.Name() + vertical[name_startindx+namelen:]
		out += vertical[:desc_startindx] + item.Desc() + vertical[desc_startindx+desclen:]
		out+= vertical
		out+= vertical
	}
	out += horisontal

	return out
}

func (inv *Inventory) Add(item Item) {
	inv.items = append(inv.items, item)
}

func (inv *Inventory) Remove() {

}

func (item ItemImpl) GetName() string {
	return item.name
}

func (item ItemImpl) GetDescription() string {
	return item.desc
}
