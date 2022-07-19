package challenge

//
// ReachingPoints
// @Description: 780. 到达终点
// @param sx
// @param sy
// @param tx
// @param ty
// @return bool
//
func ReachingPoints(sx, sy, tx, ty int) bool {
    for sx < tx && sy < ty {
        if tx < ty {
            ty %= tx
        } else {
            tx %= ty
        }
    }
    if tx < sx || ty < sy {
        return false
    }
    if sx == tx {
        return (ty-sy)%tx == 0
    } else {
        return (tx-sx)%ty == 0
    }
}
