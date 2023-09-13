var scan = chromedp.ActionFunc(func(ctx context.Context) error {
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventResponseReceived: // HL
			if strings.Contains(ev.Response.MimeType, "json") { // HL
				go func(ev *network.EventResponseReceived) {
					body, err := network.GetResponseBody(ev.RequestID).Do(ctx) // HL
					...
					var x jsonResponse
					if err := json.Unmarshal(body, &x); err == nil {
						communicate <- x // HL
					}
				}(ev)
			}
		}
	})
	return nil
})
