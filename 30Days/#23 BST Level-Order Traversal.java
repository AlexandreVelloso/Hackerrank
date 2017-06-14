import java.util.*;
import java.io.*;
class Node{
  Node left,right;
  int data;
  Node(int data){
    this.data=data;
    left=right=null;
  }
}
class Solution{
  static void levelOrder(Node root){

    System.out.print( root.data+" " );

    Queue<Node> fila = new LinkedList<>();
    if( root.left != null )
      fila.add( root.left );

    if( root.right != null )
      fila.add( root.right );

    //Write your code here
    while( ! fila.isEmpty() ){
      Node aux = fila.remove();
      System.out.print( aux.data+" " );

      if( aux.left != null )
        fila.add( aux.left );


      if( aux.right != null )
        fila.add( aux.right );

    }

    System.out.println("");
  }

  public static Node insert(Node root,int data){
    if(root==null){
      return new Node(data);
    }
    else{
      Node cur;
      if(data<=root.data){
        cur=insert(root.left,data);
        root.left=cur;
      }
      else{
        cur=insert(root.right,data);
        root.right=cur;
      }
      return root;
    }
  }
  public static void main(String args[]){
    Scanner sc=new Scanner(System.in);
    int T=sc.nextInt();
    Node root=null;
    while(T-->0){
      int data=sc.nextInt();
      root=insert(root,data);
    }
    levelOrder(root);
  }
}
