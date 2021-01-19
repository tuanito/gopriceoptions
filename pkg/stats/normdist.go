package stats

// references:
// https://github.com/chobie/go-gaussian/blob/master/gaussian.go

import (
	"fmt"
	"math"
)

var sqrt2 float64 = math.Pow(2, 0.5)
var maxNumberOfStandardDeviations float64 = 8.0

type normdist struct {
	stddev      float64
	mean        float64
	stddevsqpi  float64 // sqrt(2 x pi)
	twostddevsq float64 // 2.0 x std x std
}

var Stdnorm *normdist = NewNormdist(0.0, 1.0)

func NewNormdist(m float64, s float64) *normdist {
	n := &normdist{
		stddev: s,
		mean:   m,
	}
	n.stddevsqpi = s * math.Pow((2.0*math.Pi), 0.5)
	if s == 1.0 {
		n.twostddevsq = 2.0
	} else {
		n.twostddevsq = 2.0 * (s * s)
	}
	return n
}

func (n *normdist) String() string {
	s := fmt.Sprintf("normdist {mean: %f, stddev: %f}", n.mean, n.stddev)
	return s
}

func (n *normdist) Mean() float64 {
	return n.mean
}

func (n *normdist) Stdev() float64 {
	return n.stddev
}

func (n *normdist) Pdf(x float64) float64 {
	var expon float64
	if n.mean == 0 {
		expon = -(x * x) / n.twostddevsq
	} else {
		expon = -(math.Pow((x - n.mean), 2)) / n.twostddevsq
	}
	probDist := math.Exp(expon) / n.stddevsqpi
	return probDist
}

func (n *normdist) Cdf(x float64) float64 {
	dist := x - n.mean
	if math.Abs(dist) > maxNumberOfStandardDeviations*n.stddev {
		if x < n.mean {
			return 0.0
		} else {
			return 1.0
		}
	}
	errf := Errf(dist / (n.stddev * sqrt2))
	cdf := 0.5 * (1.0 + errf)
	return cdf
}

func Errf(z float64) float64 {
	var t float64
	t = 1.0 / (1.0 + 0.5*math.Abs(z))
	ans := 1 - t*math.Exp(-z*z-1.26551223+
		t*(1.00002368+
			t*(0.37409196+
				t*(0.09678418+
					t*(-0.18628806+
						t*(0.27886807+
							t*(-1.13520398+
								t*(1.48851587+
									t*(-0.82215223+
										t*(0.17087277))))))))))
	if z >= 0 {
		return ans
	}
	return -ans
}
