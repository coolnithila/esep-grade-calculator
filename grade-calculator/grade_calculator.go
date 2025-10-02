package esepunittests

type GradeCalculator struct {
	assignments []Grade
	exams       []Grade
	essays      []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		assignments: make([]Grade, 0),
		exams:       make([]Grade, 0),
		essays:      make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	switch gradeType {
	case Assignment:
		gc.assignments = append(gc.assignments, Grade{
			Name:  name,
			Grade: grade,
			Type:  Assignment,
		})
	case Exam:
		gc.exams = append(gc.exams, Grade{
			Name:  name,
			Grade: grade,
			Type:  Exam,
		})
	case Essay:
		gc.essays = append(gc.essays, Grade{
			Name:  name,
			Grade: grade,
			Type:  Essay,
		})
	}
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignmentAverage := computeAverage(gc.assignments)
	examAverage := computeAverage(gc.exams)
	essayAverage := computeAverage(gc.essays) // IMPORTANT: essays here

	// Weights: 50% assignments, 35% exams, 15% essays
	weighted := float64(assignmentAverage)*0.50 +
		float64(examAverage)*0.35 +
		float64(essayAverage)*0.15

	// No special rounding rule given; truncation is OK
	return int(weighted)
}

func computeAverage(grades []Grade) int {
	if len(grades) == 0 {
		return 0 // avoid divide-by-zero; empty category counts as 0
	}
	sum := 0
	for _, g := range grades { // sum the actual Grade values
		sum += g.Grade
	}
	return sum / len(grades)
}
