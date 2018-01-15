#include <iostream>
#include <cstdio>
using namespace std;

int main() {
    // Complete the code.
    
    string vet[10] = {"","one","two","three","four","five","six","seven","eight","nine"};
    
    int a,b;
    scanf("%d",&a);
    scanf("%d",&b);
    
    for( int i = a; i <= b; i++ ){
        if( 1 <= i && i <= 9 ){
            cout << vet[i] << endl;
        }else if( i % 2 == 0 ){
            printf("even\n");
        }else{
            printf("odd\n");
        }
    }
    
    return 0;
}