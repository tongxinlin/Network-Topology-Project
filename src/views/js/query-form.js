$(function(){
	var options = { 
       // target:        '#output',   // target element(s) to be updated with server response 
        beforeSubmit:  showRequest,  // pre-submit callback 
        success:       showResponse,  // post-submit callback
    }; 
 
    // bind to the form's submit event 
    $('#query-form').submit(function() { 
        $(this).ajaxSubmit(options); 
        return false; 
    }); 
});

// pre-submit callback
function showRequest(formData) {     
    return true;
} 
 
// post-submit callback 
function showResponse(responseText) { 
    //$("#query-msg").html("Processed Successfully");
    document.getElementById("query-output-example").innerHTML = "";
    var result ="<br/>" + responseText + "<br/>";
    document.getElementById("query-output").innerHTML += result;
}

$(function(){
    var options = { 
       // target:        '#output',   // target element(s) to be updated with server response 
        beforeSubmit:  showRequest2,  // pre-submit callback 
        success:       showResponse2,  // post-submit callback
    }; 
 
    // bind to the form's submit event 
    $('#reachability-form').submit(function() { 
        $(this).ajaxSubmit(options); 
        return false; 
    }); 
});

// pre-submit callback
function showRequest2(formData) {     
    return true;
} 
 
// post-submit callback 
function showResponse2(responseText) { 
    document.getElementById("assessment-output-example").innerHTML = "";
    var result = "result" + "<br/><br/>" + responseText + "<br/><br/>";
    document.getElementById("assessment-output").innerHTML += result;
}