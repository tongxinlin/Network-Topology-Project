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
  //"time"
)

const DB_HOST = "mongodb://localhost:27017"

type Input struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Src_ip      string `json:"src_ip" bson:"src_ip"`
	Dest_ip     string `json:"dest_ip" bson:"dest_ip"`
    Cost        string `json:"cost" bson:"cost"`
}

type ShortestPaths struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Src_ip      string `json:"src_ip" bson:"src_ip"`
	Dest_ip     string `json:"dest_ip" bson:"dest_ip"`
    Cost        string `json:"cost" bson:"cost"`
	Path        string `json:"path" bson:"path"`
}


// readLines reads a whole file into memory
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
  
  session, err := mgo.Dial(DB_HOST)
	defer session.Close()
	if err == nil {
		log.Println("Database connection verified")
	} else {
		log.Fatalln("Dial failed", err)
    }
  
  session.SetMode(mgo.Monotonic, true)
  
  topology := session.DB("test").C("topology")
  
  topology.RemoveAll(nil)
  topology.DropIndex("src", "dest", "cost")
    //Index
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
    

  //var entry []string
  for index,line := range lines {
    if index > 1{
        entry := strings.Fields(line)
        //cost, _ := strconv.ParseFloat(entry[2], 64)
        err = topology.Insert(&Input{ID: bson.NewObjectId(), Src_ip: entry[0], Dest_ip: entry[1], Cost: entry[2]})
        if err != nil {
            log.Fatalln("Inserting failed", err)
            panic(err)
	   }
       log.Println("Inserted", entry[0], entry[1], entry[2])
       
       
    }
  }
  
  
  var sources []string 
  err = topology.Find(nil).Distinct("src_ip", &sources)
  if err != nil {
     log.Fatal(err) 
  }
  log.Println(sources)

  var destinations []string 
  err = topology.Find(nil).Distinct("dest_ip", &destinations)
  if err != nil {
     log.Fatal(err) 
  }
  
  
  f, err := os.OpenFile("./src/tmp/input/pairs", os.O_CREATE|os.O_WRONLY, 0660)
  if err != nil {
    log.Println(err)
  }
  defer f.Close()
  

  for _, source := range sources {
      for _, destination := range destinations{
          _, err := io.WriteString(f, source + "\n" + destination + "\n")
          if err != nil {
            log.Println(err)
          }
      }
  }
  }
  
  
func WriteResultsToDB(path string) {
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
  
  session, err := mgo.Dial(DB_HOST)
  defer session.Close()
	if err == nil {
		log.Println("Database connection verified")
	} else {
		log.Fatalln("Dial failed", err)
    }
  
  session.SetMode(mgo.Monotonic, true)
  
  shortestPaths := session.DB("test").C("shortestpaths")
  
  shortestPaths.RemoveAll(nil)
  shortestPaths.DropIndex("src", "dest", "cost", "path")
    //Index
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
    

  //var entry []string
  for _,line := range lines {
        entry := strings.Fields(line)
        //cost, _ := strconv.ParseFloat(entry[2], 64)
        err = shortestPaths.Insert(&ShortestPaths{ID: bson.NewObjectId(), Src_ip: entry[0], Dest_ip: entry[1], Cost: entry[2], Path: entry[3]})
        if err != nil {
            log.Fatalln("Inserting failed", err)
            panic(err)
	   }
       log.Println("Inserted", entry[0], entry[1], entry[2], entry[3])
  }

}

func QueryShortestPaths(src string, dest string, kPaths string) (string){
  var filePath = "./src/tmp/output/output.txt"
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
  
  shortestPaths := session.DB("test").C("shortestpaths")
  
  var paths []ShortestPaths
  
  
  shortestPaths.Find(bson.M{"src_ip": src, "dest_ip": dest}).Limit(k).All(&paths)
    
  for _, path := range paths {
    log.Println(path.Cost)
    _, err := io.WriteString(file, "Source: " + path.Src_ip + " Destination: " + path.Dest_ip+ " Cost: " + path.Cost + " Path: " + path.Path + "<br>")
          if err != nil {
            log.Println(err)
            }
  }
  return filePath
  
}