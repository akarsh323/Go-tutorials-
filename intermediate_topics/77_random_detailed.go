package intermediate

import (
	"crypto/rand"
	"fmt"
	"math/big"
	mathRand "math/rand"
	"time"
)

// Topic 77: Random Numbers - Core Mental Models
// ==============================================
// This lesson breaks down how pseudo-random number generation works.
// We learn about seeds, ranges, the PRNG concept, and critical warnings
// about security and concurrency.

func main() {
	fmt.Println("=== Topic 77: Random Numbers - Understanding Randomness ===\n")

	lesson1SeedsAndDeterminism()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson2WindowShiftingTechnique()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson3DiceGameLogic()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson4CryptoRandomness()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson5ConcurrencyWarnings()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson6TeachersInsight()
}

// LESSON 1: Seeds and Determinism - The Playlist Analogy
// =====================================================
func lesson1SeedsAndDeterminism() {
	fmt.Println("LESSON 1: SEEDS AND DETERMINISM - THE PLAYLIST ANALOGY")
	fmt.Println("------------------------------------------------------\n")

	fmt.Println("THE CORE CONCEPT:")
	fmt.Println("  A Pseudo-Random Number Generator (PRNG) is like a shuffled playlist.\n")

	fmt.Println("THE ANALOGY:")
	fmt.Println("┌──────────────────────────────────────────────────────┐")
	fmt.Println("│  The Seed        = Specific \"Shuffle Order\" you choose │")
	fmt.Println("│  Random Numbers  = Songs playing one after another    │")
	fmt.Println("└──────────────────────────────────────────────────────┘\n")

	fmt.Println("SCENARIO A: Fixed Seed (Deterministic/Predictable)")
	fmt.Println("  Code: rand.Seed(42)  // Always use seed 42")
	fmt.Println("  Result:")
	fmt.Println("    • Every time you run the program, you get the SAME random numbers")
	fmt.Println("    • Like saying: \"Always play Shuffle Order #42\"\n")

	fmt.Println("  Demonstration:")
	mathRand.Seed(42)
	fmt.Printf("    Run 1: %d, %d, %d\n", mathRand.Intn(100), mathRand.Intn(100), mathRand.Intn(100))
	mathRand.Seed(42)
	fmt.Printf("    Run 2: %d, %d, %d\n", mathRand.Intn(100), mathRand.Intn(100), mathRand.Intn(100))
	fmt.Println("    (Notice: Run 1 and Run 2 are identical!)\n")

	fmt.Println("  When to use:")
	fmt.Println("    ✓ Debugging (reproduce the bug with same numbers)")
	fmt.Println("    ✓ Testing (predictable test data)")
	fmt.Println("    ✓ Games (allow players to replay with same seed)\n")

	fmt.Println("SCENARIO B: Dynamic Seed (Truly Random)")
	fmt.Println("  Code: rand.Seed(time.Now().UnixNano())")
	fmt.Println("  Result:")
	fmt.Println("    • Every millisecond, the time changes")
	fmt.Println("    • So every time you run the program, you get DIFFERENT random numbers\n")

	fmt.Println("  Demonstration:")
	mathRand.Seed(time.Now().UnixNano())
	fmt.Printf("    Run 1: %d, %d, %d\n", mathRand.Intn(100), mathRand.Intn(100), mathRand.Intn(100))
	time.Sleep(10 * time.Millisecond)
	mathRand.Seed(time.Now().UnixNano())
	fmt.Printf("    Run 2: %d, %d, %d\n", mathRand.Intn(100), mathRand.Intn(100), mathRand.Intn(100))
	fmt.Println("    (Notice: Run 1 and Run 2 are different!)\n")

	fmt.Println("  When to use:")
	fmt.Println("    ✓ Games and simulations (unpredictable gameplay)")
	fmt.Println("    ✓ Sampling and testing\n")

	fmt.Println("MODERN GO (Version 1.20+):")
	fmt.Println("  Go automatically seeds with time for you!")
	fmt.Println("  You DON'T need to write rand.Seed() anymore.")
	fmt.Println("  (But if you want deterministic behavior, you still can.)")
}

