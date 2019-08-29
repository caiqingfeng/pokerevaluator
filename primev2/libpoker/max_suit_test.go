package libpoker

import "testing"

func TestMaxSuit(t *testing.T) {
	if hand, count := maxSuit("AcKdQhJsTc6c4s"); count != 3 || hand!="AcTc6c"{
		t.Error("failed", hand, count)
	}

	if hand, count := maxSuit("AcXnQhJsTc6c4s"); count != 4 || hand!="AcTc6cXn"{
		t.Error("failed", hand, count)
	}

	if hand, count := maxSuit("5hTdXn9cKd3d2d"); count != 5 || hand!="TdKd3d2dAd"{
		t.Error("failed", hand, count)
	}

	if hand, count := maxSuit("Ts6s8s2sXn9s9h"); count != 5 || hand!="7sTs9s8s6s"{
		t.Error("failed", hand, count)
	}
}

func TestMaxFace(t *testing.T) {
	if hand := maxFace("AcKdQhJsTc6c4s"); hand!="AcKdQhJsTc"{
		t.Error("failed", hand)
	}

	if hand := maxFace("Ac2cQhJsTc6c4s"); hand!="AcQhJsTc6c"{
		t.Error("failed", hand)
	}
}
