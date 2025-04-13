package main

import (
	"fmt"
	"strconv"
	//"time"
	"os"

	"encoding/json"

	"github.com/go-rod/rod"
)

type TestData struct {
	Units []Unit `json:"units"`
}

type Unit struct {
	UnitName string `json:"unitName"`
	Pathways []Pathway `json:"pathways"` 
}

type Pathway struct {
	PathwayName string `json:"pathwayName"` 
	Questions []Question `json:"questions"` 
}

type Question struct {
	QuestionTitle string `json:"questionTitle"` 
	QuestionSubTitle string `json:"questionSubTitle"` 
	QuestionType string `json:"questionType"` 
	Options []Option `json:"options"` 
}

type Option struct {
	OptionTitle string `json:"optionTitle"` 
	Correctness bool `json:"correctness"`
}

func main() {
	var testData TestData

	file, _ := os.Create("quizData.json")
	defer file.Close()

	browser := rod.New().MustConnect()
	basepage := browser.MustPage("https://developer.android.com/courses/android-basics-compose/course")

	pathways := basepage.MustWaitStable().MustElements(".compose-pathway-link")

	// обробка усіх pathway
	for _, pathway := range pathways {
		pathwayLink := pathway.MustProperty("href").String()

		var unitNum int
		var pathwayNum int
		fmt.Sscanf(pathwayLink, "https://developer.android.com/courses/pathways/android-basics-compose-unit-%d-pathway-%d", &unitNum, &pathwayNum)
		testData.Units = append(testData.Units, Unit{})
		curUnit := &testData.Units[len(testData.Units)-1]
		curUnit.UnitName = "Unit-" + fmt.Sprint(unitNum)
		curUnit.Pathways = append(curUnit.Pathways, Pathway{})
		curPathway := &curUnit.Pathways[len(curUnit.Pathways)-1]
		curPathway.PathwayName = "Pathway-" + fmt.Sprint(pathwayNum)

		page := browser.MustPage(pathwayLink)
		page.MustElement("div.devsite-playlist--item--actions:nth-child(3) > a:nth-child(1)").MustClick()

		qs := page.MustWaitStable().MustElements("li.devsite-quiz-question")

		for _, q := range qs {
			curPathway.Questions = append(curPathway.Questions, Question{})
			curQuestion := &curPathway.Questions[len(curPathway.Questions)-1]

			curQuestion.QuestionTitle = q.MustElement("h2").MustText()
			rawSubTitle := q.MustElements("p")
			for _, st := range rawSubTitle {
				curQuestion.QuestionSubTitle = st.MustProperty("innerHTML").String()

			}
			curQuestion.QuestionType = *q.MustAttribute("data-type")

			if QuestionTitle == "Fill-in-the-blanks" {
				curQuestion.Options = append(curQuestion.Options, Option{})
				if QuestionSubTitle == "\u003cspan\u003e\n      To handle conflicts when inserting into a database, you can pass a(n) \u003c/span\u003e\u003cinput type=\"text\" name=\"question-6\" class=\"devsite-quiz-text-input\" placeholder=\"Fill in the blank...\"\u003e\u003cspan\u003e parameter, such as IGNORE, to the @Insert annotation.\n    \u003c/span\u003e" {
					curQuestion.Options[0].OptionTitle = "onConflict"
					curQuestion.Options[0].Correctness = true
				}
				if QuestionSubTitle == "\n      The ___ thread is responsible for displaying the user interface responding to user input.\n    " {
					curQuestion.Options[0].OptionTitle = "Main"
					curQuestion.Options[0].Correctness = true
				}
			}
		}

		// вибір певного варіанту та перевірка його на правильність
		for i := 0; i < 6; i++ {
			bs := page.MustWaitStable().MustElements("input[value='" + strconv.Itoa(i) + "']");
			for _, b := range bs {
				var nq, nans int

				fmt.Sscanf(b.MustProperty("name").String(), "question-%d", &nq)
				fmt.Sscanf(b.MustProperty("value").String(), "%d",  &nans)
				curPathway.Questions[nq].Options = append(curPathway.Questions[nq].Options, Option{})
				curOption := &curPathway.Questions[nq].Options[len(curPathway.Questions[nq].Options)-1]
				ot := page.MustElement("label[for=\"" + b.MustProperty("id").String() + "\"]")
				curOption.OptionTitle = ot.MustProperty("innerHTML").String()

				b.MustClick()
			}
			page.MustElement("button.button-primary").MustClick()
			cbs := page.MustWaitStable().MustElements("input.variant-success")
			for _, cb := range cbs {
				var nq, nans int
				fmt.Sscanf(cb.MustProperty("name").String(), "question-%d", &nq)
				fmt.Sscanf(cb.MustProperty("value").String(), "%d",  &nans)
				curOption := &curPathway.Questions[nq].Options[nans]
				curOption.Correctness = true
			}
			page.MustElement("button.button").MustClick()
		}

		page.Close()
	}

	data, _ := json.Marshal(testData)
	dataOutput := string(data)
	fmt.Fprintf(file, dataOutput)
}
