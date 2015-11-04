


window.onload = function (/*maxNumberOfNodes*/){   
	
	var selectSources, selectdestinations, aDestNode, aSrcNode;
	var maxNumberOfNodes = 15;
	
	for(var i = 0; i <= maxNumberOfNodes; i++) {
		selectSources = document.getElementById("src");
		console.log(selectSources);
		aSrcNode = document.createElement("OPTION");
		console.log(aSrcNode);
		selectSources.options.add(aSrcNode);
		aSrcNode.text = i;
		aSrcNode.value = i;
	}
	
	
	for(var i = 0; i <= maxNumberOfNodes; i++) {
		selectdestinations = document.getElementById("dest");
		aDestNode = document.createElement("OPTION");
		selectdestinations.options.add(aDestNode);
		aDestNode.text = i;
		aDestNode.value = i;
	}
}


// somehow query the DB to get the max number of nodes and create a drop down list 
// that holds that specifc amount