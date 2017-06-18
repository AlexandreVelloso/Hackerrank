import java.io.*;
import java.net.*;
import java.util.*;

public class Download{

  public static void main(String[] args) {
    URL url;
    InputStream is = null;
    BufferedReader br;
    String line;
    String result = "";

    try {
      url = new URL( args[0] );
      is = url.openStream();  // throws an IOException
      br = new BufferedReader(new InputStreamReader(is));

      while( (line = br.readLine() ) != null ){

        line = line.replace(">",">\n");
        System.out.println( line );
      }

    } catch (MalformedURLException mue) {
      mue.printStackTrace();
    } catch (IOException ioe) {
      ioe.printStackTrace();
    } finally {
      try {
        if (is != null) is.close();
      } catch (IOException ioe) {
        // nothing to see here
      }
    }

  }
}
