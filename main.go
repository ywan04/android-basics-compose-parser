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
		page := browser.MustPage(pathway.MustProperty("href").String())

		page.MustElement("div.devsite-playlist--item--actions:nth-child(3) > a:nth-child(1)").MustClick()

		// вибір певного варіанту та перевірка його на правильність
		for i := 0; i < 4; i++ {
			bs := page.MustWaitStable().MustElements("input[value='" + strconv.Itoa(i) + "']");
			for _, b := range bs {
				b.MustClick()
			}
			page.MustElement("button.button-primary").MustClick()
			cbs := page.MustWaitStable().MustElements("input.variant-success")
			for _, cb := range cbs {
				var nq, nans int
				fmt.Sscanf(cb.MustProperty("name").String(), "question-%d", &nq)
				fmt.Sscanf(cb.MustProperty("value").String(), "%d",  &nans)
				a[nq][nans] = true
			}
			page.MustElement("button.button").MustClick()
		}

		// виведення правильних відповідей
		for i := 0; i < 10; i++ {
			fmt.Printf("%2d:\n", i+1)
			for j := 0; j < 4; j++ {
				fmt.Printf("  - %t\n", a[i][j])
			}
		}

		page.Close()
	}

	time.Sleep(time.Hour)
}
