package handler

import "testing"

func testFactory(t *testing.T, s string, r []string) {
	segments := GetSegments(s)
	if len(segments) != len(r) {
		t.Fatalf("len(GetSegments(%s)) = %d, want %v", s, len(segments), len(r))
		return
	}
	for i, part := range r {
		if part != segments[i] {
			t.Fatalf("GetSegments(%s) = %v, want %v", s, segments, r)
			return
		}
	}
}

func TestGetSegmentsEmpty(t *testing.T) {
	testFactory(t, "", make([]string, 0))
}

func TestGetSegmentsSingleSplash(t *testing.T) {
	testFactory(t, "/", make([]string, 0))
}

func TestGetSegmentsRelativeSingleWord(t *testing.T) {
	r := []string{"test"}
	testFactory(t, "test", r)
}

func TestGetSegmentsRelativeWithEndSingleWord(t *testing.T) {
	r := []string{"test"}
	testFactory(t, "test/", r)
}

func TestGetSegmentsAbsoluteSingleWord(t *testing.T) {
	r := []string{"test"}
	testFactory(t, "/test", r)
}

func TestGetSegmentsAbsoluteWithEndSingleWord(t *testing.T) {
	r := []string{"test"}
	testFactory(t, "/test/", r)
}

func TestGetSegmentsRelativeTwoWord(t *testing.T) {
	r := []string{"test", "test"}
	testFactory(t, "test/test", r)
}

func TestGetSegmentsRelativeWithEndTwoWord(t *testing.T) {
	r := []string{"test", "test"}
	testFactory(t, "test/test/", r)
}

func TestGetSegmentsAbsoluteTwoWord(t *testing.T) {
	r := []string{"test", "test"}
	testFactory(t, "/test/test", r)
}

func TestGetSegmentsAbsoluteWithEndTwoWord(t *testing.T) {
	r := []string{"test", "test"}
	testFactory(t, "/test/test", r)
}

func TestGetSegmentsEmptyMiddle(t *testing.T) {
	r := []string{"test", "", "test"}
	testFactory(t, "/test//test", r)
}

func TestGetSegmentsEmptyEnd(t *testing.T) {
	r := []string{"test", "test", ""}
	testFactory(t, "/test/test//", r)
}
