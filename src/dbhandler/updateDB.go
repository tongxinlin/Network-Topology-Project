package dbhandler

import (
  "bufio"
  "log"
  "os"
  "io"
  "strings"
  "strconv"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

const DB_HOST = "mongodb://localhost:27017"

// Structure for topology as given in input
type Input struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Src_ip      string `json:"src_ip" bson:"src_ip"`
	Dest_ip     string `json:"dest_ip" bson:"dest_ip"`
    Cost        string `json:"cost" bson:"cost"`
}

// Structure for the shortest paths to be stored in db
type ShortestPaths struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Src_ip      string `json:"src_ip" bson:"src_ip"`
	Dest_ip     string `json:"dest_ip" bson:"dest_ip"`
    Cost        string `json:"cost" bson:"cost"`
	Path        string `json:"path" bson:"path"`
}


// Input: path to uploaded file
//
// Writes the uploaded topology to db
func WriteToDB(path string) {
  file, err := os.Open(path)
  if err != nil {
    log.Fatalln(err)
    }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  
  // Connect to mongoDB
  session, err := mgo.Dial(DB_HOST)
	defer session.Close()
	if err == nil {
		log.Println("Database connection verified")
	} else {
		log.Fatalln("Dial failed", err)
    }
  
  session.SetMode(mgo.Monotonic, true)
  
  // Get the topology table, clear and prepare it for the input
  topology := session.DB("test").C("topology")
  topology.RemoveAll(nil)
  topology.DropIndex("src", "dest", "cost")
  topologyIndex := mgo.Index{
    Key:        []string{"src", "dest", "cost"},
    Background: true,
    Unique:     false,
    DropDups:   false,
  }
  err = topology.EnsureIndex(topologyIndex)
  if err != nil {
	panic(err)
    log.Fatalln("Index ensuring failed")
  }
    

  // Read the lines from file to db
  for index,line := range lines {
    if index > 1{
     entry := strings.Fields(line)
     err = topology.Insert(&Input{ID: bson.NewObjectId(), Src_ip: entry[0], Dest_ip: entry[1], Cost: entry[2]})
     if err != nil {
        log.Fatalln("Inserting failed", err)
        panic(err)
      }         
    }
  }
  
  // Get all possible source-destination pairs from the topology
  var sources []string 
  err = topology.Find(nil).Distinct("src_ip", &sources)
  if err != nil {
     log.Fatal(err) 
  }
  var destinations []string 
  err = topology.Find(nil).Distinct("dest_ip", &destinations)
  if err != nil {
     log.Fatal(err) 
  }
  
  
  f, err := os.OpenFile("./src/tmp/input/pairs.txt", os.O_TRUNC|os.O_WRONLY, 0660)
  if err != nil {
    log.Println(err)
  }
  defer f.Close()
  
  // Write the pairs to a file for Yen's algorithm
  for _, source := range sources {
      for _, destination := range destinations{
          _, err := io.WriteString(f, source + "\n" + destination + "\n")
          if err != nil {
            log.Println(err)
          }
      }
  }
  }
  
// Write the file Yen's algorithm created to a new db  
func WriteResultsToDB() {
  file, err := os.Open("./src/tmp/output/results.txt")
  if err != nil {
    log.Fatalln(err)
  }
  defer file.Close()

  // Read the lines from a file to buffer
  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  
  session, err := mgo.Dial(DB_HOST)
  defer session.Close()
	if err == nil {
		log.Println("Database connection verified")
	} else {
		log.Fatalln("Dial failed", err)
    }
  
  session.SetMode(mgo.Monotonic, true)
  
  // Get the shortestpaths table, clear and prepare it for the input
  shortestPaths := session.DB("test").C("shortestpaths")
  shortestPaths.RemoveAll(nil)
  shortestPaths.DropIndex("src", "dest", "cost", "path")
  shortestPathsIndex := mgo.Index{
    Key:        []string{"src", "dest", "cost", "path"},
    Background: true,
    Unique:     false,
    DropDups:   false,
  }
  err = shortestPaths.EnsureIndex(shortestPathsIndex)
  if err != nil {
    panic(err)
    log.Fatalln("Index ensuring failed")
  }
    

  // Write the shortest paths to db
  for _,line := range lines {
    entry := strings.Fields(line)
    err = shortestPaths.Insert(&ShortestPaths{ID: bson.NewObjectId(), Src_ip: entry[0], Dest_ip: entry[1], Cost: entry[2], Path: entry[3]})
    if err != nil {
        log.Fatalln("Inserting failed", err)
        panic(err)
    }
  }

}

// Input: source ip (string), destination ip (string), number of shortest paths (string)
// Output: file where the result of the query is written
//
// Query for a certain source destination path for k shortest paths.
// Writes the result as html formatted output to a file
func QueryShortestPaths(src string, dest string, kPaths string) (string){
  // Set the output file
  var filePath = "./src/tmp/output/output.txt"
  
  // Convert k to an integer
  k, err := strconv.Atoi(kPaths)
  if err != nil {
        log.Println(err)
    }
  
  file, err := os.OpenFile(filePath, os.O_TRUNC|os.O_WRONLY, 0660)
  if err != nil {
    log.Println(err)
  }
  defer file.Close()
  
  session, err := mgo.Dial(DB_HOST)
  defer session.Close()
  if err == nil {
	log.Println("Database connection verified")
  } else {
    log.Fatalln("Dial failed", err)
  }
  
  session.SetMode(mgo.Monotonic, true)
  
  // Get the shortestpaths table and query it
  shortestPaths := session.DB("test").C("shortestpaths")
  var paths []ShortestPaths
  shortestPaths.Find(bson.M{"src_ip": src, "dest_ip": dest}).Limit(k).All(&paths)

  // Write the results to file
  for _, path := range paths {
    log.Println(path.Cost)
    _, err := io.WriteString(file, "Source: " + path.Src_ip + "<br>Destination: " + path.Dest_ip+ "<br>Cost: " + path.Cost + "<br>Path: " + path.Path + "<br><br>")
    if err != nil {
      log.Println(err)
    }
  }
  return filePath
}

// Input: ip (string)
// Output: file where the result of the query is written
//
// Query the neighbors of a certain ip
// Writes the result as html formatted output to a file
func NeighborsOf(node string) (string){
  // Set the output file
  var filePath = "./src/tmp/output/neighbors.txt"
  
  file, err := os.OpenFile(filePath, os.O_TRUNC|os.O_WRONLY, 0660)
  if err != nil {
    log.Println(err)
  }
  defer file.Close()
  
  session, err := mgo.Dial(DB_HOST)
  defer session.Close()
  if err == nil {
	log.Println("Database connection verified")
  } else {
    log.Fatalln("Dial failed", err)
  }
  
  session.SetMode(mgo.Monotonic, true)
  
  // Get the topology table and query it
  topology := session.DB("test").C("topology")
  var neighbors []Input
  topology.Find(bson.M{"src_ip": node}).All(&neighbors)
  
  // Write the results to file
  _, err = io.WriteString(file, "Border Node: " + node + "<br>Neighboring Nodes: ")
  if err != nil {
    log.Println(err)
  }
  for _, path := range neighbors {
    _, err := io.WriteString(file, path.Dest_ip + " ")
          if err != nil {
            log.Println(err)
            }
  }
  
  return filePath
}