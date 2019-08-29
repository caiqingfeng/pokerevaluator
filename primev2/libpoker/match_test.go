package libpoker

import "testing"

func TestMatch(t *testing.T) {
	BuildScoreTbl()
	match := Anb{"5d6dJcJh7d", 0, "Js7cKdKh3c", 0, 0}
	ProcessMatch(&match)
	if match.Result != 2 {
		// fmt.Println(alice, bob)
		t.Error("failed")
	}

	nmatch := Anb{"AsKsQsJsTs",0, "QsQhQdQcJh", 0,0}
	ProcessMatch(&nmatch)
	if nmatch.Result != 1 {
		// fmt.Println(alice, bob)
		t.Error("failed", nmatch.Alice, nmatch.Bob,
			nmatch.Result, nmatch.AliceRank, nmatch.BobRank)
	}

	mmatch := Anb{"9h6c6s9c3s8hXn",0, "6h2h9h6c6s9c3s", 0,0}
	ProcessMatch(&mmatch)
	if mmatch.Result != 1 {
		// fmt.Println(alice, bob)
		t.Error("failed", mmatch.Alice, mmatch.Bob,
			mmatch.Result, mmatch.AliceRank, mmatch.BobRank)
	}

	mmatch = Anb{"8dJd8hQd6h6cXn",0, "8cQc8dJd8hQd6h", 0,0}
	ProcessMatch(&mmatch)
	if mmatch.Result != 2 {
		// fmt.Println(alice, bob)
		t.Error("failed", mmatch.Alice, mmatch.Bob,
			mmatch.Result, mmatch.AliceRank, mmatch.BobRank)
	}

	mmatch = Anb{"4c3c4h2d3sJd3d",0, "XnQh4c3c4h2d3s", 0,0}
	ProcessMatch(&mmatch)
	if mmatch.Result != 2 {
		// fmt.Println(alice, bob)
		t.Error("failed", mmatch.Alice, mmatch.Bob,
			mmatch.Result, mmatch.AliceRank, mmatch.BobRank)
	}

	mmatch = Anb{"4cAdThTc4hXnKs",0, "TdQc4cAdThTc4h", 0,0}
	ProcessMatch(&mmatch)
	if mmatch.Result != 0 {
		// fmt.Println(alice, bob)
		t.Error("failed", mmatch.Alice, mmatch.Bob,
			mmatch.Result, mmatch.AliceRank, mmatch.BobRank)
	}

	mmatch = Anb{"6cAh6h2h6d9h8h",0, "5hKs6cAh6h2h6d", 0,0}
	ProcessMatch(&mmatch)
	if mmatch.Result != 1 {
		// fmt.Println(alice, bob)
		t.Error("failed", mmatch.Alice, mmatch.Bob,
			mmatch.Result, mmatch.AliceRank, mmatch.BobRank)
	}

	mmatch = Anb{"6cQdQsTdKd3dQh",0, "Xn3h6cQdQsTdKd", 0,0}
	ProcessMatch(&mmatch)
	if mmatch.Result != 0 {
		// fmt.Println(alice, bob)
		t.Error("failed", mmatch.Alice, mmatch.Bob,
			mmatch.Result, mmatch.AliceRank, mmatch.BobRank)
	}

	mmatch = Anb{"2s2hKh8h9s5s4h",0, "7s4s2s2hKh8h9s", 0,0}
	ProcessMatch(&mmatch)
	if mmatch.Result != 0 {
		// fmt.Println(alice, bob)
		t.Error("failed", mmatch.Alice, mmatch.Bob,
			mmatch.Result, mmatch.AliceRank, mmatch.BobRank)
	}

	mmatch = Anb{"6d9h3hKdQdKh5d",0, "4sXn6d9h3hKdQd", 0,0}
	ProcessMatch(&mmatch)
	if mmatch.Result != 0 {
		// fmt.Println(alice, bob)
		t.Error("failed", mmatch.Alice, mmatch.Bob,
			mmatch.Result, mmatch.AliceRank, mmatch.BobRank)
	}

	mmatch = Anb{"Xn9cKd3d2d9d9s",0, "5hTdXn9cKd3d2d", 0,0}
	ProcessMatch(&mmatch)
	if mmatch.Result != 1 {
		// fmt.Println(alice, bob)
		t.Error("failed", mmatch.Alice, mmatch.Bob,
			mmatch.Result, mmatch.AliceRank, mmatch.BobRank)
	}

	mmatch = Anb{"Xn2c4c8hKd8c5c",0, "6cQsXn2c4c8hKd", 0,0}
	ProcessMatch(&mmatch)
	if mmatch.Result != 1 {
		// fmt.Println(alice, bob)
		t.Error("failed", mmatch.Alice, mmatch.Bob,
			mmatch.Result, mmatch.AliceRank, mmatch.BobRank)
	}
}

