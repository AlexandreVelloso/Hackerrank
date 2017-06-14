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

  public static int getHeight(Node root){
    int height;

    if( root != null ){

      int hr = 0;
      int hl = 0;
      if( root.left != null ){
        hl = 1 + getHeight( root.left );
      }

      if( root.right != null ){
        hr = 1 + getHeight( root.right );
      }

      height = Math.max( hr, hl );

    }else{

      height = 0;
    }

    return height;
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
    int height=getHeight(root);
    System.out.println(height);
  }
}
