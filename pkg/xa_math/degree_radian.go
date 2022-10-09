package xa_math

import "math"

type DegreeRadian struct {
	Degree float64 `json:"degree"` //角度
	Radian float64 `json:"radian"` //弧度
	X0     float64 `json:"x0"`     //圆心坐标x
	Y0     float64 `json:"y0"`     //圆心坐标y
	R      float64 `json:"r"`      //半径
}

/*
	func main(){
		r := 50.0
		a := 90.0
		dr := NewDegreeRadian(a, 0, r)
		r.Radian = dr.DegreeToRadian()
		fmt.Printf("%f %f\n", dr.X1(), dr.Y1()) // 0.000000 50.000000
	}
*/
// 角度弧度互转
func NewDegreeRadian(degree, radian, r float64) *DegreeRadian {
	return &DegreeRadian{
		Degree: degree,
		Radian: radian,
		X0:     0.0,
		Y0:     0.0,
		R:      r,
	}
}

// 角度转弧度
func (dr *DegreeRadian) DegreeToRadian() float64 {
	return dr.Degree * math.Pi / 180
}

// 弧度转角度
func (dr *DegreeRadian) RadianToDegree() float64 {
	return dr.Radian * 180 / math.Pi
}

// x坐标(left)
func (dr *DegreeRadian) X1() float64 {
	return dr.X0 + dr.R*math.Cos(dr.Radian)
}

// y坐标(top)
func (dr *DegreeRadian) Y1() float64 {
	return dr.Y0 + dr.R*math.Sin(dr.Radian)
}
