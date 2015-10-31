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

void yenAlg()
{
	//Graph my_graph("../data/test_6_2");
	Graph my_graph("data/danYen");
    
    // TODO: Create some kind of struct for all wanted ip pairs
    
    // TODO: Get the wanted ips
	YenTopKShortestPathsAlg yenAlg(my_graph, my_graph.get_vertex(46),
		my_graph.get_vertex(13));

	int i=0;
    // TODO: If max k value is given, add to break the while loop
	while(yenAlg.has_next())
	{
        // TODO: Add mongoDB query here to add the entry
		++i;
		yenAlg.next()->PrintOut(cout);
	}

// 	System.out.println("Result # :"+i);
// 	System.out.println("Candidate # :"+yenAlg.get_cadidate_size());
// 	System.out.println("All generated : "+yenAlg.get_generated_path_size());

}

int main(...)
{
    // TODO: get the input (certain ips / all of them, k defined?)
	cout << "Welcome to the real world!" << endl;

	//testDijkstraGraph();
    
    // TODO: Add args to function call
	yenAlg();
}
