package staff

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/alactic/mygoproject/models/staff"
	"github.com/alactic/mygoproject/utils/connection"
	"github.com/gorilla/mux"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/couchbase/gocb.v1"
)

type Staff = staff.Staff

var bucket *gocb.Bucket = connection.Connection()

//router.HandleFunc("/staff", CreateStaffEndpoint).Methods("POST")
func CreateStaffEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Content-Type", "application/json")
	var staff Staff
	_ = json.NewDecoder(request.Body).Decode(&staff)
	id := uuid.Must(uuid.NewV4()).String()
	staff.Type = "staff"
	staff.Id = id
	_, err := bucket.Insert(id, staff, 0)
	if err != nil {
		fmt.Println("error :: ", err)
		response.WriteHeader(500)
		response.Write([]byte(`{"message": "` + err.Error() + `" }`))
		return
	}
	details := make(map[string]Staff)
	details["data"] = staff
	json.NewEncoder(response).Encode(details)
	// response.Write([]byte(`{ "id": "` + id + `"}`))
}

//router.HandleFunc("/staff/{id}", GetStaffEndpoint).Methods("GET")
func GetStaffEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	routerParams := mux.Vars(request)
	var staff Staff
	staff.Id = routerParams["id"]
	_, err := bucket.Get(routerParams["id"], &staff)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(staff)
}

//router.HandleFunc("/staff", GetStaffEndpoint).Methods("GET")
// //router.HandleFunc("/staff", GetStaffEndpoint).Methods("GET")
func GetAllStaffEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var staff []Staff
	query := gocb.NewN1qlQuery("SELECT META().id, " + bucket.Name() + ".* FROM " + bucket.Name() + " WHERE type = 'staff'")
	rows, err := bucket.ExecuteN1qlQuery(query, nil)
	if err != nil {
		fmt.Println("staffing error :: ", err)
		response.WriteHeader(500)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	var row Staff
	for rows.Next(&row) {
		staff = append(staff, row)
	}
	details := make(map[string][]Staff)
	details["data"] = staff
	fmt.Println("staff :: ", details)
	json.NewEncoder(response).Encode(details)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	fmt.Println("my file", r.FormValue("myFile"))
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("uploads", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// func ReadFile(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("File read Endpoint Hit")
// 	// fmt.Print(path.Ext(path string))
// 	if fileExists("uploads/orange.jpg") {
// 		fmt.Println("Example file exists")
// 		// var Buf bytes.Buffer
// 		// in your case file would be fileupload
// 		_, _, err := r.FormFile("uploads/orange.jpg")
// 		if err != nil {
// 			fmt.Println("working with file error :: ", err)
// 			// panic(err)
// 		} else {
// 			fmt.Println("reading file was successful ")
// 		}
// 		// defer file.Close()
// 		// name := strings.Split(header.Filename, ".")
// 		// fmt.Printf("File name %s\n", name[0])
// 		// // Copy the file data to my buffer
// 		// io.Copy(&Buf, file)
// 		// // do something with the contents...
// 		// // I normally have a struct defined and unmarshal into a struct, but this will
// 		// // work as an example
// 		// contents := Buf.String()
// 		// fmt.Println(contents)
// 		// Buf.Reset()
// 	} else {
// 		fmt.Println("Example file does not exist (or is a directory)")
// 	}
// 	// file, err := os.Open("uploads/upload.png") // For read access.
// 	// defer file.Close()
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	// reader, _ := os.Open("uploads/upload-189240647.png")
// 	// var Buf bytes.Buffer
// 	// // in your case file would be fileupload
// 	// file, header, err := r.FormFile("/uploads/orange.jpg")
// 	// if err != nil {
// 	// 	fmt.Println("working with file error :: ", err)
// 	// 	panic(err)
// 	// }
// 	// defer file.Close()
// 	// name := strings.Split(header.Filename, ".")
// 	// fmt.Printf("File name %s\n", name[0])
// 	// // Copy the file data to my buffer
// 	// io.Copy(&Buf, file)
// 	// // do something with the contents...
// 	// // I normally have a struct defined and unmarshal into a struct, but this will
// 	// // work as an example
// 	// contents := Buf.String()
// 	// fmt.Println(contents)
// 	// I reset the buffer in case I want to use it again
// 	// reduces memory allocations in more intense projects
// 	// Buf.Reset()
// 	// do something else
// 	// etc write header
// 	return
// }
// 	// var Buf bytes.Buffer
// 	// // in your case file would be fileupload
// 	// file, header, err := r.FormFile("/uploads/orange.jpg")
// 	// if err != nil {
// 	// 	fmt.Println("working with file error :: ", err)
// 	// 	panic(err)
// 	// }
// 	// defer file.Close()
// 	// name := strings.Split(header.Filename, ".")
// 	// fmt.Printf("File name %s\n", name[0])
// 	// // Copy the file data to my buffer
// 	// io.Copy(&Buf, file)
// 	// // do something with the contents...
// 	// // I normally have a struct defined and unmarshal into a struct, but this will
// 	// // work as an example
// 	// contents := Buf.String()
// 	// fmt.Println(contents)
// 	// I reset the buffer in case I want to use it again
// 	// reduces memory allocations in more intense projects
// 	// Buf.Reset()
// 	// do something else
// 	// etc write header
// 	return
// }
