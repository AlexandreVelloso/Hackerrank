using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
class Solution {

    static int theLoveLetterMystery(string s){
        int numberChange = 0;

        int i = 0;
        int j = s.Length-1;

        while( i < j ){
            numberChange += Math.Abs( s[i] - s[j] );
            i++;
            j--;
        }

        return numberChange;
    }

    static void Main(String[] args) {
        int q = Convert.ToInt32(Console.ReadLine());
        for(int a0 = 0; a0 < q; a0++){
            string s = Console.ReadLine();
            int result = theLoveLetterMystery(s);
            Console.WriteLine(result);
        }
    }
}
