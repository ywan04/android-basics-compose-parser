package main

import (
	"fmt"
	"strconv"
	"time"

	"encoding/json"

	"github.com/go-rod/rod"
)

type TestData struct {
	units []struct {
		unitName string
		pathways []struct {
			pathwayName string
			questions []struct {
				questionTitle string
				questionSubTitle string
				questionType string
				options []struct {
					optionTitle string
					correctness string
				}
			}
		}
	}
}

func main() {
	var testData TestData

	browser := rod.New().MustConnect()
	basepage := browser.MustPage("https://developer.android.com/courses/android-basics-compose/course")

	pathways := basepage.MustWaitStable().MustElements(".compose-pathway-link")

	// обробка усіх pathway
	for _, pathway := range pathways {
		pathwayLink = pathway.MustProperty("href").String()

		var unitNum int
		var pathwayNum int
		fmt.Sscanf(pathwayLink, "/courses/pathways/android-basics-compose-unit-%d-pathway-%d", &unitNum, &pathwayNum)
		// TODO: add this data to structure

		page := browser.MustPage(pathwayLink)
		page.MustElement("div.devsite-playlist--item--actions:nth-child(3) > a:nth-child(1)").MustClick()

		qs := page.MustWaitStable().MustElements("devsite-quiz-question")
		for _, q := range qs {
			qNum := q.MustProperty("data-index")
			qTitle := q.MustElement("h2").MustText()
			qSubTitle := q.MustElement("p").MustText()
		}
		// TODO: add this data to structure

		// вибір певного варіанту та перевірка його на правильність
		// TODO: вирахувати макс кількість питань
		for i := 0; i < 4; i++ {
			bs := page.MustWaitStable().MustElements("input[value='" + strconv.Itoa(i) + "']");
			for _, b := range bs {
				var nq, nans int
				fmt.Sscanf(cb.MustProperty("name").String(), "question-%d", &nq)
				fmt.Sscanf(cb.MustProperty("value").String(), "%d",  &nans)
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
