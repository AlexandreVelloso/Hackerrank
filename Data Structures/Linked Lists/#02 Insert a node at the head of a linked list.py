"""
 Insert Node at the begining of a linked list
 head input could be None as well for empty list
 Node is defined as

 class Node(object):

   def __init__(self, data=None, next_node=None):
       self.data = data
       self.next = next_node

 return back the head of the linked list in the below method.
"""

def Insert(head, data):

    if( head == None ):
      return Node( data )
    else:

      # make a aux to the head of the linked list
      aux = head
      # create the new head
      head = Node( data )
      # save prevous head after the new head
      head.next = aux

      return head
    # end else
