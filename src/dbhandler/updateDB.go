package dbhandler

import (
  "bufio"
  "log"
  "os"
  "strings"
  //"strconv"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  //"time"
)

const DB_HOST = "mongodb://localhost:27017"

type Input struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Src_ip      string
	Dest_ip     string
    Cost        string
	//Timestamp time.Time
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
  
    // Index
	topologyIndex := mgo.Index{
		Key:        []string{"src", "dest", "cost"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = topology.EnsureIndex(topologyIndex)
	if err != nil {
		panic(err)
        log.Fatalln("Index ensuring failed")
	}
    

  //var entry []string
  for index,line := range lines {
    if index > 1{
        entry := strings.Split(line, " ")
        //cost, _ := strconv.ParseFloat(entry[2], 64)
        err = topology.Insert(&Input{Src_ip: entry[0], Dest_ip: entry[1], Cost: entry[2]})
        if err != nil {
            log.Fatalln("Inserting failed")
            panic(err)
	   }
       
       
    }
  }
}