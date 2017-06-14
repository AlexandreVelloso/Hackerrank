import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;

public class Balanced_Brackets {

  public static boolean isBalanced(String expression) {
    Stack<Character> stack = new Stack<>();

    char bracket;
    for ( int i = 0; i < expression.length(); i++ ) {

      bracket = expression.charAt( i );

      switch( bracket ){

        case ')':
          if( stack.empty() || stack.pop() != '(' )
            return false;
          break;

        case ']':
          if( stack.empty() || stack.pop() != '[' )
            return false;
          break;

        case '}':
          if( stack.empty() || stack.pop() != '{' )
            return false;
          break;

        default:
          stack.push( bracket );
          break;
      }// end switch

    }// end for

    if( stack.empty() == false )
      return false;

    return true;
  }// end method isBalanced

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int t = in.nextInt();
        for (int a0 = 0; a0 < t; a0++) {

            String expression = in.next();
            System.out.println( (isBalanced(expression)) ? "YES" : "NO" );

            //if( a0 == 582 || a0 == 582 ) System.out.println(expression);
        }
    }
}
