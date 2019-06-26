Tea.context(function () {
	this.options = [];

	this.$delay(function () {
		this.$find("input[name='name']").focus();
		this.loadOptions();
	});

	this.loadOptions = function () {
		this.$get("/api/options")
			.success(function (resp) {
				this.options = buildFormFromOptions(resp.data);
			});
	};

	this.submitSuccess = function (resp) {
		alert("instance created successfully");
		window.location = "/?instance=" + resp.data.instanceId;
	};
});