// LESSON 2: Window Shifting - The Range Problem
// ==============================================
func lesson2WindowShiftingTechnique() {
	fmt.Println("LESSON 2: WINDOW SHIFTING - THE RANGE PROBLEM")
	fmt.Println("---------------------------------------------\n")

	fmt.Println("THE FUNDAMENTAL ISSUE:")
	fmt.Println("  Computers count from 0.")
	fmt.Println("  Humans count from 1.")
	fmt.Println("  This creates friction when generating ranges.\n")

	fmt.Println("HOW rand.Intn(N) WORKS:")
	fmt.Println("  rand.Intn(N) returns a random number from 0 to (N-1)")
	fmt.Println("  Examples:")
	fmt.Println("    • rand.Intn(10)  → {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}")
	fmt.Println("    • rand.Intn(100) → {0, 1, 2, ..., 98, 99}")
	fmt.Println("    • rand.Intn(6)   → {0, 1, 2, 3, 4, 5}\n")

	fmt.Println("THE DICE PROBLEM:")
	fmt.Println("  Requirements: Generate a number 1-6 (a dice roll)")
	fmt.Println("  Problem:      rand.Intn(6) gives us 0-5, not 1-6\n")

	fmt.Println("THE SOLUTION: Add a Shift")
	fmt.Println("  Add the starting point to shift the \"window\":")
	fmt.Println("  rand.Intn(6) + 1\n")

	fmt.Println("  Visualization:")
	fmt.Println("    Computer generates:  0  1  2  3  4  5")
	fmt.Println("    We add 1 to shift:   +1 +1 +1 +1 +1 +1")
	fmt.Println("    Result (1-6):        1  2  3  4  5  6\n")

	fmt.Println("PRACTICAL DEMONSTRATION:")
	fmt.Println("  Rolling 10 dice (1-6):")
	mathRand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		roll := mathRand.Intn(6) + 1
		fmt.Printf("    Roll %d: %d\n", i+1, roll)
	}
	fmt.Println()

	fmt.Println("THE GENERAL FORMULA:")
	fmt.Println("  To get a random number between min and max (inclusive):")
	fmt.Println("    rand.Intn(max - min + 1) + min\n")

	fmt.Println("  Examples:")
	fmt.Println("    Between 1-6:        rand.Intn(6 - 1 + 1) + 1 = rand.Intn(6) + 1")
	fmt.Println("    Between 10-20:      rand.Intn(20 - 10 + 1) + 10 = rand.Intn(11) + 10")
	fmt.Println("    Between 1990-2020:  rand.Intn(2020 - 1990 + 1) + 1990 = rand.Intn(31) + 1990\n")

	fmt.Println("CUSTOM RANGE EXAMPLE: Random Year (1990-2020)")
	minYear := 1990
	maxYear := 2020
	fmt.Println("  Code:")
	fmt.Printf("    randomYear := rand.Intn(maxYear - minYear + 1) + minYear\n")
	fmt.Printf("    randomYear := rand.Intn(2020 - 1990 + 1) + 1990\n")
	fmt.Printf("    randomYear := rand.Intn(31) + 1990\n\n")

	fmt.Println("  Generated years:")
	mathRand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		year := mathRand.Intn(maxYear-minYear+1) + minYear
		fmt.Printf("    Year %d: %d\n", i+1, year)
	}
	fmt.Println()

	fmt.Println("WHY THIS MATTERS:")
	fmt.Println("  ✓ Many real problems need custom ranges")
	fmt.Println("  ✓ Off-by-one errors are common bugs")
	fmt.Println("  ✓ Understanding the formula prevents these bugs")
}

// LESSON 3: Dice Game Logic - Putting It Together
// ===============================================
func lesson3DiceGameLogic() {
	fmt.Println("LESSON 3: DICE GAME LOGIC - PUTTING IT TOGETHER")
	fmt.Println("----------------------------------------------\n")

	fmt.Println("A SIMPLE DICE GAME STRUCTURE:\n")

	fmt.Println("Code Structure:")
	fmt.Println("  for {")
	fmt.Println("      // 1. Generate random numbers")
	fmt.Println("      die1 := rand.Intn(6) + 1    // Die 1: 1-6")
	fmt.Println("      die2 := rand.Intn(6) + 1    // Die 2: 1-6")
	fmt.Println("")
	fmt.Println("      // 2. Show result")
	fmt.Println("      fmt.Printf(\"You rolled: %d and %d\\n\", die1, die2)")
	fmt.Println("")
	fmt.Println("      // 3. Process result")
	fmt.Println("      total := die1 + die2")
	fmt.Println("      fmt.Printf(\"Total: %d\\n\", total)")
	fmt.Println("")
	fmt.Println("      // 4. Exit condition")
	fmt.Println("      if userWantsToExit {")
	fmt.Println("          break")
	fmt.Println("      }")
	fmt.Println("  }\n")

	fmt.Println("LOGIC BREAKDOWN:")
	fmt.Println("  Step 1: Generate Numbers")
	fmt.Println("    • rand.Intn(6) gives 0-5")
	fmt.Println("    • + 1 shifts to 1-6")
	fmt.Println("    • Do this twice for two dice\n")

	fmt.Println("  Step 2: Display")
	fmt.Println("    • Show the player what they rolled\n")

	fmt.Println("  Step 3: Process")
	fmt.Println("    • Add them up, check win conditions, etc.\n")

	fmt.Println("  Step 4: Loop Control")
	fmt.Println("    • Ask if player wants to roll again")
	fmt.Println("    • Break if they say no\n")

	fmt.Println("SIMULATION (5 rolls):")
	mathRand.Seed(time.Now().UnixNano())
	for roll := 1; roll <= 5; roll++ {
		die1 := mathRand.Intn(6) + 1
		die2 := mathRand.Intn(6) + 1
		total := die1 + die2

		fmt.Printf("  Roll %d: [%d] + [%d] = %d", roll, die1, die2, total)

		// Simple win condition (example)
		if total == 7 || total == 11 {
			fmt.Println(" ✓ LUCKY!")
		} else {
			fmt.Println()
		}
	}
}

