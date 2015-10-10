#include<iostream>
#include<stdlib.h>
using namespace std;
int main(int argc, char** argv)
{
	cout << "input1: " << argv[1] << '\n';
	cout << "input2: " << argv[2] << '\n';
    cout << "sum: " << atoi(argv[1]) + atoi(argv[2]) << '\n';
    return 100;
}
