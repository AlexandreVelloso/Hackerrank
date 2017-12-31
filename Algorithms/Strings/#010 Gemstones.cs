using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
class Solution {

    static int gemstones(string[] arr){

        // each line represents a string
        // each column a occurrence of a character: 0 = a, 1 = b,...
        int[][] occurrences = new int[ arr.Length ][];

        for( int i = 0; i < arr.Length; i++ ){
            occurrences[i] = new int[ 26 ];
        }

        // count the occurrences of each string
        for( int i = 0; i < arr.Length; i++ ){

            // in each character from arr[i]
            foreach( char c in arr[i] ){
                occurrences[i][ c-'a' ]++;
            }
        }

        // count the number of letters that appears in all strings
        int count = 0;
        for( int i = 0; i < 26; i++ ){

            // count the number of lines that have this letter
            int n = 0;
            for( int j = 0; j < arr.Length; j++ ){
                if( occurrences[j][i] != 0 ){
                    n++;
                }
            }

            // if appears in all lines, count this letter
            if( n == arr.Length ){
                count++;
            }
        }

        return count;
    }

    static void Main(String[] args) {
        int n = Convert.ToInt32(Console.ReadLine());
        string[] arr = new string[n];
        for(int arr_i = 0; arr_i < n; arr_i++){
           arr[arr_i] = Console.ReadLine();
        }
        int result = gemstones(arr);
        Console.WriteLine(result);
    }
}
