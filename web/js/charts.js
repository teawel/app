function renderChart(box, canvas) {
	var instance = echarts.init(box);
	var f = "renderChart" + canvas.type[0].toUpperCase() + canvas.type.substring(1);
	var options = window[f].call(this, canvas.chart);
	instance.setOption(options);
}

function renderChartLine(chart) {
	return {
		xAxis: {
			type: "category",
			data: chart.labels
		},
		yAxis: {
			type: "value",
			scale: true
		},
		series: chart.lines.$map(function (k, line) {
			return {
				data: line.values.$map(function (k, v) {
					return v.value;
				}),
				type: "line",
				areaStyle: line.isFilled ? {} : null,
				smooth: line.isSmooth,
				color: (line.color.length > 0) ? line.color : null,
				lineStyle: {
					width: 1.5
				}
			};
		}),
		grid: {
			top: 10,
			bottom: 20,
			right: 10
		},
		tooltip: {
			show: true,
			trigger: "axis"
		}
	};
}