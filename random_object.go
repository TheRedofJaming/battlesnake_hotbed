package hotbed

import (
	"math/rand"

	"github.com/google/uuid"
)

const MAX_BODY_LENGTH = 15

func RandomSnake(name string, available_points map[Point]bool) Snake {
	id := uuid.New().String()
	health := rand.Intn(100)
	body := randomBody(rand.Intn(MAX_BODY_LENGTH))
	return Snake{
		id,
		name,
		health,
		body,
	}
}

func randomBody(l int) (error, []Point, Point) {

}
