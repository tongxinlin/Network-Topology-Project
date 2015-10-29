// temp output till we get the DB query working. 

var outputFileReader;



var button = document.getElementById("refresh");

button.addEventListener("click", function(e) {
	readSingleFile("Network-Topology-Project/src/tmp/output/calc_result_dir.txt")
}, false);



/* 
 * checks for the current file API's on the useres browser.  
 */
function checkForFileAPIs () {
	if (window.File && window.FileReader && window.FileList && window.Blob) {
		outputFileReader = new FileReader();
		return true;
	}
	else {
	
		alert("The File API is not fully supported on your broswer");
		return false;
		
	}
}
\
/*
 * Takes in a file path and pushes the contents of the file into an array that we can 
 * place in a html tag. 
 * 
 */
function readSingleFile(filePath) {
	var output = "";
	
	if(filePath.files && filePath.files[0]) {
		outputFileReader.onload = funciton(e) {
			
			output = e.target.result;
			dsiplayTextAsHTMl(output);
		};
		
		reader.readAsText(filePath.files[0]);
	}
}

/*
 * "embeds" the text from the file as HTML in the output container on the hompapge. 
 */
function dsiplayTextAsHTMl(txt) {
	var conatiner = document.getElementById("output");
	conatiner.innterHTML = txt;
}






