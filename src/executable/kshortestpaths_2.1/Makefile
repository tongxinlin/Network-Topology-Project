CXX=c++ --std=c++0x
 CFLAGS=-c -Wall
 SOURCES=DijkstraShortestPathAlg.cpp Graph.cpp MainP.cpp YenTopKShortestPathsAlg.cpp
 OBJECTS=DijkstraShortestPathAlg.o Graph.o MainP.o YenTopKShortestPathsAlg.o
 EXECUTABLE=algorithm

 all: $(EXECUTABLE)

 $(EXECUTABLE): $(OBJECTS)
	$(CXX) $(OBJECTS) -o $(EXECUTABLE)

 $(OBJECTS): $(SOURCES)
	$(CXX) $(CFLAGS) -c $(SOURCES)

clean:
	@rm -f $(PROGRAMS) *.o core
