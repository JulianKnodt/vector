package vector

type Polygon []Vec3

type errorNotCoplanar string

func (e errorNotCoplanar) Error() string { return string(e) }

const ErrorNotCoplanar = errorNotCoplanar("Points in polygon not coplanar")

func NewPolygon(from []Vec3) (*Polygon, error) {
	if !Coplanar(from) {
		return nil, ErrorNotCoplanar
	}
	p := Polygon(from)
	return &p, nil
}

func (p Polygon) Normal(to Vec3) (dir Vec3, invAble bool) {
	return normalNoCheck(p[0], p[1], p[2])
}
