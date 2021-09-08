//https://www.hackerrank.com/challenges/queue-using-two-stacks/problem

package algorithm

import "fmt"

func queue() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	// input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// inputArray := strings.Split(strings.TrimRight(input, "\n"), " ")
	var input int
	var queue  = make([]int, 0)

	fmt.Scanf("%d", &input)

	for i:=0 ; i < input; i++ {
		var input2 int
		fmt.Scanf("%d", &input2)
		if input2 == 1 {
			var input3 int
			fmt.Scanf("%d", &input3)
			queue = append(queue, input3)
		} else if input2 == 2 {
			queue = queue[1:]
		} else if input2 == 3 {
			fmt.Println(queue[0])
		}
	}
}
