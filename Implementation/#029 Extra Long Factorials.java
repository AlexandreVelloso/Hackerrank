import java.io.*;
import java.util.*;
import java.math.BigInteger;

public class ExtraLongFactorials {

    public static String factorial( int number ){
      BigInteger fact = new BigInteger("1");

      for( int i = 2; i <= number; i++ ){

        fact = fact.multiply( new BigInteger(""+i) );
      }

      return fact.toString();
    }

    public static void main(String[] args) {
        Scanner sc = new Scanner( System.in );

        // number to calculate the factorial
        int number;

        number = sc.nextInt();

        System.out.println( factorial(number) );

    }
}
