import java.io.*;
import java.util.*;
import java.util.regex.*;

public class Solution {

  private static String regexString;
  private static Pattern p;

  public static List<String> readInput( ){
    Scanner sc = new Scanner( System.in );

    // size of the array from input
    int size = Integer.parseInt( sc.nextLine() );

    // list of names that uses gmail
    List<String> list = new LinkedList<>();

    for (int i = 0; i < size; i++) {

      String name = sc.next();
      String email = sc.next();

      if( match(email) ){
        list.add(name);
      }
    }

    return list;
  }

  /**
  * Verify in the regex if the email match with a gmail type
  */
  public static boolean match( String email ){
    Matcher m = p.matcher(email);
    return m.find();
  }

  public static void main(String[] args) {

    regexString = "[a-z]+@gmail.com";
    p = Pattern.compile(regexString);

    List<String> list = readInput();
    // sort the list
    list.sort(null);
    // show the elements from list
    list.forEach(System.out::println);

  }
}
