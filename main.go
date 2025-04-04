package main

import (
	"fmt"
	"strconv"
	"time"

	//"encoding/json"

	"github.com/go-rod/rod"
)

type TestData struct {
	units []Unit
}

type Unit struct {
	unitName string
	pathways []Pathway
}

type Pathway struct {
	pathwayName string
	questions []Question
}

type Question struct {
	questionTitle string
	questionSubTitle string
	questionType string
	options []Option
}

type Option struct {
	optionTitle string
	correctness string
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
		testData.units = append(testData.units, Unit{})
		curUnit := &testData.units[len(testData.units)-1]
		curUnit.unitName = "TodoUnitName"
		curUnit.pathways = append(curUnit.pathways, Pathway{})
		curPathway := &curUnit.pathways[len(curUnit.pathways)-1]
		curPathway.pathwayName = "TodoPathwayName"

		page := browser.MustPage(pathwayLink)
		page.MustElement("div.devsite-playlist--item--actions:nth-child(3) > a:nth-child(1)").MustClick()

		qs := page.MustWaitStable().MustElements("devsite-quiz-question")
		for _, q := range qs {
			curPathway.questions = append(curPathway.questions, Question{})
			curQuestion := &curPathway.questions[len(curPathway.questions)-1]

			curQuestion.questionTitle = q.MustElement("h2").MustText()
			curQuestion.questionSubTitle = q.MustElement("p").MustText()
			curQuestion.questionType = "TODO"
			fmt.Println(curQuestion.questionSubTitle)
		}
		// TODO: add this data to structure

		// вибір певного варіанту та перевірка його на правильність
		// TODO: вирахувати макс кількість питань
		for i := 0; i < 4; i++ {
			bs := page.MustWaitStable().MustElements("input[value='" + strconv.Itoa(i) + "']");
			for _, b := range bs {
				var nq, nans int
				fmt.Sscanf(b.MustProperty("name").String(), "question-%d", &nq)
				fmt.Sscanf(b.MustProperty("value").String(), "%d",  &nans)
				// TODO: add this data to structure

				b.MustClick()
			}
			page.MustElement("button.button-primary").MustClick()
			cbs := page.MustWaitStable().MustElements("input.variant-success")
			for _, cb := range cbs {
				var nq, nans int
				fmt.Sscanf(cb.MustProperty("name").String(), "question-%d", &nq)
				fmt.Sscanf(cb.MustProperty("value").String(), "%d",  &nans)
				// TODO: add this data to structure
			}
			page.MustElement("button.button").MustClick()
		}

		page.Close()
	}

	time.Sleep(time.Hour)
}
