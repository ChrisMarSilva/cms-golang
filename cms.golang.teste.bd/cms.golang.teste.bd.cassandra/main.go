package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gocql/gocql"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.bd.cassandra
// go get github.com/gocql/gocql
// go mod tidy

// go run main.go

func main() {

	var start time.Time

	cluster := gocql.NewCluster("localhost:9042")
	cluster.Keyspace = "example"
	cluster.Consistency = gocql.Quorum // gocql.One // gocql.Quorum // gocql.Any
	cluster.ConnectTimeout = time.Minute * 10
	cluster.ProtoVersion = 4
	cluster.CQLVersion = "3.0.0"

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal("#1: ", err)
	}
	defer session.Close()

	// func CreateSession() *gocql.Session {

	// err = session.Query(`CREATE KEYSPACE ks WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1}  AND durable_writes = true`).Exec()
	// if err != nil {
	// 	log.Fatalf("Error creating Cassandra keyspace: %s", err)
	// }
	// log.Println("Keyspace 'ks' created")

	ctx := context.Background()

	start = time.Now()
	for i := 1; i <= 1; i++ { // 10000=10.000= // 100000=100.000=
		// id, erro := gocql.RandomUUID()
		err = session.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`, "me", gocql.TimeUUID(), "hello world "+strconv.Itoa(i)).WithContext(ctx).Exec()
		if err != nil {
			log.Fatal("#2: ", err)
		}
	}
	log.Println("session.Query=", time.Since(start)) // 10.000 = 2m57.9659076s

	start = time.Now()
	ts := time.Now().UnixNano() / 1000
	batch := session.NewBatch(gocql.LoggedBatch).WithContext(ctx).WithTimestamp(ts) // gocql.NewBatch(gocql.LoggedBatch)
	stmt := "INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)"
	counter := 0
	for i := 0; i < 1; i++ { // 10000=10.000= // 100000=100.000=
		counter++
		batch.Query(stmt, "me", gocql.TimeUUID(), "hello world "+strconv.Itoa(i))
		if counter == 1000 { // gocql.BatchSizeMaximum = 65.535 // if i%1000 == 0 {
			err = session.ExecuteBatch(batch)
			if err != nil {
				log.Fatal("#2.1: ", err)
			}
			ts = time.Now().UnixNano() / 1000
			batch = session.NewBatch(gocql.LoggedBatch).WithContext(ctx).WithTimestamp(ts) // gocql.NewBatch(gocql.LoggedBatch)
			counter = 0
		}
	}
	err = session.ExecuteBatch(batch)
	if err != nil {
		log.Fatal("#2.2: ", err)
	}
	log.Println("session.ExecuteBatch=", time.Since(start)) // 10.000 = 1.4366411s

	// os.Exit(1) // go run main.go

	var id gocql.UUID
	var text string

	start = time.Now()
	scanner := session.Query(`SELECT id, text FROM tweet WHERE timeline = ? ALLOW FILTERING`, "me").WithContext(ctx).Iter().Scanner()
	var idx int = 0
	for scanner.Next() {
		err = scanner.Scan(&id, &text)
		if err != nil {
			log.Fatal("#4: ", err)
		}
		idx++
		fmt.Println("Tweet("+strconv.Itoa(idx)+"):", id, text)
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal("#5: ", err)
	}
	log.Println("session.QueryAll=", time.Since(start))

	start = time.Now()
	err = session.Query(`SELECT id, text FROM tweet WHERE timeline = ? LIMIT 1 ALLOW FILTERING`, "me").WithContext(ctx).Consistency(gocql.One).Scan(&id, &text)
	if err != nil {
		log.Fatal("#3: ", err)
	}
	fmt.Println("Tweet:", id, text)
	log.Println("session.QueryOne=", time.Since(start))

	fmt.Print("FIM")
}

/*

-- creaate keyspace
-- CREATE KEYSPACE IF NOT EXISTS example WITH replication = {'class':'SimpleStrategy','replication_factor':1};
-- CREATE KEYSPACE IF NOT EXISTS example WITH replication = {'class':'NetworkTopologyStrategy', 'DC1':'1'};

-- create table
-- CREATE TABLE IF NOT EXISTS example.tweet (id TIMEUUID, timeline TEXT, text TEXT, PRIMARY KEY (id) );
-- CREATE TABLE IF NOT EXISTS example.signature (signer TEXT, comment TEXT, timestamp BIGINT, );
-- CREATE TABLE IF NOT EXISTS example.documents ( id TIMEUUID, company TEXT, name TEXT, tags SET<TEXT>, signatures SET<frozen <signature>>, status TEXT, timestamp TIMESTAMP, PRIMARY KEY(company, id) );
-- CREATE TABLE IF NOT EXISTS example.messages ( id UUID, user_id UUID, Message text, PRIMARY KEY(id) );
-- CREATE TABLE users ( id UUID, firstname text, lastname text, age int, email text, city text, PRIMARY KEY (id) );


CREATE TABLE blogplatform.users (
    user_id int PRIMARY KEY,
    active_status int,
    password text,
    user_name text
) WITH bloom_filter_fp_chance = 0.01
    AND caching = {'keys': 'ALL', 'rows_per_partition': 'ALL'}
    AND comment = ''
    AND compaction = {'class': 'SizeTieredCompactionStrategy'}
    AND compression = {'sstable_compression': 'org.apache.cassandra.io.compress.LZ4Compressor'}
    AND crc_check_chance = 1.0
    AND dclocal_read_repair_chance = 0.1
    AND default_time_to_live = 0
    AND gc_grace_seconds = 864000
    AND max_index_interval = 2048
    AND memtable_flush_period_in_ms = 0
    AND min_index_interval = 128
    AND read_repair_chance = 0.0
    AND speculative_retry = '99.0PERCENTILE';





func main() {


 cluster := gocql.NewCluster("PublicIP", "PublicIP", "PublicIP") //replace PublicIP with the IP addresses used by your cluster.
 cluster.Consistency = gocql.Quorum
 cluster.ProtoVersion = 4
 cluster.ConnectTimeout = time.Second * 10
 cluster.Authenticator = gocql.PasswordAuthenticator{Username: "Username", Password: "Password"} //replace the username and password fields with their real settings.
 session, err := cluster.CreateSession()

 if err != nil {
 log.Println(err)
 return
 }
 defer session.Close()

 // create keyspaces
 err = session.Query("CREATE KEYSPACE IF NOT EXISTS sleep_centre WITH REPLICATION = {'class' : 'NetworkTopologyStrategy', 'AWS_VPC_US_WEST_2' : 3};").Exec()
 if err != nil {
 log.Println(err)
 return
 }

 // create table
 err = session.Query("CREATE TABLE IF NOT EXISTS sleep_centre.sleep_study (name text, study_date date, sleep_time_hours float, PRIMARY KEY (name, study_date));").Exec()
 if err != nil {
 log.Println(err)
 return
 }

 // insert some practice data
 err = session.Query("INSERT INTO sleep_centre.sleep_study (name, study_date, sleep_time_hours) VALUES ('James', '2018-01-07', 8.2);").Exec()
 err = session.Query("INSERT INTO sleep_centre.sleep_study (name, study_date, sleep_time_hours) VALUES ('James', '2018-01-08', 6.4);").Exec()
 err = session.Query("INSERT INTO sleep_centre.sleep_study (name, study_date, sleep_time_hours) VALUES ('James', '2018-01-09', 7.5);").Exec()
 err = session.Query("INSERT INTO sleep_centre.sleep_study (name, study_date, sleep_time_hours) VALUES ('Bob', '2018-01-07', 6.6);").Exec()
 err = session.Query("INSERT INTO sleep_centre.sleep_study (name, study_date, sleep_time_hours) VALUES ('Bob', '2018-01-08', 6.3);").Exec()
 err = session.Query("INSERT INTO sleep_centre.sleep_study (name, study_date, sleep_time_hours) VALUES ('Bob', '2018-01-09', 6.7);").Exec()
 err = session.Query("INSERT INTO sleep_centre.sleep_study (name, study_date, sleep_time_hours) VALUES ('Emily', '2018-01-07', 7.2);").Exec()
 err = session.Query("INSERT INTO sleep_centre.sleep_study (name, study_date, sleep_time_hours) VALUES ('Emily', '2018-01-09', 7.5);").Exec()
 if err != nil {
 log.Println(err)
 return
 }

 // Return average sleep time for James
 var sleep_time_hours float32

 sleep_time_output := session.Query("SELECT avg(sleep_time_hours) FROM sleep_centre.sleep_study WHERE name = 'James';").Iter()
 sleep_time_output.Scan(&sleep_time_hours)
 fmt.Println("Average sleep time for James was: ", sleep_time_hours, "h")

 // return average sleep time for group
 sleep_time_output = session.Query("SELECT avg(sleep_time_hours) FROM sleep_centre.sleep_study;").Iter()
 sleep_time_output.Scan(&sleep_time_hours)
 fmt.Println("Average sleep time for the group was: ", sleep_time_hours, "h")
}



func init() {
  var err error
  cluster := gocql.NewCluster("127.0.0.1")
  cluster.Keyspace = "streamdemoapi"
  Session, err = cluster.CreateSession()
  if err != nil {
    panic(err)
  }
  fmt.Println("cassandra init done")
}

func main() {
  CassandraSession := Cassandra.Session
  defer CassandraSession.Close()
  router := mux.NewRouter().StrictSlash(true)
  ...
}

func Post(w http.ResponseWriter, r *http.Request) {
  var errs []string
  var gocqlUuid gocql.UUID
  // FormToUser() is included in Users/processing.go
  // we will describe this later
  user, errs := FormToUser(r)
  // have we created a user correctly
  var created bool = false
  // if we had no errors from FormToUser, we will
  // attempt to save our data to Cassandra
  if len(errs) == 0 {
    fmt.Println("creating a new user")
    // generate a unique UUID for this user
    gocqlUuid = gocql.TimeUUID()
    // write data to Cassandra
    if err := Cassandra.Session.Query(`
      INSERT INTO users (id, firstname, lastname, email, city, age) VALUES (?, ?, ?, ?, ?, ?)`,
      gocqlUuid, user.FirstName, user.LastName, user.Email, user.City, user.Age).Exec(); err != nil {
      errs = append(errs, err.Error())
    } else {
      created = true
    }
  }
  // depending on whether we created the user, return the
  // resource ID in a JSON payload, or return our errors
  if created {
    fmt.Println("user_id", gocqlUuid)
    json.NewEncoder(w).Encode(NewUserResponse{ID: gocqlUuid})
  } else {
    fmt.Println("errors", errs)
    json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
  }
}

func Get(w http.ResponseWriter, r *http.Request) {
  var userList []User
  m := map[string]interface{}{}
  query := "SELECT id,age,firstname,lastname,city,email FROM users"
  iterable := Cassandra.Session.Query(query).Iter()
  for iterable.MapScan(m) {
    userList = append(userList, User{
      ID: m["id"].(gocql.UUID),
      Age: m["age"].(int),
      FirstName: m["firstname"].(string),
      LastName: m["lastname"].(string),
      Email: m["email"].(string),
      City: m["city"].(string),
    })
    m = map[string]interface{}{}
  }
  json.NewEncoder(w).Encode(UsersResponse{Users: userList})
}

func GetOne(w http.ResponseWriter, r *http.Request) {
  var user User
  var errs []string
  var found bool = false
  vars := mux.Vars(r)
  id := vars["user_uuid"]
  uuid, err := gocql.ParseUUID(id)
  if err != nil {
    errs = append(errs, err.Error())
  } else {
    m := map[string]interface{}{}
    query := "SELECT id,age,firstname,lastname,city,email FROM users WHERE id=? LIMIT 1"
    iterable := Cassandra.Session.Query(query, uuid).Consistency(gocql.One).Iter()
    for iterable.MapScan(m) {
      found = true
      user = User{
        ID: m["id"].(gocql.UUID),
        Age: m["age"].(int),
        FirstName: m["firstname"].(string),
        LastName: m["lastname"].(string),
        Email: m["email"].(string),
        City: m["city"].(string),
      }
    }
    if !found {
      errs = append(errs, "User not found")
    }
  }
  if found {
    json.NewEncoder(w).Encode(GetUserResponse{User: user})
  } else {
    json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
  }
}



2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
func main() {
    cluster := gocql.NewCluster("192.168.1.1", "192.168.1.2", "192.168.1.3")
    cluster.Keyspace = "example"
    cluster.Consistency = gocql.Quorum
    session, _ := cluster.CreateSession()
    defer session.Close()

    if err := session.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
        "me", gocql.TimeUUID(), "hello world").Exec(); err != nil {
        log.Fatal(err)
    }

    var id gocql.UUID
    var text string

    if err := session.Query(`SELECT id, text FROM tweet WHERE timeline = ? LIMIT 1`,
        "me").Consistency(gocql.One).Scan(&id, &text); err != nil {
        log.Fatal(err)
    }
    fmt.Println("Tweet:", id, text)

    iter := session.Query(`SELECT id, text FROM tweet WHERE timeline = ?`, "me").Iter()
    for iter.Scan(&id, &text) {
        fmt.Println("Tweet:", id, text)
    }
    if err := iter.Close(); err != nil {
        log.Fatal(err)
    }
}


type Signature struct {
	Signer    string `cql:"signer"`
	Comment   string `cql:"comment"`
	Timestamp int64  `cql:"timestamp"`
}

type Document struct {
	Company    string      `cql:"company"`
	Id         gocql.UUID  `cql:"id"`
	Name       string      `cql:"name"`
	Tags       []string    `cql:"tags"`
	Signatures []Signature `cql:"signatures"`
	Status     string      `cql:"status"`
	Timestamp  time.Time   `cql:"timestamp"`
}




var Session *gocql.Session

func initSession() {
	port := func(p string) int {
		i, err := strconv.Atoi(p)
		if err != nil {
			return 9042
		}

		return i
	}

	consistancy := func(c string) gocql.Consistency {
		gc, err := gocql.MustParseConsistency(c)
		if err != nil {
			return gocql.All
		}

		return gc
	}

	cluster := gocql.NewCluster(cassandraConfig.host)
	cluster.Port = port(cassandraConfig.port)
	cluster.Keyspace = cassandraConfig.keyspace
	cluster.Consistency = consistancy(cassandraConfig.consistancy)

	s, err := cluster.CreateSession()
	if err != nil {
		log.Printf("ERROR: fail create cassandra session, %s", err.Error())
		os.Exit(1)
	}
	Session = s
}

func clearSession() {
	Session.Close()
}



package main

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gocql/gocql"
	"github.com/mitchellh/mapstructure"
)

type Signature struct {
	Signer    string `cql:"signer"`
	Comment   string `cql:"comment"`
	Timestamp int64  `cql:"timestamp"`
}

type Document struct {
	Company    string      `cql:"company"`
	Id         gocql.UUID  `cql:"id"`
	Name       string      `cql:"name"`
	Tags       []string    `cql:"tags"`
	Signatures []Signature `cql:"signatures"`
	Status     string      `cql:"status"`
	Timestamp  time.Time   `cql:"timestamp"`
}

var Session *gocql.Session

func initSession() {
	port := func(p string) int {
		i, err := strconv.Atoi(p)
		if err != nil {
			return 9042
		}

		return i
	}

	consistancy := func(c string) gocql.Consistency {
		gc, err := gocql.MustParseConsistency(c)
		if err != nil {
			return gocql.All
		}

		return gc
	}

	cluster := gocql.NewCluster(cassandraConfig.host)
	cluster.Port = port(cassandraConfig.port)
	cluster.Keyspace = cassandraConfig.keyspace
	cluster.Consistency = consistancy(cassandraConfig.consistancy)

	s, err := cluster.CreateSession()
	if err != nil {
		log.Printf("ERROR: fail create cassandra session, %s", err.Error())
		os.Exit(1)
	}
	Session = s
}

func clearSession() {
	Session.Close()
}

func createDocument(document *Document) error {
	q := `
		INSERT INTO documents (
		    company,
		    id,
		    name,
		    tags,
		    signatures,
				status,
				timestamp
		)
		VALUES (?, ?, ?, ?, ?, ?, ?)
    	`
	err := Session.Query(q,
		document.Company,
		document.Id,
		document.Name,
		document.Tags,
		document.Signatures,
		document.Status,
		document.Timestamp).Exec()
	if err != nil {
		log.Printf("ERROR: fail create document, %s", err.Error())
	}

	return err
}

func getDocument(company string, id gocql.UUID) (*Document, error) {
	toSignatures := func(i interface{}) []Signature {
		sigs := []Signature{}
		sig := Signature{}
		for _, s := range i.([]map[string]interface{}) {
			mapstructure.Decode(s, &sig)
			sigs = append(sigs, sig)
		}

		return sigs
	}

	m := map[string]interface{}{}
	q := `
		SELECT * FROM documents
			WHERE company = ? AND id = ?
		LIMIT 1
    	`
	itr := Session.Query(q, company, id).Consistency(gocql.One).Iter()
	for itr.MapScan(m) {
		document := &Document{}
		document.Company = m["company"].(string)
		document.Id = m["id"].(gocql.UUID)
		document.Name = m["name"].(string)
		document.Tags = m["tags"].([]string)
		document.Signatures = toSignatures(m["signatures"])
		document.Status = m["status"].(string)
		document.Timestamp = m["timestamp"].(time.Time)

		log.Printf("INFO: found document, %v", document)

		return document, nil
	}

	return nil, errors.New("document not found")
}

func updateDocument(company string, id gocql.UUID, name string, status string) error {
	q := `
        	UPDATE documents
		SET name = ?, status = ?
		WHERE company = ? AND id = ?
    	`
	err := Session.Query(q, name, status, company, id).Exec()
	if err != nil {
		log.Printf("ERROR: fail update document, %s", err.Error())
		return err
	}

	return nil
}

func addTag(company string, id gocql.UUID, tag string) error {
	q := `
		UPDATE example.documents
		SET tags = tags + ?
		WHERE company = ? AND id = ?;
	`

	err := Session.Query(q, []string{tag}, company, id).Exec()
	if err != nil {
		log.Printf("ERROR: fail add tag, %s", err.Error())
		return err
	}

	return nil
}

func removeTag(company string, id gocql.UUID, tag string) error {
	q := `
		UPDATE example.documents
		SET tags = tags - ?
		WHERE company = ? AND id = ?;
	`

	err := Session.Query(q, []string{tag}, company, id).Exec()
	if err != nil {
		log.Printf("ERROR: fail remove tag, %s", err.Error())
		return err
	}

	return nil
}

func addSignature(company string, id gocql.UUID, signature Signature) error {
	q := `
		UPDATE example.documents
		SET signatures = signatures + ?
		WHERE company = ? AND id = ?;
	`

	err := Session.Query(q, []Signature{signature}, company, id).Exec()
	if err != nil {
		log.Printf("ERROR: fail add signature, %s", err.Error())
		return err
	}

	return nil
}

func removeSignature(company string, id gocql.UUID, signature Signature) error {
	q := `
		UPDATE example.documents
		SET signatures = signatures - ?
		WHERE company = ? AND id = ?;
	`

	err := Session.Query(q, []Signature{signature}, company, id).Exec()
	if err != nil {
		log.Printf("ERROR: fail remove signature, %s", err.Error())
		return err
	}

	return nil
}


package main

import (
	"github.com/gocql/gocql"
	"time"
)

func main() {
	// init cassandra session
	initSession()

	// signatures
	sig1 := Signature{
		Signer:    "eranga",
		Comment:   "valid doc",
		Timestamp: (time.Now().UnixNano() / int64(time.Millisecond)),
	}
	sig2 := Signature{
		Signer:    "herath",
		Comment:   "invalid doc",
		Timestamp: (time.Now().UnixNano() / int64(time.Millisecond)),
	}
	sig3 := Signature{
		Signer:    "lamabda",
		Comment:   "approved doc",
		Timestamp: (time.Now().UnixNano() / int64(time.Millisecond)),
	}

	// documents
	doc1 := &Document{
		Company:    "rahasak",
		Id:         gocql.TimeUUID(),
		Name:       "invoice-01",
		Tags:       []string{"inv", "ops"},
		Signatures: []Signature{sig1, sig2},
		Status:     "pending",
		Timestamp:  time.Now(),
	}
	doc2 := &Document{
		Company:    "rahasak",
		Id:         gocql.TimeUUID(),
		Name:       "order-02",
		Tags:       []string{"ord", "ops"},
		Signatures: []Signature{sig1, sig2},
		Status:     "pending",
		Timestamp:  time.Now(),
	}

	// create doc
	createDocument(doc1)
	createDocument(doc2)

	// query document
	getDocument("rahasak", doc1.Id)

	// update document
	updateDocument(doc1.Company, doc1.Id, "order-01", "approved")

	// add signature
	addSignature("rahasak", doc1.Id, sig3)

	// remove singature
	removeSignature("rahasak", doc1.Id, sig2)

	// add tag
	addTag("rahasak", doc1.Id, "invoice")

	// remove tag
	removeTag("rahasak", doc1.Id, "ops")

	// clear session
	clearSession()
}

*/
