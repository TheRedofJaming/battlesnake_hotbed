package hotbed

//  No detection for head-head collisions. Other snakes are simply declared ilegal.
func LegalMoves[PS Point | Snake](b Board, start PS) (legal []move) {
	var p Point
	if snake, ok := any(start).(Snake); ok { // if a snake is passed into the function, use it's Head directly.
		p = snake.Head
	} else {
		p = any(start).(Point)
	}
	probablePoints := map[move]Point{
		UP: Point{
			p.X,
			p.Y + 1,
		},
		DOWN: Point{
			p.X,
			p.Y - 1,
		},
		RIGHT: Point{
			p.X + 1,
			p.Y,
		},
		LEFT: Point{
			p.X - 1,
			p.Y,
		},
	}
	impossiblePoints := make(map[Point]bool)
	for _, h := range b.Hazards {
		impossiblePoints[h] = true
	}
	for _, snake := range b.Snakes {
		for _, s := range snake.Body {
			impossiblePoints[s] = true
		}
	}
	for move, p := range probablePoints {
		if !impossiblePoints[p] {
			legal = append(legal, move)
		}
	}
	return
}
