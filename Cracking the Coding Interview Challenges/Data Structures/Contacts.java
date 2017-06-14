import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;

class Node{
    public String value;
    public Node[] next;

    public Node( String value ){
      this.value = value;
      this.next = new Node[26];
    }
}

class Trie{
  Node root;

  public void addName( String name ){

    for ( int i = 0; i < name.length(); i++) {
      

    }

  }// end addName
}

public class Contacts {

  private static Set<String> contacts = new TreeSet<>();

  public static void operation( String op, String contact ){

    if( op.equals( "add" ) ){
      contacts.add( contact );
    }else{

      Iterator<String> iterator = contacts.iterator();

      // number of contacts finded
      int cont = 0;

      // for each contact
      while( iterator.hasNext() ){

        // get the contact
        String nextContact = iterator.next();

        if( nextContact.startsWith( contact) ){
          cont++;
        }else{
          break;
        }

      }// end while

      System.out.println( cont );

    }// end if

  }// end method operation

  public static void main(String[] args) {
      Scanner in = new Scanner(System.in);
      int n = in.nextInt();
      for(int a0 = 0; a0 < n; a0++){
          String op = in.next();
          String contact = in.next();

          operation( op, contact );
      }
  }
}
