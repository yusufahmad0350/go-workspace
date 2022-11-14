package main

import "fmt"

func main(){

}



func shift_left([][]int mat, int s, int row, int n, int amount) {
    datatype*  = malloc(sizeof(datatype) * n);
    for(cotemp_bufferl:= 0; col < n; col++) {
        datatype temp = mat[row][(col+amount)%s];
        temp_buffer[(col+amount)%s] = mat[row][col];
        temp_buffer[col] = temp;
    }
    memcpy(mat[row], temp_buffer, n);
    free(temp_buffer);
}

void shift_up(datatype** mat, int s, int col, int n, int amount) {
    datatype* temp_buffer = malloc(sizeof(datatype) * n);
    for(int row = 0; row < n; row++) {
        datatype temp = mat[(row+amount)%s][col];
        temp_buffer[(row+amount)%s] = mat[row][col];
        temp_buffer[row] = temp;
    }
    memcpy(&mat[0][col], temp_buffer, n);
    free(temp_buffer);
}

void cannon_mul(int p_sqrt,datatype** a, datatype** b, datatype** c, int n) {
    /* 2D matrices and n^2 sized only!*/
    int i = 0, j = 0, k = 0;
    int s = p_sqrt;
    for(i = 0; i < (s-1); i++) {
        shift_left(a, s, i, s-1, i); // Skew matrix a
    }
    for (i = 0; i < (s-1); i++) {
        shift_up(b, s, i, s-1, i); // Skew matrix b
    }
    for(k = 0; k < (s-1); k++) {
        for(i = 0; i < (s-1); i++) {
            for(j = 0; j < (s-1); j++) {
                c[i][j] += a[i][j]*b[i][j];
                shift_left(a, s, i, s-1, 1);
                shift_up(b, s, i, s-1, 1);  
            }                       
        }
    }  
}