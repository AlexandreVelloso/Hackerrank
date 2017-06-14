package main;

import "fmt";

/**
* Test if a&b is greather than max
* but, a&b need to be less than k
*/
func getMax( a, b, k, max int) int{
	if( (a&b > max) && (a&b < k) ){
		max = a&b;
	}

	return max;
}

func problem(){

	var(
		//
		n int;
		//
		k int;
		//
		max = 0;
	);

	// read the input
	fmt.Scan( &n, &k );

	// for each value of i and j, where i < j
	for i := 1; i <= n; i++ {
		for j := i+1; j <= n; j++ {

			// get the biggest valid value
			max = getMax( i, j, k, max );
		}
	}

	fmt.Println(max);
}

func main(){

	// number of test cases
	var num_tests int;

	fmt.Scan( &num_tests );

	for i := 0; i < num_tests; i++ {

		problem();
	}// fim for
}// fim main