// LESSON 4: Cryptographic Randomness - The Security Warning
// ==========================================================
func lesson4CryptoRandomness() {
	fmt.Println("LESSON 4: CRYPTOGRAPHIC RANDOMNESS - THE SECURITY WARNING")
	fmt.Println("---------------------------------------------------------\n")

	fmt.Println("TWO TYPES OF RANDOMNESS IN GO:\n")

	fmt.Println("1. MATH/RAND (Pseudo-Random, Fast, Predictable)")
	fmt.Println("   Package: math/rand")
	fmt.Println("   Speed:   FAST ⚡")
	fmt.Println("   Security: WEAK 🔓\n")

	fmt.Println("   How it works:")
	fmt.Println("     • Uses a mathematical formula")
	fmt.Println("     • If hacker knows a few numbers, can predict the next")
	fmt.Println("     • Perfect for games and simulations\n")

	fmt.Println("   Use cases:")
	fmt.Println("     ✓ Games and simulations")
	fmt.Println("     ✓ Sampling and testing")
	fmt.Println("     ✓ Anything where predictability is OK\n")

	fmt.Println("2. CRYPTO/RAND (Cryptographically Secure)")
	fmt.Println("   Package: crypto/rand")
	fmt.Println("   Speed:   SLOW 🐢")
	fmt.Println("   Security: STRONG 🔒\n")

	fmt.Println("   How it works:")
	fmt.Println("     • Uses system entropy (truly random)")
	fmt.Println("     • Mathematically impossible to predict")
	fmt.Println("     • Suitable for security-critical operations\n")

	fmt.Println("   Use cases:")
	fmt.Println("     ✓ API keys and tokens")
	fmt.Println("     ✓ Password resets")
	fmt.Println("     ✓ Security certificates")
	fmt.Println("     ✓ Anything with security implications\n")

	fmt.Println("DEMONSTRATION: Cryptographic Random Bytes")
	fmt.Println("  Generating 16 random bytes (for API key):\n")

	cryptoBytes := make([]byte, 16)
	_, err := rand.Read(cryptoBytes)
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
	} else {
		fmt.Printf("  Raw bytes:     %v\n", cryptoBytes)
		fmt.Printf("  Hex encoding:  %x\n\n", cryptoBytes)
	}

	fmt.Println("DEMONSTRATION: Cryptographic Random Number (1-100)")
	fmt.Println("  Code:")
	fmt.Println("    bigMax := big.NewInt(100)")
	fmt.Println("    randomNum, _ := rand.Int(rand.Reader, bigMax)")
	fmt.Println("    fmt.Println(randomNum)\n")

	bigMax := big.NewInt(100)
	randomNum, _ := rand.Int(rand.Reader, bigMax)
	fmt.Printf("  Result: %d (0-99)\n\n", randomNum)

	fmt.Println("THE CRITICAL WARNING:")
	fmt.Println("┌─────────────────────────────────────────────────┐")
	fmt.Println("│ NEVER use math/rand for security-critical code! │")
	fmt.Println("│ ALWAYS use crypto/rand for passwords, tokens,  │")
	fmt.Println("│ keys, and any security-sensitive operations.   │")
	fmt.Println("└─────────────────────────────────────────────────┘")
}

