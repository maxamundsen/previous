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
			PieChart(),
			LineChart(),
			BarChart(),
		),

	).Render(w)
}

func BarChart() Node {
	label_string := MakeJsArray(
		[]string{"Item1", "Item2", "Item3", "Item4"},
	)

	data_string := MakeJsArray(
		[]int{3, 5, 8, 13},
	)

	return Card(
		InlineStyle(`
			$me {
				display: flex;
				align-items: center;
				justify-content: center;
				max-height: $100;
			}
		`),
		Canvas(),
		InlineScriptf(`
			let ctx = me("canvas", me());

			new Chart(ctx, {
				type: 'bar',
				data: {
					labels: %s,
					datasets: [{
						label: 'Acquisitions by year',
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

func PieChart() Node {
	return Card(
		InlineStyle(`
			$me {
				display: flex;
				align-items: center;
				justify-content: center;
				max-height: $100;
			}
		`),
		Canvas(),
		InlineScript(`
			let ctx = me("canvas", me());

			let data = {
				labels: [
					'Red',
					'Blue',
					'Yellow'
				],
				datasets: [{
					label: 'My First Dataset',
					data: [300, 50, 100],
					backgroundColor: [
						'rgb(255, 99, 132, 50)',
						'rgb(54, 162, 235, 50)',
						'rgb(255, 205, 86, 50)'
					],
					hoverOffset: 4
				}]
			};

			let config = {
				type: 'pie',
				data: data,
			};

			new Chart(ctx, config);
		`),
	)
}

func LineChart() Node {
	return Card(
		InlineStyle(`
			$me {
				display: flex;
				align-items: center;
				justify-content: center;
				max-height: $100;
			}
		`),
		Canvas(),
		InlineScript(`
			let ctx = me("canvas", me());

			let data = {
				labels: ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'],
				datasets: [{
					label: 'Nothing Ever Happens',
					data: [65, 59, 80, 81, 56, 55, 40],
					fill: false,
					borderColor: 'rgb(75, 192, 192)',
					tension: 0.1
				}]
			};

			let config = {
				type: 'line',
				data: data,
			};

			new Chart(ctx, config);
		`),
	)
}