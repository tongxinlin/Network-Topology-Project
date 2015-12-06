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

// Input: input file (string), output file (string), file containing ip pairs (ifstream&)
//
// Calculates the shortest paths between given ip pairs
// Outputs the result to given file
void getKShortestPaths(string filename, ofstream& outFile, ifstream& vertexPairs)
{
    // Create a graph of the topology
	Graph my_graph(filename);
    // Variables for ip pairs
    int startVertex;
    int destinationVertex;
    
    vertexPairs >> startVertex;
    
    // Get shortest paths for all ip pairs
    while (vertexPairs){
        vertexPairs >> destinationVertex;
        
        // Eliminate null paths
        if (startVertex != destinationVertex){
            // Calculate the shortest paths for one ip pair
            YenTopKShortestPathsAlg yenAlg(my_graph, my_graph.get_vertex(startVertex),
                my_graph.get_vertex(destinationVertex));
            // Write all shortest paths to file
            // Modify this by adding a limit if used for large inputs
            while(yenAlg.has_next())
            {
                outFile << startVertex << " " << destinationVertex << " ";
                yenAlg.next()->PrintOut(outFile);
            }
        }
        vertexPairs >> startVertex;
    }
}


int main(int argc, const char *argv[])
{
    // Get the file to process
    string fileName = argv[1];
    ofstream outFile;
    // File to write the shortest paths to
    outFile.open("./src/tmp/output/results.txt");
    ifstream vertexPairs;
    // File to read all ip pairs from
    vertexPairs.open("./src/tmp/input/pairs.txt");
    // Calculate the shortest paths
    getKShortestPaths(fileName, outFile, vertexPairs);
    outFile.close();
    vertexPairs.close();
    return 0;
    
}
