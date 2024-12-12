package y2024d12

// type Region struct {
// 	crop  rune
// 	nodes map[Vec2D]RegionNode
// }

// func (r Region) Perimeter() int {
// 	p := 0
// 	for _, n := range r.nodes {
// 		p += n.Perimeter()
// 	}
// 	return p
// }

// func (r Region) Area() int {
// 	return len(r.nodes)
// }

// func (r Region) Price() int {
// 	return r.Perimeter() * r.Area()
// }

type Vec2D struct {
	X int
	Y int
}

type RegionNode struct {
	crop      rune
	postition Vec2D
	regionId  int
	toRight   *RegionNode
	toLeft    *RegionNode
	toUp      *RegionNode
	toBottom  *RegionNode
}

func (rn *RegionNode) LinkRight(node *RegionNode) {
	if rn.crop != node.crop {
		return
	}
	if rn.toRight != node {
		rn.toRight = node
	}
	node.LinkLeft(rn)
}
func (rn *RegionNode) LinkLeft(node *RegionNode) {
	if rn.crop != node.crop {
		return
	}
	if rn.toLeft != node {
		rn.toLeft = node
	}
	node.LinkRight(rn)
}
func (rn *RegionNode) LinkUp(node *RegionNode) {
	if rn.crop != node.crop {
		return
	}
	if rn.toUp != node {
		rn.toUp = node
	}
	node.LinkBottom(rn)
}
func (rn *RegionNode) LinkBottom(node *RegionNode) {
	if rn.crop != node.crop {
		return
	}
	if rn.toBottom != node {
		rn.toBottom = node
	}
	node.LinkUp(rn)
}

func (rn *RegionNode) PushRegionId(regionId int) {
	if rn.regionId == -1 {
		rn.regionId = regionId
	}
	if rn.toLeft != nil && rn.toLeft.regionId == -1 {
		rn.toLeft.regionId = rn.regionId
	}
	if rn.toRight != nil && rn.toRight.regionId == -1 {
		rn.toRight.regionId = rn.regionId
	}
	if rn.toUp != nil && rn.toUp.regionId == -1 {
		rn.toUp.regionId = rn.regionId
	}
	if rn.toBottom != nil && rn.toBottom.regionId == -1 {
		rn.toBottom.regionId = rn.regionId
	}
}

func (rn RegionNode) Perimeter() int {
	perimeter := 0
	if rn.toRight == nil {
		perimeter++
	}
	if rn.toLeft == nil {
		perimeter++
	}
	if rn.toUp == nil {
		perimeter++
	}
	if rn.toBottom == nil {
		perimeter++
	}
	return perimeter
}
