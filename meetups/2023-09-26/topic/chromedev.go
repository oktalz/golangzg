ctx, cancel := chromedp.NewContext(
	context.Background(),
)
defer cancel()

ctx, cancel = context.WithTimeout(ctx, 60*time.Second*60)
defer cancel()

var buf []byte

err := chromedp.Run(ctx,
	chromedp.EmulateViewport(2560, 1440), // HL
	scan,                                 // HL
	submit(initialURL, 90, &buf))         // HL
if err != nil {
	return err
}
