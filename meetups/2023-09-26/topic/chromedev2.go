var img []byte
for selectedDate.After(firstDate) {
	selectedDate = selectedDate.AddDate(0, 0, -1)
	date := selectedDate.Format("2006-01-02")
	fmt.Println("scaning date " + date)
	err = chromedp.Run(ctx,
		chromedp.EmulateViewport(2560, 1440),
		dayBefore(&img),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(path.Join("graph", date+".png"), img, 0o644); err != nil {
		log.Fatal(err)
	}
}
