import java.util.Scanner;

public class Tratar{

  public static void main(String [] args){

    Scanner sc = new Scanner( System.in );

    // as primeiras linhas sao lixo
    for ( int i = 0; i < 14; i++) {
      sc.nextLine();
    }

    String str;
    while( sc.hasNext() ){

      while( (str = sc.nextLine().replace("</a>","") ).equals("Solve Challenge") );

      if( str.length() > 0 ){

        if( str.equals("1") )
          break;

        System.out.println( "touch "+str.replace(" ","\\ ")+"."+args[0] );


      }

    }// end while

  }
}
