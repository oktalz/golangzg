func submit(urlstr string, quality int, res *[]byte, date *[]byte, yesterday *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(`//input[@id="username"]`), // HL
		chromedp.WaitVisible(`//input[@id="value"]`),
		chromedp.WaitVisible(`//span[@id="submitDataverify"]`),      // HL
		chromedp.SetValue(`#username`, `myusername`, chromedp.ByID), // HL
		chromedp.SetValue(`#value`, `mypassword`, chromedp.ByID),
		chromedp.Click(`#submitDataverify`, chromedp.ByID),
		chromedp.WaitVisible(`input[placeholder="Select date"]`),
		chromedp.Screenshot(`.nco-single-energy canvas[data-zr-dom-id="zr_0"]`, date, chromedp.NodeVisible, chromedp.ByQuery), // HL
		chromedp.Click(`span[role="img"][aria-label="left"][title="Previous day"]`, chromedp.ByQuery),                         // HL
		chromedp.Sleep(time.Second * 5),
		chromedp.Screenshot(`.nco-single-energy canvas[data-zr-dom-id="zr_0"]`, yesterday, chromedp.NodeVisible, chromedp.ByQuery),
	}
}
