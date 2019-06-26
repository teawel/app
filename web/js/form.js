function buildFormFromOptions(options) {
	var result = [];
	options.$each(function (k, option) {
		var f = "buildOption" + option.type[0].toUpperCase() + option.type.substring(1);
		var elements = window[f].apply(this, [option]);
		elements.$each(function (_, element) {
			result.push({
				"title": option.title,
				"subtitle": option.subtitle,
				"description": option.description,
				"html": element
			});
		});
	});

	//console.log(JSON.stringify(result, "", " "));
	return result;
}

function buildOptionString(option) {
	if (option.attrs == null) {
		option.attrs = [];
	}
	if (option.maxLength > 0) {
		option.attrs["maxlength"] = option.maxlength.toString();
	}
	if (option.placeholder.length > 0) {
		option.attrs["placeholder"] = option.placeholder;
	}
	if (option.value != null) {
		option.attrs["value"] = option.value;
	}
	if (option.size > 0) {
		option.attrs["size"] = option.size.toString();
	}
	option.attrs["name"] = option.namespace + "_" + option.code;
	if (option.rightLabel.length == 0) {
		var result = "<input type=\"text\" " + buildOptionAttrs(option.attrs) + "/>";
	} else {
		var result = "<div class=\"ui input right labeled\"><input type=\"text\" " + buildOptionAttrs(option.attrs) + "/><label class=\"ui label\">" + option.rightLabel + "</label></div>";
	}

	return [result];
}

function buildOptionAttrs(attrs) {
	if (attrs == null) {
		return "";
	}
	var result = [];
	for (var key in attrs) {
		if (!attrs.hasOwnProperty(key)) {
			continue;
		}
		result.push(key + "=\"" + buildOptionEscape(attrs[key]) + "\"");
	}
	return result.join(" ")
}

function buildOptionEscape(value) {
	return value.replace(/"/g, "&quot;");
}