package hotbed

import (
	"math/rand"

	"github.com/google/uuid"
)

const MAX_BODY_LENGTH = 15

func RandomSnake(name string, availabePoints []Point) Snake {
	id := uuid.New().String()
	health := rand.Intn(100)
	head, availablePoints := RandomHead(availabePoints) // mutates available_points
	body, availablePoints := RandomBody(rand.Intn(MAX_BODY_LENGTH), availablePoints)
	return Snake{
		id,
		name,
		health,
		body,
	}
}

func RandomBody(l int, availablePoints []Point) (error, []Point, Point) {

}

func RandomHead(availablePoints []Point) (Point, []Point) {
	i := rand.Intn(len(availablePoints))
	chosen := availablePoints[i]
	availablePoints[i] = availablePoints[len(availablePoints)-1] // replace 'chosen' with the last element
	availablePoints = availablePoints[:len(availablePoints)-1]   // and delete it
	return chosen, availablePoints
}

func RandomBoard(width, height int) Board {

}
