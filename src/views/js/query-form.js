$(function(){
	var options = { 
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
    document.getElementById("query-output").innerHTML = responseText;
}

$(function(){
    var options = {
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
    document.getElementById("assessment-output").innerHTML = responseText;
}