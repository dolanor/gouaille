package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	stacks := make([]int, 4)

	fillStacks(stacks)

	ended := false
	reader := bufio.NewReader(os.Stdin)
	for ended == false {
		fmt.Println(stacks)
		fmt.Print("Which stack do you choose ? [1-4]: ")
		line, _ := reader.ReadString('\n')
		line = line[:len(line)-1]
		stacknum, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Not the right format")
			continue
		}
		stacknum--
		if stacknum < 0 || stacknum > 3 {
			fmt.Println("Not the right stack number")
			continue
		}

		fmt.Print("How many matches will you take ? [1-", stacks[stacknum], "] (0 to choose another stack): ")
		line, _ = reader.ReadString('\n')
		line = line[:len(line)-1]
		matchesnum, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Not the right format")
			continue
		}
		if matchesnum < 1 || matchesnum > stacks[stacknum] {
			fmt.Println("You can't take that amount of matches")
			continue
		}
		stacks[stacknum] -= matchesnum
		if youWin(stacks) {
			fmt.Println("Congratulations, you won!")
			continue
		}

		fmt.Print("\nCPU: My Turn finally!\n\n")
		howMuchToPick := 0
		stackToPickFrom := -1
		for t := range stacks {
			for j := 1; j <= stacks[t]; j++ {
				stacks[t] -= j
				if stacks[0]^stacks[1]^stacks[2]^stacks[3] == 0 {
					howMuchToPick = j
					stackToPickFrom = t
				}
				stacks[t] += j
			}
		}
		if stackToPickFrom == -1 {
			fmt.Println("Damn it. I think I lost")
			stackToPickFrom, howMuchToPick = pickRandomely(stacks)
		}

		stacks[stackToPickFrom] -= howMuchToPick
		fmt.Println("I took ", howMuchToPick, " from the stack #", stackToPickFrom+1)
		if youWin(stacks) {
			fmt.Println("Too bad, you lost! CPU won")
			continue
		}

	}

}

func pickRandomely(stacks []int) (int, int) {
	i := 0
	for stacks[i] == 0 {
		i = rand.Intn(3)
	}
	num := rand.Intn(stacks[i])
	if num == 0 {
		num++
	}

	return i, num
}

func youWin(stacks []int) bool {
	sum := 0
	for _, i := range stacks {
		sum += i
	}
	if sum == 0 {
		fillStacks(stacks)
		return true
	}
	return false
}

func fillStacks(stacks []int) {
	stacks[0] = 1
	stacks[1] = 3
	stacks[2] = 5
	stacks[3] = 7
}