func TestMatches(t *testing.T) {
	BuildScoreTbl()
	mmatch := Anb{"As2h7sXn5hTs6s",0, "6h3hAs2h7sXn5h", 0,0}
	ProcessMatch(&mmatch)
	if mmatch.Result != 2 {
		// fmt.Println(alice, bob)
		t.Error("failed", mmatch.Alice, mmatch.Bob,
			mmatch.Result, mmatch.AliceRank, mmatch.BobRank)
	}

	mmatch = Anb{"7sXn3sJsTs8s3d",0, "Ks2h7sXn3sJsTs", 0,0}
	ProcessMatch(&mmatch)
	if mmatch.Result != 1 {
		// fmt.Println(alice, bob)
		t.Error("failed", mmatch.Alice, mmatch.Bob,
			mmatch.Result, mmatch.AliceRank, mmatch.BobRank)
	}

	mmatch = Anb{"8s2sXn9s9h7h7s",0, "Ts6s8s2sXn9s9h", 0,0}
	ProcessMatch(&mmatch)
	if mmatch.Result != 2 {
		// fmt.Println(alice, bob)
		t.Error("failed", mmatch.Alice, mmatch.Bob,
			mmatch.Result, mmatch.AliceRank, mmatch.BobRank)
	}
}

func TestMatch1(t *testing.T) {
	BuildScoreTbl()
	alice := EvaluateHandStr("5d6dJcJh7d")
	bob := EvaluateHandStr("Js7cKdKh3c")
	if !(alice > bob) {
		// fmt.Println(alice, bob)
		t.Error("failed")
	}
}

func TestMatch2(t *testing.T) {
	BuildScoreTbl()
	alice := EvaluateHandStr("6d6s3s3d6hJsQs")
	bob := EvaluateHandStr("8h7c6d6s3s3d6h")
	if alice != bob {
		// fmt.Println(alice, bob)
		t.Error("failed")
	}
}

func TestMatch3(t *testing.T) {
	BuildScoreTbl()
	alice := EvaluateHandStr("8c5s8dAhAs7d5d")
	bob := EvaluateHandStr("6s3d8c5s8dAhAs")
	if !(alice < bob) {
		// fmt.Println(alice, bob)
		t.Error("failed")
	}
}

func TestMatch4(t *testing.T) {
	BuildScoreTbl()
	alice := EvaluateHandStr("TsTc4d4cTdAh3h")
	bob := EvaluateHandStr("2c4hTsTc4d4cTd")
	if !(alice == bob) {
		// fmt.Println(alice, bob)
		t.Error("failed")
	}
}

func TestMatch5(t *testing.T) {
	BuildScoreTbl()
	alice := EvaluateHandStr("KdQs4hKh7c7d8c")
	bob := EvaluateHandStr("4sQcKdQs4hKh7c")
	if !(alice > bob) {
		// fmt.Println(alice, bob)
		t.Error("failed")
	}
}

func TestMatch6(t *testing.T) {
	BuildScoreTbl()
	alice := EvaluateHandStr("As8dKh8hJhAc5d")
	bob := EvaluateHandStr("JdAhAs8dKh8hJh")
	if !(alice > bob) {
		// fmt.Println(alice, bob)
		t.Error("failed")
	}
}

func TestMatch7(t *testing.T) {
	BuildScoreTbl()
	alice := EvaluateHandStr("5h4d5c5d9d9hAc")
	bob := EvaluateHandStr("9s2s5h4d5c5d9d")
	if !(alice == bob) {
		// fmt.Println(alice, bob)
		t.Error("failed")
	}
}

func TestMatch8(t *testing.T) {
	BuildScoreTbl()
	alice := EvaluateHandStr("9cKs9hTdJcKc7s")
	bob := EvaluateHandStr("KhTc9cKs9hTdJc")
	if !(alice > bob) {
		// fmt.Println(alice, bob)
		t.Error("failed")
	}
	alice = EvaluateHandStr("3h6c5dAsQd")
	bob = EvaluateHandStr("Qs7d7hQcTs")
	if alice < bob {
		// fmt.Println(alice, bob)
		t.Error("failed", alice, bob)
	}
}