#include<iostream>
//#include<stdlib.h>
#include<fstream>

using namespace std;
int main(int argc, char** argv)
{
	ifstream ifs;
	ifs.open(argv[1]);
	ofstream ofs;
	ofs.open("./src/tmp/output/output.txt");
	char c;
	while(ifs >> c){
		ofs << c;
	}
	ifs.close();
	ofs.close();
	cout << "./src/tmp/output/output.txt";
    return 0;
}
