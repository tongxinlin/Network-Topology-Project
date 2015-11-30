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

void getKShortestPaths(string filename, ofstream& outFile, ifstream& vertexPairs)
{
	//Graph my_graph("../data/test_6_2");
	Graph my_graph(filename);
    int startVertex;
    int destinationVertex;
        
    vertexPairs >> startVertex;
    
    while (vertexPairs){
        vertexPairs >> destinationVertex;
                
        if (startVertex != destinationVertex){
            YenTopKShortestPathsAlg yenAlg(my_graph, my_graph.get_vertex(startVertex),
                my_graph.get_vertex(destinationVertex));

            while(yenAlg.has_next())
            {
                outFile << startVertex << " " << destinationVertex << " ";
                yenAlg.next()->PrintOut(outFile);
            }
        }
        vertexPairs >> startVertex;
    }
}
// 	System.out.println("Result # :"+i);
// 	System.out.println("Candidate # :"+yenAlg.get_cadidate_size());
// 	System.out.println("All generated : "+yenAlg.get_generated_path_size());


int main(int argc, const char *argv[])
{
    string fileName = argv[1];
    
    ofstream outFile;
    outFile.open("./src/tmp/output/results.txt");
    
    ifstream vertexPairs;
    vertexPairs.open("./src/tmp/input/pairs");
    getKShortestPaths(fileName, outFile, vertexPairs);
    outFile.close();
    vertexPairs.close();
    cout << "./src/tmp/output/results.txt";
	//testDijkstraGraph();
    
}
