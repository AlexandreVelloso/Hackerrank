import re

def fun(s):
    # return True if s is a valid email, else return False

    #
    # (\w | - | _ )+       -> one or more letters, digits, - or _
    # \@                   -> one @
    # ([a-z]|[A-Z]|[0-9])+ -> one or more letters or digits
    # \.                   -> one .
    # \w{1,3}              -> one, two or three letters or digits
    # $                    -> end of string
    pattern = re.compile("(\w|-|_)+\@([a-z]|[A-Z]|[0-9])+\.\w{1,3}$")

    return pattern.match( s )

def filter_mail(emails):
    return list(filter(fun, emails))

if __name__ == '__main__':
    n = int(input())
    emails = []
    for _ in range(n):
        emails.append(input())

filtered_emails = filter_mail(emails)
filtered_emails.sort()
print(filtered_emails)
