using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
class Solution {

    static string funnyString(string s){

        string result = "Funny";
        int size = s.Length;

        // Complete this function
        for( int i = 1; i < size; i++ ){

            int num1 = Math.Abs(s[i] - s[i-1]);
            int num2 = Math.Abs(s[size-i-1] - s[size-i]);

            if( num1 != num2 ){
                result = "Not Funny";
                break;
            }
        }

        return result;
    }

    static void Main(String[] args) {
        int q = Convert.ToInt32(Console.ReadLine());
        for(int a0 = 0; a0 < q; a0++){
            string s = Console.ReadLine();
            string result = funnyString(s);
            Console.WriteLine(result);
        }
    }
}
