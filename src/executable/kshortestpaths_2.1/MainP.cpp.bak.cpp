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


#include <bsoncxx/builder/stream/document.hpp>
#include <bsoncxx/types.hpp>
#include <bsoncxx/json.hpp>

#include <mongocxx/client.hpp>
#include <mongocxx/instance.hpp>
#include <mongocxx/options/find.hpp>

using bsoncxx::builder::stream::document;
using bsoncxx::builder::stream::open_document;
using bsoncxx::builder::stream::close_document;
using bsoncxx::builder::stream::open_array;
using bsoncxx::builder::stream::close_array;
using bsoncxx::builder::stream::finalize;

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

void getKShortestPaths(string filename, ofstream& outFile)
{
	//Graph my_graph("../data/test_6_2");
	Graph my_graph(filename);
    int startVertex;
    int destinationVertex;
    bool exists = false;
    
    string start;
    string end;
    
    mongocxx::instance inst{};
    mongocxx::client conn{};

    auto db = conn["test"];
    
    db["shortestpaths"].drop();
    //document index_spec{};
    //index_spec << "cost" << '-1';
    db["shortestpaths"].create_index(document{} << "cost" << "-1" << finalize, {});
    
    auto source = db["topology.src_ip"].find({});
    for (auto&& doc : source){
        start = bsoncxx::to_json(doc);
        startVertex = std::stoi(start);
        
        auto dest = db["topology.dest_ip"].find({});
        for (auto&& field : dest){
            end = bsoncxx::to_json(field);
            destinationVertex = std::stoi(end);
            // TODO: Get the wanted ips
            document filter;
            filter << "source" << start << "dest" << end;
            auto dest = db["shortestpaths"].find(filter);
            for (auto&& doc : dest){
                exists = true;
                continue;
            }
            
            if (startVertex != destinationVertex && !exists){
                YenTopKShortestPathsAlg yenAlg(my_graph, my_graph.get_vertex(startVertex),
                    my_graph.get_vertex(destinationVertex));

                while(yenAlg.has_next())
                {
                    // TODO: Add mongoDB query here to add the entry
                    yenAlg.next()->PrintOut(outFile);
                    outFile << '\n';
                }
            }
            exists = false;
        }
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
    getKShortestPaths(fileName, outFile);
    outFile.close();
    cout << "./src/tmp/output/output.txt";
	//testDijkstraGraph();
    
}
