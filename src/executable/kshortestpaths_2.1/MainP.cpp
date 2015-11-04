/************************************************************************/
/* $Id: MainP.cpp 65 2010-09-08 06:48:36Z yan.qi.asu $                                                                 */
/************************************************************************/

#include <limits>
#include <set>
#include <map>
#include <queue>
#include <string>
#include <vector>
#include <fstream>
#include <iostream>
#include <algorithm>
#include "GraphElements.h"
#include "Graph.h"
#include "DijkstraShortestPathAlg.h"
#include "YenTopKShortestPathsAlg.h"

using namespace std;


void testDijkstraGraph()
{
	Graph* my_graph_pt = new Graph("data/danYen");
	DijkstraShortestPathAlg shortest_path_alg(my_graph_pt);
	BasePath* result =
		shortest_path_alg.get_shortest_path(
			my_graph_pt->get_vertex(46), my_graph_pt->get_vertex(13));
	result->PrintOut(cout);
}

void getKShortestPaths(int k, int startVertex, int destinationVertex, string filename, ofstream& outFile)
{
	//Graph my_graph("../data/test_6_2");
	Graph my_graph(filename);
    
    // TODO: Create some kind of struct for all wanted ip pairs
    
    // TODO: Get the wanted ips
	YenTopKShortestPathsAlg yenAlg(my_graph, my_graph.get_vertex(startVertex),
		my_graph.get_vertex(destinationVertex));

	int i=0;
    // TODO: If max k value is given, add to break the while loop
	while(yenAlg.has_next() && i < k)
	{
        // TODO: Add mongoDB query here to add the entry
		++i;
		yenAlg.next()->PrintOut(outFile);
        outFile << '\n';
	}

// 	System.out.println("Result # :"+i);
// 	System.out.println("Candidate # :"+yenAlg.get_cadidate_size());
// 	System.out.println("All generated : "+yenAlg.get_generated_path_size());

}

int main(int argc, const char *argv[])
{
    string fileName = argv[1];
    int k = stoi(argv[2]);
    ifstream inFile;
    inFile.open(fileName);
    int amountOfNodes;
    inFile >> amountOfNodes;
    inFile.close();
    
    ofstream outFile;
    outFile.open("./src/tmp/output/output.txt");
    
    for (int start = 0; start < amountOfNodes; start++){
        for (int destination = 0; destination < amountOfNodes; destination++){
        	getKShortestPaths(k, start, destination, fileName, outFile);
        }
    }
    outFile.close();
    cout << "./src/tmp/output/output.txt";
	//testDijkstraGraph();
    
}
