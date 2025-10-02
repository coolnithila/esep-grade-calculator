package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()
	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()
	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

// Was originally expected to be F, but math is ~96.9 → A.
func TestGetGradeF(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 95, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 91, Essay)

	actual_value := gradeCalculator.GetFinalGrade()
	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

// --- Extra tests for boundaries & coverage ---

func TestBoundaryAt90_IsA(t *testing.T) {
	expected_value := "A"
	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("a1", 90, Assignment)
	gradeCalculator.AddGrade("e1", 90, Exam)
	gradeCalculator.AddGrade("s1", 90, Essay)
	if got := gradeCalculator.GetFinalGrade(); got != expected_value {
		t.Errorf("Expected '%s'; got '%s'", expected_value, got)
	}
}

func TestBoundaryAt80_IsB(t *testing.T) {
	expected_value := "B"
	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("a1", 80, Assignment)
	gradeCalculator.AddGrade("e1", 80, Exam)
	gradeCalculator.AddGrade("s1", 80, Essay)
	if got := gradeCalculator.GetFinalGrade(); got != expected_value {
		t.Errorf("Expected '%s'; got '%s'", expected_value, got)
	}
}

func TestEmptyEssays_CountsAsZero(t *testing.T) {
	// 0.5*100 + 0.35*100 + 0.15*0 = 85 → B
	expected_value := "B"
	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("a1", 100, Assignment)
	gradeCalculator.AddGrade("e1", 100, Exam)
	if got := gradeCalculator.GetFinalGrade(); got != expected_value {
		t.Errorf("Expected '%s' with no essays; got '%s'", expected_value, got)
	}
}

func TestGradeTypeString_NotEmpty(t *testing.T) {
	expected_value := true
	actual_value := Assignment.String() != "" && Exam.String() != "" && Essay.String() != ""
	if expected_value != actual_value {
		t.Errorf("Expected GradeType.String() to be non-empty for all types; got %v", actual_value)
	}
}

func TestGetGradeC_AllSeventies(t *testing.T) {
	expected_value := "C"
	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("a1", 70, Assignment)
	gradeCalculator.AddGrade("e1", 70, Exam)
	gradeCalculator.AddGrade("s1", 70, Essay)
	if got := gradeCalculator.GetFinalGrade(); got != expected_value {
		t.Errorf("Expected %s, got %s", expected_value, got)
	}
}

func TestGetGradeD_AllSixties(t *testing.T) {
	expected_value := "D"
	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("a1", 60, Assignment)
	gradeCalculator.AddGrade("e1", 60, Exam)
	gradeCalculator.AddGrade("s1", 60, Essay)
	if got := gradeCalculator.GetFinalGrade(); got != expected_value {
		t.Errorf("Expected %s, got %s", expected_value, got)
	}
}

func TestGetGradeF_AllZeros(t *testing.T) {
	expected_value := "F"
	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("a1", 0, Assignment)
	gradeCalculator.AddGrade("e1", 0, Exam)
	gradeCalculator.AddGrade("s1", 0, Essay)
	if got := gradeCalculator.GetFinalGrade(); got != expected_value {
		t.Errorf("Expected %s, got %s", expected_value, got)
	}
}
