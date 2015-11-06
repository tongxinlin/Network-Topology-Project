#include<iostream>
//#include<stdlib.h>
#include<fstream>

using namespace std;
int main(int argc, char** argv)
{
	ofstream ofs;
	ofs.open("./src/tmp/output/output.txt");
	ofs << "uploaded file name: " <<argv[1] << " ";
	ofs << "dest: " << argv[2] << " ";
	ofs << "src: " << argv[3] << " ";
	ofs << "kpaths: " << argv[4] << " ";
	ofs.close();
	cout << "./src/tmp/output/output.txt";
    return 0;
}
