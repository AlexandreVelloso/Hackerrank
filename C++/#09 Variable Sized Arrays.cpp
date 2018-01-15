#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;


int main() {
    /* Enter your code here. Read input from STDIN. Print output to STDOUT */   
    int** matrix;
    int num_rows, num_queries;
    int size_column;
    int x,y;
    
    cin >> num_rows >> num_queries;
    
    matrix = new int*[num_rows];
    
    // create and fill the matrix
    for( int i = 0; i < num_rows; i++ ){
        cin >> size_column;
        
        matrix[i] = new int[ size_column ];
        
        // fill a matrix row
        for( int j = 0; j < size_column; j++ ){
            cin >> matrix[i][j];
        }
    }
    
    // for each query
    for( int i = 0; i < num_queries; i++ ){
        cin >> x >> y;
        cout << matrix[x][y] << endl;
    }
    
    return 0;
}