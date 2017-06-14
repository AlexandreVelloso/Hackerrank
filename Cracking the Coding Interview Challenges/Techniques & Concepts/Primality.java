import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;

public class Primality {

    public static boolean isPrime( int n ){
        int maxDivisor = (int)Math.sqrt( n );

        int numDivisors = 2;

        for ( int i = 2; i <= maxDivisor; i++ ){
            if( n%i == 0 )
                numDivisors++;
        }

        return ( (numDivisors == 2 && n != 1) ? true : false );
    }

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int p = in.nextInt();
        for(int a0 = 0; a0 < p; a0++){
            int n = in.nextInt();

            System.out.println( isPrime(n) ? "Prime" : "Not prime" );
        }
    }
}
