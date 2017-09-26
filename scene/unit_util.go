package scene

import "math"

func moveToTarget(u *unitBase, target uniter) {
	ux, uy := u.GetPosition()
	tx, ty := target.GetPosition()

	// calculate which way to go
	// move speed is temporary
	dx, dy := tx-ux, ty-uy
	newx := (float64)(u.moveSpeed) / getDistance(ux, uy, tx, ty) * (float64)(dx)
	newy := (float64)(u.moveSpeed) / getDistance(ux, uy, tx, ty) * (float64)(dy)
	u.sprite.SetPosition(u.sprite.GetPosition().X+(int)(newx), u.sprite.GetPosition().Y+(int)(newy))
}

func canAttackToTarget(u *unitBase, target uniter) bool {
	ux, uy := u.GetPosition()
	tx, ty := target.GetPosition()

	if (float64)(u.attackinfo.attackRange) >= getDistance(ux, uy, tx, ty) {
		return true
	}
	return false
}

func getDistance(ax, ay, bx, by float32) float64 {
	dx, dy := ax-bx, ay-by
	return math.Sqrt((float64)(dx*dx + dy*dy))
}

func getDistanceBetweenUnit(u1, u2 uniter) float64 {
	ax, ay := u1.GetPosition()
	bx, by := u2.GetPosition()
	return getDistance(ax, ay, bx, by)
}

func nearestUnit(u *unitBase, enemies map[string]uniter) uniter {
	var (
		distance float64
		retID    string
	)
	for i, v := range enemies {
		if !v.IsSpawned() {
			continue
		}
		d := getDistanceBetweenUnit(u, v)
		if distance == 0 {
			distance = d
			retID = i
			continue
		}
		if distance > d {
			distance = d
			retID = i
		}
	}

	if retID == "" {
		return nil
	}
	return enemies[retID]
}
