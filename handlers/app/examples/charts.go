package examples

import (
	. "previous/ui"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"previous/middleware"

	"net/http"
)

func ChartHandler(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	session := middleware.GetSession(r)

	AppLayout("Chart.js Example", LAYOUT_SECTION_EXAMPLES, *identity, session,
		Card(
			P(Text("Since we have the ability to \"inline\" Javascript inside our UI components, we can build reusable JS-dependent charts. This example uses the Chart.js library.")),
		),

		Br(),

		Grid2x2(
			BarChart(),
			BarChart(),
			BarChart(),
			BarChart(),
		),

	).Render(w)
}

func BarChart() Node {
	label_string := MakeJsArray(
		[]string{"test", "hello", "item3", "spiral"},
	)

	data_string := MakeJsArray(
		[]int{3, 5, 8, 13},
	)

	return Card(
		Canvas(),
		InlineScriptf(`
			let ctx = me("canvas", me());

			new Chart(ctx, {
				type: 'bar',
				data: {
					labels: %s,
					datasets: [{
						data: %s,
						borderWidth: 1
					}]
				},
			});`,
			label_string,
			data_string,
		),
	)
}
