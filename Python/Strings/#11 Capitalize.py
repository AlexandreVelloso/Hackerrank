def capitalize(string):
    return(' '.join(( newString.capitalize() for newString in string.split(' ') )))
# end capitalize

if __name__ == '__main__':
    string = input()
    capitalized_string = capitalize(string)
    print(capitalized_string)
