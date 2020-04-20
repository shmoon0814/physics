package vector3

import (
	"errors"
	"physics/typedef"
)

var errVectorCreateFail = errors.New("Vector Param not valid. nil or only three param")

type vector3 struct {
	x, y, z, pad typedef.Real
}

//주소를 넘긴다.
//할당 받을때는 vector3 := NewVector3(1,2,3) 이렇게 써야한다
func NewVector3(coordinate ...typedef.Real) (*vector3, error) {

	if coordinate != nil && len(coordinate) == 3 {
		return &vector3{x: coordinate[0], y: coordinate[1], z: coordinate[2]}, nil
	} else if coordinate == nil || len(coordinate) == 0 {
		return &vector3{x: 0, y: 0, z: 0}, nil
	} else {
		return nil, errVectorCreateFail
	}
}

func (v *vector3) Invert() {
	v.x = -v.x
	v.y = -v.y
	v.z = -v.z
}
