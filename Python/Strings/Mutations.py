def mutate_string(string, position, character):

    # transform the string to list
    string = list( string )

    # change the character
    string[ position ] = character

    # transform the list to string
    return "".join( string )

if __name__ == '__main__':
    s = input()
    i, c = input().split()
    s_new = mutate_string(s, int(i), c)
    print(s_new)
