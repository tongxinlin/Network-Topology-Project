$(function(){
	var options = { 
       // target:        '#output',   // target element(s) to be updated with server response 
        beforeSubmit:  showRequest,  // pre-submit callback 
        success:       showResponse,  // post-submit callback
    }; 
 
    // bind to the form's submit event 
    $('#form').submit(function() { 
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
    $("#msg").html("processed successfully");
	$("#output").html(responseText);
}