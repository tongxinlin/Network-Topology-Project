CXX=g++
 CFLAGS=-c --std=c++11 -Wall -I/usr/local/mongo-cxx-driver/include/mongocxx/v0.3 -I/usr/local/mongo-cxx-driver/include/libmongoc-1.0 -I/usr/local/mongo-cxx-driver/include/bsoncxx/v0.3 -I/usr/local/mongo-cxx-driver/include/libbson-1.0
 LDFLAGS=-L/usr/local/mongo-cxx-driver/lib -lmongocxx -lbsoncxx
 SOURCES=DijkstraShortestPathAlg.cpp Graph.cpp MainP.cpp YenTopKShortestPathsAlg.cpp
 OBJECTS=DijkstraShortestPathAlg.o Graph.o MainP.o YenTopKShortestPathsAlg.o
 EXECUTABLE=algorithm

 all: $(EXECUTABLE)

 $(EXECUTABLE): $(OBJECTS)
	$(CXX) $(OBJECTS) -o $(EXECUTABLE) $(LDFLAGS)

 $(OBJECTS): $(SOURCES)
	$(CXX) $(CFLAGS) -c $(SOURCES)
