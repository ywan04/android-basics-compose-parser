package main

import (
	"fmt"
	"strconv"
	//"time"

	//"encoding/json"

	"github.com/go-rod/rod"
)

type TestData struct {
	Units []Unit `json: "units"`
}

type Unit struct {
	UnitName string `json: "unitName"`
	Pathways []Pathway `json: "pathways"` 
}

type Pathway struct {
	PathwayName string `json: "pathwayName"` 
	Questions []Question `json: "questions"` 
}

type Question struct {
	QuestionTitle string `json: "questionTitle"` 
	QuestionSubTitle string `json: "questionSubTitle"` 
	QuestionType string `json: "questionType"` 
	Options []Option `json: "options"` 
}

type Option struct {
	OptionTitle string `json: "optionTitle"` 
	Correctness string `json: "correctness"`
}

func main() {
	var testData TestData

	browser := rod.New().MustConnect()
	basepage := browser.MustPage("https://developer.android.com/courses/android-basics-compose/course")

	pathways := basepage.MustWaitStable().MustElements(".compose-pathway-link")

	// обробка усіх pathway
	for _, pathway := range pathways {
		pathwayLink := pathway.MustProperty("href").String()

		var unitNum int
		var pathwayNum int
		fmt.Sscanf(pathwayLink, "/courses/pathways/android-basics-compose-unit-%d-pathway-%d", &unitNum, &pathwayNum)
		// TODO: add this data to structure
		testData.Units = append(testData.Units, Unit{})
		curUnit := &testData.Units[len(testData.Units)-1]
		curUnit.UnitName = "TodoUnitName"
		curUnit.Pathways = append(curUnit.Pathways, Pathway{})
		curPathway := &curUnit.Pathways[len(curUnit.Pathways)-1]
		curPathway.PathwayName = "TodoPathwayName"

		page := browser.MustPage(pathwayLink)
		page.MustElement("div.devsite-playlist--item--actions:nth-child(3) > a:nth-child(1)").MustClick()

		qs := page.MustWaitStable().MustElements("li.devsite-quiz-question")

		for _, q := range qs {
			curPathway.Questions = append(curPathway.Questions, Question{})
			curQuestion := &curPathway.Questions[len(curPathway.Questions)-1]

			curQuestion.QuestionTitle = q.MustElement("h2").MustText()
			rawSubTitle := q.MustElements("p")
			for _, st := range rawSubTitle {
				curQuestion.QuestionSubTitle = st.MustText()
			}
			curQuestion.QuestionType = "TODO"
		}
		// TODO: add this data to structure

		// вибір певного варіанту та перевірка його на правильність
		// TODO: вирахувати макс кількість питань
		for i := 0; i < 6; i++ {
			bs := page.MustWaitStable().MustElements("input[value='" + strconv.Itoa(i) + "']");
			for _, b := range bs {
				var nq, nans int

				fmt.Sscanf(b.MustProperty("name").String(), "question-%d", &nq)
				fmt.Sscanf(b.MustProperty("value").String(), "%d",  &nans)
				// TODO: add this data to structure
				curPathway.Questions[nq].Options = append(curPathway.Questions[nq].Options, Option{})
				curOption := curPathway.Questions[nq].Options[len(curPathway.Questions[nq].Options)-1]
				curOption.OptionTitle = "TODO"

				b.MustClick()
			}
			page.MustElement("button.button-primary").MustClick()
			cbs := page.MustWaitStable().MustElements("input.variant-success")
			for _, cb := range cbs {
				var nq, nans int
				fmt.Sscanf(cb.MustProperty("name").String(), "question-%d", &nq)
				fmt.Sscanf(cb.MustProperty("value").String(), "%d",  &nans)
				// TODO: add this data to structure
				fmt.Println(nans)
				curOption := curPathway.Questions[nq].Options[nans]
				curOption.Correctness = "true"
				fmt.Println(curPathway.Questions[nq].Options)
			}
			page.MustElement("button.button").MustClick()
		}

		page.Close()
		break
	}

	fmt.Println(testData)
	//data, _ := json.Marshal(testData)
	//dataOutput := string(data)
	//fmt.Println(dataOutput)
}
