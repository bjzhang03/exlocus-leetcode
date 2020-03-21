package medium_0593

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	if len(p1) > 0 && len(p2) > 0 && len(p3) > 0 && len(p4) > 0 {
		po1 := *&point{p1[0], p1[1]}
		po2 := *&point{p2[0], p2[1]}
		po3 := *&point{p3[0], p3[1]}
		po4 := *&point{p4[0], p4[1]}
		if solve(po1, po2, po3, po4) {
			return true
		}
		if solve(po1, po2, po4, po3) {
			return true
		}
		if solve(po1, po3, po2, po4) {
			return true
		}
		if solve(po1, po3, po4, po2) {
			return true
		}
		if solve(po1, po4, po2, po3) {
			return true
		}
		if solve(po1, po4, po3, po2) {
			return true
		}
	}
	return false
}

func solve(p1, p2, p3, p4 point) bool {
	// 两队对边相等
	if distanceSqure(p1, p2) > 0 &&
		distanceSqure(p1, p2) == distanceSqure(p3, p4) &&
		distanceSqure(p2, p3) == distanceSqure(p1, p4) &&
		distanceSqure(p1, p2) == distanceSqure(p2, p3) &&
		distanceSqure(p1, p4) == distanceSqure(p2, p3) {
		// 直角三角形
		if distanceSqure(p1, p2)+distanceSqure(p2, p3) == distanceSqure(p1, p3) &&
			distanceSqure(p2, p3)+distanceSqure(p3, p4) == distanceSqure(p2, p4) {
			return true
		}
	}
	return false
}

type point struct {
	x, y int
}

func distanceSqure(p1, p2 point) int {
	return (p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)
}
