package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id string
	title string
	location string
	summary string
}

// Scrape Indeed by a term
func Scrape(term string) {
	var baseURL string = "https://kr.indeed.com/jobs?q="+term+"&limit=50"

	jobs := []extractedJob{}
	// 여러 페이지에서 배열을 받아오기 떄문에
	c := make(chan []extractedJob)
	totalPages := getPages(baseURL)

	for i :=0; i<totalPages; i++ {
		go getPage(i, baseURL, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <- c
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Done, extracted: ", len(jobs))
}

func getPage(page int, url string, mainC chan<- []extractedJob) {
	// 한 페이지에 해당하는 jobs array
	jobs := []extractedJob{}
	c := make(chan extractedJob)
	pageURL := url + "&start=" + fmt.Sprint(page*50)
	fmt.Println(pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i:=0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	
	mainC <- jobs
} 

func extractJob (card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := CleanString(card.Find(".title>a").Text())
	location := CleanString(card.Find(".sjcl").Text())
	summary := CleanString(card.Find(".summary").Text())
	c <- extractedJob{
		id:id,
		title:title,
		location:location,
		summary:summary}
	
}

func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	// 메모리 누수 발생
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {

		pages = s.Find("a").Length()
	})

	return pages
}

// save jobs to csv file
func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Link","Title","Location","Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk="+job.id, job.title, job.location, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}

// CleanString
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}