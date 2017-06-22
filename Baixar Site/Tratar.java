import java.util.Scanner;

public class Tratar{

  /**
  * Contar o numero de casas de um numero
  */
  public static int contCasas( int num ){

    int cont = 1;

    while( num > 9 ){
      num /= 10;
      cont++;
    }

    return cont;
  }

  /**
  * Se eu tenho o numero 20, e quero ele com 3 caracteres
  * eu tenho que arrumar ele para 020, isso que esse metodo faz
  */
  public static String arrumar( int num, int numCasas ){
    String arrumado = "";

    for (int i = contCasas(num); i < numCasas; i++) {
      arrumado += "0";
    }

    arrumado += num;

    return arrumado;
  }

  /**
  * Args[0] e' o nome da extensao do arquivo
  * Args[1] e' o ultimo numero baixado
  */
  public static void main(String [] args){

    Scanner sc = new Scanner( System.in );

    // as primeiras linhas sao lixo
    for ( int i = 0; i < 14; i++) {
      sc.nextLine();
    }

    // numero do ultimo exercicio baixado
    int num = Integer.parseInt( args[1] );
    int numCasas = args[1].length();

    String str;
    while( sc.hasNext() ){

      while( (str = sc.nextLine().replace("</a>","") ).equals("Solve Challenge") );

      if( str.length() > 0 ){

        if( str.equals("1") )
          break;


          String name = str.replace(" ","\\ ").replace("!","\\!").replace("(","\\(").replace(")","\\)");

        System.out.printf( "touch \\#%s\\ %s.%s\n", arrumar(++num, numCasas), name, args[0] );


      }

    }// end while

  }
}
