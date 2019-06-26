Tea.context(function () {
	this.instances = [];
	this.currentInstance = null;
	this.currentDashboard = null;
	this.info = null;

	this.$delay(function () {
		this.loadInstances();
		this.loadInfo();
	});

	this.loadInstances = function () {
		this.$get("/instances")
			.success(function (resp) {
				this.instances = resp.data.instances;

				if (this.instances.length > 0) {
					var instanceId = this.queryArg("instance");
					if (instanceId.length == 0) {
						instanceId = this.instances[0].id;
					}

					this.$get("/instance")
						.params({
							"instance": instanceId
						})
						.success(function (resp) {
							this.currentInstance = resp.data;
							var dashboardId = this.queryArg("dashboard");
							if (dashboardId.length == 0) {
								if (this.currentInstance.dashboards.length > 0) {
									this.currentDashboard = this.currentInstance.dashboards[0];
									this.$delay(function () {
										this.renderDashboard();
									});
								}
								return
							}
							if (dashboardId.length > 0) {
								this.currentDashboard = this.currentInstance.dashboards.$find(function (k, v) {
									return v.id == dashboardId;
								});
								this.$delay(function () {
									this.renderDashboard();
								});
							}
						});
				}
			});
	};

	/**
	 * dashboard
	 */
	this.shouldUpdateDashboardVersion = "";

	this.renderDashboard = function () {
		var that = this;
		var template = this.info.dashboardTemplates.$find(function (k, v) {
			return v.id == that.currentDashboard.id;
		});
		if (template != null && this.currentDashboard.version != template.version) {
			this.shouldUpdateDashboardVersion = this.currentDashboard.version + "->" + template.version;
		}

		this.$get("/dashboard")
			.params({
				"instance": this.currentInstance.id,
				"dashboard": this.currentDashboard.id
			})
			.success(function (resp) {
				var dashboard = resp.data;
				var that = this;
				dashboard.charts.$each(function (k, canvas) {
					var box = that.$find("#chart-box-" + canvas.id + " .canvas");
					box[0].style.height = (canvas.heightPercent * 180 * window.innerWidth / 1024) + "px";
					that.$delay(function () {
						renderChart(box[0], canvas);
					});
				});
			});
	};

	this.loadInfo = function () {
		this.$get("/api/all")
			.success(function (resp) {
				this.info = resp.data;
			});
	};

	this.updateInstance = function () {
		if (!window.confirm("Confirm to update dashboards?")) {
			return;
		}
		this.$get("/instance/update")
			.params({
				"instance": this.currentInstance.id
			})
			.success(function () {
				window.location = "/?instance=" + this.currentInstance.id;
			});
	};

	this.runOperation = function (operation) {
		if (!window.confirm("Confirm to run the operation?")) {
			return;
		}
		var params = [];
		if (this.info.options != null && this.currentInstance.options != null) {
			var that = this;
			this.info.options.$each(function (k, option) {
				var code = option.code;
				if (typeof (that.currentInstance.options[code]) != "undefined") {
					params.push(code + "=" + escape(that.currentInstance.options[code]));
				}
			});
		}
		this.$get("/api/run/" + operation + "?" + params.join("&"))
			.success(function (resp) {
				var output = resp.data.output;
				if (output.length == 0) {
					alert("[no output]");
				} else {
					alert("output: " + output);
				}
			});
	};

	this.queryArg = function (name) {
		var query = window.location.search;
		var markerIndex = query.indexOf("?");
		if (markerIndex < 0) {
			return "";
		}
		query = query.substring(1);
		var params = query.split("&");
		for (var i = 0; i < params.length; i++) {
			var pieces = params[i].split("=", 2);
			if (pieces[0] == name) {
				if (pieces.length == 1) {
					return "";
				}
				return pieces[1];
			}
		}
		return "";
	};

	this.notImplement = function () {
		alert("The feature has not been implemented yet.");
	};
});