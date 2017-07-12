'''
 Print elements of a linked list on console
 head input could be None as well for empty list
 Node is defined as

class Node(object):

    def __init__(self, data=None, next_node=None):
        self.data = data
        self.next = next_node
    # end method __init__

# end class
'''

def print_list(head):

    while( head != None ):
        print( head.data )
        head = head.next

# end method print_list
