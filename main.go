package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/chekalskiy/compilebox/compilebox"
	"os"
)

type CodeSubmission struct {
	Language string   `json:"language"`
	Code     string   `json:"code"`
	Stdins   []string `json:"stdins"`
}

func (s CodeSubmission) String() string {
	return fmt.Sprintf("( <CodeSubmission> {Language: %s, Code: Hidden, Stdins: %s} )", s.Language, s.Stdins)
}

type ExecutionResult struct {
	Stdouts []string           `json:"stdouts"`
	Message compilebox.Message `json:"message"`
}

// type LanguagesResponse struct {
// 	Languages map[string]compilebox.Language `json:"languages"`
// }

var box compilebox.Interface

func main() {
	port := getEnv("COMPILEBOX_PORT", "31337")

	languages := make(map[string]compilebox.Language, 0)
	languages["c++"] = compilebox.Language{
		Compiler: "g++ -o /usercode/a.out -std=c++14 -Wall -fexceptions",
		SourceFile: "file.cpp",
		OptionalExecutable: "/usercode/a.out",
		CommentPrefix: "//",
	}

	languages["c#"] = compilebox.Language{
		Compiler: "dmcs",
		SourceFile: "file.cs",
		OptionalExecutable: "mono /usercode/file.exe",
		CommentPrefix: "//",
	}

	languages["python"] = compilebox.Language{
		Compiler: "python3",
		SourceFile: "file.py",
		OptionalExecutable: "",
		CommentPrefix: "#",
	}

	box = compilebox.New(languages)

	http.HandleFunc("/languages/", getLangs)
	http.HandleFunc("/eval/", evalCode)

	log.Println("testbox listening on " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	log.Printf("Environment variable %s not found, setting to %s", key, fallback)
	os.Setenv(key, fallback)
	return fallback
}

func evalCode(w http.ResponseWriter, r *http.Request) {
	log.Println("Received code subimssion...")
	decoder := json.NewDecoder(r.Body)
	var submission CodeSubmission
	err := decoder.Decode(&submission)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	// fmt.Printf("...along with %d stdin inputs\n", len(submission.Stdins))
	fmt.Println(submission)
	stdouts, msg := box.EvalWithStdins(submission.Language, submission.Code, submission.Stdins)
	log.Println(stdouts, msg)

	if len(stdouts) == 0 {
		log.Println("Code produced no output")
		stdouts = append(stdouts, "ZERO OUTPUTS")
	}

	buf, _ := json.MarshalIndent(ExecutionResult{
		Stdouts: stdouts,
		Message: msg,
	}, "", "   ")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}

func getLangs(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received languages request...")
	workingLangs := make(map[string]compilebox.Language)

	// make a list of currently supported languages
	for k, v := range box.LanguageMap {
		if v.Disabled != "true" {
			workingLangs[k] = v
		}
	}

	fmt.Printf("currently supporting %d of %d known languages\n", len(workingLangs), len(box.LanguageMap))

	// add boilerplate and comment info
	// log.Println(workingLangs)

	// encode language list
	buf, _ := json.MarshalIndent(workingLangs, "", "   ")

	// write working language list back to client
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}
