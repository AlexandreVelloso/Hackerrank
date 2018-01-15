#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;


int main() {
    /* Enter your code here. Read input from STDIN. Print output to STDOUT */   
    int size;
    cin >> size;
    
    int vet[size];
    for( int i = 0; i < size; i++ ){
        cin >> vet[i];
    }
    
    // print in the reverse order
    for( int i = size-1; i >= 0; i-- ){
        cout << vet[i] << " ";
    }
    
    return 0;
}