$(function(){
	var options = { 
       // target:        '#output',   // target element(s) to be updated with server response 
        beforeSubmit:  showRequest,  // pre-submit callback 
        success:       showResponse,  // post-submit callback
    }; 
 
    // bind to the form's submit event 
    $('#index-form').submit(function() { 
        $(this).ajaxSubmit(options); 
        return false; 
    }); 
});

// pre-submit callback
function showRequest() {     
    return true;
} 
 
// post-submit callback 
function showResponse() { 
    $("#index-msg").html("    Uploaded Successfully");
}
