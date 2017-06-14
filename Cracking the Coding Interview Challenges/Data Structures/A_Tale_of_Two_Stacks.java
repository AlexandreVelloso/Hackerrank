import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;

class MyQueue <Object>{

  ArrayList<Object> queue = new ArrayList<>();

  public void enqueue( Object obj ){
    queue.add( obj );
  }

  public Object dequeue(){

    Object obj = queue.get(0);
    queue.remove(0);

    return obj;
  }

  public Object peek(){

    return queue.get(0);
  }

}

public class A_Tale_of_Two_Stacks {
    public static void main(String[] args) {
        MyQueue<Integer> queue = new MyQueue<Integer>();

        Scanner scan = new Scanner(System.in);
        int n = scan.nextInt();

        for (int i = 0; i < n; i++) {
            int operation = scan.nextInt();
            if (operation == 1) { // enqueue
              queue.enqueue(scan.nextInt());
            } else if (operation == 2) { // dequeue
              queue.dequeue();
            } else if (operation == 3) { // print/peek
              System.out.println(queue.peek());
            }
        }
        scan.close();
    }
}
