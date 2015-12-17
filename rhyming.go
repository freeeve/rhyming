package main

import (
	"bufio"
	"compress/gzip"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/wfreeman/rhyming/Godeps/_workspace/src/github.com/daaku/go.httpgzip"
	"github.com/wfreeman/rhyming/Godeps/_workspace/src/github.com/dotcypress/phonetics"
)

type Dict struct {
	Strings []string
	Words   map[int32]Word
}

type WordJson struct {
	Word             string
	Soundex          string
	Syllables        int8
	PronunciationStr string
	RhymesWith2      []string
	RhymesWith3      []string
	RhymesWith4      []string
	RhymesWith5      []string
}

type Word struct {
	Soundex       string
	Syllables     int8
	Pronunciation Symbols
	RhymesWith2   []byte
	RhymesWith3   []byte
	RhymesWith4   []byte
	RhymesWith5   []byte
}

type Symbols []Symbol

func (d *Dict) Get(word string) (int32, Word, bool) {
	idx := sort.SearchStrings(dict.Strings, word)
	if idx >= 0 && word == dict.Strings[idx] {
		return int32(idx), d.Words[int32(idx)], true
	} else {
		return -1, Word{}, false
	}
}

func (d *Dict) Add(word string, w Word) {
	d.Strings = append(dict.Strings, word)
	d.Set(int32(len(dict.Strings)-1), w)
}

func (d *Dict) Set(idx int32, w Word) {
	d.Words[idx] = w
}

func createRhymes() {
	f, err := os.Open("mhyph2.txt")
	br := bufio.NewReader(f)
	defer f.Close()
	//temp := map[string]struct{}{}
	for err == nil {
		line := ""
		line, err = br.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		arr := strings.Split(line, "Ĩ")
		word := strings.Join(arr, "")

		word = strings.ToLower(word)
		//temp[word] = struct{}{}
	}

	f2, err := os.Open("cmudict-0.7b")
	br = bufio.NewReader(f2)
	defer f2.Close()
	for err == nil {
		line := ""
		line, err = br.ReadString('\n')
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, ";") {
			continue
		}
		if line == "" {
			break
		}
		arr := strings.Split(line, "  ")
		word := arr[0]
		word = strings.ToLower(word)
		//if _, ok := temp[word]; ok {
		w := Word{}
		w.Pronunciation = []Symbol{}
		pronounce := strings.Split(arr[1], " ")
		for _, x := range pronounce {
			sym, err := ParseSymbol(x)
			if err != nil {
				panic(err)
			}
			w.Pronunciation = append(w.Pronunciation, sym)
		}
		w.Soundex = phonetics.EncodeSoundex(word)
		w.Syllables = 0
		for _, v := range w.Pronunciation {
			if v.Type() == Vowel {
				w.Syllables++
			}
		}
		dict.Add(word, w)
		//}
	}

	count := 0
	for k, v := range dict.Words {
		for pk, pv := range dict.Words {
			touched := false
			if k != pk && len(v.Pronunciation) > 5 && len(pv.Pronunciation) > 5 && Equals(v.Pronunciation[len(v.Pronunciation)-5:], pv.Pronunciation[len(pv.Pronunciation)-5:]) {
				v.RhymesWith5 = AppendBytes(v.RhymesWith5, pk)
				touched = true
			} else if k != pk && len(v.Pronunciation) > 4 && len(pv.Pronunciation) > 4 && Equals(v.Pronunciation[len(v.Pronunciation)-4:], pv.Pronunciation[len(pv.Pronunciation)-4:]) {
				v.RhymesWith4 = AppendBytes(v.RhymesWith4, pk)
				touched = true
			} else if k != pk && len(v.Pronunciation) > 3 && len(pv.Pronunciation) > 3 && Equals(v.Pronunciation[len(v.Pronunciation)-3:], pv.Pronunciation[len(pv.Pronunciation)-3:]) {
				v.RhymesWith3 = AppendBytes(v.RhymesWith3, pk)
				touched = true
			} else if k != pk && len(v.Pronunciation) > 2 && len(pv.Pronunciation) > 2 && Equals(v.Pronunciation[len(v.Pronunciation)-2:], pv.Pronunciation[len(pv.Pronunciation)-2:]) {
				v.RhymesWith2 = AppendBytes(v.RhymesWith2, pk)
				touched = true
			}
			if touched {
				dict.Words[k] = v
			}
		}
		count++
		if count%1000 == 0 {
			fmt.Println(fmt.Sprintf("%d out of %d", count, len(dict.Words)))
		}
	}

	//w.Flush()

	of, err := os.Create("rhymes.gob.gz")
	if err != nil {
		panic(err)
	}
	defer of.Close()
	oz := gzip.NewWriter(of)
	out := gob.NewEncoder(oz)
	out.Encode(dict)
	oz.Flush()
	oz.Close()
}

func Equals(a, b []Symbol) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Json(idx int32, w Word) WordJson {
	wj := WordJson{}
	wj.Word = dict.Strings[idx]
	wj.Syllables = w.Syllables
	wj.Soundex = w.Soundex
	wj.PronunciationStr = ""
	for _, x := range w.Pronunciation {
		wj.PronunciationStr += x.String()
	}
	for _, x := range DecodeBytes(w.RhymesWith5) {
		wj.RhymesWith5 = append(wj.RhymesWith5, dict.Strings[x])
	}
	for _, x := range DecodeBytes(w.RhymesWith4) {
		wj.RhymesWith4 = append(wj.RhymesWith4, dict.Strings[x])
	}
	for _, x := range DecodeBytes(w.RhymesWith3) {
		wj.RhymesWith3 = append(wj.RhymesWith3, dict.Strings[x])
	}
	for _, x := range DecodeBytes(w.RhymesWith2) {
		wj.RhymesWith2 = append(wj.RhymesWith2, dict.Strings[x])
	}
	return wj
}

var dict Dict

func main() {
	dict = Dict{Words: map[int32]Word{}}
	f, err := os.Open("rhymes.gob.gz")
	if err != nil {
		createRhymes()
	} else {
		z, _ := gzip.NewReader(f)
		g := gob.NewDecoder(z)
		err = g.Decode(&dict)
		if err != nil {
			panic(err)
		}
	}
	idx, quartile, _ := dict.Get("quartile")
	fmt.Println(Json(int32(idx), quartile))

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", defaultHandler)
	serveMux.HandleFunc("/search", searchHandler)

	panic(http.ListenAndServe(":"+os.Getenv("PORT"), httpgzip.NewHandler(serveMux)))
}

func defaultHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	body, _ := ioutil.ReadFile("public/index.html")
	w.Write(body)
}

func searchHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := req.URL.Query()["q"][0]
	query = strings.ToLower(query)
	idx, word, ok := dict.Get(query)
	var err error
	if ok {
		err = json.NewEncoder(w).Encode(Json(idx, word))
	} else {
		err = json.NewEncoder(w).Encode(WordJson{})
	}
	if err != nil {
		log.Println("error writing search response:", err)
	}
}
