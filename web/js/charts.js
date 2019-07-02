function renderChart(box, canvas) {
	var f = "renderChart" + canvas.type[0].toUpperCase() + canvas.type.substring(1);
	if (typeof (window[f]) != "function") {
		console.error("no chart render function called '" + f + "'");
		return;
	}
	window[f].call(this, box, canvas.chart);
}

function renderChartLine(box, chart) {
	if (chart.series == null) {
		return;
	}
	var instance = echarts.init(box);
	var options = {
		series: chart.series.$map(function (k, series) {
			return {
				name: series.name,
				data: series.values.$map(function (k, v) {
					return v.value;
				}),
				type: "line",
				areaStyle: series.isFilled ? {} : null,
				smooth: series.isSmooth,
				color: (series.color.length > 0) ? series.color : null,
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
		},
		animation: null
	};
	buildChartAxis(chart, options);
	instance.setOption(options);
}

function renderChartPie(box, chart) {
	if (chart.series == null) {
		return;
	}
	var instance = echarts.init(box);
	var options = {
		series: chart.series.$map(function (k, series) {
			var center = ["50%", "50%"];
			if (series.center != null) {
				center = [series.center.x, series.center.y];
			}
			var radius = series.radius;
			if (radius.length == 0) {
				radius = "55%";
			}
			return {
				name: series.name,
				type: "pie",
				radius: radius,
				center: center,
				data: series.values.$map(function (k, value) {
					return {
						value: value.value,
						name: value.label
					};
				})
			};
		}),
		tooltip: {
			show: true,
			trigger: "item"
		},
		animation: null
	};

	instance.setOption(options);
}

function renderChartRadar(box, chart) {
	if (chart.series == null) {
		return;
	}
	var instance = echarts.init(box);
	var center = ["50%", "50%"];
	if (chart.center != null) {
		center = [chart.center.x, chart.center.y];
	}
	var radius = chart.radius;
	if (radius.length == 0) {
		radius = "55%";
	}
	var options = {
		radar: {
			center: center,
			radius: radius,
			indicator: chart.categories.$map(function (k, category) {
				return {
					name: category.name,
					max: category.max
				};
			})
		},
		series: chart.series.$map(function (k, series) {
			return {
				name: series.name,
				type: "radar",
				data: series.subCategories.$map(function (k, subCategory) {
					return {
						name: subCategory.name,
						value: subCategory.values
					};
				})
			};
		}),
		tooltip: {
			show: true,
			trigger: "item"
		},
		animation: null
	};
	instance.setOption(options);
}

function renderChartFunnel(box, chart) {
	if (chart.series == null) {
		return;
	}
	var instance = echarts.init(box);
	var options = {
		series: chart.series.$map(function (k, series) {
			return {
				top: "10%",
				bottom: "10%",
				left: "10%",
				right: "10%",
				name: series.name,
				type: "funnel",
				label: {
					show: true,
					position: "inside"
				},
				data: series.values.$map(function (k, value) {
					return {
						"name": value.label,
						"value": value.value
					};
				})
			};
		}),
		tooltip: {
			show: true,
			trigger: "item"
		},
		animation: null
	};
	instance.setOption(options);
}

function renderChartScatter(box, chart) {
	if (chart.series == null) {
		return;
	}
	var instance = echarts.init(box);
	var options = {
		series: chart.series.$map(function (k, series) {
			var sizeFunc = 10;
			if (series.size.length > 0) {
				if (isFuncString(series.size)) {
					var f = funcString(series.size);
					f.series = series;
					sizeFunc = function (value, data) {
						return f.call(this, f.series.values[data.dataIndex]);
					};
				} else {
					sizeFunc = series.size
				}
			}
			return {
				name: series.name,
				data: series.values.$map(function (k, value) {
					return value.value;
				}),
				type: "scatter",
				label: {
					emphasis: {
						show: true,
						formatter: function (param) {
							return series.values[param.dataIndex].label;
						},
						position: "bottom"
					}
				},
				symbolSize: sizeFunc
			};
		}),
		grid: {
			top: "10%",
			bottom: "10%"
		},
		tooltip: {
			show: true,
			trigger: "item"
		},
		animation: null
	};
	buildChartAxis(chart, options);
	instance.setOption(options);
}

function renderChartBar(box, chart) {
	if (chart.series == null) {
		return;
	}
	var instance = echarts.init(box);
	var options = {
		series: chart.series.$map(function (k, series) {
			var barWidth = series.width;
			if (barWidth == null || barWidth.toString().length == 0) {
				barWidth = 10;
			}
			var stack = series.stack;
			if (stack != null && stack.length == 0) {
				stack = null;
			}
			return {
				name: series.name,
				data: series.values.$map(function (k, v) {
					return v.value;
				}),
				type: "bar",
				color: (series.color.length > 0) ? series.color : null,
				barWidth: barWidth,
				stack: stack
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
		},
		animation: null
	};

	buildChartAxis(chart, options);
	instance.setOption(options);
}

function renderChartGauge(box, chart) {
	if (chart.series == null) {
		return;
	}
	var instance = echarts.init(box);
	var options = {
		series: chart.series.$map(function (k, series) {
			var option = {
				name: series.name,
				type: "gauge",
				data: series.values.$map(function (k, v) {
					return {
						value: v.value,
						name: v.label
					};
				})
			};
			if (chartIsNotEmpty(series.max)) {
				option["max"] = series.max;
			}
			if (chartIsNotEmpty(series.min)) {
				option["min"] = series.min;
			}
			if (chartIsNotEmpty(series.radius)) {
				option["radius"] = series.radius;
			}
			if (chartIsNotEmpty(series.center)) {
				option["center"] = [series.center.x, series.center.y];
			}
			if (chartIsNotEmpty(series.splitNumber)) {
				option["splitNumber"] = series.splitNumber;
			}
			buildChartAxisLine(series.axisLine, option);
			buildChartAxisTick(series.axisTick, option);
			buildChartPointer(series.pointer, option);
			buildChartSplitLine(series.splitLine, option);

			option["detail"] = {};
			if (isFuncString(series.detail)) {
				option["detail"]["formatter"] = funcString(series.detail);
			} else if (chartIsNotEmpty(series.detail)) {
				option["detail"]["formatter"] = series.detail;
			}
			buildChartTextStyle(series.detailStyle, option.detail);
			return option;
		}),
		animation: null
	};
	instance.setOption(options);
}

function renderChartTable(box, chart) {
	console.log(JSON.stringify(chart, "", "  "));
	var table = "<table class=\"ui table\">";
	if (chart.cols != null) {
		var header = "";
		if (chart.cols.$any(function (k, col) {
			return col.header.length > 0;
		})) {
			chart.cols.$each(function (k, col) {
				if (col.header.length == 0) {
					header += "<th>&nbsp;</th>";
				} else {
					header += "<th>" + col.header + "</th>";
				}
			});
		}
		if (header.length > 0) {
			table += "<thead><tr>" + header + "</tr></thead>";
		}
	}
	if (chart.rows != null) {
		chart.rows.$each(function (k, row) {
			var tr = "";
			if (row.values != null) {
				row.values.$each(function (k, v) {
					var width = "";
					if (chart.cols != null && chart.cols.length > k) {
						var col = chart.cols[k];
						if (chartIsNotEmpty(col.width)) {
							width = col.width;
						}
					}

					if (width.length > 0) {
						tr += "<td style=\"width:" + width + "\">" + v + "</td>";
					} else {
						tr += "<td>" + v + "</td>";
					}
				});
			}
			table += "<tr>" + tr + "</tr>"
		});
	}
	table += "</table>";
	box.innerHTML = table;
}

function renderChartEchart(box, chart) {
	var instance = echarts.init(box);
	var code = chart.code;
	eval('(function () {  chart = instance; ' + code + ' })()');
}

function renderChartValue(box, chart) {
	if (chart.value == null) {
		return;
	}
	box.innerHTML = "<div class=\"ui statistic single-value\">" +
		"<div class=\"value\" style=\"color:" + chart.valueColor + "\">" + chart.value.value + "</div>" +
		"<div class=\"label\" style=\"color:" + chart.labelColor + "\">" + chart.value.label + "</div>" +
		"</div>";
}

/**
 * clock
 */
var clockTimer = null;
var clockSeconds = 0;

function renderChartClock(box, chart) {
	var canvasId = "canvas-" + Math.random();
	setTimeout(function () {
		var timestamp = chart.timestamp;
		var diff = new Date().getTime() / 1000 - timestamp;
		var options = {
			rimColour: "#ccc",
			colour: "rgba(255, 0, 0, 0.2)",
			rim: 2,
			markerType: "dot",
			markerDisplay: true,
			addHours: parseInt(diff / 3600, 10),
			addMinutes: parseInt(diff % 3600 / 60, 10),
			addSeconds: parseInt(diff % 60, 10)
		};
		var myClock = new clock(canvasId, options);
		try {
			myClock.start();
		} catch (e) {
		}

		if (clockTimer != null) {
			clearInterval(clockTimer);
			clockSeconds = 0;
		}
		clockTimer = setInterval(function () {
			showClockTime(chart, timestamp);
		}, 1000);

		showClockTime(chart, timestamp);
	});
	box.innerHTML = "<div style=\"position: relative\"> \
				<canvas id=\"" + canvasId + "\" style=\"width:20em;display: block; margin: 0 auto\"></canvas> \
				<div style='text-align:center;margin-top:1em' id='" + "time-" + chart.id + "'></div></div>";
}

function showClockTime(chart, timestamp) {
	clockSeconds++;

	var date = new Date((timestamp + clockSeconds) * 1000);
	var hour = date.getHours().toString();
	var minute = date.getMinutes().toString();
	var second = date.getSeconds().toString();
	var time = ((hour.length == 1) ? "0" + hour : hour) + ":" + ((minute.length == 1) ? "0" + minute : minute) + ":" + ((second.length == 1) ? "0" + second : second);
	Tea.element("#time-" + chart.id).html(time);
}

/**
 * url frame
 */
function renderChartUrl(box, chart) {
	box.innerHTML = '<iframe src="' + chart.url + '" frameborder="0" scrolling="yes"></iframe>';
}

/**
 *
 * html
 */
function renderChartHtml(box, chart) {
	box.innerHTML = chart.html;
}


/**
 * utils
 */
function isFuncString(s) {
	return s != null && s.startsWith("func$")
}

function funcString(s) {
	var f = null;
	eval("f = " + s.substring(5));
	return f;
}

function buildChartAxis(chart, options) {
	options["xAxis"] = {
		type: "category",
		data: chart.labels
	};
	options["yAxis"] = {
		type: "value"
	};
	if (chart.axis != null) {
		if (chart.axis.reverse) {
			options.xAxis["type"] = "value";
			delete (options.xAxis["data"]);

			options.yAxis["type"] = "category";
			options.yAxis["data"] = chart.labels;
		}

		["x", "y"].$each(function (k, axisIndex) {
			if (chart.axis[axisIndex] != null) {
				var current = options.xAxis;
				if (axisIndex == "y") {
					current = options.yAxis;
				}
				if (chart.axis[axisIndex].type.length > 0) {
					current["type"] = chart.axis[axisIndex].type;
				}
				if (chart.axis[axisIndex].max != null && chart.axis[axisIndex].max.toString().length > 0) {
					current["max"] = chart.axis[axisIndex].max;
				}
				if (chart.axis[axisIndex].labels != null) {
					current["data"] = chart.axis[axisIndex].labels;
				}
				if (chart.axis[axisIndex].name != null && chart.axis[axisIndex].name.length > 0) {
					current["name"] = chart.axis[axisIndex].name;
				}
			}
		});


		["x", "y"].$each(function (index, axisIndex) {
			if (options[axisIndex + "Axis"].name != null && options[axisIndex + "Axis"].name.length > 0) {
				var width = chartLettersWidth(options[axisIndex + "Axis"].name) + 10;
				if (options["grid"] == null) {
					if (axisIndex == "x") {
						options["grid"] = {
							right: width
						};
					} else {
						options["grid"] = {
							top: 20
						};
					}
				} else {
					if (axisIndex == "x") {
						if (options["grid"].right == null) {
							options["grid"].right = width;
						} else {
							options["grid"].right += width;
						}
					} else {
						if (options["grid"].top == null) {
							options["grid"].top = 20;
						} else {
							options["grid"].top += 20;
						}
					}
				}
			}
		});
	}
}

function buildChartAxisLine(axisLine, options) {
	if (axisLine == null) {
		return;
	}
	if (chartIsNotEmpty(axisLine.width)) {
		if (options["axisLine"] == null || options["axisLine"]["lineStyle"] == null) {
			options["axisLine"] = {
				lineStyle: {
					width: axisLine.width
				}
			};
		} else {
			options["axisLine"]["lineStyle"]["width"] = axisLine.width;
		}
	}
}

function buildChartAxisTick(axisTick, options) {
	if (axisTick == null) {
		return;
	}
	if (chartIsNotEmpty(axisTick.length)) {
		if (options["axisTick"] == null) {
			options["axisTick"] = {
				length: axisTick.length
			};
		} else {
			options["axisTick"]["length"] = axisTick.length;
		}
	}
	if (chartIsNotEmpty(axisTick.width)) {
		if (options["axisTick"] == null || options["axisTick"]["lineStyle"] == null) {
			options["axisTick"] = {
				lineStyle: {
					width: axisTick.width
				}
			};
		} else {
			options["axisTick"]["lineStyle"]["width"] = axisTick.width;
		}
	}
	if (chartIsNotEmpty(axisTick.color)) {
		if (options["axisTick"] == null || options["axisTick"]["lineStyle"] == null) {
			options["axisTick"] = {
				lineStyle: {
					color: axisTick.color
				}
			};
		} else {
			options["axisTick"]["lineStyle"]["color"] = axisTick.color;
		}
	}
}

function buildChartSplitLine(splitLine, options) {
	if (splitLine == null) {
		return;
	}
	if (chartIsNotEmpty(splitLine.length)) {
		if (options["splitLine"] == null) {
			options["splitLine"] = {
				length: splitLine.length
			};
		} else {
			options["splitLine"]["length"] = splitLine.length;
		}
	}
	if (chartIsNotEmpty(splitLine.width)) {
		if (options["splitLine"] == null || options["splitLine"]["lineStyle"] == null) {
			options["splitLine"] = {
				lineStyle: {
					width: splitLine.width
				}
			};
		} else {
			options["splitLine"]["lineStyle"]["width"] = splitLine.width;
		}
	}
	if (chartIsNotEmpty(splitLine.color)) {
		if (options["splitLine"] == null || options["splitLine"]["lineStyle"] == null) {
			options["splitLine"] = {
				lineStyle: {
					color: splitLine.color
				}
			};
		} else {
			options["splitLine"]["lineStyle"]["color"] = splitLine.color;
		}
	}
}

function buildChartPointer(pointer, options) {
	if (pointer == null) {
		return;
	}
	if (chartIsNotEmpty(pointer.width)) {
		if (options["pointer"] == null) {
			options["pointer"] = {
				width: pointer.width
			};
		} else {
			options["pointer"]["width"] = pointer.width;
		}
	}
	if (chartIsNotEmpty(pointer.length)) {
		if (options["pointer"] == null) {
			options["pointer"] = {
				length: pointer.length
			};
		} else {
			options["pointer"]["length"] = pointer.length;
		}
	}
}

function buildChartTextStyle(textStyle, options) {
	if (textStyle == null) {
		return;
	}
	if (chartIsNotEmpty(textStyle.color)) {
		options["color"] = textStyle.color;
	}
	if (chartIsNotEmpty(textStyle.fontSize)) {
		options["fontSize"] = textStyle.fontSize;
	}
	if (chartIsNotEmpty(textStyle.fontFamily)) {
		options["fontFamily"] = textStyle.fontFamily;
	}
	if (chartIsNotEmpty(textStyle.fontWeight)) {
		options["fontWeight"] = textStyle.fontWeight;
	}
}

function chartLettersWidth(characters) {
	if (characters == null || characters.length == 0) {
		return 0;
	}
	var span = document.createElement("span");
	if (span.innerText != null) {
		span.innerText = characters;
	} else if (span.textContent != null) {
		span.textContent = characters;
	}
	span.style.cssText = "font-family:'sans-serif';font-size:16px;visibility:hidden";
	document.body.appendChild(span);
	var result = span.offsetWidth;
	document.body.removeChild(span);
	return result;
}

function chartIsNotEmpty(v) {
	return v != null && v.toString().length > 0;
}