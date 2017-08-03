import sys

"""
Node is defined as
self.left (the left child of the node)
self.right (the right child of the node)
self.data (the value of the node)
"""
def preOrder(root):
    #Write your code here
    if( not root is None ):
        sys.stdout.write( str(root.data)+" " )
        #print(root.data, end=' ')
        preOrder( root.left )
        preOrder( root.right )