// LESSON 5: Concurrency and Race Conditions
// =========================================
func lesson5ConcurrencyWarnings() {
	fmt.Println("LESSON 5: CONCURRENCY AND RACE CONDITIONS")
	fmt.Println("----------------------------------------\n")

	fmt.Println("THE CONCURRENCY TRAP:")
	fmt.Println("  When using Go routines (goroutines), you run multiple things at once.\n")

	fmt.Println("THE PROBLEM:")
	fmt.Println("  The default math/rand generator is GLOBAL (shared by all goroutines).")
	fmt.Println("  If two goroutines ask for a random number at the EXACT SAME time,")
	fmt.Println("  they might clash → Race Condition → Corrupted data.\n")

	fmt.Println("VISUAL EXAMPLE:\n")
	fmt.Println("  Timeline (both goroutines running):")
	fmt.Println("    Goroutine A: [Request random] ← Time 1.000ms")
	fmt.Println("    Goroutine B:          [Request random] ← Time 1.001ms")
	fmt.Println("    Problem: They might interfere with each other!\n")

	fmt.Println("SOLUTION 1: Use sync.Mutex (Lock access)")
	fmt.Println("  Code pattern:")
	fmt.Println("    var mu sync.Mutex")
	fmt.Println("")
	fmt.Println("    go func() {")
	fmt.Println("        mu.Lock()  // Only one goroutine at a time")
	fmt.Println("        num := rand.Intn(100)")
	fmt.Println("        mu.Unlock()")
	fmt.Println("    }()\n")

	fmt.Println("  Effect:")
	fmt.Println("    • Goroutines take turns (serialized)")
	fmt.Println("    • No race condition")
	fmt.Println("    • Slower due to lock overhead\n")

	fmt.Println("SOLUTION 2: Give each goroutine its own generator")
	fmt.Println("  Code pattern:")
	fmt.Println("    go func() {")
	fmt.Println("        localRand := mathRand.New(mathRand.NewSource(time.Now().UnixNano()))")
	fmt.Println("        num := localRand.Intn(100)")
	fmt.Println("    }()\n")

	fmt.Println("  Effect:")
	fmt.Println("    • Each goroutine has its own PRNG")
	fmt.Println("    • No contention or locking")
	fmt.Println("    • Much faster\n")

	fmt.Println("BEST PRACTICE:")
	fmt.Println("  If you have heavy concurrent random number generation:")
	fmt.Println("    → Use Solution 2 (separate generators)")
	fmt.Println("  If you have light concurrent access:")
	fmt.Println("    → Solution 1 (mutex) is simpler to implement")
}

// LESSON 6: Teacher's Insight - Mental Models Summary
// =================================================
func lesson6TeachersInsight() {
	fmt.Println("LESSON 6: TEACHER'S INSIGHT - MENTAL MODELS SUMMARY")
	fmt.Println("---------------------------------------------------\n")

	fmt.Println("YOU NOW UNDERSTAND:")
	fmt.Println("  1. Seeding = Choosing the Shuffle Order")
	fmt.Println("     • Fixed seed = Deterministic (good for debugging)")
	fmt.Println("     • Time-based seed = Random (good for games)\n")

	fmt.Println("  2. Ranges = Window Shifting Problem")
	fmt.Println("     • Formula: rand.Intn(max - min + 1) + min")
	fmt.Println("     • Off-by-one errors are common\n")

	fmt.Println("  3. Two Types of Random")
	fmt.Println("     • math/rand = Fast but predictable (games)")
	fmt.Println("     • crypto/rand = Slow but secure (passwords)\n")

	fmt.Println("  4. Concurrency Requires Care")
	fmt.Println("     • Global generator = Race condition risk")
	fmt.Println("     • Solutions: Mutex or separate generators\n")

	fmt.Println("THE CORE PRINCIPLE:")
	fmt.Println("┌──────────────────────────────────────────────────────┐")
	fmt.Println("│ Randomness is a TOOL with specific purposes:        │")
	fmt.Println("│ • Use math/rand for everything not security-related │")
	fmt.Println("│ • Use crypto/rand for anything with secrets         │")
	fmt.Println("│ • Be aware of concurrency implications              │")
	fmt.Println("└──────────────────────────────────────────────────────┘\n")

	fmt.Println("COMMON MISTAKES TO AVOID:")
	fmt.Println("  ✗ Using math/rand for passwords")
	fmt.Println("  ✗ Forgetting to shift ranges (off-by-one bugs)")
	fmt.Println("  ✗ Ignoring race conditions in concurrent code")
	fmt.Println("  ✗ Not testing edge cases (1, max values)")
	fmt.Println("  ✗ Seeding multiple times in a loop (wastes entropy)\n")

	fmt.Println("YOU'RE NOW READY FOR:")
	fmt.Println("  ✓ Building games and simulations")
	fmt.Println("  ✓ Generating test data")
	fmt.Println("  ✓ Creating secure tokens and keys")
	fmt.Println("  ✓ Working with concurrent random operations")
	fmt.Println("  ✓ Debugging reproducible test failures")
}
