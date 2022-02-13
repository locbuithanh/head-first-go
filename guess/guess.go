// Chalenge user to guess a random number
package guess

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	t := rand.Intn(100) + 1

	for i := 0; i < 10; i++ {
		fmt.Print("Input: ")
		r := bufio.NewReader(os.Stdin)
		s, err := r.ReadString('\n')
		if err != nil {
			log.Panic(err)
		}
		tt, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			log.Fatal(err)
		}
		if tt < t {
			fmt.Println("Low")
		} else if tt > t {
			fmt.Println("High")
		} else {
			fmt.Println("Right")
			break
		}
		if i == 9 {
			fmt.Println("Out of chances", t)
		}
	}

}
